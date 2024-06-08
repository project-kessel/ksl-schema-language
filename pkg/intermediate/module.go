package intermediate

import (
	"encoding/json"
	"io"
	"os"

	"project-kessel.org/ksl-schema-language/internal/semantic"
)

type Module struct {
	Name                 string                 `json:"name"`
	Imports              []string               `json:"imports,omitempty"`
	Types                []*Type                `json:"types"`
	ExtensionDefinitions []*ExtensionDefinition `json:"defined_extensions"`
}

type Type struct {
	Name       string      `json:"name"`
	Visibility string      `json:"visibility,omitempty"`
	Relations  []*Relation `json:"relations"`
}

type Relation struct {
	Name       string        `json:"name"`
	Visibility string        `json:"visibility,omitempty"`
	Body       *RelationBody `json:"body"`
}

type RelationBody struct {
	Kind        string        `json:"kind"`
	Module      string        `json:"module,omitempty"`
	Type        string        `json:"type,omitempty"`
	Cardinality string        `json:"cardinality,omitempty"`
	Relation    string        `json:"relation,omitempty"`
	SubRelation string        `json:"sub_relation,omitempty"`
	Left        *RelationBody `json:"left,omitempty"`
	Right       *RelationBody `json:"right,omitempty"`
}

type ExtensionReference struct {
	Module string            `json:"module,omitempty"`
	Name   string            `json:"name"`
	Params map[string]string `json:"params,omitempty"`
}

type ExtensionDefinition struct {
	Name       string         `json:"name"`
	Visibility string         `json:"visibility,omitempty"`
	Params     []string       `json:"params,omitempty"`
	Types      []*DynamicType `json:"types"`
}

type DynamicType struct {
	Name      *DynamicName       `json:"name"`
	Relations []*DynamicRelation `json:"relations"`
}

type DynamicRelation struct {
	Name *DynamicName         `json:"name"`
	Body *DynamicRelationBody `json:"body"`
}

type DynamicRelationBody struct {
	Kind        string               `json:"kind"`
	Module      string               `json:"module,omitempty"`
	Type        string               `json:"type,omitempty"`
	Cardinality string               `json:"cardinality,omitempty"`
	Relation    *DynamicName         `json:"relation,omitempty"`
	SubRelation *DynamicName         `json:"sub_relation,omitempty"`
	Left        *DynamicRelationBody `json:"left,omitempty"`
	Right       *DynamicRelationBody `json:"right,omitempty"`
}

type DynamicName struct {
	Kind     string         `json:"kind"`
	Param    string         `json:"param,omitempty"`
	Value    string         `json:"value,omitempty"`
	Segments []*DynamicName `json:"segments,omitempty"`
}

func LoadFile(filename string) (*Module, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	return Load(file)
}

func Load(reader io.Reader) (*Module, error) {
	module := &Module{}
	data, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, module)

	return module, err
}

func Store(module *Module, writer io.Writer) error {
	encoder := json.NewEncoder(writer)
	return encoder.Encode(module)
}

func (b *RelationBody) ToSemantic() (semantic.RelationExpression, error) {
	switch b.Kind {
	case "self":
		cardinality, err := toCardinality(b.Cardinality)
		if err != nil {
			return nil, err
		}

		return semantic.NewSelfRelationExpression(semantic.NewTypeReference(b.Module, b.Type), cardinality), nil
	case "reference":
		return semantic.NewReferenceRelationExpression(b.Relation, nil), nil
	case "nested_reference":
		return semantic.NewReferenceRelationExpression(b.Relation, &b.SubRelation), nil
	case "union":
	case "intersect":
	case "except":
		left, err := b.Left.ToSemantic()
		if err != nil {
			return nil, err
		}
		right, err := b.Right.ToSemantic()
		if err != nil {
			return nil, err
		}

		return semantic.NewSetRelationExpression(b.Kind, left, right), nil
	}

	return nil, semantic.ErrSymbolNotFound
}

func (r *Relation) ToSemantic() (*semantic.Relation, error) {
	v, err := toVisibility(r.Visibility)
	if err != nil {
		return nil, err
	}

	b, err := r.Body.ToSemantic()
	if err != nil {
		return nil, err
	}

	return semantic.NewRelation(r.Name, v, b)
}

func (t *Type) ToSemantic() (*semantic.Type, error) {
	v, err := toVisibility(t.Visibility)
	if err != nil {
		return nil, err
	}

	st := semantic.NewType(t.Name, v)

	for _, r := range t.Relations {
		sr, err := r.ToSemantic()
		if err != nil {
			return nil, err
		}

		if err := st.AddRelation(sr); err != nil {
			return nil, err
		}
	}

	return st, nil
}

func (m *Module) ToSemantic() (*semantic.Module, error) {
	sm := semantic.NewModule(m.Name, m.Imports)
	for _, t := range m.Types {
		st, err := t.ToSemantic()
		if err != nil {
			return nil, err
		}

		if err = sm.AddType(st); err != nil {
			return nil, err
		}
	}

	return sm, nil
}

func toVisibility(v string) (semantic.Visibility, error) {
	switch v {
	case "private":
		return semantic.VisibilityPrivate, nil
	case "internal":
		return semantic.VisibilityInternal, nil
	case "public", "": //Default to public
		return semantic.VisibilityPublic, nil
	}

	return semantic.VisibilityPrivate, semantic.ErrSymbolNotFound
}

func toCardinality(v string) (semantic.Cardinality, error) {
	switch v {
	case "AtMostOne":
		return semantic.CardinalityAtMostOne, nil
	case "ExactlyOne":
		return semantic.CardinalityExactlyOne, nil
	case "AtLeastOne":
		return semantic.CardinalityAtLeastOne, nil
	case "Any", "":
		return semantic.CardinalityAny, nil
	}

	return semantic.CardinalityAny, semantic.ErrSymbolNotFound
}
