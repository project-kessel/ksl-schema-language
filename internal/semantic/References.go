package semantic

import "fmt"

type extensionDescriptor struct {
	moduleName    string
	extensionName string //Might be changed to strings to force generated names to be materialized / might just be transparently materialized when evaluating an extension
}

func (d extensionDescriptor) Resolve(inType *Type) (*Extension, error) {
	module := inType.module

	if d.moduleName != "" {
		var found bool
		module, found = inType.module.schema.modules[d.moduleName]
		if !found {
			return nil, fmt.Errorf("Module %s, %w", d.moduleName, ErrSymbolNotFound)
		}
	}

	resolvedExtension, found := module.extensions[d.extensionName]
	if !found {
		return nil, fmt.Errorf("Type %s.%s: %w", d.moduleName, d.extensionName, ErrSymbolNotFound)
	}

	if module == inType.module {
		if !resolvedExtension.visibility.ModuleVisible {
			return nil, fmt.Errorf("Type %s.%s: %w", d.moduleName, d.extensionName, ErrSymbolNotAccessible)
		}
	} else {
		if !resolvedExtension.visibility.SchemaVisible {
			return nil, fmt.Errorf("Type %s.%s: %w", d.moduleName, d.extensionName, ErrSymbolNotAccessible)
		}
	}

	return resolvedExtension, nil
}

type ExtensionReference struct {
	descriptor extensionDescriptor
	parentType *Type
	instance   *Extension
}

func NewExtensionReference(parentType *Type, moduleName string, extensionName string) *ExtensionReference {
	return &ExtensionReference{
		parentType: parentType,
		descriptor: extensionDescriptor{
			moduleName:    moduleName,
			extensionName: extensionName,
		},
	}
}

func (r *ExtensionReference) IsResolved() bool {
	return r.instance != nil
}

func (r *ExtensionReference) Extension() *Extension {
	return r.instance
}

func (r *ExtensionReference) Resolve() error {
	resolvedExtension, err := r.descriptor.Resolve(r.parentType)
	if err != nil {
		return err
	}

	r.instance = resolvedExtension
	return nil
}

type typeDescriptor struct {
	moduleName string
	typeName   string //Might be changed to strings to force generated names to be materialized / might just be transparently materialized when evaluating an extension
}

func (d typeDescriptor) Resolve(inType *Type) (*Type, error) {
	module := inType.module

	if d.moduleName != "" {
		var found bool
		module, found = inType.module.schema.modules[d.moduleName]
		if !found {
			return nil, fmt.Errorf("Module %s, %w", d.moduleName, ErrSymbolNotFound)
		}
	}

	resolvedType, found := module.types[d.typeName]
	if !found {
		return nil, fmt.Errorf("Type %s.%s: %w", d.moduleName, d.typeName, ErrSymbolNotFound)
	}

	if module == inType.module {
		if !resolvedType.visibility.ModuleVisible {
			return nil, fmt.Errorf("Type %s.%s: %w", d.moduleName, d.typeName, ErrSymbolNotAccessible)
		}
	} else {
		if !resolvedType.visibility.SchemaVisible {
			return nil, fmt.Errorf("Type %s.%s: %w", d.moduleName, d.typeName, ErrSymbolNotAccessible)
		}
	}

	return resolvedType, nil
}

type TypeReference struct {
	descriptor  typeDescriptor
	instance    *Type
	subRelation string
	all         bool
}

func NewTypeReference(moduleName string, typeName string, subRelation string, all bool) *TypeReference {
	return &TypeReference{
		descriptor: typeDescriptor{
			moduleName: moduleName,
			typeName:   typeName,
		},
		subRelation: subRelation,
		all:         all,
	}
}

func (r *TypeReference) IsResolved() bool {
	return r.instance != nil
}

func (r *TypeReference) Extension() *Type {
	return r.instance
}

func (r *TypeReference) Resolve(fromType *Type) error {
	resolvedType, err := r.descriptor.Resolve(fromType)
	if err != nil {
		return err
	}

	r.instance = resolvedType
	return nil
}

type relationDescriptor struct {
	moduleName   string
	typeName     Name
	relationName Name
}

func (d relationDescriptor) Resolve(inType *Type) (*Relation, error) {
	resolvedTypeName := d.typeName.String()

	module := inType.module

	if d.moduleName != "" {
		var found bool
		module, found = inType.module.schema.modules[d.moduleName]
		if !found {
			return nil, fmt.Errorf("Module %s, %w", d.moduleName, ErrSymbolNotFound)
		}
	}

	resolvedType, found := module.types[resolvedTypeName]
	if !found {
		return nil, fmt.Errorf("Type %s.%s: %w", d.moduleName, resolvedTypeName, ErrSymbolNotFound)
	}

	if module == inType.module {
		if !resolvedType.visibility.ModuleVisible {
			return nil, fmt.Errorf("Type %s.%s: %w", d.moduleName, resolvedTypeName, ErrSymbolNotAccessible)
		}
	} else {
		if !resolvedType.visibility.SchemaVisible {
			return nil, fmt.Errorf("Type %s.%s: %w", d.moduleName, resolvedTypeName, ErrSymbolNotAccessible)
		}
	}

	resolvedRelationName := d.relationName.String()

	relation, found := resolvedType.relations[resolvedRelationName]
	if !found {
		return nil, fmt.Errorf("Relation %s in type %s.%s: %w", resolvedRelationName, resolvedType.module.name, resolvedType.name, ErrSymbolNotFound)
	}

	if resolvedType == inType {
		//Same type, relations are always accessible
		return relation, nil
	} else if resolvedType.module == inType.module {
		if !relation.visibility.ModuleVisible {
			return nil, fmt.Errorf("Relation %s in type %s.%s: %w", resolvedRelationName, resolvedType.module.name, resolvedType.name, ErrSymbolNotAccessible)
		}
	} else {
		if !relation.visibility.SchemaVisible {
			return nil, fmt.Errorf("Relation %s in type %s.%s: %w", resolvedRelationName, resolvedType.module.name, resolvedType.name, ErrSymbolNotAccessible)
		}
	}

	return relation, nil
}

type RelationReference struct {
	parentType *Type
	descriptor relationDescriptor
	instance   *Relation
}

func NewRelationReference(parentType *Type, moduleName string, typeName Name, relationName Name) *RelationReference {
	return &RelationReference{
		parentType: parentType,
		descriptor: relationDescriptor{
			moduleName:   moduleName,
			typeName:     typeName,
			relationName: relationName,
		},
	}
}

func (r *RelationReference) IsResolved() bool {
	return r.instance != nil
}

func (r *RelationReference) Relation() *Relation {
	return r.instance
}

func (r *RelationReference) Resolve() error {
	relation, err := r.descriptor.Resolve(r.parentType)
	if err != nil {
		return err
	}

	r.instance = relation
	return nil
}
