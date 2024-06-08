package semantic

import "fmt"

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
