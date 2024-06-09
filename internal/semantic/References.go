package semantic

import (
	"fmt"
)

type extensionDescriptor struct {
	moduleName    string
	extensionName string //Might be changed to strings to force generated names to be materialized / might just be transparently materialized when evaluating an extension
}

func (d extensionDescriptor) Resolve(frommodule *Module) (*Extension, error) {
	inmodule := frommodule
	if d.moduleName != "" {
		var found bool
		inmodule, found = frommodule.schema.modules[d.moduleName]
		if !found {
			return nil, fmt.Errorf("Module %s, %w", d.moduleName, ErrSymbolNotFound)
		}
	}

	resolvedExtension, found := inmodule.extensions[d.extensionName]
	if !found {
		return nil, fmt.Errorf("Type %s.%s: %w", d.moduleName, d.extensionName, ErrSymbolNotFound)
	}

	if inmodule != frommodule {
		if !resolvedExtension.visibility.SchemaVisible {
			return nil, fmt.Errorf("Type %s.%s: %w", d.moduleName, d.extensionName, ErrSymbolNotAccessible)
		}
	}

	return resolvedExtension, nil
}

type ExtensionReference struct {
	descriptor extensionDescriptor
	module     *Module
	onType     *Type
	relation   *Relation
	params     map[string]string
}

func NewExtensionReference(moduleName string, extensionName string, params map[string]string) *ExtensionReference {
	return &ExtensionReference{
		descriptor: extensionDescriptor{
			moduleName:    moduleName,
			extensionName: extensionName,
		},
		params: params,
	}
}

func (r *ExtensionReference) Apply() error {
	e, err := r.descriptor.Resolve(r.module)
	if err != nil {
		return err
	}

	params := make(map[string]string)
	for key, value := range r.params {
		params[key] = value
	}

	return e.Apply(r.module, r.onType, r.relation, params)
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
	typeName     string
	relationName string
}

func (d relationDescriptor) Resolve(inType *Type) (*Relation, error) {
	resolvedTypeName := d.typeName

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

	resolvedRelationName := d.relationName

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

func NewRelationReference(parentType *Type, moduleName string, typeName string, relationName string) *RelationReference {
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
