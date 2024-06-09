package semantic

import (
	"fmt"

	"github.com/authzed/spicedb/pkg/schemadsl/compiler"
)

type Schema struct {
	modules map[string]*Module
}

func NewSchema() *Schema {
	return &Schema{modules: map[string]*Module{}}
}

func (s *Schema) AddModule(module *Module) error {
	resolvedName := module.Name()
	if _, found := s.modules[resolvedName]; found {
		return fmt.Errorf("%s: %w", resolvedName, ErrSymbolExists)
	}

	s.modules[resolvedName] = module
	module.schema = s
	return nil
}

func (s *Schema) ToZanzibar() ([]compiler.SchemaDefinition, error) {
	namespaceNames := map[string]bool{} //Track names used for Zanzibar namespaces
	elements := []compiler.SchemaDefinition{}

	for _, module := range s.modules {
		namespaces, err := module.ToZanzibar()
		if err != nil {
			return elements, err
		}

		for _, namespace := range namespaces {
			if namespaceNames[namespace.Name] {
				return elements, fmt.Errorf("Zanzibar namespace %s: %w", namespace.Name, ErrSymbolExists)
			} else {
				elements = append(elements, namespace)
				namespaceNames[namespace.Name] = true
			}
		}
	}

	return elements, nil
}
