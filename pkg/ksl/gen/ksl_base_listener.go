// Code generated from /home/wscalf/Projects/project-kessel/ksl-schema-language/pkg/ksl/ksl.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // ksl

import "github.com/antlr4-go/antlr/v4"

// BasekslListener is a complete listener for a parse tree produced by kslParser.
type BasekslListener struct{}

var _ kslListener = &BasekslListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasekslListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasekslListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasekslListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasekslListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterName is called when production name is entered.
func (s *BasekslListener) EnterName(ctx *NameContext) {}

// ExitName is called when production name is exited.
func (s *BasekslListener) ExitName(ctx *NameContext) {}

// EnterFile is called when production file is entered.
func (s *BasekslListener) EnterFile(ctx *FileContext) {}

// ExitFile is called when production file is exited.
func (s *BasekslListener) ExitFile(ctx *FileContext) {}

// EnterVersion is called when production version is entered.
func (s *BasekslListener) EnterVersion(ctx *VersionContext) {}

// ExitVersion is called when production version is exited.
func (s *BasekslListener) ExitVersion(ctx *VersionContext) {}

// EnterNamespace is called when production namespace is entered.
func (s *BasekslListener) EnterNamespace(ctx *NamespaceContext) {}

// ExitNamespace is called when production namespace is exited.
func (s *BasekslListener) ExitNamespace(ctx *NamespaceContext) {}

// EnterImport_stmt is called when production import_stmt is entered.
func (s *BasekslListener) EnterImport_stmt(ctx *Import_stmtContext) {}

// ExitImport_stmt is called when production import_stmt is exited.
func (s *BasekslListener) ExitImport_stmt(ctx *Import_stmtContext) {}

// EnterDeclaration is called when production declaration is entered.
func (s *BasekslListener) EnterDeclaration(ctx *DeclarationContext) {}

// ExitDeclaration is called when production declaration is exited.
func (s *BasekslListener) ExitDeclaration(ctx *DeclarationContext) {}

// EnterTypeExpr is called when production typeExpr is entered.
func (s *BasekslListener) EnterTypeExpr(ctx *TypeExprContext) {}

// ExitTypeExpr is called when production typeExpr is exited.
func (s *BasekslListener) ExitTypeExpr(ctx *TypeExprContext) {}

// EnterTypeReference is called when production typeReference is entered.
func (s *BasekslListener) EnterTypeReference(ctx *TypeReferenceContext) {}

// ExitTypeReference is called when production typeReference is exited.
func (s *BasekslListener) ExitTypeReference(ctx *TypeReferenceContext) {}

// EnterExtensionParam is called when production extensionParam is entered.
func (s *BasekslListener) EnterExtensionParam(ctx *ExtensionParamContext) {}

// ExitExtensionParam is called when production extensionParam is exited.
func (s *BasekslListener) ExitExtensionParam(ctx *ExtensionParamContext) {}

// EnterExtensionParams is called when production extensionParams is entered.
func (s *BasekslListener) EnterExtensionParams(ctx *ExtensionParamsContext) {}

// ExitExtensionParams is called when production extensionParams is exited.
func (s *BasekslListener) ExitExtensionParams(ctx *ExtensionParamsContext) {}

// EnterExtensionReference is called when production extensionReference is entered.
func (s *BasekslListener) EnterExtensionReference(ctx *ExtensionReferenceContext) {}

// ExitExtensionReference is called when production extensionReference is exited.
func (s *BasekslListener) ExitExtensionReference(ctx *ExtensionReferenceContext) {}

// EnterRelation is called when production relation is entered.
func (s *BasekslListener) EnterRelation(ctx *RelationContext) {}

// ExitRelation is called when production relation is exited.
func (s *BasekslListener) ExitRelation(ctx *RelationContext) {}

// EnterOR is called when production OR is entered.
func (s *BasekslListener) EnterOR(ctx *ORContext) {}

// ExitOR is called when production OR is exited.
func (s *BasekslListener) ExitOR(ctx *ORContext) {}

// EnterReference is called when production Reference is entered.
func (s *BasekslListener) EnterReference(ctx *ReferenceContext) {}

// ExitReference is called when production Reference is exited.
func (s *BasekslListener) ExitReference(ctx *ReferenceContext) {}

// EnterAnd is called when production And is entered.
func (s *BasekslListener) EnterAnd(ctx *AndContext) {}

// ExitAnd is called when production And is exited.
func (s *BasekslListener) ExitAnd(ctx *AndContext) {}

// EnterSelf is called when production Self is entered.
func (s *BasekslListener) EnterSelf(ctx *SelfContext) {}

// ExitSelf is called when production Self is exited.
func (s *BasekslListener) ExitSelf(ctx *SelfContext) {}

// EnterSubRelation is called when production SubRelation is entered.
func (s *BasekslListener) EnterSubRelation(ctx *SubRelationContext) {}

// ExitSubRelation is called when production SubRelation is exited.
func (s *BasekslListener) ExitSubRelation(ctx *SubRelationContext) {}

// EnterUnless is called when production Unless is entered.
func (s *BasekslListener) EnterUnless(ctx *UnlessContext) {}

// ExitUnless is called when production Unless is exited.
func (s *BasekslListener) ExitUnless(ctx *UnlessContext) {}

// EnterParen is called when production Paren is entered.
func (s *BasekslListener) EnterParen(ctx *ParenContext) {}

// ExitParen is called when production Paren is exited.
func (s *BasekslListener) ExitParen(ctx *ParenContext) {}

// EnterParamNames is called when production paramNames is entered.
func (s *BasekslListener) EnterParamNames(ctx *ParamNamesContext) {}

// ExitParamNames is called when production paramNames is exited.
func (s *BasekslListener) ExitParamNames(ctx *ParamNamesContext) {}

// EnterExtension is called when production extension is entered.
func (s *BasekslListener) EnterExtension(ctx *ExtensionContext) {}

// ExitExtension is called when production extension is exited.
func (s *BasekslListener) ExitExtension(ctx *ExtensionContext) {}

// EnterDynamicType is called when production dynamicType is entered.
func (s *BasekslListener) EnterDynamicType(ctx *DynamicTypeContext) {}

// ExitDynamicType is called when production dynamicType is exited.
func (s *BasekslListener) ExitDynamicType(ctx *DynamicTypeContext) {}

// EnterDynamicRelation is called when production dynamicRelation is entered.
func (s *BasekslListener) EnterDynamicRelation(ctx *DynamicRelationContext) {}

// ExitDynamicRelation is called when production dynamicRelation is exited.
func (s *BasekslListener) ExitDynamicRelation(ctx *DynamicRelationContext) {}

// EnterLiteral is called when production Literal is entered.
func (s *BasekslListener) EnterLiteral(ctx *LiteralContext) {}

// ExitLiteral is called when production Literal is exited.
func (s *BasekslListener) ExitLiteral(ctx *LiteralContext) {}

// EnterVariable is called when production Variable is entered.
func (s *BasekslListener) EnterVariable(ctx *VariableContext) {}

// ExitVariable is called when production Variable is exited.
func (s *BasekslListener) ExitVariable(ctx *VariableContext) {}

// EnterTemplate is called when production Template is entered.
func (s *BasekslListener) EnterTemplate(ctx *TemplateContext) {}

// ExitTemplate is called when production Template is exited.
func (s *BasekslListener) ExitTemplate(ctx *TemplateContext) {}

// EnterDynamicUnless is called when production DynamicUnless is entered.
func (s *BasekslListener) EnterDynamicUnless(ctx *DynamicUnlessContext) {}

// ExitDynamicUnless is called when production DynamicUnless is exited.
func (s *BasekslListener) ExitDynamicUnless(ctx *DynamicUnlessContext) {}

// EnterDynamicOR is called when production DynamicOR is entered.
func (s *BasekslListener) EnterDynamicOR(ctx *DynamicORContext) {}

// ExitDynamicOR is called when production DynamicOR is exited.
func (s *BasekslListener) ExitDynamicOR(ctx *DynamicORContext) {}

// EnterDynamicParen is called when production DynamicParen is entered.
func (s *BasekslListener) EnterDynamicParen(ctx *DynamicParenContext) {}

// ExitDynamicParen is called when production DynamicParen is exited.
func (s *BasekslListener) ExitDynamicParen(ctx *DynamicParenContext) {}

// EnterDynamicSelf is called when production DynamicSelf is entered.
func (s *BasekslListener) EnterDynamicSelf(ctx *DynamicSelfContext) {}

// ExitDynamicSelf is called when production DynamicSelf is exited.
func (s *BasekslListener) ExitDynamicSelf(ctx *DynamicSelfContext) {}

// EnterDynamicAnd is called when production DynamicAnd is entered.
func (s *BasekslListener) EnterDynamicAnd(ctx *DynamicAndContext) {}

// ExitDynamicAnd is called when production DynamicAnd is exited.
func (s *BasekslListener) ExitDynamicAnd(ctx *DynamicAndContext) {}

// EnterDynamicReference is called when production DynamicReference is entered.
func (s *BasekslListener) EnterDynamicReference(ctx *DynamicReferenceContext) {}

// ExitDynamicReference is called when production DynamicReference is exited.
func (s *BasekslListener) ExitDynamicReference(ctx *DynamicReferenceContext) {}

// EnterDynamicSubRelation is called when production DynamicSubRelation is entered.
func (s *BasekslListener) EnterDynamicSubRelation(ctx *DynamicSubRelationContext) {}

// ExitDynamicSubRelation is called when production DynamicSubRelation is exited.
func (s *BasekslListener) ExitDynamicSubRelation(ctx *DynamicSubRelationContext) {}
