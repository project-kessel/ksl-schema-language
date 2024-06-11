// Code generated from /home/kooshy/Projects/ksl-schema-language/pkg/ksl/ksl.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // ksl

import "github.com/antlr4-go/antlr/v4"

// kslListener is a complete listener for a parse tree produced by kslParser.
type kslListener interface {
	antlr.ParseTreeListener

	// EnterFile is called when entering the file production.
	EnterFile(c *FileContext)

	// EnterVersion is called when entering the version production.
	EnterVersion(c *VersionContext)

	// EnterModule is called when entering the module production.
	EnterModule(c *ModuleContext)

	// EnterImport_stmt is called when entering the import_stmt production.
	EnterImport_stmt(c *Import_stmtContext)

	// EnterDeclaration is called when entering the declaration production.
	EnterDeclaration(c *DeclarationContext)

	// EnterTypeExpr is called when entering the typeExpr production.
	EnterTypeExpr(c *TypeExprContext)

	// EnterTypeReference is called when entering the typeReference production.
	EnterTypeReference(c *TypeReferenceContext)

	// EnterDynamicTypeReference is called when entering the dynamicTypeReference production.
	EnterDynamicTypeReference(c *DynamicTypeReferenceContext)

	// EnterExtensionParam is called when entering the extensionParam production.
	EnterExtensionParam(c *ExtensionParamContext)

	// EnterExtensionParams is called when entering the extensionParams production.
	EnterExtensionParams(c *ExtensionParamsContext)

	// EnterExtensionReference is called when entering the extensionReference production.
	EnterExtensionReference(c *ExtensionReferenceContext)

	// EnterRelation is called when entering the relation production.
	EnterRelation(c *RelationContext)

	// EnterOR is called when entering the OR production.
	EnterOR(c *ORContext)

	// EnterReference is called when entering the Reference production.
	EnterReference(c *ReferenceContext)

	// EnterAnd is called when entering the And production.
	EnterAnd(c *AndContext)

	// EnterSelf is called when entering the Self production.
	EnterSelf(c *SelfContext)

	// EnterSubRelation is called when entering the SubRelation production.
	EnterSubRelation(c *SubRelationContext)

	// EnterUnless is called when entering the Unless production.
	EnterUnless(c *UnlessContext)

	// EnterParen is called when entering the Paren production.
	EnterParen(c *ParenContext)

	// EnterParamNames is called when entering the paramNames production.
	EnterParamNames(c *ParamNamesContext)

	// EnterExtension is called when entering the extension production.
	EnterExtension(c *ExtensionContext)

	// EnterDynamicType is called when entering the dynamicType production.
	EnterDynamicType(c *DynamicTypeContext)

	// EnterDynamicRelation is called when entering the dynamicRelation production.
	EnterDynamicRelation(c *DynamicRelationContext)

	// EnterLiteral is called when entering the Literal production.
	EnterLiteral(c *LiteralContext)

	// EnterVariable is called when entering the Variable production.
	EnterVariable(c *VariableContext)

	// EnterTemplate is called when entering the Template production.
	EnterTemplate(c *TemplateContext)

	// EnterDynamicUnless is called when entering the DynamicUnless production.
	EnterDynamicUnless(c *DynamicUnlessContext)

	// EnterDynamicOR is called when entering the DynamicOR production.
	EnterDynamicOR(c *DynamicORContext)

	// EnterDynamicParen is called when entering the DynamicParen production.
	EnterDynamicParen(c *DynamicParenContext)

	// EnterDynamicSelf is called when entering the DynamicSelf production.
	EnterDynamicSelf(c *DynamicSelfContext)

	// EnterDynamicAnd is called when entering the DynamicAnd production.
	EnterDynamicAnd(c *DynamicAndContext)

	// EnterDynamicReference is called when entering the DynamicReference production.
	EnterDynamicReference(c *DynamicReferenceContext)

	// EnterDynamicSubRelation is called when entering the DynamicSubRelation production.
	EnterDynamicSubRelation(c *DynamicSubRelationContext)

	// ExitFile is called when exiting the file production.
	ExitFile(c *FileContext)

	// ExitVersion is called when exiting the version production.
	ExitVersion(c *VersionContext)

	// ExitModule is called when exiting the module production.
	ExitModule(c *ModuleContext)

	// ExitImport_stmt is called when exiting the import_stmt production.
	ExitImport_stmt(c *Import_stmtContext)

	// ExitDeclaration is called when exiting the declaration production.
	ExitDeclaration(c *DeclarationContext)

	// ExitTypeExpr is called when exiting the typeExpr production.
	ExitTypeExpr(c *TypeExprContext)

	// ExitTypeReference is called when exiting the typeReference production.
	ExitTypeReference(c *TypeReferenceContext)

	// ExitDynamicTypeReference is called when exiting the dynamicTypeReference production.
	ExitDynamicTypeReference(c *DynamicTypeReferenceContext)

	// ExitExtensionParam is called when exiting the extensionParam production.
	ExitExtensionParam(c *ExtensionParamContext)

	// ExitExtensionParams is called when exiting the extensionParams production.
	ExitExtensionParams(c *ExtensionParamsContext)

	// ExitExtensionReference is called when exiting the extensionReference production.
	ExitExtensionReference(c *ExtensionReferenceContext)

	// ExitRelation is called when exiting the relation production.
	ExitRelation(c *RelationContext)

	// ExitOR is called when exiting the OR production.
	ExitOR(c *ORContext)

	// ExitReference is called when exiting the Reference production.
	ExitReference(c *ReferenceContext)

	// ExitAnd is called when exiting the And production.
	ExitAnd(c *AndContext)

	// ExitSelf is called when exiting the Self production.
	ExitSelf(c *SelfContext)

	// ExitSubRelation is called when exiting the SubRelation production.
	ExitSubRelation(c *SubRelationContext)

	// ExitUnless is called when exiting the Unless production.
	ExitUnless(c *UnlessContext)

	// ExitParen is called when exiting the Paren production.
	ExitParen(c *ParenContext)

	// ExitParamNames is called when exiting the paramNames production.
	ExitParamNames(c *ParamNamesContext)

	// ExitExtension is called when exiting the extension production.
	ExitExtension(c *ExtensionContext)

	// ExitDynamicType is called when exiting the dynamicType production.
	ExitDynamicType(c *DynamicTypeContext)

	// ExitDynamicRelation is called when exiting the dynamicRelation production.
	ExitDynamicRelation(c *DynamicRelationContext)

	// ExitLiteral is called when exiting the Literal production.
	ExitLiteral(c *LiteralContext)

	// ExitVariable is called when exiting the Variable production.
	ExitVariable(c *VariableContext)

	// ExitTemplate is called when exiting the Template production.
	ExitTemplate(c *TemplateContext)

	// ExitDynamicUnless is called when exiting the DynamicUnless production.
	ExitDynamicUnless(c *DynamicUnlessContext)

	// ExitDynamicOR is called when exiting the DynamicOR production.
	ExitDynamicOR(c *DynamicORContext)

	// ExitDynamicParen is called when exiting the DynamicParen production.
	ExitDynamicParen(c *DynamicParenContext)

	// ExitDynamicSelf is called when exiting the DynamicSelf production.
	ExitDynamicSelf(c *DynamicSelfContext)

	// ExitDynamicAnd is called when exiting the DynamicAnd production.
	ExitDynamicAnd(c *DynamicAndContext)

	// ExitDynamicReference is called when exiting the DynamicReference production.
	ExitDynamicReference(c *DynamicReferenceContext)

	// ExitDynamicSubRelation is called when exiting the DynamicSubRelation production.
	ExitDynamicSubRelation(c *DynamicSubRelationContext)
}
