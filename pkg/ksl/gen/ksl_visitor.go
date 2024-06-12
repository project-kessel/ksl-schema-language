// Code generated from /home/kooshy/Projects/ksl-schema-language/pkg/ksl/ksl.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // ksl

import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by kslParser.
type kslVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by kslParser#file.
	VisitFile(ctx *FileContext) interface{}

	// Visit a parse tree produced by kslParser#version.
	VisitVersion(ctx *VersionContext) interface{}

	// Visit a parse tree produced by kslParser#module.
	VisitModule(ctx *ModuleContext) interface{}

	// Visit a parse tree produced by kslParser#import_stmt.
	VisitImport_stmt(ctx *Import_stmtContext) interface{}

	// Visit a parse tree produced by kslParser#declaration.
	VisitDeclaration(ctx *DeclarationContext) interface{}

	// Visit a parse tree produced by kslParser#typeExpr.
	VisitTypeExpr(ctx *TypeExprContext) interface{}

	// Visit a parse tree produced by kslParser#typeReference.
	VisitTypeReference(ctx *TypeReferenceContext) interface{}

	// Visit a parse tree produced by kslParser#dynamicTypeReference.
	VisitDynamicTypeReference(ctx *DynamicTypeReferenceContext) interface{}

	// Visit a parse tree produced by kslParser#extensionParam.
	VisitExtensionParam(ctx *ExtensionParamContext) interface{}

	// Visit a parse tree produced by kslParser#extensionParams.
	VisitExtensionParams(ctx *ExtensionParamsContext) interface{}

	// Visit a parse tree produced by kslParser#extensionReference.
	VisitExtensionReference(ctx *ExtensionReferenceContext) interface{}

	// Visit a parse tree produced by kslParser#relation.
	VisitRelation(ctx *RelationContext) interface{}

	// Visit a parse tree produced by kslParser#OR.
	VisitOR(ctx *ORContext) interface{}

	// Visit a parse tree produced by kslParser#Reference.
	VisitReference(ctx *ReferenceContext) interface{}

	// Visit a parse tree produced by kslParser#And.
	VisitAnd(ctx *AndContext) interface{}

	// Visit a parse tree produced by kslParser#Self.
	VisitSelf(ctx *SelfContext) interface{}

	// Visit a parse tree produced by kslParser#SubRelation.
	VisitSubRelation(ctx *SubRelationContext) interface{}

	// Visit a parse tree produced by kslParser#Unless.
	VisitUnless(ctx *UnlessContext) interface{}

	// Visit a parse tree produced by kslParser#Paren.
	VisitParen(ctx *ParenContext) interface{}

	// Visit a parse tree produced by kslParser#paramNames.
	VisitParamNames(ctx *ParamNamesContext) interface{}

	// Visit a parse tree produced by kslParser#extension.
	VisitExtension(ctx *ExtensionContext) interface{}

	// Visit a parse tree produced by kslParser#dynamicType.
	VisitDynamicType(ctx *DynamicTypeContext) interface{}

	// Visit a parse tree produced by kslParser#dynamicRelation.
	VisitDynamicRelation(ctx *DynamicRelationContext) interface{}

	// Visit a parse tree produced by kslParser#Literal.
	VisitLiteral(ctx *LiteralContext) interface{}

	// Visit a parse tree produced by kslParser#Variable.
	VisitVariable(ctx *VariableContext) interface{}

	// Visit a parse tree produced by kslParser#Template.
	VisitTemplate(ctx *TemplateContext) interface{}

	// Visit a parse tree produced by kslParser#DynamicUnless.
	VisitDynamicUnless(ctx *DynamicUnlessContext) interface{}

	// Visit a parse tree produced by kslParser#DynamicOR.
	VisitDynamicOR(ctx *DynamicORContext) interface{}

	// Visit a parse tree produced by kslParser#DynamicParen.
	VisitDynamicParen(ctx *DynamicParenContext) interface{}

	// Visit a parse tree produced by kslParser#DynamicSelf.
	VisitDynamicSelf(ctx *DynamicSelfContext) interface{}

	// Visit a parse tree produced by kslParser#DynamicAnd.
	VisitDynamicAnd(ctx *DynamicAndContext) interface{}

	// Visit a parse tree produced by kslParser#DynamicReference.
	VisitDynamicReference(ctx *DynamicReferenceContext) interface{}

	// Visit a parse tree produced by kslParser#DynamicSubRelation.
	VisitDynamicSubRelation(ctx *DynamicSubRelationContext) interface{}
}
