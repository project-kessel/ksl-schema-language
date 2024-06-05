package semantic

type Type struct {
	name       Name
	visibility Visibility
	module     *Module
	relations  map[string]*Relation
	extensions []*ExtensionReference
}

func NewType(name Name, visiblity Visibility) *Type {
	return &Type{name: name, visibility: visiblity, relations: map[string]*Relation{}}
}

func (t *Type) Name() Name {
	return t.name
}
