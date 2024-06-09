package main

import (
	"fmt"

	"github.com/authzed/spicedb/pkg/schemadsl/generator"
	"project-kessel.org/ksl-schema-language/internal/semantic"
	"project-kessel.org/ksl-schema-language/pkg/intermediate"
)

func main() {
	rbac, err := intermediate.LoadFile("samples/rbac.json")
	if err != nil {
		fmt.Printf("Error opening file: %s\n", err)
	}

	inventory, err := intermediate.LoadFile("samples/inventory.json")
	if err != nil {
		fmt.Printf("Error opening file: %s\n", err)
	}

	schema := semantic.NewSchema()
	sm, err := rbac.ToSemantic()
	if err != nil {
		fmt.Printf("Error converting %s to semantic model: %s\n", sm.Name(), err)
		return
	}
	err = schema.AddModule(sm)
	if err != nil {
		fmt.Printf("Error adding module %s: %s", sm.Name(), err)
	}

	sm, err = inventory.ToSemantic()
	if err != nil {
		fmt.Printf("Error converting %s to semantic model: %s\n", sm.Name(), err)
		return
	}
	err = schema.AddModule(sm)
	if err != nil {
		fmt.Printf("Error adding module %s: %s", sm.Name(), err)
	}

	definitions, err := schema.ToZanzibar()
	if err != nil {
		fmt.Printf("Error generating Zanzibar model: %s", err)
	}

	source, _, err := generator.GenerateSchema(definitions)
	if err != nil {
		fmt.Printf("Error generating SpiceDB output: %s", err)
	}

	fmt.Println(source)
}
