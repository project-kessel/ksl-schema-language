# KSL: Kessel Schema Language

**Kessel Schema Language (KSL)** is a domain-specific language for defining authorization schemas. It compiles to [SpiceDB](https://github.com/authzed/spicedb) `.zed` schema files, which power the access-control engine in [Project Kessel](https://github.com/project-kessel).

This guide explains the core concepts behind KSL and how to write KSL files. It is an **explanation** document — focused on building understanding rather than walking through a specific task. If you are looking to migrate an existing service, see the [Migrate from RBAC v1 to RBAC v2](https://project-kessel.github.io/docs/building-with-kessel/how-to/migrate-from-rbac-v1-to-v2/) how-to guide instead.

---

## Before you begin

KSL sits on top of the Kessel RBAC model. Before diving into KSL syntax, it helps to be comfortable with these concepts:

| Topic | Where to read |
|---|---|
| How roles, role bindings, and permissions work in Kessel RBAC | [Role-based access control concepts](https://project-kessel.github.io/docs/building-with-kessel/concepts/rbac/) |
| How the authorization graph and permission checks work | [Relationships and permissions](https://project-kessel.github.io/docs/building-with-kessel/concepts/relationships-permissions/) |
| How V1 and V2 RBAC differ | [Coming from RBAC v1](https://project-kessel.github.io/docs/building-with-kessel/concepts/coming-from-rbac-v1/) |
| The insights-rbac service architecture and data model | [insights-rbac architecture](https://github.com/project-kessel/insights-rbac/blob/master/docs/ARCHITECTURE.md) |
| How role bindings are stored and replicated to SpiceDB | [Role bindings internals](https://github.com/project-kessel/insights-rbac/blob/master/docs/role_bindings.md) |

---

## Table of Contents

1. [Why KSL exists](#why-ksl-exists)
2. [Core concepts](#core-concepts)
3. [Anatomy of a KSL file](#anatomy-of-a-ksl-file)
4. [Types](#types)
5. [Relations](#relations)
6. [Cardinality](#cardinality)
7. [Importing other namespaces](#importing-other-namespaces)
8. [Extensions: reusable templates](#extensions-reusable-templates)
9. [Applying extensions](#applying-extensions)
10. [Building a schema](#building-a-schema)
11. [Understanding the output](#understanding-the-output)
12. [A real-world example](#a-real-world-example)
13. [Quick reference](#quick-reference)

---

## Why KSL exists

SpiceDB's native `.zed` schema format is expressive but flat — there is no namespacing, no imports, and no code reuse. When many teams share one authorization schema, you end up with a single enormous file that is hard to own, hard to read, and full of copy-pasted patterns.

KSL solves this by adding:

- **Namespaces** — each service owns its own slice of the schema
- **Imports** — namespaces can reference types from other namespaces
- **Extensions** — reusable, parameterized templates that inject relations into types
- **Visibility** — control which types and relations are accessible to other namespaces

KSL files compile down to a single SpiceDB `.zed` schema that the authorization engine evaluates at permission-check time.

---

## Core concepts

Three building blocks make up every KSL schema:

| Concept | What it is |
|---|---|
| **Namespace** | A named container for a set of types. Corresponds roughly to a service (e.g., `rbac`, `drive`, `comments`). |
| **Type** | A category of resource that can have permissions applied to it (e.g., `workspace`, `document`, `role`). |
| **Relation** | A named connection from one type to another, or a computed permission derived from other relations. |

The authorization question SpiceDB answers is always: _"Does subject X have relation Y on object Z?"_ KSL describes exactly how to compute that answer.

---

## V2 permission naming

Before writing any KSL, it is worth understanding how V2 permission names are formed from V1 permissions, since this convention appears throughout KSL schema files.

V1 permissions use a three-part `app:resource:verb` format (e.g., `drive:document:read`). V2 names are derived from V1 by:

1. Replacing colons with underscores: `drive:document:read` → `drive_document_read`
2. Changing the verb to avoid collisions with the raw V1 name:
   - `read` → `view`
   - `write` → `edit`

So `drive:document:read` becomes `drive_document_view`, and `drive:document:write` becomes `drive_document_edit`.

The V1 name (`drive_document_read`) is still stored on the role for backward compatibility — the V2 name (`drive_document_view`) is what Kessel checks at runtime. The KSL extension system wires these together automatically.

---

## Anatomy of a KSL file

Every KSL file follows the same structure:

```
version 0.1
namespace <name>

import <other-namespace>    // optional, repeat as needed

// type definitions, extension definitions, extension invocations
```

A complete, minimal example — a document service that borrows its workspace model from the shared RBAC foundation:

```ksl
version 0.1
namespace drive
import rbac

public type document {
    private relation workspace: [ExactlyOne rbac.workspace]
    public relation view: workspace.drive_document_view
}
```

**Rules:**
- Every file must start with `version 0.1` and a `namespace` declaration.
- A file defines exactly one namespace.
- `import` statements come before any type definitions.
- Line comments use `//`.

---

## Types

Use a `type` block to define a category of resource that can have permissions applied to it:

```ksl
public type workspace {
    relation parent: [AtMostOne workspace]
    relation member: [Any principal or group.member]
}
```

### Visibility

Visibility controls which other namespaces can reference a type:

| Keyword | Within this namespace | From other namespaces |
|---|---|---|
| `public` | Yes | Yes |
| `internal` | Yes | No |
| `private` | No | No |

When omitted, the compiler defaults to `public`.

```ksl
internal type role_binding {   // only usable within this namespace
    relation subject: [Any principal]
    relation granted: [Any role]
}
```

---

## Relations

A **relation** is a named edge that either connects objects of one type to another, or computes a derived permission from other relations. The syntax is:

```ksl
[visibility] relation <name>: <body>
```

There are four kinds of relation body:

### 1. Self — a direct connection to typed objects

Use `[cardinality? type (or type)*]` to declare which types of objects can be linked via this relation.

```ksl
relation subject: [Any principal or group.member]
relation workspace: [ExactlyOne rbac.workspace]
relation parent: [AtMostOne workspace]
```

The square brackets hold a list of allowed types, separated by `or`. See [Cardinality](#cardinality) for the `Any`, `ExactlyOne`, etc. keywords.

### 2. Reference — an alias for another relation on the same type

```ksl
relation read: workspace
```

This says "`read` resolves to the same set of objects as `workspace`."

### 3. Sub-relation — traverse to another type's relation

```ksl
relation view: workspace.drive_document_view
```

This says "`view` is computed by following the `workspace` link to the linked workspace object, and then evaluating `drive_document_view` on that workspace." In other words: a user can `view` this document if they have `drive_document_view` on the document's workspace.

### 4. Boolean algebra — combine the above

Use `or`, `and`, and `unless` to build composite permissions:

```ksl
// union: member directly, OR a member through a group
relation member: subject or group.member

// intersection: user must have BOTH permissions simultaneously
relation comments_comment_view: drive_document_view and comments_comment_view_assigned

// exclusion: granted access, but not if explicitly banned
relation effective: direct or parent.effective unless banned
```

Parentheses are supported for grouping:

```ksl
relation view: user_grant.view or (parent.view unless parent.blocked)
```

### The `bool` pseudo-type

`bool` is a shorthand that means "any principal in the system" (a wildcard). Use it when a role permission applies to anyone who holds that role, rather than a specific named principal:

```ksl
relation drive_document_view: [bool]   // any principal who holds this role
```

---

## Cardinality

Inside self-relation brackets, an optional cardinality keyword constrains how many objects can be linked:

| Keyword | Meaning |
|---|---|
| `Any` | Zero or more (default when omitted) |
| `AtMostOne` | Zero or one |
| `ExactlyOne` | Exactly one |
| `AtLeastOne` | One or more |

```ksl
relation workspace: [ExactlyOne rbac.workspace]   // a document belongs to exactly one workspace
relation parent: [AtMostOne workspace]            // a workspace may or may not have a parent
relation member: [Any principal or group.member]  // any number of members
```

Cardinality is captured in the intermediate representation and available to validation tooling. It does not currently produce SpiceDB-level enforcement.

---

## Importing other namespaces

To reference a type defined in another namespace, declare an import and use `namespace.type` syntax:

```ksl
version 0.1
namespace drive
import rbac

public type document {
    private relation workspace: [ExactlyOne rbac.workspace]
}
```

The imported namespace (`rbac`) must be compiled alongside this file — you cannot import a namespace that is not part of the same build invocation. Only `public` types from another namespace can be imported; `internal` and `private` types are not accessible from outside their namespace.

---

## Extensions: reusable templates

Extensions are parameterized templates that inject relations into types when applied. They are the primary code-reuse mechanism in KSL, and the way the shared RBAC foundation injects permissions into your service's types.

### Defining an extension

```ksl
[visibility] extension <name>(<param1>, <param2>, ...) {
    type <typeName> {
        [allow_duplicates] [visibility] relation <dynamicName>: <dynamicBody>
    }
    ...
}
```

An extension can target multiple types at once. Relation names and bodies can reference parameters using `${param}` substitution. Built-in implicit variables are also available:

| Variable | Value |
|---|---|
| `${NAMESPACE}` | The namespace invoking the extension |
| `${TYPE}` | The type the extension is applied to (when applied within a type) |
| `${RELATION}` | The relation the extension is applied to (when applied on a relation) |
| `${MODULE}` | Same as `${NAMESPACE}` |

For names composed of multiple parts, use backtick template syntax:

```ksl
relation `${app}_${resource}_${verb}`: [bool]
```

### A simple extension example

```ksl
public extension workspace_permission(full_name) {
    type role {
        relation ${full_name}: [bool]
        allow_duplicates relation `${MODULE}_all_all`: [bool]
    }
    type workspace {
        public relation ${full_name}: user_grant.${full_name} or parent.${full_name}
    }
}
```

When applied from the `drive` namespace with `full_name: 'drive_document_view'`, this injects:

- Into `role`: a `drive_document_view` relation (directly assignable via `[bool]`) and a shared wildcard `drive_all_all` (`${MODULE}` resolves to `drive`, the calling namespace)
- Into `workspace`: a `drive_document_view` permission computed from the workspace's role bindings and its parent workspace

### `allow_duplicates`

When multiple extension calls each try to add the same relation name to a type (for example, a shared wildcard like `drive_all_all` that every `drive` permission contributes), KSL would normally raise a "symbol already exists" error. The `allow_duplicates` modifier tells the extension system to silently skip the relation if it already exists:

```ksl
allow_duplicates relation `${MODULE}_all_all`: [bool]
```

This is essential for wildcard catch-all relations that are shared across many permissions.

---

## Applying extensions

Apply an extension with `@namespace.extensionName(param: 'value', ...)`.

Extensions can be applied in two places:

### At the namespace (file) level

The most common usage. Place the invocation as a top-level statement ending with a semicolon:

```ksl
version 0.1
namespace drive
import rbac

// "drive:document:read" → V2 name "drive_document_view" (read becomes view)
@rbac.add_v1_based_permission(app:'drive', resource:'document', verb:'read', v2_perm:'drive_document_view');

// "drive:document:write" → V2 name "drive_document_edit" (write becomes edit)
@rbac.add_v1_based_permission(app:'drive', resource:'document', verb:'write', v2_perm:'drive_document_edit');
```

Each call injects relations into the types listed in the extension definition (e.g., `role`, `role_binding`, `workspace`), all within the `rbac` namespace.

### Inside a type body

Apply an extension within a type when you also want the extension to interact with that type's own relations:

```ksl
public type document {
    private relation workspace: [ExactlyOne rbac.workspace]

    @rbac.workspace_permission(full_name:'drive_document_view')
    public relation view: workspace.drive_document_view
}
```

Here, the `@rbac.workspace_permission` call injects relations into `rbac/role` and `rbac/workspace`, while `view` is also defined directly on `document`.

---

## Building a schema

The `ksl` CLI compiles `.ksl` files (and optionally pre-compiled `.json` intermediate files) into a SpiceDB `.zed` schema.

### Install

```bash
go install github.com/project-kessel/ksl-schema-language/cmd/ksl@latest
```

### Compile a single KSL file to intermediate JSON

```bash
ksl -c myfile.ksl
# output: myfile.json
```

### Build a complete SpiceDB schema from multiple files

Pass all your `.ksl` and `.json` files together. KSL resolves imports and applies extensions across all of them in a single pass:

```bash
ksl -o schema.zed rbac.ksl drive.ksl comments.ksl permissions.json
```

If `-o` is omitted, output is written to `schema.zed`.

### Typical project layout

```
schemas/src/
  rbac.ksl
  drive.ksl
  comments.ksl
  ...
  permissions.json       # generated file containing extension calls

schemas/schema.zed       # generated output — do not edit by hand
```

---

## Understanding the output

KSL compiles each type to a SpiceDB `definition`. The namespace and type name are joined with `/`:

| KSL | SpiceDB |
|---|---|
| `namespace rbac` + `type workspace` | `definition rbac/workspace` |
| `namespace drive` + `type document` | `definition drive/document` |

Every KSL relation produces **two** SpiceDB entries:

1. A `relation t_<name>: ...` that stores the raw links (which objects are connected)
2. A `permission <name> = ...` that computes the result from those links

```zed
definition rbac/workspace {
    relation t_parent: rbac/workspace
    permission parent = t_parent

    relation t_user_grant: rbac/role_binding
    permission user_grant = t_user_grant

    permission drive_document_view = t_user_grant->drive_document_view + t_parent->drive_document_view
}
```

**What this means in plain terms:**

- `relation t_parent: rbac/workspace` — stores the raw "this workspace has a parent" link. The `t_` prefix marks it as a raw data relation, not a queryable permission.
- `permission parent = t_parent` — exposes the parent link as a queryable permission (SpiceDB requires this two-step split).
- `permission drive_document_view = t_user_grant->drive_document_view + t_parent->drive_document_view` — a workspace grants `drive_document_view` if **either**: one of its role bindings (`t_user_grant`) grants that permission, **or** its parent workspace grants it (`t_parent->drive_document_view`). This is how permissions propagate up the workspace hierarchy — granting access at a parent workspace automatically covers all workspaces nested inside it.

### The `->` (arrow) operator

The `->` in SpiceDB means: "follow this relation link to the connected object, then evaluate this permission on that object."

So `t_user_grant->drive_document_view` means: "find the role bindings linked via `t_user_grant`, and ask each one: do you grant `drive_document_view`?" If any role binding in the workspace's set grants it, the permission is satisfied.

Similarly, `t_parent->drive_document_view` means: "find the parent workspace linked via `t_parent`, and ask it: do you grant `drive_document_view`?" Because the parent workspace has the same rule, this question cascades up the entire workspace tree until reaching the root. A permission granted at the top-level workspace covers every workspace and every document nested inside it.

This `->` pattern appears everywhere that KSL uses sub-relation syntax — `workspace.drive_document_view` in KSL becomes `t_workspace->drive_document_view` in SpiceDB.

### KSL operator translation

| KSL | SpiceDB |
|---|---|
| `or` | `+` (union) |
| `and` | `&` (intersection) |
| `unless` | `-` (exclusion) |
| `workspace.drive_document_view` | `t_workspace->drive_document_view` (follow link, then check) |
| `[bool]` | `rbac/principal:*` (any principal) |
| `[Any A or B.member]` | `A \| B#member` |

---

## A real-world example

This walkthrough shows how a document-sharing service adds its permissions to the shared RBAC schema, using the same patterns found across all services that integrate with Kessel.

Imagine two services: **Drive**, which manages documents, and **Comments**, which manages comment threads on those documents. Both need permissions that flow through workspaces (think: shared folders), and Comments has a cross-service dependency — you should only be able to read comments on a document you can already read.

> **Background:** In RBAC v1, permissions are expressed as `app:resource:verb` strings (e.g., `drive:document:read`). V2 maps these to named graph permissions using the `app_resource_verb` format with updated verbs (`read`→`view`, `write`→`edit`). For the full picture of why, see [Coming from RBAC v1](https://project-kessel.github.io/docs/building-with-kessel/concepts/coming-from-rbac-v1/).

### Step 1 — The RBAC foundation (rbac.ksl)

The foundation namespace defines the core types and reusable extensions that every service builds on:

```ksl
version 0.1
namespace rbac

public type principal {}

internal type role {
    relation all_all_all: [bool]   // super-admin wildcard
    relation child: [Any role]     // role inheritance links
}

internal type role_binding {
    relation subject: [Any principal or group.member]
    relation role:    [Any role]
}

public type workspace {
    relation parent:  [AtMostOne workspace]
    relation binding: [Any role_binding]
}

// Maps a V1 permission (app:resource:verb) to a V2 permission name.
// Injects relations into role, role_binding, and workspace.
public extension add_v1_based_permission(app, resource, verb, v2_perm) {
    type role {
        allow_duplicates private relation `${app}_all_all`:         [bool]
        allow_duplicates private relation `${app}_${resource}_all`: [bool]
        allow_duplicates private relation `${app}_all_${verb}`:     [bool]
        private relation `${app}_${resource}_${verb}`:              [bool]
        relation `${v2_perm}`: `${app}_${resource}_${verb}` or `${app}_${resource}_all` or `${app}_all_${verb}` or `${app}_all_all` or all_all_all or child.`${v2_perm}`
    }
    type role_binding {
        relation `${v2_perm}`: subject and role.`${v2_perm}`
    }
    type workspace {
        relation `${v2_perm}`: binding.`${v2_perm}` or parent.`${v2_perm}`
    }
}

// Grants a new permission only when the user has two other permissions simultaneously.
public extension add_contingent_permission(first, second, contingent) {
    type workspace {
        relation `${contingent}`: `${first}` and `${second}`
    }
}
```

### Step 2 — The Drive service (drive.ksl)

The Drive service registers its document permissions, mapping V1 names to V2:

```ksl
version 0.1
namespace drive
import rbac

// "drive:document:read" → V2 name "drive_document_view"
@rbac.add_v1_based_permission(app:'drive', resource:'document', verb:'read', v2_perm:'drive_document_view');

// "drive:document:write" → V2 name "drive_document_edit"
@rbac.add_v1_based_permission(app:'drive', resource:'document', verb:'write', v2_perm:'drive_document_edit');
```

### Step 3 — The Comments service (comments.ksl)

The Comments service registers its own permission, then adds a cross-service constraint:

```ksl
version 0.1
namespace comments
import rbac

// "comments:comment:read" → intermediate V2 name "comments_comment_view_assigned".
// The "_assigned" suffix signals this alone is not the final checkable permission.
@rbac.add_v1_based_permission(app:'comments', resource:'comment', verb:'read', v2_perm:'comments_comment_view_assigned');

// The final "comments_comment_view" permission requires BOTH:
//   - drive_document_view (the user can read the underlying document)
//   - comments_comment_view_assigned (the user's role grants comment access)
// This prevents comment threads from being visible on documents
// the user has lost read access to.
@rbac.add_contingent_permission(first:'drive_document_view', second:'comments_comment_view_assigned', contingent:'comments_comment_view');
```

### What gets generated

**In `rbac/role`** (produced by the `drive.ksl` extension calls):

```zed
permission drive_document_view =
    drive_document_read            // role has exact V1 "drive:document:read"
    + drive_document_all           // role has "drive:document:*" (any verb)
    + drive_all_read               // role has "drive:*:read" (any resource)
    + drive_all_all                // role has "drive:*:*" (any resource and verb)
    + all_all_all                  // role is a super-admin (*:*:*)
    + t_child->drive_document_view // inherited from a parent role
```

**What this means in plain terms:** A role grants `drive_document_view` if it was assigned the exact `drive:document:read` V1 permission **or** any of the three wildcard forms that cover it (`drive:document:*`, `drive:*:read`, `drive:*:*`) **or** the global super-admin permission **or** the role inherits from a parent role that grants it.

The `t_child->drive_document_view` entry is role inheritance in action: `t_child` links a role to any roles it is nested under. The `->` follows that link to the parent role and checks if it grants `drive_document_view`. This means a role automatically inherits all permissions of any role it is a child of — exactly like how an "Editor" role inheriting from a "Viewer" role works in Google Drive.

**In `rbac/workspace`** (produced by the `comments.ksl` extension call):

```zed
permission comments_comment_view = (drive_document_view & comments_comment_view_assigned)
```

**What this means in plain terms:** A workspace only grants `comments_comment_view` if the user simultaneously has **both** `drive_document_view` (they can read the underlying document) **and** `comments_comment_view_assigned` (their role explicitly grants comment access). This intersection means a user who has been removed from a document — losing `drive_document_view` — automatically loses the ability to see its comment threads, even if their role still includes `comments_comment_view_assigned`.

**The full permission check in practice:** When a user tries to read comments on a document in a workspace, SpiceDB evaluates:

1. Does the user have `comments_comment_view` on the document's workspace?
2. That requires `drive_document_view` AND `comments_comment_view_assigned` on the same workspace.
3. Each of those propagates up through the workspace tree via `t_parent->`, so a permission granted at a parent workspace covers all nested workspaces and every document inside them.
4. `drive_document_view` is satisfied if the user's role bindings include `drive:document:read` or any of its wildcard forms, or if they inherit the permission from a parent role.

---

## Quick reference

### File structure

```ksl
version 0.1
namespace <name>
import <namespace>

[visibility] type <name> { ... }
[visibility] extension <name>(<params>) { ... }
@<namespace>.<extension>(<params>);
```

### V2 permission naming

| V1 format | V2 format | Example |
|---|---|---|
| `app:resource:read` | `app_resource_view` | `drive:document:read` → `drive_document_view` |
| `app:resource:write` | `app_resource_edit` | `drive:document:write` → `drive_document_edit` |

### Visibility keywords

| Keyword | Within namespace | Across namespaces |
|---|---|---|
| `public` | Yes | Yes |
| `internal` | Yes | No |
| `private` | No | No |

### Cardinality

| Keyword | Meaning |
|---|---|
| `Any` | 0 or more (default) |
| `AtMostOne` | 0 or 1 |
| `ExactlyOne` | Exactly 1 |
| `AtLeastOne` | 1 or more |

### Relation body syntax

| Form | Syntax | Meaning |
|---|---|---|
| Self (direct link) | `[Any TypeA or TypeB.rel]` | Store a direct object link |
| Reference | `someRelation` | Alias for another relation |
| Sub-relation | `someRelation.anotherRelation` | Follow link, evaluate permission on the target |
| Union | `a or b` | Either satisfies |
| Intersection | `a and b` | Both must be satisfied |
| Exclusion | `a unless b` | a, minus anything also in b |

### Extension variables

| Variable | Value |
|---|---|
| `${NAMESPACE}` | Invoking namespace |
| `${TYPE}` | Enclosing type (when applied within a type) |
| `${RELATION}` | Enclosing relation (when applied on a relation) |
| `${MODULE}` | Same as `${NAMESPACE}` |

### CLI commands

```bash
ksl -c file.ksl              # compile to intermediate JSON
ksl -o out.zed a.ksl b.json  # build full SpiceDB schema
```

### KSL → SpiceDB mapping

| KSL | SpiceDB | Plain meaning |
|---|---|---|
| `namespace drive` + `type document` | `definition drive/document` | Resource type scoped to a namespace |
| `relation foo: ...` | `relation t_foo: ...` + `permission foo = t_foo` | Raw link + queryable permission |
| `[bool]` | `rbac/principal:*` | Any principal (wildcard) |
| `or` | `+` | Union: either satisfies |
| `and` | `&` | Intersection: both required |
| `unless` | `-` | Exclusion: subtract a set |
| `a.b` (sub-relation) | `t_a->b` | Follow link `a`, evaluate `b` on the target |

### Escaping reserved keywords

If a relation name conflicts with a KSL keyword (e.g., `version`, `namespace`), prefix it with `#`:

```ksl
relation #version: [ExactlyOne lockversion]   // stored as "version" in the output
```
