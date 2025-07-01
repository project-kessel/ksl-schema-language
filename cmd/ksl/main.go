package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"flag"

	"github.com/authzed/spicedb/pkg/schemadsl/generator"
	"github.com/project-kessel/ksl-schema-language/internal/semantic"
	"github.com/project-kessel/ksl-schema-language/pkg/intermediate"
	"github.com/project-kessel/ksl-schema-language/pkg/ksl"
)

func main() {
	toCompile := flag.String("c", "", "A KSL source file to be compiled to an intermediate representation")
	output := flag.String("o", "", "The output file to write.")

	flag.Parse()
	if toCompile != nil && *toCompile == "" {
		toCompile = nil
	}
	if output != nil && *output == "" {
		output = nil
	}

	if toCompile != nil {
		err := compileToIL(*toCompile, output)
		if err != nil {
			fmt.Printf("Failed to compile: " + err.Error())
			os.Exit(1)
		}
		return
	}

	err := build(flag.Args(), output)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		os.Exit(1)
	}
}

func build(sources []string, output *string) error {
	if len(sources) == 0 {
		return errors.New("no source files specified")
	}

	schema := semantic.NewSchema()

	for _, source := range sources {
		var ir *intermediate.Namespace
		var err error
		switch filepath.Ext(source) {
		case ".ksl":
			ir, err = ksl.Compile(openRead(source))
			if err != nil {
				if errors.Is(err, ksl.ErrSyntaxError) {
					return fmt.Errorf("%w:\nfile %s", err, source)
				}
				return err
			}
		case ".json":
			ir, err = intermediate.Load(openRead(source))
			if err != nil {
				return err
			}
		}

		namespace, err := ir.ToSemantic()
		if err != nil {
			return err
		}

		err = schema.AddNamespace(namespace)
		if err != nil {
			return err
		}
	}

	err := schema.ApplyExtensions()
	if err != nil {
		return err
	}

	definitions, err := schema.ToZanzibar()
	if err != nil {
		return err
	}

	generated, _, err := generator.GenerateSchema(definitions)
	if err != nil {
		return err
	}

	var file io.Writer
	if output != nil {
		file = openWrite(*output)
	} else {
		file = openWrite("schema.zed")
	}

	_, err = io.WriteString(file, generated)
	return err
}

func compileToIL(source string, output *string) error {
	ir, err := ksl.Compile(openRead(source))
	if err != nil {
		return err
	}

	var file io.Writer
	if output == nil {
		file = openWrite(changeExtension(source, ".json"))
	} else {
		file = openWrite(*output)
	}

	return intermediate.Store(ir, file)
}

func changeExtension(path, newExt string) string {
	ext := filepath.Ext(path)
	path = strings.TrimSuffix(path, ext)
	return path + newExt
}

func openRead(file string) io.Reader {
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}

	return f
}

func openWrite(file string) io.Writer {
	f, err := os.OpenFile(file, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		panic(err)
	}

	return f
}
