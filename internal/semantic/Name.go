package semantic

import "strings"

type Name interface {
	String(params map[string]string) string
}

type LiteralName string

func (n LiteralName) String(params map[string]string) string {
	return string(n)
}

type TemplatedName []Name

func (e *Extension) NewTemplatedName(segments []Name) TemplatedName {
	name := TemplatedName(segments)
	return name
}

func (n TemplatedName) String(params map[string]string) string {
	var sb strings.Builder

	for _, segment := range n {
		sb.WriteString(segment.String(params))
	}

	return sb.String()
}

type VarName struct {
	Name string
}

func (e *Extension) NewVarName(varName string) *VarName {
	return &VarName{Name: varName}
}

func (n *VarName) String(params map[string]string) string {
	return params[n.Name]
}
