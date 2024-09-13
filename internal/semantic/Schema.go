package semantic

import (
	"fmt"

	"github.com/authzed/spicedb/pkg/schemadsl/compiler"
)

type Schema struct {
	namespaces map[string]*Namespace
}

func NewSchema() *Schema {
	return &Schema{namespaces: map[string]*Namespace{}}
}

func (s *Schema) AddNamespace(namespace *Namespace) error {
	resolvedName := namespace.Name()
	if _, found := s.namespaces[resolvedName]; found {
		return fmt.Errorf("%s: %w", resolvedName, ErrSymbolExists)
	}

	s.namespaces[resolvedName] = namespace
	namespace.schema = s
	return nil
}

func (s *Schema) ApplyExtensions() error {
	for _, ns := range s.namespaces {
		err := ns.ApplyExtensions()
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *Schema) ToZanzibar() ([]compiler.SchemaDefinition, error) {
	namespaceNames := map[string]bool{} //Track names used for Zanzibar namespaces
	elements := []compiler.SchemaDefinition{}

	for _, ns := range s.namespaces {
		namespaces, err := ns.ToZanzibar()
		if err != nil {
			return elements, err
		}

		for _, ns := range namespaces {
			if namespaceNames[ns.Name] {
				return elements, fmt.Errorf("zanzibar namespace %s: %w", ns.Name, ErrSymbolExists)
			} else {
				elements = append(elements, ns)
				namespaceNames[ns.Name] = true
			}
		}
	}

	return elements, nil
}
