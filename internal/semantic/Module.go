package semantic

import (
	"fmt"

	core "github.com/authzed/spicedb/pkg/proto/core/v1"
)

type Module struct {
	name       string
	schema     *Schema
	imports    map[string]string
	types      map[string]*Type
	extensions map[string]*Extension
}

type Import struct {
	Name  string
	Alias string
}

func (m *Module) Name() string {
	return m.name
}

func NewModule(name string, imports []string) *Module {
	m := &Module{name: name, imports: map[string]string{}, types: map[string]*Type{}, extensions: map[string]*Extension{}}

	for _, i := range imports {
		m.imports[i] = i
	}

	return m
}

func (m *Module) ApplyExtensions() error {
	//Should handle extensions directly on the module
	for _, t := range m.types {
		err := t.ApplyExtensions()
		if err != nil {
			return err
		}
	}

	return nil
}

func (m *Module) AddType(t *Type) error {
	if _, found := m.types[t.name]; found {
		return fmt.Errorf("module %s, type %s: %w", m.name, t.name, ErrSymbolExists)
	}

	m.types[t.name] = t
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

func (m *Module) ToZanzibar() ([]*core.NamespaceDefinition, error) {
	namespaces := []*core.NamespaceDefinition{}

	for _, t := range m.types {
		namespace, err := t.ToZanzibar()
		if err != nil {
			return namespaces, err
		}

		namespaces = append(namespaces, namespace)
	}

	return namespaces, nil
}
