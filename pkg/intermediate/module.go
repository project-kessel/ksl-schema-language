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
	Name       string                `json:"name"`
	Visibility string                `json:"visibility,omitempty"`
	Body       *RelationBody         `json:"body"`
	Extensions []*ExtensionReference `json:"extensions,omitempty"`
}

type RelationBody struct {
	Kind        string           `json:"kind"`
	Types       []*TypeReference `json:"types"`
	Cardinality string           `json:"cardinality,omitempty"`
	Relation    string           `json:"relation,omitempty"`
	SubRelation string           `json:"sub_relation,omitempty"`
	Left        *RelationBody    `json:"left,omitempty"`
	Right       *RelationBody    `json:"right,omitempty"`
}

type ExtensionReference struct {
	Module string            `json:"module,omitempty"`
	Name   string            `json:"name"`
	Params map[string]string `json:"params,omitempty"`
}

type TypeReference struct {
	Module      string `json:"module,omitempty"`
	Name        string `json:"name"`
	SubRelation string `json:"sub_relation,omitempty"`
	All         bool   `json:"all,omitempty"`
}

type ExtensionDefinition struct {
	Name       string         `json:"name"`
	Visibility string         `json:"visibility,omitempty"`
	Params     []string       `json:"params,omitempty"`
	Types      []*DynamicType `json:"types"`
}

type DynamicType struct {
	Name       *DynamicName       `json:"name"`
	Visibility string             `json:"visibility"`
	Relations  []*DynamicRelation `json:"relations"`
}

type DynamicRelation struct {
	Name             *DynamicName        `json:"name"`
	Visibility       string              `json:"visibility"`
	IgnoreDuplicates bool                `json:"ignoreduplicates,omitempty"`
	Body             DynamicRelationBody `json:"body"`
}

type DynamicRelationBody struct {
	Kind        string               `json:"kind"`
	Types       []*TypeReference     `json:"types"`
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

		types := make([]*semantic.TypeReference, 0)
		for _, t := range b.Types {
			types = append(types, semantic.NewTypeReference(t.Module, t.Name, t.SubRelation, t.All))
		}

		return semantic.NewSelfRelationExpression(types, cardinality), nil
	case "reference":
		return semantic.NewReferenceRelationExpression(b.Relation, nil), nil
	case "nested_reference":
		return semantic.NewReferenceRelationExpression(b.Relation, &b.SubRelation), nil
	case "union", "intersect", "except":
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

func (r *Relation) ToSemantic(t *semantic.Type) (*semantic.Relation, error) {
	v, err := toVisibility(r.Visibility)
	if err != nil {
		return nil, err
	}

	b, err := r.Body.ToSemantic()
	if err != nil {
		return nil, err
	}

	sr, err := semantic.NewRelation(r.Name, t, v, b, nil)
	if err != nil {
		return nil, err
	}

	for _, e := range r.Extensions {
		sr.AddExtension(e.ToSemantic())
	}

	return sr, nil
}

func (t *Type) ToSemantic(m *semantic.Module) (*semantic.Type, error) {
	v, err := toVisibility(t.Visibility)
	if err != nil {
		return nil, err
	}

	st := semantic.NewType(t.Name, m, v)

	for _, r := range t.Relations {
		sr, err := r.ToSemantic(st)
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
		st, err := t.ToSemantic(sm)
		if err != nil {
			return nil, err
		}

		if err = sm.AddType(st); err != nil {
			return nil, err
		}
	}

	for _, e := range m.ExtensionDefinitions {
		converted, err := e.ToSemantic()
		if err != nil {
			return nil, err
		}

		err = sm.AddExtension(converted)
		if err != nil {
			return nil, err
		}
	}

	return sm, nil
}

func (e *ExtensionReference) ToSemantic() *semantic.ExtensionReference {
	return semantic.NewExtensionReference(e.Module, e.Name, e.Params)
}

func (e *ExtensionDefinition) ToSemantic() (*semantic.Extension, error) {
	visibility, err := toVisibility(e.Visibility)
	if err != nil {
		return nil, err
	}

	types := []*semantic.DynamicType{}
	for _, dt := range e.Types {
		t, err := dt.ToSemantic()
		if err != nil {
			return nil, err
		}

		types = append(types, t)
	}
	return semantic.NewExtension(e.Name, visibility, types), nil
}

func (dt *DynamicType) ToSemantic() (*semantic.DynamicType, error) {
	name, err := dt.Name.ToSemantic()
	if err != nil {
		return nil, err
	}

	visibility, err := toVisibility(dt.Visibility)
	if err != nil {
		return nil, err
	}

	relations := []*semantic.DynamicRelation{}
	for _, r := range dt.Relations {
		sr, err := r.ToSemantic()
		if err != nil {
			return nil, err
		}

		relations = append(relations, sr)
	}

	return &semantic.DynamicType{Name: name, Visibility: visibility, Relations: relations}, nil
}

func (dr *DynamicRelation) ToSemantic() (*semantic.DynamicRelation, error) {
	name, err := dr.Name.ToSemantic()
	if err != nil {
		return nil, err
	}

	visibility, err := toVisibility(dr.Visibility)
	if err != nil {
		return nil, err
	}

	body, err := dr.Body.ToSemantic()
	if err != nil {
		return nil, err
	}

	return &semantic.DynamicRelation{Name: name, Visibility: visibility, Body: body, IgnoreDuplicates: dr.IgnoreDuplicates}, nil
}

func (dr *DynamicRelationBody) ToSemantic() (semantic.DynamicRelationBody, error) {
	switch dr.Kind {
	case "self":
		cardinality, err := toCardinality(dr.Cardinality)
		if err != nil {
			return nil, err
		}

		types := make([]*semantic.TypeReference, 0)
		for _, t := range dr.Types {
			types = append(types, semantic.NewTypeReference(t.Module, t.Name, t.SubRelation, t.All))
		}

		return semantic.NewSelfRelationExpression(types, cardinality), nil
	case "reference":
		relation, err := dr.Relation.ToSemantic()
		if err != nil {
			return nil, err
		}
		return &semantic.DynamicReferenceRelationExpression{Relation: relation}, nil
	case "nested_reference":
		relation, err := dr.Relation.ToSemantic()
		if err != nil {
			return nil, err
		}
		var subrelation semantic.Name
		if dr.SubRelation != nil {
			subrelation, err = dr.SubRelation.ToSemantic()
			if err != nil {
				return nil, err
			}
		}
		return &semantic.DynamicReferenceRelationExpression{Relation: relation, SubRelation: subrelation}, nil
	case "union", "intersect", "except":
		left, err := dr.Left.ToSemantic()
		if err != nil {
			return nil, err
		}
		right, err := dr.Right.ToSemantic()
		if err != nil {
			return nil, err
		}

		return semantic.DynamicSetRelationExpression{Kind: dr.Kind, Left: left, Right: right}, nil
	}

	return nil, semantic.ErrSymbolNotFound
}

func (dn *DynamicName) ToSemantic() (semantic.Name, error) {
	switch dn.Kind {
	case "literal":
		return semantic.LiteralName(dn.Value), nil
	case "template":
		segments := []semantic.Name{}
		for _, segment := range dn.Segments {
			converted, err := segment.ToSemantic()
			if err != nil {
				return nil, err
			}

			segments = append(segments, converted)
		}
		return semantic.TemplatedName(segments), nil
	case "param":
		return &semantic.VarName{Name: dn.Param}, nil
	}

	return nil, semantic.ErrSymbolNotFound
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
