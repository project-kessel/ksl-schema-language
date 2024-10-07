package semantic

import (
	"fmt"
)

type extensionDescriptor struct {
	namespaceName string
	extensionName string //Might be changed to strings to force generated names to be materialized / might just be transparently materialized when evaluating an extension
}

func (d extensionDescriptor) Resolve(fromNS *Namespace) (*Extension, error) {
	inNS := fromNS
	if d.namespaceName != "" {
		var found bool
		inNS, found = fromNS.schema.namespaces.Get(d.namespaceName)
		if !found {
			return nil, fmt.Errorf("Namespace %s, %w", d.namespaceName, ErrSymbolNotFound)
		}
	}

	resolvedExtension, found := inNS.extensions.Get(d.extensionName)
	if !found {
		return nil, fmt.Errorf("Type %s.%s: %w", d.namespaceName, d.extensionName, ErrSymbolNotFound)
	}

	if inNS != fromNS {
		if !resolvedExtension.visibility.SchemaVisible {
			return nil, fmt.Errorf("Type %s.%s: %w", d.namespaceName, d.extensionName, ErrSymbolNotAccessible)
		}
	}

	return resolvedExtension, nil
}

type ExtensionReference struct {
	descriptor extensionDescriptor
	namespace  *Namespace
	onType     *Type
	relation   *Relation
	params     map[string]string
}

func NewExtensionReference(namespaceName string, extensionName string, params map[string]string) *ExtensionReference {
	return &ExtensionReference{
		descriptor: extensionDescriptor{
			namespaceName: namespaceName,
			extensionName: extensionName,
		},
		params: params,
	}
}

func (r *ExtensionReference) Apply() error {
	e, err := r.descriptor.Resolve(r.namespace)
	if err != nil {
		return err
	}

	params := make(map[string]string)
	for key, value := range r.params {
		params[key] = value
	}

	return e.Apply(r.namespace, r.onType, r.relation, params)
}

type typeDescriptor struct {
	namespaceName string
	typeName      string //Might be changed to strings to force generated names to be materialized / might just be transparently materialized when evaluating an extension
}

func (d typeDescriptor) Resolve(inType *Type) (*Type, error) {
	namespace := inType.namespace

	if d.namespaceName != "" {
		var found bool
		namespace, found = inType.namespace.schema.namespaces.Get(d.namespaceName)
		if !found {
			return nil, fmt.Errorf("Namespace %s, %w", d.namespaceName, ErrSymbolNotFound)
		}
	}

	resolvedType, found := namespace.types.Get(d.typeName)
	if !found {
		return nil, fmt.Errorf("Type %s.%s: %w", namespace.name, d.typeName, ErrSymbolNotFound)
	}

	if namespace == inType.namespace {
		if !resolvedType.visibility.NamespaceVisible {
			return nil, fmt.Errorf("Type %s.%s: %w", namespace.name, d.typeName, ErrSymbolNotAccessible)
		}
	} else {
		if !resolvedType.visibility.SchemaVisible {
			return nil, fmt.Errorf("Type %s.%s: %w", namespace.name, d.typeName, ErrSymbolNotAccessible)
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

func NewTypeReference(namespaceName string, typeName string, subRelation string, all bool) *TypeReference {
	return &TypeReference{
		descriptor: typeDescriptor{
			namespaceName: namespaceName,
			typeName:      typeName,
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
	namespaceName string
	typeName      string
	relationName  string
}

func (d relationDescriptor) Resolve(inType *Type) (*Relation, error) {
	resolvedTypeName := d.typeName

	namespace := inType.namespace

	if d.namespaceName != "" {
		var found bool
		namespace, found = inType.namespace.schema.namespaces.Get(d.namespaceName)
		if !found {
			return nil, fmt.Errorf("Namespace %s, %w", d.namespaceName, ErrSymbolNotFound)
		}
	}

	resolvedType, found := namespace.types.Get(resolvedTypeName)
	if !found {
		return nil, fmt.Errorf("Type %s.%s: %w", d.namespaceName, resolvedTypeName, ErrSymbolNotFound)
	}

	if namespace == inType.namespace {
		if !resolvedType.visibility.NamespaceVisible {
			return nil, fmt.Errorf("Type %s.%s: %w", d.namespaceName, resolvedTypeName, ErrSymbolNotAccessible)
		}
	} else {
		if !resolvedType.visibility.SchemaVisible {
			return nil, fmt.Errorf("Type %s.%s: %w", d.namespaceName, resolvedTypeName, ErrSymbolNotAccessible)
		}
	}

	resolvedRelationName := d.relationName

	relation, found := resolvedType.relations.Get(resolvedRelationName)
	if !found {
		return nil, fmt.Errorf("Relation %s in type %s.%s: %w", resolvedRelationName, resolvedType.namespace.name, resolvedType.name, ErrSymbolNotFound)
	}

	if resolvedType == inType {
		//Same type, relations are always accessible
		return relation, nil
	} else if resolvedType.namespace == inType.namespace {
		if !relation.visibility.NamespaceVisible {
			return nil, fmt.Errorf("Relation %s in type %s.%s: %w", resolvedRelationName, resolvedType.namespace.name, resolvedType.name, ErrSymbolNotAccessible)
		}
	} else {
		if !relation.visibility.SchemaVisible {
			return nil, fmt.Errorf("Relation %s in type %s.%s: %w", resolvedRelationName, resolvedType.namespace.name, resolvedType.name, ErrSymbolNotAccessible)
		}
	}

	return relation, nil
}

type RelationReference struct {
	parentType *Type
	descriptor relationDescriptor
	instance   *Relation
}

func NewRelationReference(parentType *Type, namespaceName string, typeName string, relationName string) *RelationReference {
	return &RelationReference{
		parentType: parentType,
		descriptor: relationDescriptor{
			namespaceName: namespaceName,
			typeName:      typeName,
			relationName:  relationName,
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
