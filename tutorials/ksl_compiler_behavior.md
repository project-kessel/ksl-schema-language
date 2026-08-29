# KSL Compiler Behavior & Special Features

Important behaviors and magic that happens during KSL compilation.

## 🔮 Special Type Mappings

### `[bool]` → `rbac/principal:*`

**What you write:**
```ksl
relation permission: [bool]
```

**What gets generated:**
```zed
relation t_permission: rbac/principal:*
```

**Why:** The compiler hard-codes `[bool]` to mean "any principal" (universal identity type).

**Requirements:**
- You **must** have a `principal` type in the `rbac` namespace
- This is not configurable (yet)

**Source:** `/pkg/ksl/compiler.go` lines 15-19

## 🏷️ Naming Conventions & Transformations

### Relation Name Prefixes

**What you write:**
```ksl
relation owner: [ExactlyOne user]
```

**What gets generated:**
```zed
permission owner = t_owner
relation t_owner: namespace/user
```

**Rule:** All internal relations get `t_` prefix to avoid naming conflicts.

### SpiceDB Naming Rules

Generated names must match: `^[a-z][a-z0-9_]{1,62}[a-z0-9]$`

**✅ Valid:** `owner`, `can_read`, `t_owner`, `global_all_permissions`
**❌ Invalid:** `_all_all`, `Owner`, `can-read`, `123permission`

## 🔄 Visibility Behavior

### Default Visibility

**What you write:**
```ksl
type document {
    relation owner: [ExactlyOne user]  # No visibility specified
}
```

**What it becomes:** `public` (default visibility)

### Internal Relations Still Generate Permissions

Even `private` and `internal` relations generate SpiceDB permissions - visibility is a KSL concept, not enforced by SpiceDB.

## 🎭 Extension System Magic

### Built-in Template Variables

When extensions are applied, the compiler automatically injects these variables:

- `${NAMESPACE}` - The namespace where the extension is being applied
- `${TYPE}` - The type where the extension is being applied
- `${RELATION}` - The relation where the extension is being applied
- `${permission_name}` - Your custom parameter (example)

**⚠️ Important:** `${MODULE}` was removed in a language refactor and no longer works. Use `${NAMESPACE}` instead.

### Template Expansion Example

**Extension definition:**
```ksl
public extension workspace_permission(permission_name) {
    type role {
        relation ${permission_name}: [bool]
        relation `mod_${NAMESPACE}_type_${TYPE}_all`: [bool]
        relation `mod_${NAMESPACE}_all_rel_${RELATION}`: [bool]
    }
}
```

**When applied to:**
```ksl
namespace inventory
type server {
    @rbac.workspace_permission(permission_name:'servers_read')
    relation can_read: workspace.servers_read
}
```

**Variables get set to:**
- `${NAMESPACE}` = `"inventory"`
- `${TYPE}` = `"server"`
- `${RELATION}` = `"can_read"`
- `${permission_name}` = `"servers_read"`

**Generated relations:**
```ksl
relation servers_read: [bool]                      # ${permission_name}
relation mod_inventory_type_server_all: [bool]     # mod_${NAMESPACE}_type_${TYPE}_all
relation mod_inventory_all_rel_can_read: [bool]    # mod_${NAMESPACE}_all_rel_${RELATION}
```

### `allow_duplicates` Keyword

**What it does:** Prevents errors when the same relation is generated multiple times.

**Example Problem:**
```ksl
type role {
    // Without allow_duplicates, this fails if applied twice to same type
    relation global_all_permissions: [bool]
}
```

**Solution:**
```ksl
type role {
    // With allow_duplicates, second application is silently ignored
    allow_duplicates relation global_all_permissions: [bool]
}
```

**When you need it:**
- Extensions that add the same "global" permissions to multiple resources
- Template relations that might be generated multiple times
- Shared permission patterns across different extension applications

**Code behavior:** If a relation with the same name already exists, the duplicate is ignored instead of causing an error.

### Variable Scoping & Availability

**When variables are set:**
- Extension application triggers variable injection
- Each `@extension()` call gets its own variable context

**Variable availability by context:**

| Context | NAMESPACE | TYPE | RELATION |
|---------|-----------|------|----------|
| **Extension on namespace** | ✅ | ❌ | ❌ |
| **Extension on type** | ✅ | ✅ | ❌ |
| **Extension on relation** | ✅ | ✅ | ✅ |

**Example showing context differences:**
```ksl
namespace inventory

// Namespace-level extension: only NAMESPACE available
@rbac.namespace_permission()
type server {
    // Type-level extension: NAMESPACE/TYPE available
    @rbac.type_permission()

    // Relation-level extension: ALL variables available
    @rbac.workspace_permission(permission_name:'servers_read')
    relation can_read: workspace.servers_read
}
```

## ⚠️ Cardinality Constraints

### Not Enforced by SpiceDB

**Important:** Cardinality constraints are **documentation only**:

```ksl
relation owner: [ExactlyOne user]  # KSL constraint
```

SpiceDB will happily store multiple owners - your application must enforce the `ExactlyOne` rule.

## 🔧 Reserved Keywords

### Escaping with `#`

**What you write:**
```ksl
relation #version: [ExactlyOne user]  # 'version' is reserved
```

**What gets generated:**
```zed
relation version: namespace/user  # # is stripped
```

## 🚨 Common Gotchas

### 1. Missing `principal` Type
```ksl
# ❌ This will fail:
relation perm: [bool]  # No principal type defined

# ✅ This works:
public type principal { }
relation perm: [bool]
```

### 2. Invalid Template Names
```ksl
# ❌ Generates invalid names:
relation `_${NAMESPACE}_all`: [bool]  # Starts with underscore

# ✅ Valid naming:
relation `mod_${NAMESPACE}_all`: [bool]  # Starts with letter
```

### 3. Cross-Namespace Without Import
```ksl
# ❌ This will fail:
relation owner: [ExactlyOne other_namespace.user]  # No import

# ✅ This works:
import other_namespace
relation owner: [ExactlyOne other_namespace.user]
```

### 4. Extension Parameter Quoting
```ksl
# ✅ Both work:
@rbac.permission(permission_name:'read_access')
@rbac.permission(permission_name:"read_access")

# ❌ This fails:
@rbac.permission(permission_name:read_access)  # Missing quotes
```

### 5. Using Outdated `${MODULE}` Variable
```ksl
# ❌ This fails (legacy variable removed):
relation `${MODULE}_all_permissions`: [bool]

# ✅ Use current variable:
relation `${NAMESPACE}_all_permissions`: [bool]
```

**Note:** Many sample files in the repository still use `${MODULE}` but this is outdated and won't work.

## 🎯 Best Practices

1. **Always define `principal`** if using `[bool]`
2. **Use letter prefixes** in templates: `mod_${NAMESPACE}` not `${NAMESPACE}_`
3. **Import before cross-namespace** references
4. **Quote extension parameters** with single or double quotes
5. **Remember cardinality is documentation** - implement validation in your app

This behavior is based on the current KSL compiler implementation and may change in future versions.
