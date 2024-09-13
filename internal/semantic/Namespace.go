package semantic

import (
	"fmt"

	core "github.com/authzed/spicedb/pkg/proto/core/v1"
)

type Namespace struct {
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

func (m *Namespace) Name() string {
	return m.name
}

func NewNamespace(name string, imports []string) *Namespace {
	m := &Namespace{name: name, imports: map[string]string{}, types: map[string]*Type{}, extensions: map[string]*Extension{}}

	for _, i := range imports {
		m.imports[i] = i
	}

	return m
}

func (m *Namespace) ApplyExtensions() error {
	//Should handle extensions directly on the namespace
	for _, t := range m.types {
		err := t.ApplyExtensions()
		if err != nil {
			return err
		}
	}

	return nil
}

func (ns *Namespace) AddType(t *Type) error {
	if _, found := ns.types[t.name]; found {
		return fmt.Errorf("namespace %s, type %s: %w", ns.name, t.name, ErrSymbolExists)
	}

	ns.types[t.name] = t
	t.namespace = ns
	return nil
}

func (m *Namespace) AddExtension(e *Extension) error {
	if _, found := m.extensions[e.name]; found {
		return fmt.Errorf("namespace %s, extension %s: %w", m.name, e.name, ErrSymbolExists)
	}

	m.extensions[e.name] = e
	e.namespace = m

	return nil
}

func (m *Namespace) ToZanzibar() ([]*core.NamespaceDefinition, error) {
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
