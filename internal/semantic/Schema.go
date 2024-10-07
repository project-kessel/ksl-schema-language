package semantic

import (
	"fmt"

	"github.com/authzed/spicedb/pkg/schemadsl/compiler"
	"github.com/project-kessel/ksl-schema-language/internal/util"
)

type Schema struct {
	namespaces *util.OrderedMap[*Namespace]
}

func NewSchema() *Schema {
	return &Schema{namespaces: util.NewOrderedMap[*Namespace]()}
}

func (s *Schema) AddNamespace(namespace *Namespace) error {
	resolvedName := namespace.Name()
	if _, found := s.namespaces.Get(resolvedName); found {
		return fmt.Errorf("%s: %w", resolvedName, ErrSymbolExists)
	}

	s.namespaces.Add(resolvedName, namespace)
	namespace.schema = s
	return nil
}

func (s *Schema) ApplyExtensions() error {
	return s.namespaces.Iterate(func(s string, ns *Namespace) error {
		return ns.ApplyExtensions()
	})
}

func (s *Schema) ToZanzibar() ([]compiler.SchemaDefinition, error) {
	namespaceNames := map[string]bool{} //Track names used for Zanzibar namespaces
	elements := []compiler.SchemaDefinition{}

	err := s.namespaces.Iterate(func(s string, ns *Namespace) error {
		namespaces, err := ns.ToZanzibar()
		if err != nil {
			return err
		}

		for _, ns := range namespaces {
			if namespaceNames[ns.Name] {
				return fmt.Errorf("zanzibar namespace %s: %w", ns.Name, ErrSymbolExists)
			} else {
				elements = append(elements, ns)
				namespaceNames[ns.Name] = true
			}
		}

		return nil
	})

	return elements, err
}
