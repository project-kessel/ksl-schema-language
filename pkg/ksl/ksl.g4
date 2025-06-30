grammar ksl;

// Keywords
VERSIONNUM: [0-9] RESOLVE [0-9]+;
RESOLVE: '.';
ACCESS: PUBLIC | INTERNAL | PRIVATE;
PUBLIC: 'public';
INTERNAL: 'internal';
PRIVATE: 'private';
TYPE: 'type';
RELATION: 'relation';
IMPORT: 'import';
EXTENSION: 'extension';
CARDINALITY: ATMOSTONE | EXACTLYONE | ATLEASTONE | ANY;
ATMOSTONE: 'AtMostOne';
EXACTLYONE: 'ExactlyOne';
ATLEASTONE: 'AtLeastOne';
ANY: 'Any';
AS: 'as';
AND: 'and';
OR: 'or';
UNLESS: 'unless';
ALLOW_DUPLICATES: 'allow_duplicates';
EXPAND: ':';
LBRACE: '{';
RBRACE: '}';
EXTENSION_CALL: '@';
LPAREN: '(';
RPAREN: ')';
LSQUARE: '[';
RSQARE: ']';
VARREF: '$';
TEMPLATE_DELIM: '`';
STRING_DELIM: '\'';
ARG_DELIM: ',';
DECL_END: ';';

// Tokens
NAME: [a-zA-Z_][a-zA-Z_0-9]*;
COMMENT: '//' ~[\r\n]* -> skip;
WS: [ \r\n\t]+ -> skip;

file: version namespace import_stmt* declaration+;

version: NAME VERSIONNUM;
namespace: NAME NAME;
import_stmt: IMPORT NAME;
declaration: typeExpr DECL_END? | extension DECL_END? | extensionReference DECL_END;

typeExpr: extensionReference* ACCESS? TYPE NAME LBRACE relation* RBRACE;
typeReference: NAME (RESOLVE NAME)*;

extensionParam: NAME EXPAND STRING_DELIM value=~STRING_DELIM STRING_DELIM;
extensionParams: extensionParam (ARG_DELIM extensionParam)*;
extensionReference: EXTENSION_CALL typeReference LPAREN extensionParams? RPAREN;
relation: extensionReference* ACCESS? RELATION NAME EXPAND relationBody;
relationBody: LSQUARE CARDINALITY? typeReference (OR typeReference)* RSQARE #Self
    | NAME #Reference
    | relationName=NAME RESOLVE subrelationName=NAME #SubRelation
    | LPAREN relationBody RPAREN #Paren
    | relationBody AND relationBody #And
    | relationBody OR relationBody #OR
    | relationBody UNLESS relationBody #Unless;

paramNames: NAME (ARG_DELIM NAME)*;
extension: ACCESS? EXTENSION NAME LPAREN paramNames? RPAREN LBRACE dynamicType+ RBRACE;

dynamicType: ACCESS? TYPE dynamicName LBRACE dynamicRelation* RBRACE;

dynamicRelation: ALLOW_DUPLICATES? ACCESS? RELATION dynamicName EXPAND dynamicBody;

dynamicName: NAME #Literal
    | VARREF LBRACE NAME RBRACE #Variable
    | TEMPLATE_DELIM dynamicName+ TEMPLATE_DELIM #Template;

dynamicBody: LSQUARE CARDINALITY? typeReference (OR typeReference)* RSQARE #DynamicSelf
    | dynamicName #DynamicReference
    | dynamicName RESOLVE dynamicName #DynamicSubRelation
    | LPAREN dynamicBody RPAREN #DynamicParen
    | dynamicBody AND dynamicBody #DynamicAnd
    | dynamicBody OR dynamicBody #DynamicOR
    | dynamicBody UNLESS dynamicBody #DynamicUnless;

