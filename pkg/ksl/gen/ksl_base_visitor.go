// Code generated from /home/wscalf/Projects/project-kessel/ksl-schema-language/pkg/ksl/ksl.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // ksl

import "github.com/antlr4-go/antlr/v4"

type BasekslVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BasekslVisitor) VisitName(ctx *NameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasekslVisitor) VisitFile(ctx *FileContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasekslVisitor) VisitVersion(ctx *VersionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasekslVisitor) VisitNamespace(ctx *NamespaceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasekslVisitor) VisitImport_stmt(ctx *Import_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasekslVisitor) VisitDeclaration(ctx *DeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasekslVisitor) VisitTypeExpr(ctx *TypeExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasekslVisitor) VisitTypeReference(ctx *TypeReferenceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasekslVisitor) VisitExtensionParam(ctx *ExtensionParamContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasekslVisitor) VisitExtensionParams(ctx *ExtensionParamsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasekslVisitor) VisitExtensionReference(ctx *ExtensionReferenceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasekslVisitor) VisitRelation(ctx *RelationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasekslVisitor) VisitOR(ctx *ORContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasekslVisitor) VisitReference(ctx *ReferenceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasekslVisitor) VisitAnd(ctx *AndContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasekslVisitor) VisitSelf(ctx *SelfContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasekslVisitor) VisitSubRelation(ctx *SubRelationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasekslVisitor) VisitUnless(ctx *UnlessContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasekslVisitor) VisitParen(ctx *ParenContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasekslVisitor) VisitParamNames(ctx *ParamNamesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasekslVisitor) VisitExtension(ctx *ExtensionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasekslVisitor) VisitDynamicType(ctx *DynamicTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasekslVisitor) VisitDynamicRelation(ctx *DynamicRelationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasekslVisitor) VisitLiteral(ctx *LiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasekslVisitor) VisitVariable(ctx *VariableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasekslVisitor) VisitTemplate(ctx *TemplateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasekslVisitor) VisitDynamicUnless(ctx *DynamicUnlessContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasekslVisitor) VisitDynamicOR(ctx *DynamicORContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasekslVisitor) VisitDynamicParen(ctx *DynamicParenContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasekslVisitor) VisitDynamicSelf(ctx *DynamicSelfContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasekslVisitor) VisitDynamicAnd(ctx *DynamicAndContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasekslVisitor) VisitDynamicReference(ctx *DynamicReferenceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasekslVisitor) VisitDynamicSubRelation(ctx *DynamicSubRelationContext) interface{} {
	return v.VisitChildren(ctx)
}
