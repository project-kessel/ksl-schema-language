package semantic

import (
	"fmt"

	"github.com/authzed/spicedb/pkg/namespace"
	core "github.com/authzed/spicedb/pkg/proto/core/v1"
)

type Type struct {
	name       string
	visibility Visibility
	module     *Module
	relations  map[string]*Relation
	extensions []*ExtensionReference
}

func NewType(name string, visiblity Visibility) *Type {
	return &Type{name: name, visibility: visiblity, relations: map[string]*Relation{}, extensions: []*ExtensionReference{}}
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
	t.extensions = append(t.extensions, e)
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
	//SpiceDB-specific - in our own Zanzibar model, we'd keep the module and type name separate until we reached the output formatter
	return fmt.Sprintf("%s/%s", t.module.name, t.name)
}
