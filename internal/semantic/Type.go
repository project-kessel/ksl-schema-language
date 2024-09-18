package semantic

import (
	"fmt"

	"github.com/authzed/spicedb/pkg/namespace"
	core "github.com/authzed/spicedb/pkg/proto/core/v1"
)

type Type struct {
	name       string
	visibility Visibility
	namespace  *Namespace
	relations  map[string]*Relation
	extensions []*ExtensionReference
}

func NewType(name string, m *Namespace, visiblity Visibility) *Type {
	return &Type{name: name, namespace: m, visibility: visiblity, relations: map[string]*Relation{}, extensions: []*ExtensionReference{}}
}

func (t *Type) AddRelation(r *Relation) error {
	if _, found := t.relations[r.name]; found {
		return fmt.Errorf("Type %s, relation %s: %w", t.name, r.name, ErrSymbolExists)
	}

	t.relations[r.name] = r
	r.inType = t

	return nil
}

func (t *Type) AddExtension(e *ExtensionReference) {
	e.namespace = t.namespace
	e.onType = t

	t.extensions = append(t.extensions, e)
}

func (t *Type) ApplyExtensions() error {
	for _, e := range t.extensions {
		if err := e.Apply(); err != nil {
			return err
		}
	}

	for _, r := range t.relations {
		err := r.ApplyExtensions()
		if err != nil {
			return err
		}
	}

	return nil
}

func (t *Type) ToZanzibar() (*core.NamespaceDefinition, error) {
	namespace := namespace.Namespace(t.SpiceDBName()) //SpiceDB-specific

	for _, relation := range t.relations {
		rels, err := relation.ToZanzibar()

		if err != nil {
			return namespace, err
		}

		namespace.Relation = append(namespace.Relation, rels...)
	}

	return namespace, nil
}

func (t *Type) SpiceDBName() string {
	//SpiceDB-specific - in our own Zanzibar model, we'd keep the namespace and type name separate until we reached the output formatter
	return fmt.Sprintf("%s/%s", t.namespace.name, t.name)
}
