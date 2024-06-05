package semantic

import "fmt"

type Module struct {
	name       string
	schema     *Schema
	types      map[string]*Type
	extensions map[string]*Extension
}

func (m *Module) Name() string {
	return m.name
}

func NewModule(name string) *Module {
	return &Module{name: name}
}

func (m *Module) AddType(t *Type) error {
	resolvedName := t.Name().String()
	if _, found := m.types[resolvedName]; found {
		return fmt.Errorf("module %s, type %s: %w", m.name, resolvedName, ErrSymbolExists)
	}

	m.types[resolvedName] = t
	t.module = m
	return nil
}

func (m *Module) AddExtension(e *Extension) error {
	if _, found := m.extensions[e.name]; found {
		return fmt.Errorf("module %s, extension %s: %w", m.name, e.name, ErrSymbolExists)
	}

	m.extensions[e.name] = e
	e.module = m

	return nil
}
