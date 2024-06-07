package intermediate

import (
	"encoding/json"
	"io"
	"os"
)

type Module struct {
	Name                 string                `json:"name"`
	Imports              []string              `json:"imports,omitempty"`
	Types                []Type                `json:"types"`
	ExtensionDefinitions []ExtensionDefinition `json:"defined_extensions"`
}

type Type struct {
	Name       string     `json:"name"`
	Visibility string     `json:"visibility,omitempty"`
	Relations  []Relation `json:"relations"`
}

type Relation struct {
	Name       string       `json:"name"`
	Visibility string       `json:"visibility,omitempty"`
	Body       RelationBody `json:"body"`
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
	Name       string        `json:"name"`
	Visibility string        `json:"visibility,omitempty"`
	Params     []string      `json:"params,omitempty"`
	Types      []DynamicType `json:"types"`
}

type DynamicType struct {
	Name      DynamicName       `json:"name"`
	Relations []DynamicRelation `json:"relations"`
}

type DynamicRelation struct {
	Name DynamicName         `json:"name"`
	Body DynamicRelationBody `json:"body"`
}

type DynamicRelationBody struct {
	Kind        string               `json:"kind"`
	Module      string               `json:"module,omitempty"`
	Type        string               `json:"type,omitempty"`
	Cardinality string               `json:"cardinality,omitempty"`
	Relation    DynamicName          `json:"relation,omitempty"`
	SubRelation DynamicName          `json:"sub_relation,omitempty"`
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
