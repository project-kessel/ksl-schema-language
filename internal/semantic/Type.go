package semantic

import (
	"fmt"

	"github.com/authzed/spicedb/pkg/namespace"
	core "github.com/authzed/spicedb/pkg/proto/core/v1"
	"github.com/project-kessel/ksl-schema-language/internal/util"
)

type Type struct {
	name       string
	visibility Visibility
	namespace  *Namespace
	relations  *util.OrderedMap[*Relation]
	extensions []*ExtensionReference
}

func NewType(name string, m *Namespace, visiblity Visibility) *Type {
	return &Type{name: name, namespace: m, visibility: visiblity, relations: util.NewOrderedMap[*Relation](), extensions: []*ExtensionReference{}}
}

func (t *Type) AddRelation(r *Relation) error {
	if _, found := t.relations.Get(r.name); found {
		return fmt.Errorf("Type %s, relation %s: %w", t.name, r.name, ErrSymbolExists)
	}

	t.relations.Add(r.name, r)
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

	return t.relations.Iterate(func(s string, r *Relation) error {
		return r.ApplyExtensions()
	})
}

func (t *Type) ToZanzibar() (*core.NamespaceDefinition, error) {
	namespace := namespace.Namespace(t.SpiceDBName()) //SpiceDB-specific

	err := t.relations.Iterate(func(s string, r *Relation) error {
		rels, err := r.ToZanzibar()

		if err != nil {
			return err
		}

		namespace.Relation = append(namespace.Relation, rels...)
		return nil
	})

	return namespace, err
}

func (t *Type) SpiceDBName() string {
	//SpiceDB-specific - in our own Zanzibar model, we'd keep the namespace and type name separate until we reached the output formatter
	return fmt.Sprintf("%s/%s", t.namespace.name, t.name)
}
