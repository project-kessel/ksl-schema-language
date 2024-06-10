package main

import (
	"fmt"
	"io"
	"os"

	"github.com/authzed/spicedb/pkg/schemadsl/generator"
	"project-kessel.org/ksl-schema-language/internal/semantic"
	"project-kessel.org/ksl-schema-language/pkg/ksl"
)

func main() {
	rbac, err := ksl.Compile(open("samples/rbac.ksl"))
	if err != nil {
		fmt.Printf("Error opening file: %s\n", err)
		return
	}

	inventory, err := ksl.Compile(open("samples/inventory.ksl"))
	if err != nil {
		fmt.Printf("Error opening file: %s\n", err)
		return
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
		return
	}

	sm, err = inventory.ToSemantic()
	if err != nil {
		fmt.Printf("Error converting %s to semantic model: %s\n", sm.Name(), err)
		return
	}
	err = schema.AddModule(sm)
	if err != nil {
		fmt.Printf("Error adding module %s: %s", sm.Name(), err)
		return
	}

	err = schema.ApplyExtensions()
	if err != nil {
		fmt.Printf("Error applying extensions: %s", err)
		return
	}

	definitions, err := schema.ToZanzibar()
	if err != nil {
		fmt.Printf("Error generating Zanzibar model: %s", err)
		return
	}

	source, _, err := generator.GenerateSchema(definitions)
	if err != nil {
		fmt.Printf("Error generating SpiceDB output: %s", err)
	}

	fmt.Println(source)
}

func open(file string) io.Reader {
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}

	return f
}
