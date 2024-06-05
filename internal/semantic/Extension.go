package semantic

import (
	"fmt"
	"strings"
)

type Extension struct {
	name       string
	visibility Visibility
	varNames   []*VarName
	types      map[string]*Type
	module     *Module
}

type TemplatedName []Name

func (e *Extension) NewTemplatedName(segments []Name) TemplatedName {
	name := TemplatedName(segments)
	return name
}

func (n TemplatedName) String() string {
	var sb strings.Builder

	for _, segment := range n {
		sb.WriteString(segment.String())
	}

	return sb.String()
}

type VarName struct {
	name  string
	value string
}

func (e *Extension) NewVarName(variable string) *VarName {
	name := &VarName{name: variable}
	e.varNames = append(e.varNames, name)
	return name
}

func (n *VarName) String() string {
	return string(n.value)
}

func (n *VarName) Update(vars map[string]string) error {
	value, found := vars[n.name]
	if !found {
		return fmt.Errorf("variable %s: %w", n.name, ErrSymbolNotFound)
	}

	n.value = value
	return nil
}
