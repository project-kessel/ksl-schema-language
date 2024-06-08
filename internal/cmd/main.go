package main

import (
	"fmt"

	"project-kessel.org/ksl-schema-language/internal/semantic"
	"project-kessel.org/ksl-schema-language/pkg/intermediate"
)

func main() {
	rbac, err := intermediate.LoadFile("samples/rbac.json")
	if err != nil {
		fmt.Printf("Error opening file: %s\n", err)
	}

	inventory, err := intermediate.LoadFile("samples/inventory.json")

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
}
