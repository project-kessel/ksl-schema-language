package semantic

import "fmt"

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
