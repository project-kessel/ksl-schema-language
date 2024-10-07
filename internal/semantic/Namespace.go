package semantic

import (
	"fmt"

	core "github.com/authzed/spicedb/pkg/proto/core/v1"
	"github.com/project-kessel/ksl-schema-language/internal/util"
)

type Namespace struct {
	name          string
	schema        *Schema
	imports       *util.OrderedMap[bool]
	types         *util.OrderedMap[*Type]
	extensions    *util.OrderedMap[*Extension]
	extensionRefs []*ExtensionReference
}

type Import struct {
	Name  string
	Alias string
}

func (m *Namespace) Name() string {
	return m.name
}

func NewNamespace(name string, imports []string) *Namespace {
	m := &Namespace{name: name, imports: util.NewOrderedMap[bool](), types: util.NewOrderedMap[*Type](), extensions: util.NewOrderedMap[*Extension](), extensionRefs: []*ExtensionReference{}}

	for _, i := range imports {
		m.imports.Add(i, true)
	}

	return m
}

func (m *Namespace) ApplyExtensions() error {
	for _, e := range m.extensionRefs {
		if err := e.Apply(); err != nil {
			return err
		}
	}

	return m.types.Iterate(func(s string, t *Type) error {
		return t.ApplyExtensions()
	})
}

func (ns *Namespace) AddType(t *Type) error {
	if _, found := ns.types.Get(t.name); found {
		return fmt.Errorf("namespace %s, type %s: %w", ns.name, t.name, ErrSymbolExists)
	}

	ns.types.Add(t.name, t)
	t.namespace = ns
	return nil
}

func (m *Namespace) AddExtension(e *Extension) error {
	if _, found := m.extensions.Get(e.name); found {
		return fmt.Errorf("namespace %s, extension %s: %w", m.name, e.name, ErrSymbolExists)
	}

	m.extensions.Add(e.name, e)
	e.namespace = m

	return nil
}

func (m *Namespace) AddExtensionReference(e *ExtensionReference) {
	e.namespace = m
	m.extensionRefs = append(m.extensionRefs, e)
}

func (m *Namespace) ToZanzibar() ([]*core.NamespaceDefinition, error) {
	namespaces := []*core.NamespaceDefinition{}

	err := m.types.Iterate(func(s string, t *Type) error {
		namespace, err := t.ToZanzibar()
		if err != nil {
			return err
		}

		namespaces = append(namespaces, namespace)
		return nil
	})

	return namespaces, err
}
