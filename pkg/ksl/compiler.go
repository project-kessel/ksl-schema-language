package ksl

import (
	"io"

	"github.com/antlr4-go/antlr/v4"
	"project-kessel.org/ksl-schema-language/pkg/intermediate"
	parser "project-kessel.org/ksl-schema-language/pkg/ksl/gen"
)

var (
	BooleanType *intermediate.TypeReference = &intermediate.TypeReference{ //Module and Name should be customizable, maybe through the language, maybe through a compiler flag
		Module: "iam",
		Name:   "user",
		All:    true,
	}
)

func Compile(r io.Reader) (*intermediate.Module, error) {
	data, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}

	source := antlr.NewInputStream(string(data))
	lexer := parser.NewkslLexer(source)
	tokens := antlr.NewCommonTokenStream(lexer, antlr.LexerDefaultTokenChannel)
	interpreter := parser.NewkslParser(tokens)

	_ = interpreter.File()

	return &intermediate.Module{}, nil
}
