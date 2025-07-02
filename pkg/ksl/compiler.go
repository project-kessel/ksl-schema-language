package ksl

import (
	"fmt"
	"io"
	"slices"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	"github.com/project-kessel/ksl-schema-language/pkg/intermediate"
	parser "github.com/project-kessel/ksl-schema-language/pkg/ksl/gen"
)

var (
	BooleanType *intermediate.TypeReference = &intermediate.TypeReference{ //Namespace and Name should be customizable, maybe through the language, maybe through a compiler flag
		Namespace: "rbac",
		Name:      "principal",
		All:       true,
	}
)

type CustomErrorListener struct {
	*antlr.DefaultErrorListener
	Errors []string
}

func (c *CustomErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{},
	line, column int, msg string, e antlr.RecognitionException) {
	errorMessage := fmt.Sprintf("line %d:%d %s", line, column, msg)
	c.Errors = append(c.Errors, errorMessage)
}

func Compile(r io.Reader) (*intermediate.Namespace, error) {
	data, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}

	source := antlr.NewInputStream(string(data))
	lexer := parser.NewkslLexer(source)
	tokens := antlr.NewCommonTokenStream(lexer, antlr.LexerDefaultTokenChannel)
	interpreter := parser.NewkslParser(tokens)

	errorListener := &CustomErrorListener{}
	interpreter.RemoveErrorListeners()
	interpreter.AddErrorListener(errorListener)

	file := interpreter.File()

	if len(errorListener.Errors) > 0 {
		errs := strings.Join(errorListener.Errors, "\n")

		return nil, fmt.Errorf("%w:\n%s", ErrSyntaxError, errs)
	}

	converter := &converter{}
	return converter.fileToNamespace(file)
}

type converter struct {
	imports []string
}

func (c *converter) fileToNamespace(f parser.IFileContext) (*intermediate.Namespace, error) {
	name := getStrippedName(f.Namespace().Name())

	imports := []string{}
	for _, i := range f.AllImport_stmt() {
		imports = append(imports, getStrippedName(i.Name()))
	}

	c.imports = imports

	types := []*intermediate.Type{}
	extensions := []*intermediate.ExtensionDefinition{}
	extensionsRefs := []*intermediate.ExtensionReference{}

	for _, d := range f.AllDeclaration() {
		if exp := d.TypeExpr(); exp != nil {
			t, err := c.typeExprToType(exp)
			if err != nil {
				return nil, err
			}

			types = append(types, t)
		}

		if exp := d.Extension(); exp != nil {
			e := c.extensionExprToExtension(exp)

			extensions = append(extensions, e)
		}

		if exp := d.ExtensionReference(); exp != nil {
			e := c.extensionRefExprToExtensionRef(exp)

			extensionsRefs = append(extensionsRefs, e)
		}
	}

	return &intermediate.Namespace{Name: name, Imports: imports, Types: types, ExtensionDefinitions: extensions, ExtensionReferences: extensionsRefs}, nil
}

func (c *converter) typeExprToType(t parser.ITypeExprContext) (*intermediate.Type, error) {
	name := getStrippedName(t.Name())
	access := "public"
	if accessExpr := t.ACCESS(); accessExpr != nil {
		access = accessExpr.GetText()
	}

	extensions := []*intermediate.ExtensionReference{}
	for _, extensionExpr := range t.AllExtensionReference() {
		extensions = append(extensions, c.extensionRefExprToExtensionRef(extensionExpr))
	}

	relations := []*intermediate.Relation{}
	for _, relationExpr := range t.AllRelation() {
		relation, err := c.relationExprToRelation(relationExpr)
		if err != nil {
			return nil, err
		}

		relations = append(relations, relation)
	}

	return &intermediate.Type{Name: name, Visibility: access, Extensions: extensions, Relations: relations}, nil
}

func (c *converter) relationExprToRelation(r parser.IRelationContext) (*intermediate.Relation, error) {
	name := getStrippedName(r.Name())
	access := "public"
	if accessExpr := r.ACCESS(); accessExpr != nil {
		access = accessExpr.GetText()
	}

	body := c.relationBodyExprToRelation(r.RelationBody())

	extensionRefs := []*intermediate.ExtensionReference{}
	for _, extensionRefExpr := range r.AllExtensionReference() {
		extensionRef := c.extensionRefExprToExtensionRef(extensionRefExpr)

		extensionRefs = append(extensionRefs, extensionRef)
	}

	return &intermediate.Relation{Name: name, Visibility: access, Body: body, Extensions: extensionRefs}, nil
}

func (c *converter) relationBodyExprToRelation(r parser.IRelationBodyContext) *intermediate.RelationBody {
	switch body := r.(type) {
	case *parser.SelfContext:
		cardinality := "Any"
		if c := body.CARDINALITY(); c != nil {
			cardinality = c.GetText()
		}

		typeRefs := make([]*intermediate.TypeReference, 0)
		for _, t := range body.AllTypeReference() {
			typeRefs = append(typeRefs, c.typeReferenceExprToTypeReference(t))
		}

		return &intermediate.RelationBody{
			Kind:        "self",
			Cardinality: cardinality,
			Types:       typeRefs,
		}
	case *parser.ReferenceContext:
		relation := getStrippedName(body.Name())
		return &intermediate.RelationBody{
			Kind:     "reference",
			Relation: relation,
		}
	case *parser.SubRelationContext:
		relation := body.GetRelationName().GetText()
		subrelation := body.GetSubrelationName().GetText()
		return &intermediate.RelationBody{
			Kind:        "nested_reference",
			Relation:    relation,
			SubRelation: subrelation,
		}
	case *parser.ParenContext:
		return c.relationBodyExprToRelation(body.RelationBody())
	case *parser.AndContext:
		left := c.relationBodyExprToRelation(body.RelationBody(0))
		right := c.relationBodyExprToRelation(body.RelationBody(1))
		return &intermediate.RelationBody{
			Kind:  "intersect",
			Left:  left,
			Right: right,
		}
	case *parser.ORContext:
		left := c.relationBodyExprToRelation(body.RelationBody(0))
		right := c.relationBodyExprToRelation(body.RelationBody(1))
		return &intermediate.RelationBody{
			Kind:  "union",
			Left:  left,
			Right: right,
		}
	case *parser.UnlessContext:
		left := c.relationBodyExprToRelation(body.RelationBody(0))
		right := c.relationBodyExprToRelation(body.RelationBody(1))
		return &intermediate.RelationBody{
			Kind:  "except",
			Left:  left,
			Right: right,
		}
	}
	return nil
}

func (c *converter) extensionRefExprToExtensionRef(e parser.IExtensionReferenceContext) *intermediate.ExtensionReference {
	typeRef := c.typeReferenceExprToTypeReference(e.TypeReference())

	params := map[string]string{}
	if paramsExpr := e.ExtensionParams(); paramsExpr != nil {
		for _, paramExpr := range paramsExpr.AllExtensionParam() {
			name := getStrippedName(paramExpr.Name())
			value := paramExpr.GetValue().GetText()

			params[name] = value
		}
	}

	return &intermediate.ExtensionReference{Namespace: typeRef.Namespace, Name: typeRef.Name, Params: params}
}

func (c *converter) extensionExprToExtension(t parser.IExtensionContext) *intermediate.ExtensionDefinition {
	name := getStrippedName(t.Name())
	access := "public"
	if accessExpr := t.ACCESS(); accessExpr != nil {
		access = accessExpr.GetText()
	}

	params := []string{}
	if paramsExpr := t.ParamNames(); paramsExpr != nil {
		for _, paramExpr := range paramsExpr.AllName() {
			params = append(params, getStrippedName(paramExpr))
		}
	}

	types := []*intermediate.DynamicType{}
	for _, dynamicTypeExpr := range t.AllDynamicType() {
		types = append(types, c.dynamicTypeExprToDynamicType(dynamicTypeExpr))
	}

	return &intermediate.ExtensionDefinition{Name: name, Visibility: access, Params: params, Types: types}
}

func (c *converter) dynamicTypeExprToDynamicType(t parser.IDynamicTypeContext) *intermediate.DynamicType {
	name := c.dynamicNameExprToDynamicName(t.DynamicName())
	access := "public"
	if accessExpr := t.ACCESS(); accessExpr != nil {
		access = accessExpr.GetText()
	}

	relations := []*intermediate.DynamicRelation{}
	for _, relationExpr := range t.AllDynamicRelation() {
		relations = append(relations, c.dynamicRelationExprToDynamicRelation(relationExpr))
	}

	return &intermediate.DynamicType{Name: name, Visibility: access, Relations: relations}
}

func (c *converter) dynamicRelationExprToDynamicRelation(r parser.IDynamicRelationContext) *intermediate.DynamicRelation {
	name := c.dynamicNameExprToDynamicName(r.DynamicName())
	access := "public"
	if accessExpr := r.ACCESS(); accessExpr != nil {
		access = accessExpr.GetText()
	}

	allowdups := false
	if r.ALLOW_DUPLICATES() != nil {
		allowdups = true
	}

	body := c.dynamicRelationBodyExprToDynamicRelationBody(r.DynamicBody())

	return &intermediate.DynamicRelation{Name: name, Visibility: access, IgnoreDuplicates: allowdups, Body: *body}
}

func (c *converter) dynamicRelationBodyExprToDynamicRelationBody(b parser.IDynamicBodyContext) *intermediate.DynamicRelationBody {
	switch body := b.(type) {
	case *parser.DynamicSelfContext:
		cardinality := "Any"
		if c := body.CARDINALITY(); c != nil {
			cardinality = c.GetText()
		}

		typeRefs := make([]*intermediate.TypeReference, 0)
		for _, t := range body.AllTypeReference() {
			typeRefs = append(typeRefs, c.typeReferenceExprToTypeReference(t))
		}

		return &intermediate.DynamicRelationBody{
			Kind:        "self",
			Cardinality: cardinality,
			Types:       typeRefs,
		}
	case *parser.DynamicReferenceContext:
		relation := c.dynamicNameExprToDynamicName(body.DynamicName())
		return &intermediate.DynamicRelationBody{
			Kind:     "reference",
			Relation: relation,
		}
	case *parser.DynamicSubRelationContext:
		relation := body.DynamicName(0)
		subrelation := body.DynamicName(1)
		return &intermediate.DynamicRelationBody{
			Kind:        "nested_reference",
			Relation:    c.dynamicNameExprToDynamicName(relation),
			SubRelation: c.dynamicNameExprToDynamicName(subrelation),
		}
	case *parser.DynamicParenContext:
		return c.dynamicRelationBodyExprToDynamicRelationBody(body.DynamicBody())
	case *parser.DynamicAndContext:
		left := c.dynamicRelationBodyExprToDynamicRelationBody(body.DynamicBody(0))
		right := c.dynamicRelationBodyExprToDynamicRelationBody(body.DynamicBody(1))
		return &intermediate.DynamicRelationBody{
			Kind:  "intersect",
			Left:  left,
			Right: right,
		}
	case *parser.DynamicORContext:
		left := c.dynamicRelationBodyExprToDynamicRelationBody(body.DynamicBody(0))
		right := c.dynamicRelationBodyExprToDynamicRelationBody(body.DynamicBody(1))
		return &intermediate.DynamicRelationBody{
			Kind:  "union",
			Left:  left,
			Right: right,
		}
	case *parser.DynamicUnlessContext:
		left := c.dynamicRelationBodyExprToDynamicRelationBody(body.DynamicBody(0))
		right := c.dynamicRelationBodyExprToDynamicRelationBody(body.DynamicBody(1))
		return &intermediate.DynamicRelationBody{
			Kind:  "except",
			Left:  left,
			Right: right,
		}
	}
	return nil
}

func (c *converter) dynamicNameExprToDynamicName(n parser.IDynamicNameContext) *intermediate.DynamicName {
	switch body := n.(type) {
	case *parser.LiteralContext:
		return &intermediate.DynamicName{
			Kind:  "literal",
			Value: body.GetText(),
		}
	case *parser.VariableContext:
		return &intermediate.DynamicName{
			Kind:  "param",
			Param: getStrippedName(body.Name()),
		}
	case *parser.TemplateContext:
		segments := []*intermediate.DynamicName{}
		for _, segmentExpr := range body.AllDynamicName() {
			segments = append(segments, c.dynamicNameExprToDynamicName(segmentExpr))
		}
		return &intermediate.DynamicName{
			Kind:     "template",
			Segments: segments,
		}
	}

	return nil
}

func (c *converter) typeReferenceExprToTypeReference(t parser.ITypeReferenceContext) *intermediate.TypeReference {
	var namespaceName, typeName, subRelation string

	segments := t.AllName()
	first := getStrippedName(segments[0])
	if slices.Contains(c.imports, first) {
		namespaceName = first
		typeName = getStrippedName(segments[1])
		if len(segments) > 2 {
			subRelation = getStrippedName(segments[2])
		}
	} else {
		namespaceName = ""
		typeName = getStrippedName(segments[0])
		if len(segments) > 1 {
			subRelation = getStrippedName(segments[1])
		}
	}

	if namespaceName == "" && typeName == "bool" && subRelation == "" {
		return BooleanType
	}

	return &intermediate.TypeReference{Namespace: namespaceName, Name: typeName, SubRelation: subRelation}
}

func getStrippedName(n parser.INameContext) string {
	if n.FORCED_NAME() != nil {
		forcedName := n.FORCED_NAME().GetText()
		return forcedName[1:]
	}

	return n.NAME().GetText()
}
