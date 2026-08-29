# Cardinality Constraints Explained with Real Examples

## The Key Point: These are BUSINESS RULES for your application

Think of cardinality as **how many relationships are allowed**:

## 1. ExactlyOne - Must have exactly 1

**Example: Document Owner**
```
✅ VALID:   document:123 owner user:alice
❌ INVALID: document:123 has no owner (missing required owner)
❌ INVALID: document:123 owner user:alice, user:bob (too many owners)
```

**Real-world meaning**: Every document must have exactly one owner, no more, no less.

## 2. Any - Can have 0, 1, 2, 3, ... (unlimited)

**Example: Document Editor**
```
✅ VALID: document:123 has no editors (document can exist without editors)
✅ VALID: document:123 editor user:alice
✅ VALID: document:123 editor user:alice, user:bob, user:charlie
```

**Real-world meaning**: Documents can have any number of editors, including zero.

## 3. AtMostOne - Can have 0 or 1 (but not more)

**Example: Document Reviewer**
```
✅ VALID: document:123 has no reviewer (review is optional)
✅ VALID: document:123 reviewer user:diana
❌ INVALID: document:123 reviewer user:diana, user:eve (can't have 2 reviewers)
```

**Real-world meaning**: Documents can optionally have one reviewer, but never multiple reviewers.

## 4. AtLeastOne - Must have 1 or more

**Example: Document Approver**
```
❌ INVALID: document:123 has no approvers (approval is required)
✅ VALID: document:123 approver user:frank
✅ VALID: document:123 approver user:frank, user:george, user:helen
```

**Real-world meaning**: Documents must have at least one approver to be valid.

## Summary Table

| Cardinality | Min | Max | Examples |
|-------------|-----|-----|----------|
| `ExactlyOne` | 1 | 1 | Owner, Primary Contact, CEO |
| `Any` | 0 | ∞ | Editors, Viewers, Tags |
| `AtMostOne` | 0 | 1 | Reviewer, Assigned To, Current Status |
| `AtLeastOne` | 1 | ∞ | Approvers, Required Skills, Team Members |

## In Your Application Code

```python
def assign_owner(doc_id, user_id):
    # ExactlyOne: Remove old owner, add new one
    old_owners = get_relationships(doc_id, "owner")
    for owner in old_owners:
        delete_relationship(doc_id, "owner", owner)
    add_relationship(doc_id, "owner", user_id)

def add_editor(doc_id, user_id):
    # Any: Just add, no validation needed
    add_relationship(doc_id, "editor", user_id)

def set_reviewer(doc_id, user_id):
    # AtMostOne: Remove existing reviewer first
    old_reviewers = get_relationships(doc_id, "reviewer")
    for reviewer in old_reviewers:
        delete_relationship(doc_id, "reviewer", reviewer)
    if user_id:  # Can be None (no reviewer)
        add_relationship(doc_id, "reviewer", user_id)

def add_approver(doc_id, user_id):
    # AtLeastOne: Add approver, but validate at least one exists
    add_relationship(doc_id, "approver", user_id)
    # When removing, ensure at least one remains:
    # if len(get_relationships(doc_id, "approver")) < 2:
    #     raise Error("Cannot remove last approver")
```

## Why You Don't See Difference in SpiceDB Schema

SpiceDB doesn't enforce these constraints - it just stores relationships. Your application must implement the business logic to respect these cardinality rules.
