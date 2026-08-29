# Minimal KSL Learning Guide

Learn **ALL KSL syntax** with just 4 files! Each file teaches unique concepts with zero redundancy.

## 📚 Complete Learning Path (30 minutes total)

### **File 1: All Basic Syntax** (10 min)
- `01_complete_basic_syntax.ksl` + `.zed`
- **Teaches**: Types, all cardinalities, all visibility modifiers, all relation expressions

### **File 2: Advanced Relations & Cross-Namespace** (10 min)
- `02_advanced_relations.ksl` + `02_cross_namespace.ksl` → `02_advanced_and_cross_namespace.zed`
- **Teaches**: Nested relations, imports, cross-namespace references

### **File 3: Extension System** (10 min)
- `rbac_foundation.ksl` + `03_extension_usage.ksl` → `03_complete_extensions.zed`
- **Teaches**: Templates, dynamic generation, `${variables}`, extension usage

### **File 4: Real-World Complete Example** (bonus)
- `rbac_foundation.ksl` + `04_complete_real_world.ksl` → `04_complete_real_world.zed`
- **Teaches**: Everything combined in a practical enterprise scenario

## 🎯 What Each File Covers

| Syntax Feature | File 1 | File 2 | File 3 | File 4 |
|----------------|--------|--------|--------|--------|
| **Basic Types** | ✅ | | | ✅ |
| **All Cardinality** (ExactlyOne, Any, AtMostOne, AtLeastOne) | ✅ | | | ✅ |
| **Visibility** (public, internal, private) | ✅ | | | ✅ |
| **Relations** (direct, union, intersection, exclusion) | ✅ | | | ✅ |
| **Nested Relations** (team.member) | | ✅ | | ✅ |
| **Cross-Namespace** (import, namespace.type) | | ✅ | | ✅ |
| **Extensions** (templates, ${variables}) | | | ✅ | ✅ |
| **Reserved Keywords** (#version) | | | | ✅ |

## 🚀 Quick Compilation

```bash
# Tutorial 1 (single file)
../bin/ksl -o 01_complete_basic_syntax.zed 01_complete_basic_syntax.ksl

# Tutorial 2 (multiple source files)
../bin/ksl -o 02_advanced_and_cross_namespace.zed 02_advanced_relations.ksl 02_cross_namespace.ksl

# Tutorial 3 (extension system)
../bin/ksl -o 03_complete_extensions.zed rbac_foundation.ksl 03_extension_usage.ksl

# Tutorial 4 (real-world example using shared RBAC foundation)
../bin/ksl -o 04_complete_real_world.zed rbac_foundation.ksl 04_complete_real_world.ksl

# All examples already compiled - just read the .zed files!
```

## 📖 Documentation Reference

**Essential Reading:**
- `ksl_compiler_behavior.md` - **CRITICAL**: Special behaviors, magic mappings, gotchas

**Cardinality Deep-Dive:**
- `cardinality_explained.md` - When to use each cardinality type

## ✅ After Reading These 4 Files You'll Know:

- Every KSL syntax feature
- When to use each cardinality constraint
- How to organize multi-namespace schemas
- How to create reusable authorization patterns with extensions
- Real-world enterprise authorization modeling

**Total learning time: ~30 minutes to master the entire language!**
