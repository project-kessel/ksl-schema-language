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

	file := interpreter.File()
	return fileToModule(file)
}

func fileToModule(f parser.IFileContext) (*intermediate.Module, error) {
	name := f.Module().NAME().GetText()

	imports := []string{}
	for _, i := range f.AllImport_stmt() {
		imports = append(imports, i.NAME().GetText())
	}

	types := []*intermediate.Type{}
	extensions := []*intermediate.ExtensionDefinition{}
	for _, d := range f.AllDeclaration() {
		if exp := d.TypeExpr(); exp != nil {
			t, err := typeExprToType(exp)
			if err != nil {
				return nil, err
			}

			types = append(types, t)
		}

		if exp := d.Extension(); exp != nil {
			e := extensionExprToExtension(exp)

			extensions = append(extensions, e)
		}
	}

	return &intermediate.Module{Name: name, Imports: imports, Types: types, ExtensionDefinitions: extensions}, nil
}

func typeExprToType(t parser.ITypeExprContext) (*intermediate.Type, error) {
	name := t.NAME().GetText()
	access := "public"
	if accessExpr := t.ACCESS(); accessExpr != nil {
		access = accessExpr.GetText()
	}

	relations := []*intermediate.Relation{}
	for _, relationExpr := range t.AllRelation() {
		relation, err := relationExprToRelation(relationExpr)
		if err != nil {
			return nil, err
		}

		relations = append(relations, relation)
	}

	return &intermediate.Type{Name: name, Visibility: access, Relations: relations}, nil
}

func relationExprToRelation(r parser.IRelationContext) (*intermediate.Relation, error) {
	name := r.NAME().GetText()
	access := "public"
	if accessExpr := r.ACCESS(); accessExpr != nil {
		access = accessExpr.GetText()
	}

	body := relationBodyExprToRelation(r.RelationBody())

	extensionRefs := []*intermediate.ExtensionReference{}
	for _, extensionRefExpr := range r.AllExtensionReference() {
		extensionRef := extensionRefExprToExtensionRef(extensionRefExpr)

		extensionRefs = append(extensionRefs, extensionRef)
	}

	return &intermediate.Relation{Name: name, Visibility: access, Body: body, Extensions: extensionRefs}, nil
}

func relationBodyExprToRelation(r parser.IRelationBodyContext) *intermediate.RelationBody {
	switch body := r.(type) {
	case *parser.SelfContext:
		cardinality := "Any"
		if c := body.CARDINALITY(); c != nil {
			cardinality = c.GetText()
		}
		typeRef := typeReferenceExprToTypeReference(body.TypeReference())
		return &intermediate.RelationBody{
			Kind:        "self",
			Cardinality: cardinality,
			Type:        *typeRef,
		}
	case *parser.ReferenceContext:
		relation := body.NAME().GetText()
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
		return relationBodyExprToRelation(body.RelationBody())
	case *parser.AndContext:
		left := relationBodyExprToRelation(body.RelationBody(0))
		right := relationBodyExprToRelation(body.RelationBody(1))
		return &intermediate.RelationBody{
			Kind:  "intersect",
			Left:  left,
			Right: right,
		}
	case *parser.ORContext:
		left := relationBodyExprToRelation(body.RelationBody(0))
		right := relationBodyExprToRelation(body.RelationBody(1))
		return &intermediate.RelationBody{
			Kind:  "union",
			Left:  left,
			Right: right,
		}
	case *parser.UnlessContext:
		left := relationBodyExprToRelation(body.RelationBody(0))
		right := relationBodyExprToRelation(body.RelationBody(1))
		return &intermediate.RelationBody{
			Kind:  "except",
			Left:  left,
			Right: right,
		}
	}
	return nil
}

func extensionRefExprToExtensionRef(e parser.IExtensionReferenceContext) *intermediate.ExtensionReference {
	typeRef := typeReferenceExprToTypeReference(e.TypeReference())

	params := map[string]string{}
	if paramsExpr := e.ExtensionParams(); paramsExpr != nil {
		for _, paramExpr := range paramsExpr.AllExtensionParam() {
			name := paramExpr.NAME().GetText()
			value := paramExpr.GetValue().GetText()

			params[name] = value
		}
	}

	return &intermediate.ExtensionReference{Module: typeRef.Module, Name: typeRef.Name, Params: params}
}

func extensionExprToExtension(t parser.IExtensionContext) *intermediate.ExtensionDefinition {
	name := t.NAME().GetText()
	access := "public"
	if accessExpr := t.ACCESS(); accessExpr != nil {
		access = accessExpr.GetText()
	}

	params := []string{}
	if paramsExpr := t.ParamNames(); paramsExpr != nil {
		for _, paramExpr := range paramsExpr.AllNAME() {
			params = append(params, paramExpr.GetText())
		}
	}

	types := []*intermediate.DynamicType{}
	for _, dynamicTypeExpr := range t.AllDynamicType() {
		types = append(types, dynamicTypeExprToDynamicType(dynamicTypeExpr))
	}

	return &intermediate.ExtensionDefinition{Name: name, Visibility: access, Params: params, Types: types}
}

func dynamicTypeExprToDynamicType(t parser.IDynamicTypeContext) *intermediate.DynamicType {
	name := dynamicNameExprToDynamicName(t.DynamicName())
	access := "public"
	if accessExpr := t.ACCESS(); accessExpr != nil {
		access = accessExpr.GetText()
	}

	relations := []*intermediate.DynamicRelation{}
	for _, relationExpr := range t.AllDynamicRelation() {
		relations = append(relations, dynamicRelationExprToDynamicRelation(relationExpr))
	}

	return &intermediate.DynamicType{Name: name, Visibility: access, Relations: relations}
}

func dynamicRelationExprToDynamicRelation(r parser.IDynamicRelationContext) *intermediate.DynamicRelation {
	name := dynamicNameExprToDynamicName(r.DynamicName())
	access := "public"
	if accessExpr := r.ACCESS(); accessExpr != nil {
		access = accessExpr.GetText()
	}

	allowdups := false
	if r.ALLOW_DUPLICATES() != nil {
		allowdups = true
	}

	body := dynamicRelationBodyExprToDynamicRelationBody(r.DynamicBody())

	return &intermediate.DynamicRelation{Name: name, Visibility: access, IgnoreDuplicates: allowdups, Body: *body}
}

func dynamicRelationBodyExprToDynamicRelationBody(b parser.IDynamicBodyContext) *intermediate.DynamicRelationBody {
	switch body := b.(type) {
	case *parser.DynamicSelfContext:
		cardinality := "Any"
		if c := body.CARDINALITY(); c != nil {
			cardinality = c.GetText()
		}
		typeRef := typeReferenceExprToTypeReference(body.TypeReference())
		return &intermediate.DynamicRelationBody{
			Kind:        "self",
			Cardinality: cardinality,
			Type:        *typeRef,
		}
	case *parser.DynamicReferenceContext:
		relation := dynamicNameExprToDynamicName(body.DynamicName())
		return &intermediate.DynamicRelationBody{
			Kind:     "reference",
			Relation: relation,
		}
	case *parser.DynamicSubRelationContext:
		relation := body.DynamicName(0)
		subrelation := body.DynamicName(1)
		return &intermediate.DynamicRelationBody{
			Kind:        "nested_reference",
			Relation:    dynamicNameExprToDynamicName(relation),
			SubRelation: dynamicNameExprToDynamicName(subrelation),
		}
	case *parser.DynamicParenContext:
		return dynamicRelationBodyExprToDynamicRelationBody(body.DynamicBody())
	case *parser.DynamicAndContext:
		left := dynamicRelationBodyExprToDynamicRelationBody(body.DynamicBody(0))
		right := dynamicRelationBodyExprToDynamicRelationBody(body.DynamicBody(1))
		return &intermediate.DynamicRelationBody{
			Kind:  "intersect",
			Left:  left,
			Right: right,
		}
	case *parser.DynamicORContext:
		left := dynamicRelationBodyExprToDynamicRelationBody(body.DynamicBody(0))
		right := dynamicRelationBodyExprToDynamicRelationBody(body.DynamicBody(1))
		return &intermediate.DynamicRelationBody{
			Kind:  "union",
			Left:  left,
			Right: right,
		}
	case *parser.DynamicUnlessContext:
		left := dynamicRelationBodyExprToDynamicRelationBody(body.DynamicBody(0))
		right := dynamicRelationBodyExprToDynamicRelationBody(body.DynamicBody(1))
		return &intermediate.DynamicRelationBody{
			Kind:  "except",
			Left:  left,
			Right: right,
		}
	}
	return nil
}

func dynamicNameExprToDynamicName(n parser.IDynamicNameContext) *intermediate.DynamicName {
	switch body := n.(type) {
	case *parser.LiteralContext:
		return &intermediate.DynamicName{
			Kind:  "literal",
			Value: body.GetText(),
		}
	case *parser.VariableContext:
		return &intermediate.DynamicName{
			Kind:  "param",
			Param: body.NAME().GetText(),
		}
	case *parser.TemplateContext:
		segments := []*intermediate.DynamicName{}
		for _, segmentExpr := range body.AllDynamicName() {
			segments = append(segments, dynamicNameExprToDynamicName(segmentExpr))
		}
		return &intermediate.DynamicName{
			Kind:     "template",
			Segments: segments,
		}
	}

	return nil
}

func typeReferenceExprToTypeReference(t parser.ITypeReferenceContext) *intermediate.TypeReference {
	moduleName := ""
	if m := t.GetModuleName(); m != nil {
		moduleName = m.GetText()
	}

	typeName := t.GetTypeName().GetText()
	if moduleName == "" && typeName == "bool" {
		return BooleanType
	}

	return &intermediate.TypeReference{Module: moduleName, Name: typeName}
}
