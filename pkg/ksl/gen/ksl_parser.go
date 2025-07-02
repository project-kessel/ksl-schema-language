// Code generated from /home/wscalf/Projects/project-kessel/ksl-schema-language/pkg/ksl/ksl.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // ksl

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type kslParser struct {
	*antlr.BaseParser
}

var KslParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func kslParserInit() {
	staticData := &KslParserStaticData
	staticData.LiteralNames = []string{
		"", "'#'", "", "'version'", "", "'.'", "'namespace'", "", "'public'",
		"'internal'", "'private'", "'type'", "'relation'", "'import'", "'extension'",
		"", "'AtMostOne'", "'ExactlyOne'", "'AtLeastOne'", "'Any'", "'as'",
		"'and'", "'or'", "'unless'", "'allow_duplicates'", "':'", "'{'", "'}'",
		"'@'", "'('", "')'", "'['", "']'", "'$'", "'`'", "'''", "','", "';'",
	}
	staticData.SymbolicNames = []string{
		"", "KW_ESC", "FORCED_NAME", "VERSION", "VERSIONNUM", "RESOLVE", "NAMESPACE",
		"ACCESS", "PUBLIC", "INTERNAL", "PRIVATE", "TYPE", "RELATION", "IMPORT",
		"EXTENSION", "CARDINALITY", "ATMOSTONE", "EXACTLYONE", "ATLEASTONE",
		"ANY", "AS", "AND", "OR", "UNLESS", "ALLOW_DUPLICATES", "EXPAND", "LBRACE",
		"RBRACE", "EXTENSION_CALL", "LPAREN", "RPAREN", "LSQUARE", "RSQARE",
		"VARREF", "TEMPLATE_DELIM", "STRING_DELIM", "ARG_DELIM", "DECL_END",
		"NAME", "COMMENT", "WS",
	}
	staticData.RuleNames = []string{
		"name", "file", "version", "namespace", "import_stmt", "declaration",
		"typeExpr", "typeReference", "extensionParam", "extensionParams", "extensionReference",
		"relation", "relationBody", "paramNames", "extension", "dynamicType",
		"dynamicRelation", "dynamicName", "dynamicBody",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 40, 287, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1,
		5, 1, 44, 8, 1, 10, 1, 12, 1, 47, 9, 1, 1, 1, 4, 1, 50, 8, 1, 11, 1, 12,
		1, 51, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5,
		3, 5, 65, 8, 5, 1, 5, 1, 5, 3, 5, 69, 8, 5, 1, 5, 1, 5, 1, 5, 3, 5, 74,
		8, 5, 1, 6, 5, 6, 77, 8, 6, 10, 6, 12, 6, 80, 9, 6, 1, 6, 3, 6, 83, 8,
		6, 1, 6, 1, 6, 1, 6, 1, 6, 5, 6, 89, 8, 6, 10, 6, 12, 6, 92, 9, 6, 1, 6,
		1, 6, 1, 7, 1, 7, 1, 7, 5, 7, 99, 8, 7, 10, 7, 12, 7, 102, 9, 7, 1, 8,
		1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 5, 9, 113, 8, 9, 10, 9,
		12, 9, 116, 9, 9, 1, 10, 1, 10, 1, 10, 1, 10, 3, 10, 122, 8, 10, 1, 10,
		1, 10, 1, 11, 5, 11, 127, 8, 11, 10, 11, 12, 11, 130, 9, 11, 1, 11, 3,
		11, 133, 8, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12,
		3, 12, 143, 8, 12, 1, 12, 1, 12, 1, 12, 5, 12, 148, 8, 12, 10, 12, 12,
		12, 151, 9, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12,
		1, 12, 1, 12, 1, 12, 3, 12, 164, 8, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1,
		12, 1, 12, 1, 12, 1, 12, 1, 12, 5, 12, 175, 8, 12, 10, 12, 12, 12, 178,
		9, 12, 1, 13, 1, 13, 1, 13, 5, 13, 183, 8, 13, 10, 13, 12, 13, 186, 9,
		13, 1, 14, 3, 14, 189, 8, 14, 1, 14, 1, 14, 1, 14, 1, 14, 3, 14, 195, 8,
		14, 1, 14, 1, 14, 1, 14, 4, 14, 200, 8, 14, 11, 14, 12, 14, 201, 1, 14,
		1, 14, 1, 15, 3, 15, 207, 8, 15, 1, 15, 1, 15, 1, 15, 1, 15, 5, 15, 213,
		8, 15, 10, 15, 12, 15, 216, 9, 15, 1, 15, 1, 15, 1, 16, 3, 16, 221, 8,
		16, 1, 16, 3, 16, 224, 8, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 17,
		1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 4, 17, 239, 8, 17, 11,
		17, 12, 17, 240, 1, 17, 1, 17, 3, 17, 245, 8, 17, 1, 18, 1, 18, 1, 18,
		3, 18, 250, 8, 18, 1, 18, 1, 18, 1, 18, 5, 18, 255, 8, 18, 10, 18, 12,
		18, 258, 9, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18,
		1, 18, 1, 18, 1, 18, 3, 18, 271, 8, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1,
		18, 1, 18, 1, 18, 1, 18, 1, 18, 5, 18, 282, 8, 18, 10, 18, 12, 18, 285,
		9, 18, 1, 18, 0, 2, 24, 36, 19, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20,
		22, 24, 26, 28, 30, 32, 34, 36, 0, 2, 2, 0, 2, 2, 38, 38, 1, 0, 35, 35,
		308, 0, 38, 1, 0, 0, 0, 2, 40, 1, 0, 0, 0, 4, 53, 1, 0, 0, 0, 6, 56, 1,
		0, 0, 0, 8, 59, 1, 0, 0, 0, 10, 73, 1, 0, 0, 0, 12, 78, 1, 0, 0, 0, 14,
		95, 1, 0, 0, 0, 16, 103, 1, 0, 0, 0, 18, 109, 1, 0, 0, 0, 20, 117, 1, 0,
		0, 0, 22, 128, 1, 0, 0, 0, 24, 163, 1, 0, 0, 0, 26, 179, 1, 0, 0, 0, 28,
		188, 1, 0, 0, 0, 30, 206, 1, 0, 0, 0, 32, 220, 1, 0, 0, 0, 34, 244, 1,
		0, 0, 0, 36, 270, 1, 0, 0, 0, 38, 39, 7, 0, 0, 0, 39, 1, 1, 0, 0, 0, 40,
		41, 3, 4, 2, 0, 41, 45, 3, 6, 3, 0, 42, 44, 3, 8, 4, 0, 43, 42, 1, 0, 0,
		0, 44, 47, 1, 0, 0, 0, 45, 43, 1, 0, 0, 0, 45, 46, 1, 0, 0, 0, 46, 49,
		1, 0, 0, 0, 47, 45, 1, 0, 0, 0, 48, 50, 3, 10, 5, 0, 49, 48, 1, 0, 0, 0,
		50, 51, 1, 0, 0, 0, 51, 49, 1, 0, 0, 0, 51, 52, 1, 0, 0, 0, 52, 3, 1, 0,
		0, 0, 53, 54, 5, 3, 0, 0, 54, 55, 5, 4, 0, 0, 55, 5, 1, 0, 0, 0, 56, 57,
		5, 6, 0, 0, 57, 58, 3, 0, 0, 0, 58, 7, 1, 0, 0, 0, 59, 60, 5, 13, 0, 0,
		60, 61, 3, 0, 0, 0, 61, 9, 1, 0, 0, 0, 62, 64, 3, 12, 6, 0, 63, 65, 5,
		37, 0, 0, 64, 63, 1, 0, 0, 0, 64, 65, 1, 0, 0, 0, 65, 74, 1, 0, 0, 0, 66,
		68, 3, 28, 14, 0, 67, 69, 5, 37, 0, 0, 68, 67, 1, 0, 0, 0, 68, 69, 1, 0,
		0, 0, 69, 74, 1, 0, 0, 0, 70, 71, 3, 20, 10, 0, 71, 72, 5, 37, 0, 0, 72,
		74, 1, 0, 0, 0, 73, 62, 1, 0, 0, 0, 73, 66, 1, 0, 0, 0, 73, 70, 1, 0, 0,
		0, 74, 11, 1, 0, 0, 0, 75, 77, 3, 20, 10, 0, 76, 75, 1, 0, 0, 0, 77, 80,
		1, 0, 0, 0, 78, 76, 1, 0, 0, 0, 78, 79, 1, 0, 0, 0, 79, 82, 1, 0, 0, 0,
		80, 78, 1, 0, 0, 0, 81, 83, 5, 7, 0, 0, 82, 81, 1, 0, 0, 0, 82, 83, 1,
		0, 0, 0, 83, 84, 1, 0, 0, 0, 84, 85, 5, 11, 0, 0, 85, 86, 3, 0, 0, 0, 86,
		90, 5, 26, 0, 0, 87, 89, 3, 22, 11, 0, 88, 87, 1, 0, 0, 0, 89, 92, 1, 0,
		0, 0, 90, 88, 1, 0, 0, 0, 90, 91, 1, 0, 0, 0, 91, 93, 1, 0, 0, 0, 92, 90,
		1, 0, 0, 0, 93, 94, 5, 27, 0, 0, 94, 13, 1, 0, 0, 0, 95, 100, 3, 0, 0,
		0, 96, 97, 5, 5, 0, 0, 97, 99, 3, 0, 0, 0, 98, 96, 1, 0, 0, 0, 99, 102,
		1, 0, 0, 0, 100, 98, 1, 0, 0, 0, 100, 101, 1, 0, 0, 0, 101, 15, 1, 0, 0,
		0, 102, 100, 1, 0, 0, 0, 103, 104, 3, 0, 0, 0, 104, 105, 5, 25, 0, 0, 105,
		106, 5, 35, 0, 0, 106, 107, 8, 1, 0, 0, 107, 108, 5, 35, 0, 0, 108, 17,
		1, 0, 0, 0, 109, 114, 3, 16, 8, 0, 110, 111, 5, 36, 0, 0, 111, 113, 3,
		16, 8, 0, 112, 110, 1, 0, 0, 0, 113, 116, 1, 0, 0, 0, 114, 112, 1, 0, 0,
		0, 114, 115, 1, 0, 0, 0, 115, 19, 1, 0, 0, 0, 116, 114, 1, 0, 0, 0, 117,
		118, 5, 28, 0, 0, 118, 119, 3, 14, 7, 0, 119, 121, 5, 29, 0, 0, 120, 122,
		3, 18, 9, 0, 121, 120, 1, 0, 0, 0, 121, 122, 1, 0, 0, 0, 122, 123, 1, 0,
		0, 0, 123, 124, 5, 30, 0, 0, 124, 21, 1, 0, 0, 0, 125, 127, 3, 20, 10,
		0, 126, 125, 1, 0, 0, 0, 127, 130, 1, 0, 0, 0, 128, 126, 1, 0, 0, 0, 128,
		129, 1, 0, 0, 0, 129, 132, 1, 0, 0, 0, 130, 128, 1, 0, 0, 0, 131, 133,
		5, 7, 0, 0, 132, 131, 1, 0, 0, 0, 132, 133, 1, 0, 0, 0, 133, 134, 1, 0,
		0, 0, 134, 135, 5, 12, 0, 0, 135, 136, 3, 0, 0, 0, 136, 137, 5, 25, 0,
		0, 137, 138, 3, 24, 12, 0, 138, 23, 1, 0, 0, 0, 139, 140, 6, 12, -1, 0,
		140, 142, 5, 31, 0, 0, 141, 143, 5, 15, 0, 0, 142, 141, 1, 0, 0, 0, 142,
		143, 1, 0, 0, 0, 143, 144, 1, 0, 0, 0, 144, 149, 3, 14, 7, 0, 145, 146,
		5, 22, 0, 0, 146, 148, 3, 14, 7, 0, 147, 145, 1, 0, 0, 0, 148, 151, 1,
		0, 0, 0, 149, 147, 1, 0, 0, 0, 149, 150, 1, 0, 0, 0, 150, 152, 1, 0, 0,
		0, 151, 149, 1, 0, 0, 0, 152, 153, 5, 32, 0, 0, 153, 164, 1, 0, 0, 0, 154,
		164, 3, 0, 0, 0, 155, 156, 3, 0, 0, 0, 156, 157, 5, 5, 0, 0, 157, 158,
		3, 0, 0, 0, 158, 164, 1, 0, 0, 0, 159, 160, 5, 29, 0, 0, 160, 161, 3, 24,
		12, 0, 161, 162, 5, 30, 0, 0, 162, 164, 1, 0, 0, 0, 163, 139, 1, 0, 0,
		0, 163, 154, 1, 0, 0, 0, 163, 155, 1, 0, 0, 0, 163, 159, 1, 0, 0, 0, 164,
		176, 1, 0, 0, 0, 165, 166, 10, 3, 0, 0, 166, 167, 5, 21, 0, 0, 167, 175,
		3, 24, 12, 4, 168, 169, 10, 2, 0, 0, 169, 170, 5, 22, 0, 0, 170, 175, 3,
		24, 12, 3, 171, 172, 10, 1, 0, 0, 172, 173, 5, 23, 0, 0, 173, 175, 3, 24,
		12, 2, 174, 165, 1, 0, 0, 0, 174, 168, 1, 0, 0, 0, 174, 171, 1, 0, 0, 0,
		175, 178, 1, 0, 0, 0, 176, 174, 1, 0, 0, 0, 176, 177, 1, 0, 0, 0, 177,
		25, 1, 0, 0, 0, 178, 176, 1, 0, 0, 0, 179, 184, 3, 0, 0, 0, 180, 181, 5,
		36, 0, 0, 181, 183, 3, 0, 0, 0, 182, 180, 1, 0, 0, 0, 183, 186, 1, 0, 0,
		0, 184, 182, 1, 0, 0, 0, 184, 185, 1, 0, 0, 0, 185, 27, 1, 0, 0, 0, 186,
		184, 1, 0, 0, 0, 187, 189, 5, 7, 0, 0, 188, 187, 1, 0, 0, 0, 188, 189,
		1, 0, 0, 0, 189, 190, 1, 0, 0, 0, 190, 191, 5, 14, 0, 0, 191, 192, 3, 0,
		0, 0, 192, 194, 5, 29, 0, 0, 193, 195, 3, 26, 13, 0, 194, 193, 1, 0, 0,
		0, 194, 195, 1, 0, 0, 0, 195, 196, 1, 0, 0, 0, 196, 197, 5, 30, 0, 0, 197,
		199, 5, 26, 0, 0, 198, 200, 3, 30, 15, 0, 199, 198, 1, 0, 0, 0, 200, 201,
		1, 0, 0, 0, 201, 199, 1, 0, 0, 0, 201, 202, 1, 0, 0, 0, 202, 203, 1, 0,
		0, 0, 203, 204, 5, 27, 0, 0, 204, 29, 1, 0, 0, 0, 205, 207, 5, 7, 0, 0,
		206, 205, 1, 0, 0, 0, 206, 207, 1, 0, 0, 0, 207, 208, 1, 0, 0, 0, 208,
		209, 5, 11, 0, 0, 209, 210, 3, 34, 17, 0, 210, 214, 5, 26, 0, 0, 211, 213,
		3, 32, 16, 0, 212, 211, 1, 0, 0, 0, 213, 216, 1, 0, 0, 0, 214, 212, 1,
		0, 0, 0, 214, 215, 1, 0, 0, 0, 215, 217, 1, 0, 0, 0, 216, 214, 1, 0, 0,
		0, 217, 218, 5, 27, 0, 0, 218, 31, 1, 0, 0, 0, 219, 221, 5, 24, 0, 0, 220,
		219, 1, 0, 0, 0, 220, 221, 1, 0, 0, 0, 221, 223, 1, 0, 0, 0, 222, 224,
		5, 7, 0, 0, 223, 222, 1, 0, 0, 0, 223, 224, 1, 0, 0, 0, 224, 225, 1, 0,
		0, 0, 225, 226, 5, 12, 0, 0, 226, 227, 3, 34, 17, 0, 227, 228, 5, 25, 0,
		0, 228, 229, 3, 36, 18, 0, 229, 33, 1, 0, 0, 0, 230, 245, 3, 0, 0, 0, 231,
		232, 5, 33, 0, 0, 232, 233, 5, 26, 0, 0, 233, 234, 3, 0, 0, 0, 234, 235,
		5, 27, 0, 0, 235, 245, 1, 0, 0, 0, 236, 238, 5, 34, 0, 0, 237, 239, 3,
		34, 17, 0, 238, 237, 1, 0, 0, 0, 239, 240, 1, 0, 0, 0, 240, 238, 1, 0,
		0, 0, 240, 241, 1, 0, 0, 0, 241, 242, 1, 0, 0, 0, 242, 243, 5, 34, 0, 0,
		243, 245, 1, 0, 0, 0, 244, 230, 1, 0, 0, 0, 244, 231, 1, 0, 0, 0, 244,
		236, 1, 0, 0, 0, 245, 35, 1, 0, 0, 0, 246, 247, 6, 18, -1, 0, 247, 249,
		5, 31, 0, 0, 248, 250, 5, 15, 0, 0, 249, 248, 1, 0, 0, 0, 249, 250, 1,
		0, 0, 0, 250, 251, 1, 0, 0, 0, 251, 256, 3, 14, 7, 0, 252, 253, 5, 22,
		0, 0, 253, 255, 3, 14, 7, 0, 254, 252, 1, 0, 0, 0, 255, 258, 1, 0, 0, 0,
		256, 254, 1, 0, 0, 0, 256, 257, 1, 0, 0, 0, 257, 259, 1, 0, 0, 0, 258,
		256, 1, 0, 0, 0, 259, 260, 5, 32, 0, 0, 260, 271, 1, 0, 0, 0, 261, 271,
		3, 34, 17, 0, 262, 263, 3, 34, 17, 0, 263, 264, 5, 5, 0, 0, 264, 265, 3,
		34, 17, 0, 265, 271, 1, 0, 0, 0, 266, 267, 5, 29, 0, 0, 267, 268, 3, 36,
		18, 0, 268, 269, 5, 30, 0, 0, 269, 271, 1, 0, 0, 0, 270, 246, 1, 0, 0,
		0, 270, 261, 1, 0, 0, 0, 270, 262, 1, 0, 0, 0, 270, 266, 1, 0, 0, 0, 271,
		283, 1, 0, 0, 0, 272, 273, 10, 3, 0, 0, 273, 274, 5, 21, 0, 0, 274, 282,
		3, 36, 18, 4, 275, 276, 10, 2, 0, 0, 276, 277, 5, 22, 0, 0, 277, 282, 3,
		36, 18, 3, 278, 279, 10, 1, 0, 0, 279, 280, 5, 23, 0, 0, 280, 282, 3, 36,
		18, 2, 281, 272, 1, 0, 0, 0, 281, 275, 1, 0, 0, 0, 281, 278, 1, 0, 0, 0,
		282, 285, 1, 0, 0, 0, 283, 281, 1, 0, 0, 0, 283, 284, 1, 0, 0, 0, 284,
		37, 1, 0, 0, 0, 285, 283, 1, 0, 0, 0, 33, 45, 51, 64, 68, 73, 78, 82, 90,
		100, 114, 121, 128, 132, 142, 149, 163, 174, 176, 184, 188, 194, 201, 206,
		214, 220, 223, 240, 244, 249, 256, 270, 281, 283,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// kslParserInit initializes any static state used to implement kslParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewkslParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func KslParserInit() {
	staticData := &KslParserStaticData
	staticData.once.Do(kslParserInit)
}

// NewkslParser produces a new parser instance for the optional input antlr.TokenStream.
func NewkslParser(input antlr.TokenStream) *kslParser {
	KslParserInit()
	this := new(kslParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &KslParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "ksl.g4"

	return this
}

// kslParser tokens.
const (
	kslParserEOF              = antlr.TokenEOF
	kslParserKW_ESC           = 1
	kslParserFORCED_NAME      = 2
	kslParserVERSION          = 3
	kslParserVERSIONNUM       = 4
	kslParserRESOLVE          = 5
	kslParserNAMESPACE        = 6
	kslParserACCESS           = 7
	kslParserPUBLIC           = 8
	kslParserINTERNAL         = 9
	kslParserPRIVATE          = 10
	kslParserTYPE             = 11
	kslParserRELATION         = 12
	kslParserIMPORT           = 13
	kslParserEXTENSION        = 14
	kslParserCARDINALITY      = 15
	kslParserATMOSTONE        = 16
	kslParserEXACTLYONE       = 17
	kslParserATLEASTONE       = 18
	kslParserANY              = 19
	kslParserAS               = 20
	kslParserAND              = 21
	kslParserOR               = 22
	kslParserUNLESS           = 23
	kslParserALLOW_DUPLICATES = 24
	kslParserEXPAND           = 25
	kslParserLBRACE           = 26
	kslParserRBRACE           = 27
	kslParserEXTENSION_CALL   = 28
	kslParserLPAREN           = 29
	kslParserRPAREN           = 30
	kslParserLSQUARE          = 31
	kslParserRSQARE           = 32
	kslParserVARREF           = 33
	kslParserTEMPLATE_DELIM   = 34
	kslParserSTRING_DELIM     = 35
	kslParserARG_DELIM        = 36
	kslParserDECL_END         = 37
	kslParserNAME             = 38
	kslParserCOMMENT          = 39
	kslParserWS               = 40
)

// kslParser rules.
const (
	kslParserRULE_name               = 0
	kslParserRULE_file               = 1
	kslParserRULE_version            = 2
	kslParserRULE_namespace          = 3
	kslParserRULE_import_stmt        = 4
	kslParserRULE_declaration        = 5
	kslParserRULE_typeExpr           = 6
	kslParserRULE_typeReference      = 7
	kslParserRULE_extensionParam     = 8
	kslParserRULE_extensionParams    = 9
	kslParserRULE_extensionReference = 10
	kslParserRULE_relation           = 11
	kslParserRULE_relationBody       = 12
	kslParserRULE_paramNames         = 13
	kslParserRULE_extension          = 14
	kslParserRULE_dynamicType        = 15
	kslParserRULE_dynamicRelation    = 16
	kslParserRULE_dynamicName        = 17
	kslParserRULE_dynamicBody        = 18
)

// INameContext is an interface to support dynamic dispatch.
type INameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NAME() antlr.TerminalNode
	FORCED_NAME() antlr.TerminalNode

	// IsNameContext differentiates from other interfaces.
	IsNameContext()
}

type NameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNameContext() *NameContext {
	var p = new(NameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kslParserRULE_name
	return p
}

func InitEmptyNameContext(p *NameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kslParserRULE_name
}

func (*NameContext) IsNameContext() {}

func NewNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NameContext {
	var p = new(NameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = kslParserRULE_name

	return p
}

func (s *NameContext) GetParser() antlr.Parser { return s.parser }

func (s *NameContext) NAME() antlr.TerminalNode {
	return s.GetToken(kslParserNAME, 0)
}

func (s *NameContext) FORCED_NAME() antlr.TerminalNode {
	return s.GetToken(kslParserFORCED_NAME, 0)
}

func (s *NameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kslListener); ok {
		listenerT.EnterName(s)
	}
}

func (s *NameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kslListener); ok {
		listenerT.ExitName(s)
	}
}

func (s *NameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case kslVisitor:
		return t.VisitName(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *kslParser) Name() (localctx INameContext) {
	localctx = NewNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, kslParserRULE_name)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(38)
		_la = p.GetTokenStream().LA(1)

		if !(_la == kslParserFORCED_NAME || _la == kslParserNAME) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFileContext is an interface to support dynamic dispatch.
type IFileContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Version() IVersionContext
	Namespace() INamespaceContext
	AllImport_stmt() []IImport_stmtContext
	Import_stmt(i int) IImport_stmtContext
	AllDeclaration() []IDeclarationContext
	Declaration(i int) IDeclarationContext

	// IsFileContext differentiates from other interfaces.
	IsFileContext()
}

type FileContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFileContext() *FileContext {
	var p = new(FileContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kslParserRULE_file
	return p
}

func InitEmptyFileContext(p *FileContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kslParserRULE_file
}

func (*FileContext) IsFileContext() {}

func NewFileContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FileContext {
	var p = new(FileContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = kslParserRULE_file

	return p
}

func (s *FileContext) GetParser() antlr.Parser { return s.parser }

func (s *FileContext) Version() IVersionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVersionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVersionContext)
}

func (s *FileContext) Namespace() INamespaceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INamespaceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INamespaceContext)
}

func (s *FileContext) AllImport_stmt() []IImport_stmtContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IImport_stmtContext); ok {
			len++
		}
	}

	tst := make([]IImport_stmtContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IImport_stmtContext); ok {
			tst[i] = t.(IImport_stmtContext)
			i++
		}
	}

	return tst
}

func (s *FileContext) Import_stmt(i int) IImport_stmtContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IImport_stmtContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IImport_stmtContext)
}

func (s *FileContext) AllDeclaration() []IDeclarationContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDeclarationContext); ok {
			len++
		}
	}

	tst := make([]IDeclarationContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDeclarationContext); ok {
			tst[i] = t.(IDeclarationContext)
			i++
		}
	}

	return tst
}

func (s *FileContext) Declaration(i int) IDeclarationContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclarationContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclarationContext)
}

func (s *FileContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FileContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FileContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kslListener); ok {
		listenerT.EnterFile(s)
	}
}

func (s *FileContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kslListener); ok {
		listenerT.ExitFile(s)
	}
}

func (s *FileContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case kslVisitor:
		return t.VisitFile(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *kslParser) File() (localctx IFileContext) {
	localctx = NewFileContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, kslParserRULE_file)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(40)
		p.Version()
	}
	{
		p.SetState(41)
		p.Namespace()
	}
	p.SetState(45)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == kslParserIMPORT {
		{
			p.SetState(42)
			p.Import_stmt()
		}

		p.SetState(47)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(49)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&268454016) != 0) {
		{
			p.SetState(48)
			p.Declaration()
		}

		p.SetState(51)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IVersionContext is an interface to support dynamic dispatch.
type IVersionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	VERSION() antlr.TerminalNode
	VERSIONNUM() antlr.TerminalNode

	// IsVersionContext differentiates from other interfaces.
	IsVersionContext()
}

type VersionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVersionContext() *VersionContext {
	var p = new(VersionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kslParserRULE_version
	return p
}

func InitEmptyVersionContext(p *VersionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kslParserRULE_version
}

func (*VersionContext) IsVersionContext() {}

func NewVersionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VersionContext {
	var p = new(VersionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = kslParserRULE_version

	return p
}

func (s *VersionContext) GetParser() antlr.Parser { return s.parser }

func (s *VersionContext) VERSION() antlr.TerminalNode {
	return s.GetToken(kslParserVERSION, 0)
}

func (s *VersionContext) VERSIONNUM() antlr.TerminalNode {
	return s.GetToken(kslParserVERSIONNUM, 0)
}

func (s *VersionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VersionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VersionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kslListener); ok {
		listenerT.EnterVersion(s)
	}
}

func (s *VersionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kslListener); ok {
		listenerT.ExitVersion(s)
	}
}

func (s *VersionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case kslVisitor:
		return t.VisitVersion(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *kslParser) Version() (localctx IVersionContext) {
	localctx = NewVersionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, kslParserRULE_version)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(53)
		p.Match(kslParserVERSION)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(54)
		p.Match(kslParserVERSIONNUM)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// INamespaceContext is an interface to support dynamic dispatch.
type INamespaceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NAMESPACE() antlr.TerminalNode
	Name() INameContext

	// IsNamespaceContext differentiates from other interfaces.
	IsNamespaceContext()
}

type NamespaceContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNamespaceContext() *NamespaceContext {
	var p = new(NamespaceContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kslParserRULE_namespace
	return p
}

func InitEmptyNamespaceContext(p *NamespaceContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kslParserRULE_namespace
}

func (*NamespaceContext) IsNamespaceContext() {}

func NewNamespaceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NamespaceContext {
	var p = new(NamespaceContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = kslParserRULE_namespace

	return p
}

func (s *NamespaceContext) GetParser() antlr.Parser { return s.parser }

func (s *NamespaceContext) NAMESPACE() antlr.TerminalNode {
	return s.GetToken(kslParserNAMESPACE, 0)
}

func (s *NamespaceContext) Name() INameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *NamespaceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NamespaceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NamespaceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kslListener); ok {
		listenerT.EnterNamespace(s)
	}
}

func (s *NamespaceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kslListener); ok {
		listenerT.ExitNamespace(s)
	}
}

func (s *NamespaceContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case kslVisitor:
		return t.VisitNamespace(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *kslParser) Namespace() (localctx INamespaceContext) {
	localctx = NewNamespaceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, kslParserRULE_namespace)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(56)
		p.Match(kslParserNAMESPACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(57)
		p.Name()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IImport_stmtContext is an interface to support dynamic dispatch.
type IImport_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IMPORT() antlr.TerminalNode
	Name() INameContext

	// IsImport_stmtContext differentiates from other interfaces.
	IsImport_stmtContext()
}

type Import_stmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyImport_stmtContext() *Import_stmtContext {
	var p = new(Import_stmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kslParserRULE_import_stmt
	return p
}

func InitEmptyImport_stmtContext(p *Import_stmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kslParserRULE_import_stmt
}

func (*Import_stmtContext) IsImport_stmtContext() {}

func NewImport_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Import_stmtContext {
	var p = new(Import_stmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = kslParserRULE_import_stmt

	return p
}

func (s *Import_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Import_stmtContext) IMPORT() antlr.TerminalNode {
	return s.GetToken(kslParserIMPORT, 0)
}

func (s *Import_stmtContext) Name() INameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *Import_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Import_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Import_stmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kslListener); ok {
		listenerT.EnterImport_stmt(s)
	}
}

func (s *Import_stmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kslListener); ok {
		listenerT.ExitImport_stmt(s)
	}
}

func (s *Import_stmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case kslVisitor:
		return t.VisitImport_stmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *kslParser) Import_stmt() (localctx IImport_stmtContext) {
	localctx = NewImport_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, kslParserRULE_import_stmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(59)
		p.Match(kslParserIMPORT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(60)
		p.Name()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDeclarationContext is an interface to support dynamic dispatch.
type IDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TypeExpr() ITypeExprContext
	DECL_END() antlr.TerminalNode
	Extension() IExtensionContext
	ExtensionReference() IExtensionReferenceContext

	// IsDeclarationContext differentiates from other interfaces.
	IsDeclarationContext()
}

type DeclarationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeclarationContext() *DeclarationContext {
	var p = new(DeclarationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kslParserRULE_declaration
	return p
}

func InitEmptyDeclarationContext(p *DeclarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kslParserRULE_declaration
}

func (*DeclarationContext) IsDeclarationContext() {}

func NewDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclarationContext {
	var p = new(DeclarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = kslParserRULE_declaration

	return p
}

func (s *DeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclarationContext) TypeExpr() ITypeExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeExprContext)
}

func (s *DeclarationContext) DECL_END() antlr.TerminalNode {
	return s.GetToken(kslParserDECL_END, 0)
}

func (s *DeclarationContext) Extension() IExtensionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExtensionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExtensionContext)
}

func (s *DeclarationContext) ExtensionReference() IExtensionReferenceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExtensionReferenceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExtensionReferenceContext)
}

func (s *DeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kslListener); ok {
		listenerT.EnterDeclaration(s)
	}
}

func (s *DeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kslListener); ok {
		listenerT.ExitDeclaration(s)
	}
}

func (s *DeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case kslVisitor:
		return t.VisitDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *kslParser) Declaration() (localctx IDeclarationContext) {
	localctx = NewDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, kslParserRULE_declaration)
	var _la int

	p.SetState(73)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(62)
			p.TypeExpr()
		}
		p.SetState(64)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == kslParserDECL_END {
			{
				p.SetState(63)
				p.Match(kslParserDECL_END)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(66)
			p.Extension()
		}
		p.SetState(68)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == kslParserDECL_END {
			{
				p.SetState(67)
				p.Match(kslParserDECL_END)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(70)
			p.ExtensionReference()
		}
		{
			p.SetState(71)
			p.Match(kslParserDECL_END)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITypeExprContext is an interface to support dynamic dispatch.
type ITypeExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TYPE() antlr.TerminalNode
	Name() INameContext
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	AllExtensionReference() []IExtensionReferenceContext
	ExtensionReference(i int) IExtensionReferenceContext
	ACCESS() antlr.TerminalNode
	AllRelation() []IRelationContext
	Relation(i int) IRelationContext

	// IsTypeExprContext differentiates from other interfaces.
	IsTypeExprContext()
}

type TypeExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeExprContext() *TypeExprContext {
	var p = new(TypeExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kslParserRULE_typeExpr
	return p
}

func InitEmptyTypeExprContext(p *TypeExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kslParserRULE_typeExpr
}

func (*TypeExprContext) IsTypeExprContext() {}

func NewTypeExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeExprContext {
	var p = new(TypeExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = kslParserRULE_typeExpr

	return p
}

func (s *TypeExprContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeExprContext) TYPE() antlr.TerminalNode {
	return s.GetToken(kslParserTYPE, 0)
}

func (s *TypeExprContext) Name() INameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *TypeExprContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(kslParserLBRACE, 0)
}

func (s *TypeExprContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(kslParserRBRACE, 0)
}

func (s *TypeExprContext) AllExtensionReference() []IExtensionReferenceContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExtensionReferenceContext); ok {
			len++
		}
	}

	tst := make([]IExtensionReferenceContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExtensionReferenceContext); ok {
			tst[i] = t.(IExtensionReferenceContext)
			i++
		}
	}

	return tst
}

func (s *TypeExprContext) ExtensionReference(i int) IExtensionReferenceContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExtensionReferenceContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExtensionReferenceContext)
}

func (s *TypeExprContext) ACCESS() antlr.TerminalNode {
	return s.GetToken(kslParserACCESS, 0)
}

func (s *TypeExprContext) AllRelation() []IRelationContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IRelationContext); ok {
			len++
		}
	}

	tst := make([]IRelationContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IRelationContext); ok {
			tst[i] = t.(IRelationContext)
			i++
		}
	}

	return tst
}

func (s *TypeExprContext) Relation(i int) IRelationContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRelationContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRelationContext)
}

func (s *TypeExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kslListener); ok {
		listenerT.EnterTypeExpr(s)
	}
}

func (s *TypeExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kslListener); ok {
		listenerT.ExitTypeExpr(s)
	}
}

func (s *TypeExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case kslVisitor:
		return t.VisitTypeExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *kslParser) TypeExpr() (localctx ITypeExprContext) {
	localctx = NewTypeExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, kslParserRULE_typeExpr)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(78)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == kslParserEXTENSION_CALL {
		{
			p.SetState(75)
			p.ExtensionReference()
		}

		p.SetState(80)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(82)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == kslParserACCESS {
		{
			p.SetState(81)
			p.Match(kslParserACCESS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(84)
		p.Match(kslParserTYPE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(85)
		p.Name()
	}
	{
		p.SetState(86)
		p.Match(kslParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(90)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&268439680) != 0 {
		{
			p.SetState(87)
			p.Relation()
		}

		p.SetState(92)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(93)
		p.Match(kslParserRBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITypeReferenceContext is an interface to support dynamic dispatch.
type ITypeReferenceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllName() []INameContext
	Name(i int) INameContext
	AllRESOLVE() []antlr.TerminalNode
	RESOLVE(i int) antlr.TerminalNode

	// IsTypeReferenceContext differentiates from other interfaces.
	IsTypeReferenceContext()
}

type TypeReferenceContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeReferenceContext() *TypeReferenceContext {
	var p = new(TypeReferenceContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kslParserRULE_typeReference
	return p
}

func InitEmptyTypeReferenceContext(p *TypeReferenceContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kslParserRULE_typeReference
}

func (*TypeReferenceContext) IsTypeReferenceContext() {}

func NewTypeReferenceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeReferenceContext {
	var p = new(TypeReferenceContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = kslParserRULE_typeReference

	return p
}

func (s *TypeReferenceContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeReferenceContext) AllName() []INameContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(INameContext); ok {
			len++
		}
	}

	tst := make([]INameContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(INameContext); ok {
			tst[i] = t.(INameContext)
			i++
		}
	}

	return tst
}

func (s *TypeReferenceContext) Name(i int) INameContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INameContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *TypeReferenceContext) AllRESOLVE() []antlr.TerminalNode {
	return s.GetTokens(kslParserRESOLVE)
}

func (s *TypeReferenceContext) RESOLVE(i int) antlr.TerminalNode {
	return s.GetToken(kslParserRESOLVE, i)
}

func (s *TypeReferenceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeReferenceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeReferenceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kslListener); ok {
		listenerT.EnterTypeReference(s)
	}
}

func (s *TypeReferenceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kslListener); ok {
		listenerT.ExitTypeReference(s)
	}
}

func (s *TypeReferenceContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case kslVisitor:
		return t.VisitTypeReference(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *kslParser) TypeReference() (localctx ITypeReferenceContext) {
	localctx = NewTypeReferenceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, kslParserRULE_typeReference)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(95)
		p.Name()
	}
	p.SetState(100)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == kslParserRESOLVE {
		{
			p.SetState(96)
			p.Match(kslParserRESOLVE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(97)
			p.Name()
		}

		p.SetState(102)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExtensionParamContext is an interface to support dynamic dispatch.
type IExtensionParamContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetValue returns the value token.
	GetValue() antlr.Token

	// SetValue sets the value token.
	SetValue(antlr.Token)

	// Getter signatures
	Name() INameContext
	EXPAND() antlr.TerminalNode
	AllSTRING_DELIM() []antlr.TerminalNode
	STRING_DELIM(i int) antlr.TerminalNode

	// IsExtensionParamContext differentiates from other interfaces.
	IsExtensionParamContext()
}

type ExtensionParamContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	value  antlr.Token
}

func NewEmptyExtensionParamContext() *ExtensionParamContext {
	var p = new(ExtensionParamContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kslParserRULE_extensionParam
	return p
}

func InitEmptyExtensionParamContext(p *ExtensionParamContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kslParserRULE_extensionParam
}

func (*ExtensionParamContext) IsExtensionParamContext() {}

func NewExtensionParamContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExtensionParamContext {
	var p = new(ExtensionParamContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = kslParserRULE_extensionParam

	return p
}

func (s *ExtensionParamContext) GetParser() antlr.Parser { return s.parser }

func (s *ExtensionParamContext) GetValue() antlr.Token { return s.value }

func (s *ExtensionParamContext) SetValue(v antlr.Token) { s.value = v }

func (s *ExtensionParamContext) Name() INameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *ExtensionParamContext) EXPAND() antlr.TerminalNode {
	return s.GetToken(kslParserEXPAND, 0)
}

func (s *ExtensionParamContext) AllSTRING_DELIM() []antlr.TerminalNode {
	return s.GetTokens(kslParserSTRING_DELIM)
}

func (s *ExtensionParamContext) STRING_DELIM(i int) antlr.TerminalNode {
	return s.GetToken(kslParserSTRING_DELIM, i)
}

func (s *ExtensionParamContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExtensionParamContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExtensionParamContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kslListener); ok {
		listenerT.EnterExtensionParam(s)
	}
}

func (s *ExtensionParamContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kslListener); ok {
		listenerT.ExitExtensionParam(s)
	}
}

func (s *ExtensionParamContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case kslVisitor:
		return t.VisitExtensionParam(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *kslParser) ExtensionParam() (localctx IExtensionParamContext) {
	localctx = NewExtensionParamContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, kslParserRULE_extensionParam)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(103)
		p.Name()
	}
	{
		p.SetState(104)
		p.Match(kslParserEXPAND)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(105)
		p.Match(kslParserSTRING_DELIM)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(106)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*ExtensionParamContext).value = _lt

		_la = p.GetTokenStream().LA(1)

		if _la <= 0 || _la == kslParserSTRING_DELIM {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*ExtensionParamContext).value = _ri
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(107)
		p.Match(kslParserSTRING_DELIM)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExtensionParamsContext is an interface to support dynamic dispatch.
type IExtensionParamsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExtensionParam() []IExtensionParamContext
	ExtensionParam(i int) IExtensionParamContext
	AllARG_DELIM() []antlr.TerminalNode
	ARG_DELIM(i int) antlr.TerminalNode

	// IsExtensionParamsContext differentiates from other interfaces.
	IsExtensionParamsContext()
}

type ExtensionParamsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExtensionParamsContext() *ExtensionParamsContext {
	var p = new(ExtensionParamsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kslParserRULE_extensionParams
	return p
}

func InitEmptyExtensionParamsContext(p *ExtensionParamsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kslParserRULE_extensionParams
}

func (*ExtensionParamsContext) IsExtensionParamsContext() {}

func NewExtensionParamsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExtensionParamsContext {
	var p = new(ExtensionParamsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = kslParserRULE_extensionParams

	return p
}

func (s *ExtensionParamsContext) GetParser() antlr.Parser { return s.parser }

func (s *ExtensionParamsContext) AllExtensionParam() []IExtensionParamContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExtensionParamContext); ok {
			len++
		}
	}

	tst := make([]IExtensionParamContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExtensionParamContext); ok {
			tst[i] = t.(IExtensionParamContext)
			i++
		}
	}

	return tst
}

func (s *ExtensionParamsContext) ExtensionParam(i int) IExtensionParamContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExtensionParamContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExtensionParamContext)
}

func (s *ExtensionParamsContext) AllARG_DELIM() []antlr.TerminalNode {
	return s.GetTokens(kslParserARG_DELIM)
}

func (s *ExtensionParamsContext) ARG_DELIM(i int) antlr.TerminalNode {
	return s.GetToken(kslParserARG_DELIM, i)
}

func (s *ExtensionParamsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExtensionParamsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExtensionParamsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kslListener); ok {
		listenerT.EnterExtensionParams(s)
	}
}

func (s *ExtensionParamsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kslListener); ok {
		listenerT.ExitExtensionParams(s)
	}
}

func (s *ExtensionParamsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case kslVisitor:
		return t.VisitExtensionParams(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *kslParser) ExtensionParams() (localctx IExtensionParamsContext) {
	localctx = NewExtensionParamsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, kslParserRULE_extensionParams)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(109)
		p.ExtensionParam()
	}
	p.SetState(114)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == kslParserARG_DELIM {
		{
			p.SetState(110)
			p.Match(kslParserARG_DELIM)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(111)
			p.ExtensionParam()
		}

		p.SetState(116)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExtensionReferenceContext is an interface to support dynamic dispatch.
type IExtensionReferenceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EXTENSION_CALL() antlr.TerminalNode
	TypeReference() ITypeReferenceContext
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	ExtensionParams() IExtensionParamsContext

	// IsExtensionReferenceContext differentiates from other interfaces.
	IsExtensionReferenceContext()
}

type ExtensionReferenceContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExtensionReferenceContext() *ExtensionReferenceContext {
	var p = new(ExtensionReferenceContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kslParserRULE_extensionReference
	return p
}

func InitEmptyExtensionReferenceContext(p *ExtensionReferenceContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kslParserRULE_extensionReference
}

func (*ExtensionReferenceContext) IsExtensionReferenceContext() {}

func NewExtensionReferenceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExtensionReferenceContext {
	var p = new(ExtensionReferenceContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = kslParserRULE_extensionReference

	return p
}

func (s *ExtensionReferenceContext) GetParser() antlr.Parser { return s.parser }

func (s *ExtensionReferenceContext) EXTENSION_CALL() antlr.TerminalNode {
	return s.GetToken(kslParserEXTENSION_CALL, 0)
}

func (s *ExtensionReferenceContext) TypeReference() ITypeReferenceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeReferenceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeReferenceContext)
}

func (s *ExtensionReferenceContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(kslParserLPAREN, 0)
}

func (s *ExtensionReferenceContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(kslParserRPAREN, 0)
}

func (s *ExtensionReferenceContext) ExtensionParams() IExtensionParamsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExtensionParamsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExtensionParamsContext)
}

func (s *ExtensionReferenceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExtensionReferenceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExtensionReferenceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kslListener); ok {
		listenerT.EnterExtensionReference(s)
	}
}

func (s *ExtensionReferenceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kslListener); ok {
		listenerT.ExitExtensionReference(s)
	}
}

func (s *ExtensionReferenceContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case kslVisitor:
		return t.VisitExtensionReference(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *kslParser) ExtensionReference() (localctx IExtensionReferenceContext) {
	localctx = NewExtensionReferenceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, kslParserRULE_extensionReference)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(117)
		p.Match(kslParserEXTENSION_CALL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(118)
		p.TypeReference()
	}
	{
		p.SetState(119)
		p.Match(kslParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(121)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == kslParserFORCED_NAME || _la == kslParserNAME {
		{
			p.SetState(120)
			p.ExtensionParams()
		}

	}
	{
		p.SetState(123)
		p.Match(kslParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRelationContext is an interface to support dynamic dispatch.
type IRelationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	RELATION() antlr.TerminalNode
	Name() INameContext
	EXPAND() antlr.TerminalNode
	RelationBody() IRelationBodyContext
	AllExtensionReference() []IExtensionReferenceContext
	ExtensionReference(i int) IExtensionReferenceContext
	ACCESS() antlr.TerminalNode

	// IsRelationContext differentiates from other interfaces.
	IsRelationContext()
}

type RelationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRelationContext() *RelationContext {
	var p = new(RelationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kslParserRULE_relation
	return p
}

func InitEmptyRelationContext(p *RelationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kslParserRULE_relation
}

func (*RelationContext) IsRelationContext() {}

func NewRelationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RelationContext {
	var p = new(RelationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = kslParserRULE_relation

	return p
}

func (s *RelationContext) GetParser() antlr.Parser { return s.parser }

func (s *RelationContext) RELATION() antlr.TerminalNode {
	return s.GetToken(kslParserRELATION, 0)
}

func (s *RelationContext) Name() INameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *RelationContext) EXPAND() antlr.TerminalNode {
	return s.GetToken(kslParserEXPAND, 0)
}

func (s *RelationContext) RelationBody() IRelationBodyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRelationBodyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRelationBodyContext)
}

func (s *RelationContext) AllExtensionReference() []IExtensionReferenceContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExtensionReferenceContext); ok {
			len++
		}
	}

	tst := make([]IExtensionReferenceContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExtensionReferenceContext); ok {
			tst[i] = t.(IExtensionReferenceContext)
			i++
		}
	}

	return tst
}

func (s *RelationContext) ExtensionReference(i int) IExtensionReferenceContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExtensionReferenceContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExtensionReferenceContext)
}

func (s *RelationContext) ACCESS() antlr.TerminalNode {
	return s.GetToken(kslParserACCESS, 0)
}

func (s *RelationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RelationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RelationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kslListener); ok {
		listenerT.EnterRelation(s)
	}
}

func (s *RelationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kslListener); ok {
		listenerT.ExitRelation(s)
	}
}

func (s *RelationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case kslVisitor:
		return t.VisitRelation(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *kslParser) Relation() (localctx IRelationContext) {
	localctx = NewRelationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, kslParserRULE_relation)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(128)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == kslParserEXTENSION_CALL {
		{
			p.SetState(125)
			p.ExtensionReference()
		}

		p.SetState(130)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(132)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == kslParserACCESS {
		{
			p.SetState(131)
			p.Match(kslParserACCESS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(134)
		p.Match(kslParserRELATION)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(135)
		p.Name()
	}
	{
		p.SetState(136)
		p.Match(kslParserEXPAND)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(137)
		p.relationBody(0)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRelationBodyContext is an interface to support dynamic dispatch.
type IRelationBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsRelationBodyContext differentiates from other interfaces.
	IsRelationBodyContext()
}

type RelationBodyContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRelationBodyContext() *RelationBodyContext {
	var p = new(RelationBodyContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kslParserRULE_relationBody
	return p
}

func InitEmptyRelationBodyContext(p *RelationBodyContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kslParserRULE_relationBody
}

func (*RelationBodyContext) IsRelationBodyContext() {}

func NewRelationBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RelationBodyContext {
	var p = new(RelationBodyContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = kslParserRULE_relationBody

	return p
}

func (s *RelationBodyContext) GetParser() antlr.Parser { return s.parser }

func (s *RelationBodyContext) CopyAll(ctx *RelationBodyContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *RelationBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RelationBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ORContext struct {
	RelationBodyContext
}

func NewORContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ORContext {
	var p = new(ORContext)

	InitEmptyRelationBodyContext(&p.RelationBodyContext)
	p.parser = parser
	p.CopyAll(ctx.(*RelationBodyContext))

	return p
}

func (s *ORContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ORContext) AllRelationBody() []IRelationBodyContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IRelationBodyContext); ok {
			len++
		}
	}

	tst := make([]IRelationBodyContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IRelationBodyContext); ok {
			tst[i] = t.(IRelationBodyContext)
			i++
		}
	}

	return tst
}

func (s *ORContext) RelationBody(i int) IRelationBodyContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRelationBodyContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRelationBodyContext)
}

func (s *ORContext) OR() antlr.TerminalNode {
	return s.GetToken(kslParserOR, 0)
}

func (s *ORContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kslListener); ok {
		listenerT.EnterOR(s)
	}
}

func (s *ORContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kslListener); ok {
		listenerT.ExitOR(s)
	}
}

func (s *ORContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case kslVisitor:
		return t.VisitOR(s)

	default:
		return t.VisitChildren(s)
	}
}

type ReferenceContext struct {
	RelationBodyContext
}

func NewReferenceContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ReferenceContext {
	var p = new(ReferenceContext)

	InitEmptyRelationBodyContext(&p.RelationBodyContext)
	p.parser = parser
	p.CopyAll(ctx.(*RelationBodyContext))

	return p
}

func (s *ReferenceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReferenceContext) Name() INameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *ReferenceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kslListener); ok {
		listenerT.EnterReference(s)
	}
}

func (s *ReferenceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kslListener); ok {
		listenerT.ExitReference(s)
	}
}

func (s *ReferenceContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case kslVisitor:
		return t.VisitReference(s)

	default:
		return t.VisitChildren(s)
	}
}

type AndContext struct {
	RelationBodyContext
}

func NewAndContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AndContext {
	var p = new(AndContext)

	InitEmptyRelationBodyContext(&p.RelationBodyContext)
	p.parser = parser
	p.CopyAll(ctx.(*RelationBodyContext))

	return p
}

func (s *AndContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AndContext) AllRelationBody() []IRelationBodyContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IRelationBodyContext); ok {
			len++
		}
	}

	tst := make([]IRelationBodyContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IRelationBodyContext); ok {
			tst[i] = t.(IRelationBodyContext)
			i++
		}
	}

	return tst
}

func (s *AndContext) RelationBody(i int) IRelationBodyContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRelationBodyContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRelationBodyContext)
}

func (s *AndContext) AND() antlr.TerminalNode {
	return s.GetToken(kslParserAND, 0)
}

func (s *AndContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kslListener); ok {
		listenerT.EnterAnd(s)
	}
}

func (s *AndContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kslListener); ok {
		listenerT.ExitAnd(s)
	}
}

func (s *AndContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case kslVisitor:
		return t.VisitAnd(s)

	default:
		return t.VisitChildren(s)
	}
}

type SelfContext struct {
	RelationBodyContext
}

func NewSelfContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SelfContext {
	var p = new(SelfContext)

	InitEmptyRelationBodyContext(&p.RelationBodyContext)
	p.parser = parser
	p.CopyAll(ctx.(*RelationBodyContext))

	return p
}

func (s *SelfContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelfContext) LSQUARE() antlr.TerminalNode {
	return s.GetToken(kslParserLSQUARE, 0)
}

func (s *SelfContext) AllTypeReference() []ITypeReferenceContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITypeReferenceContext); ok {
			len++
		}
	}

	tst := make([]ITypeReferenceContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITypeReferenceContext); ok {
			tst[i] = t.(ITypeReferenceContext)
			i++
		}
	}

	return tst
}

func (s *SelfContext) TypeReference(i int) ITypeReferenceContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeReferenceContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeReferenceContext)
}

func (s *SelfContext) RSQARE() antlr.TerminalNode {
	return s.GetToken(kslParserRSQARE, 0)
}

func (s *SelfContext) CARDINALITY() antlr.TerminalNode {
	return s.GetToken(kslParserCARDINALITY, 0)
}

func (s *SelfContext) AllOR() []antlr.TerminalNode {
	return s.GetTokens(kslParserOR)
}

func (s *SelfContext) OR(i int) antlr.TerminalNode {
	return s.GetToken(kslParserOR, i)
}

func (s *SelfContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kslListener); ok {
		listenerT.EnterSelf(s)
	}
}

func (s *SelfContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kslListener); ok {
		listenerT.ExitSelf(s)
	}
}

func (s *SelfContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case kslVisitor:
		return t.VisitSelf(s)

	default:
		return t.VisitChildren(s)
	}
}

type SubRelationContext struct {
	RelationBodyContext
	relationName    INameContext
	subrelationName INameContext
}

func NewSubRelationContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SubRelationContext {
	var p = new(SubRelationContext)

	InitEmptyRelationBodyContext(&p.RelationBodyContext)
	p.parser = parser
	p.CopyAll(ctx.(*RelationBodyContext))

	return p
}

func (s *SubRelationContext) GetRelationName() INameContext { return s.relationName }

func (s *SubRelationContext) GetSubrelationName() INameContext { return s.subrelationName }

func (s *SubRelationContext) SetRelationName(v INameContext) { s.relationName = v }

func (s *SubRelationContext) SetSubrelationName(v INameContext) { s.subrelationName = v }

func (s *SubRelationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SubRelationContext) RESOLVE() antlr.TerminalNode {
	return s.GetToken(kslParserRESOLVE, 0)
}

func (s *SubRelationContext) AllName() []INameContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(INameContext); ok {
			len++
		}
	}

	tst := make([]INameContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(INameContext); ok {
			tst[i] = t.(INameContext)
			i++
		}
	}

	return tst
}

func (s *SubRelationContext) Name(i int) INameContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INameContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *SubRelationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kslListener); ok {
		listenerT.EnterSubRelation(s)
	}
}

func (s *SubRelationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kslListener); ok {
		listenerT.ExitSubRelation(s)
	}
}

func (s *SubRelationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case kslVisitor:
		return t.VisitSubRelation(s)

	default:
		return t.VisitChildren(s)
	}
}

type UnlessContext struct {
	RelationBodyContext
}

func NewUnlessContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *UnlessContext {
	var p = new(UnlessContext)

	InitEmptyRelationBodyContext(&p.RelationBodyContext)
	p.parser = parser
	p.CopyAll(ctx.(*RelationBodyContext))

	return p
}

func (s *UnlessContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnlessContext) AllRelationBody() []IRelationBodyContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IRelationBodyContext); ok {
			len++
		}
	}

	tst := make([]IRelationBodyContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IRelationBodyContext); ok {
			tst[i] = t.(IRelationBodyContext)
			i++
		}
	}

	return tst
}

func (s *UnlessContext) RelationBody(i int) IRelationBodyContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRelationBodyContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRelationBodyContext)
}

func (s *UnlessContext) UNLESS() antlr.TerminalNode {
	return s.GetToken(kslParserUNLESS, 0)
}

func (s *UnlessContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kslListener); ok {
		listenerT.EnterUnless(s)
	}
}

func (s *UnlessContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kslListener); ok {
		listenerT.ExitUnless(s)
	}
}

func (s *UnlessContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case kslVisitor:
		return t.VisitUnless(s)

	default:
		return t.VisitChildren(s)
	}
}

type ParenContext struct {
	RelationBodyContext
}

func NewParenContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParenContext {
	var p = new(ParenContext)

	InitEmptyRelationBodyContext(&p.RelationBodyContext)
	p.parser = parser
	p.CopyAll(ctx.(*RelationBodyContext))

	return p
}

func (s *ParenContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParenContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(kslParserLPAREN, 0)
}

func (s *ParenContext) RelationBody() IRelationBodyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRelationBodyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRelationBodyContext)
}

func (s *ParenContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(kslParserRPAREN, 0)
}

func (s *ParenContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kslListener); ok {
		listenerT.EnterParen(s)
	}
}

func (s *ParenContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kslListener); ok {
		listenerT.ExitParen(s)
	}
}

func (s *ParenContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case kslVisitor:
		return t.VisitParen(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *kslParser) RelationBody() (localctx IRelationBodyContext) {
	return p.relationBody(0)
}

func (p *kslParser) relationBody(_p int) (localctx IRelationBodyContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewRelationBodyContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IRelationBodyContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 24
	p.EnterRecursionRule(localctx, 24, kslParserRULE_relationBody, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(163)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 15, p.GetParserRuleContext()) {
	case 1:
		localctx = NewSelfContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(140)
			p.Match(kslParserLSQUARE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(142)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == kslParserCARDINALITY {
			{
				p.SetState(141)
				p.Match(kslParserCARDINALITY)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(144)
			p.TypeReference()
		}
		p.SetState(149)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == kslParserOR {
			{
				p.SetState(145)
				p.Match(kslParserOR)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(146)
				p.TypeReference()
			}

			p.SetState(151)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(152)
			p.Match(kslParserRSQARE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewReferenceContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(154)
			p.Name()
		}

	case 3:
		localctx = NewSubRelationContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(155)

			var _x = p.Name()

			localctx.(*SubRelationContext).relationName = _x
		}
		{
			p.SetState(156)
			p.Match(kslParserRESOLVE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(157)

			var _x = p.Name()

			localctx.(*SubRelationContext).subrelationName = _x
		}

	case 4:
		localctx = NewParenContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(159)
			p.Match(kslParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(160)
			p.relationBody(0)
		}
		{
			p.SetState(161)
			p.Match(kslParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(176)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 17, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(174)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 16, p.GetParserRuleContext()) {
			case 1:
				localctx = NewAndContext(p, NewRelationBodyContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, kslParserRULE_relationBody)
				p.SetState(165)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
					goto errorExit
				}
				{
					p.SetState(166)
					p.Match(kslParserAND)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(167)
					p.relationBody(4)
				}

			case 2:
				localctx = NewORContext(p, NewRelationBodyContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, kslParserRULE_relationBody)
				p.SetState(168)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
					goto errorExit
				}
				{
					p.SetState(169)
					p.Match(kslParserOR)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(170)
					p.relationBody(3)
				}

			case 3:
				localctx = NewUnlessContext(p, NewRelationBodyContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, kslParserRULE_relationBody)
				p.SetState(171)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
					goto errorExit
				}
				{
					p.SetState(172)
					p.Match(kslParserUNLESS)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(173)
					p.relationBody(2)
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(178)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 17, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IParamNamesContext is an interface to support dynamic dispatch.
type IParamNamesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllName() []INameContext
	Name(i int) INameContext
	AllARG_DELIM() []antlr.TerminalNode
	ARG_DELIM(i int) antlr.TerminalNode

	// IsParamNamesContext differentiates from other interfaces.
	IsParamNamesContext()
}

type ParamNamesContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParamNamesContext() *ParamNamesContext {
	var p = new(ParamNamesContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kslParserRULE_paramNames
	return p
}

func InitEmptyParamNamesContext(p *ParamNamesContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kslParserRULE_paramNames
}

func (*ParamNamesContext) IsParamNamesContext() {}

func NewParamNamesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParamNamesContext {
	var p = new(ParamNamesContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = kslParserRULE_paramNames

	return p
}

func (s *ParamNamesContext) GetParser() antlr.Parser { return s.parser }

func (s *ParamNamesContext) AllName() []INameContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(INameContext); ok {
			len++
		}
	}

	tst := make([]INameContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(INameContext); ok {
			tst[i] = t.(INameContext)
			i++
		}
	}

	return tst
}

func (s *ParamNamesContext) Name(i int) INameContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INameContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *ParamNamesContext) AllARG_DELIM() []antlr.TerminalNode {
	return s.GetTokens(kslParserARG_DELIM)
}

func (s *ParamNamesContext) ARG_DELIM(i int) antlr.TerminalNode {
	return s.GetToken(kslParserARG_DELIM, i)
}

func (s *ParamNamesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParamNamesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParamNamesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kslListener); ok {
		listenerT.EnterParamNames(s)
	}
}

func (s *ParamNamesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kslListener); ok {
		listenerT.ExitParamNames(s)
	}
}

func (s *ParamNamesContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case kslVisitor:
		return t.VisitParamNames(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *kslParser) ParamNames() (localctx IParamNamesContext) {
	localctx = NewParamNamesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, kslParserRULE_paramNames)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(179)
		p.Name()
	}
	p.SetState(184)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == kslParserARG_DELIM {
		{
			p.SetState(180)
			p.Match(kslParserARG_DELIM)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(181)
			p.Name()
		}

		p.SetState(186)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExtensionContext is an interface to support dynamic dispatch.
type IExtensionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EXTENSION() antlr.TerminalNode
	Name() INameContext
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	ACCESS() antlr.TerminalNode
	ParamNames() IParamNamesContext
	AllDynamicType() []IDynamicTypeContext
	DynamicType(i int) IDynamicTypeContext

	// IsExtensionContext differentiates from other interfaces.
	IsExtensionContext()
}

type ExtensionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExtensionContext() *ExtensionContext {
	var p = new(ExtensionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kslParserRULE_extension
	return p
}

func InitEmptyExtensionContext(p *ExtensionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kslParserRULE_extension
}

func (*ExtensionContext) IsExtensionContext() {}

func NewExtensionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExtensionContext {
	var p = new(ExtensionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = kslParserRULE_extension

	return p
}

func (s *ExtensionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExtensionContext) EXTENSION() antlr.TerminalNode {
	return s.GetToken(kslParserEXTENSION, 0)
}

func (s *ExtensionContext) Name() INameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *ExtensionContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(kslParserLPAREN, 0)
}

func (s *ExtensionContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(kslParserRPAREN, 0)
}

func (s *ExtensionContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(kslParserLBRACE, 0)
}

func (s *ExtensionContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(kslParserRBRACE, 0)
}

func (s *ExtensionContext) ACCESS() antlr.TerminalNode {
	return s.GetToken(kslParserACCESS, 0)
}

func (s *ExtensionContext) ParamNames() IParamNamesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParamNamesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParamNamesContext)
}

func (s *ExtensionContext) AllDynamicType() []IDynamicTypeContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDynamicTypeContext); ok {
			len++
		}
	}

	tst := make([]IDynamicTypeContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDynamicTypeContext); ok {
			tst[i] = t.(IDynamicTypeContext)
			i++
		}
	}

	return tst
}

func (s *ExtensionContext) DynamicType(i int) IDynamicTypeContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDynamicTypeContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDynamicTypeContext)
}

func (s *ExtensionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExtensionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExtensionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kslListener); ok {
		listenerT.EnterExtension(s)
	}
}

func (s *ExtensionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kslListener); ok {
		listenerT.ExitExtension(s)
	}
}

func (s *ExtensionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case kslVisitor:
		return t.VisitExtension(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *kslParser) Extension() (localctx IExtensionContext) {
	localctx = NewExtensionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, kslParserRULE_extension)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(188)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == kslParserACCESS {
		{
			p.SetState(187)
			p.Match(kslParserACCESS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(190)
		p.Match(kslParserEXTENSION)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(191)
		p.Name()
	}
	{
		p.SetState(192)
		p.Match(kslParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(194)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == kslParserFORCED_NAME || _la == kslParserNAME {
		{
			p.SetState(193)
			p.ParamNames()
		}

	}
	{
		p.SetState(196)
		p.Match(kslParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(197)
		p.Match(kslParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(199)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == kslParserACCESS || _la == kslParserTYPE {
		{
			p.SetState(198)
			p.DynamicType()
		}

		p.SetState(201)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(203)
		p.Match(kslParserRBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDynamicTypeContext is an interface to support dynamic dispatch.
type IDynamicTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TYPE() antlr.TerminalNode
	DynamicName() IDynamicNameContext
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	ACCESS() antlr.TerminalNode
	AllDynamicRelation() []IDynamicRelationContext
	DynamicRelation(i int) IDynamicRelationContext

	// IsDynamicTypeContext differentiates from other interfaces.
	IsDynamicTypeContext()
}

type DynamicTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDynamicTypeContext() *DynamicTypeContext {
	var p = new(DynamicTypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kslParserRULE_dynamicType
	return p
}

func InitEmptyDynamicTypeContext(p *DynamicTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kslParserRULE_dynamicType
}

func (*DynamicTypeContext) IsDynamicTypeContext() {}

func NewDynamicTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DynamicTypeContext {
	var p = new(DynamicTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = kslParserRULE_dynamicType

	return p
}

func (s *DynamicTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *DynamicTypeContext) TYPE() antlr.TerminalNode {
	return s.GetToken(kslParserTYPE, 0)
}

func (s *DynamicTypeContext) DynamicName() IDynamicNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDynamicNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDynamicNameContext)
}

func (s *DynamicTypeContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(kslParserLBRACE, 0)
}

func (s *DynamicTypeContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(kslParserRBRACE, 0)
}

func (s *DynamicTypeContext) ACCESS() antlr.TerminalNode {
	return s.GetToken(kslParserACCESS, 0)
}

func (s *DynamicTypeContext) AllDynamicRelation() []IDynamicRelationContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDynamicRelationContext); ok {
			len++
		}
	}

	tst := make([]IDynamicRelationContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDynamicRelationContext); ok {
			tst[i] = t.(IDynamicRelationContext)
			i++
		}
	}

	return tst
}

func (s *DynamicTypeContext) DynamicRelation(i int) IDynamicRelationContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDynamicRelationContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDynamicRelationContext)
}

func (s *DynamicTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DynamicTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DynamicTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kslListener); ok {
		listenerT.EnterDynamicType(s)
	}
}

func (s *DynamicTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kslListener); ok {
		listenerT.ExitDynamicType(s)
	}
}

func (s *DynamicTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case kslVisitor:
		return t.VisitDynamicType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *kslParser) DynamicType() (localctx IDynamicTypeContext) {
	localctx = NewDynamicTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, kslParserRULE_dynamicType)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(206)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == kslParserACCESS {
		{
			p.SetState(205)
			p.Match(kslParserACCESS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(208)
		p.Match(kslParserTYPE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(209)
		p.DynamicName()
	}
	{
		p.SetState(210)
		p.Match(kslParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(214)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&16781440) != 0 {
		{
			p.SetState(211)
			p.DynamicRelation()
		}

		p.SetState(216)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(217)
		p.Match(kslParserRBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDynamicRelationContext is an interface to support dynamic dispatch.
type IDynamicRelationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	RELATION() antlr.TerminalNode
	DynamicName() IDynamicNameContext
	EXPAND() antlr.TerminalNode
	DynamicBody() IDynamicBodyContext
	ALLOW_DUPLICATES() antlr.TerminalNode
	ACCESS() antlr.TerminalNode

	// IsDynamicRelationContext differentiates from other interfaces.
	IsDynamicRelationContext()
}

type DynamicRelationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDynamicRelationContext() *DynamicRelationContext {
	var p = new(DynamicRelationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kslParserRULE_dynamicRelation
	return p
}

func InitEmptyDynamicRelationContext(p *DynamicRelationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kslParserRULE_dynamicRelation
}

func (*DynamicRelationContext) IsDynamicRelationContext() {}

func NewDynamicRelationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DynamicRelationContext {
	var p = new(DynamicRelationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = kslParserRULE_dynamicRelation

	return p
}

func (s *DynamicRelationContext) GetParser() antlr.Parser { return s.parser }

func (s *DynamicRelationContext) RELATION() antlr.TerminalNode {
	return s.GetToken(kslParserRELATION, 0)
}

func (s *DynamicRelationContext) DynamicName() IDynamicNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDynamicNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDynamicNameContext)
}

func (s *DynamicRelationContext) EXPAND() antlr.TerminalNode {
	return s.GetToken(kslParserEXPAND, 0)
}

func (s *DynamicRelationContext) DynamicBody() IDynamicBodyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDynamicBodyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDynamicBodyContext)
}

func (s *DynamicRelationContext) ALLOW_DUPLICATES() antlr.TerminalNode {
	return s.GetToken(kslParserALLOW_DUPLICATES, 0)
}

func (s *DynamicRelationContext) ACCESS() antlr.TerminalNode {
	return s.GetToken(kslParserACCESS, 0)
}

func (s *DynamicRelationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DynamicRelationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DynamicRelationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kslListener); ok {
		listenerT.EnterDynamicRelation(s)
	}
}

func (s *DynamicRelationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kslListener); ok {
		listenerT.ExitDynamicRelation(s)
	}
}

func (s *DynamicRelationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case kslVisitor:
		return t.VisitDynamicRelation(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *kslParser) DynamicRelation() (localctx IDynamicRelationContext) {
	localctx = NewDynamicRelationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, kslParserRULE_dynamicRelation)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(220)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == kslParserALLOW_DUPLICATES {
		{
			p.SetState(219)
			p.Match(kslParserALLOW_DUPLICATES)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	p.SetState(223)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == kslParserACCESS {
		{
			p.SetState(222)
			p.Match(kslParserACCESS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(225)
		p.Match(kslParserRELATION)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(226)
		p.DynamicName()
	}
	{
		p.SetState(227)
		p.Match(kslParserEXPAND)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(228)
		p.dynamicBody(0)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDynamicNameContext is an interface to support dynamic dispatch.
type IDynamicNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsDynamicNameContext differentiates from other interfaces.
	IsDynamicNameContext()
}

type DynamicNameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDynamicNameContext() *DynamicNameContext {
	var p = new(DynamicNameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kslParserRULE_dynamicName
	return p
}

func InitEmptyDynamicNameContext(p *DynamicNameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kslParserRULE_dynamicName
}

func (*DynamicNameContext) IsDynamicNameContext() {}

func NewDynamicNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DynamicNameContext {
	var p = new(DynamicNameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = kslParserRULE_dynamicName

	return p
}

func (s *DynamicNameContext) GetParser() antlr.Parser { return s.parser }

func (s *DynamicNameContext) CopyAll(ctx *DynamicNameContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *DynamicNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DynamicNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type VariableContext struct {
	DynamicNameContext
}

func NewVariableContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VariableContext {
	var p = new(VariableContext)

	InitEmptyDynamicNameContext(&p.DynamicNameContext)
	p.parser = parser
	p.CopyAll(ctx.(*DynamicNameContext))

	return p
}

func (s *VariableContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableContext) VARREF() antlr.TerminalNode {
	return s.GetToken(kslParserVARREF, 0)
}

func (s *VariableContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(kslParserLBRACE, 0)
}

func (s *VariableContext) Name() INameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *VariableContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(kslParserRBRACE, 0)
}

func (s *VariableContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kslListener); ok {
		listenerT.EnterVariable(s)
	}
}

func (s *VariableContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kslListener); ok {
		listenerT.ExitVariable(s)
	}
}

func (s *VariableContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case kslVisitor:
		return t.VisitVariable(s)

	default:
		return t.VisitChildren(s)
	}
}

type LiteralContext struct {
	DynamicNameContext
}

func NewLiteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LiteralContext {
	var p = new(LiteralContext)

	InitEmptyDynamicNameContext(&p.DynamicNameContext)
	p.parser = parser
	p.CopyAll(ctx.(*DynamicNameContext))

	return p
}

func (s *LiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralContext) Name() INameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *LiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kslListener); ok {
		listenerT.EnterLiteral(s)
	}
}

func (s *LiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kslListener); ok {
		listenerT.ExitLiteral(s)
	}
}

func (s *LiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case kslVisitor:
		return t.VisitLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

type TemplateContext struct {
	DynamicNameContext
}

func NewTemplateContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TemplateContext {
	var p = new(TemplateContext)

	InitEmptyDynamicNameContext(&p.DynamicNameContext)
	p.parser = parser
	p.CopyAll(ctx.(*DynamicNameContext))

	return p
}

func (s *TemplateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TemplateContext) AllTEMPLATE_DELIM() []antlr.TerminalNode {
	return s.GetTokens(kslParserTEMPLATE_DELIM)
}

func (s *TemplateContext) TEMPLATE_DELIM(i int) antlr.TerminalNode {
	return s.GetToken(kslParserTEMPLATE_DELIM, i)
}

func (s *TemplateContext) AllDynamicName() []IDynamicNameContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDynamicNameContext); ok {
			len++
		}
	}

	tst := make([]IDynamicNameContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDynamicNameContext); ok {
			tst[i] = t.(IDynamicNameContext)
			i++
		}
	}

	return tst
}

func (s *TemplateContext) DynamicName(i int) IDynamicNameContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDynamicNameContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDynamicNameContext)
}

func (s *TemplateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kslListener); ok {
		listenerT.EnterTemplate(s)
	}
}

func (s *TemplateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kslListener); ok {
		listenerT.ExitTemplate(s)
	}
}

func (s *TemplateContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case kslVisitor:
		return t.VisitTemplate(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *kslParser) DynamicName() (localctx IDynamicNameContext) {
	localctx = NewDynamicNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, kslParserRULE_dynamicName)
	var _alt int

	p.SetState(244)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case kslParserFORCED_NAME, kslParserNAME:
		localctx = NewLiteralContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(230)
			p.Name()
		}

	case kslParserVARREF:
		localctx = NewVariableContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(231)
			p.Match(kslParserVARREF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(232)
			p.Match(kslParserLBRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(233)
			p.Name()
		}
		{
			p.SetState(234)
			p.Match(kslParserRBRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case kslParserTEMPLATE_DELIM:
		localctx = NewTemplateContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(236)
			p.Match(kslParserTEMPLATE_DELIM)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(238)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = 1
		for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			switch _alt {
			case 1:
				{
					p.SetState(237)
					p.DynamicName()
				}

			default:
				p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
				goto errorExit
			}

			p.SetState(240)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 26, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}
		{
			p.SetState(242)
			p.Match(kslParserTEMPLATE_DELIM)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDynamicBodyContext is an interface to support dynamic dispatch.
type IDynamicBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsDynamicBodyContext differentiates from other interfaces.
	IsDynamicBodyContext()
}

type DynamicBodyContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDynamicBodyContext() *DynamicBodyContext {
	var p = new(DynamicBodyContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kslParserRULE_dynamicBody
	return p
}

func InitEmptyDynamicBodyContext(p *DynamicBodyContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kslParserRULE_dynamicBody
}

func (*DynamicBodyContext) IsDynamicBodyContext() {}

func NewDynamicBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DynamicBodyContext {
	var p = new(DynamicBodyContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = kslParserRULE_dynamicBody

	return p
}

func (s *DynamicBodyContext) GetParser() antlr.Parser { return s.parser }

func (s *DynamicBodyContext) CopyAll(ctx *DynamicBodyContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *DynamicBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DynamicBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type DynamicUnlessContext struct {
	DynamicBodyContext
}

func NewDynamicUnlessContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DynamicUnlessContext {
	var p = new(DynamicUnlessContext)

	InitEmptyDynamicBodyContext(&p.DynamicBodyContext)
	p.parser = parser
	p.CopyAll(ctx.(*DynamicBodyContext))

	return p
}

func (s *DynamicUnlessContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DynamicUnlessContext) AllDynamicBody() []IDynamicBodyContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDynamicBodyContext); ok {
			len++
		}
	}

	tst := make([]IDynamicBodyContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDynamicBodyContext); ok {
			tst[i] = t.(IDynamicBodyContext)
			i++
		}
	}

	return tst
}

func (s *DynamicUnlessContext) DynamicBody(i int) IDynamicBodyContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDynamicBodyContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDynamicBodyContext)
}

func (s *DynamicUnlessContext) UNLESS() antlr.TerminalNode {
	return s.GetToken(kslParserUNLESS, 0)
}

func (s *DynamicUnlessContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kslListener); ok {
		listenerT.EnterDynamicUnless(s)
	}
}

func (s *DynamicUnlessContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kslListener); ok {
		listenerT.ExitDynamicUnless(s)
	}
}

func (s *DynamicUnlessContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case kslVisitor:
		return t.VisitDynamicUnless(s)

	default:
		return t.VisitChildren(s)
	}
}

type DynamicORContext struct {
	DynamicBodyContext
}

func NewDynamicORContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DynamicORContext {
	var p = new(DynamicORContext)

	InitEmptyDynamicBodyContext(&p.DynamicBodyContext)
	p.parser = parser
	p.CopyAll(ctx.(*DynamicBodyContext))

	return p
}

func (s *DynamicORContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DynamicORContext) AllDynamicBody() []IDynamicBodyContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDynamicBodyContext); ok {
			len++
		}
	}

	tst := make([]IDynamicBodyContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDynamicBodyContext); ok {
			tst[i] = t.(IDynamicBodyContext)
			i++
		}
	}

	return tst
}

func (s *DynamicORContext) DynamicBody(i int) IDynamicBodyContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDynamicBodyContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDynamicBodyContext)
}

func (s *DynamicORContext) OR() antlr.TerminalNode {
	return s.GetToken(kslParserOR, 0)
}

func (s *DynamicORContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kslListener); ok {
		listenerT.EnterDynamicOR(s)
	}
}

func (s *DynamicORContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kslListener); ok {
		listenerT.ExitDynamicOR(s)
	}
}

func (s *DynamicORContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case kslVisitor:
		return t.VisitDynamicOR(s)

	default:
		return t.VisitChildren(s)
	}
}

type DynamicParenContext struct {
	DynamicBodyContext
}

func NewDynamicParenContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DynamicParenContext {
	var p = new(DynamicParenContext)

	InitEmptyDynamicBodyContext(&p.DynamicBodyContext)
	p.parser = parser
	p.CopyAll(ctx.(*DynamicBodyContext))

	return p
}

func (s *DynamicParenContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DynamicParenContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(kslParserLPAREN, 0)
}

func (s *DynamicParenContext) DynamicBody() IDynamicBodyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDynamicBodyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDynamicBodyContext)
}

func (s *DynamicParenContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(kslParserRPAREN, 0)
}

func (s *DynamicParenContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kslListener); ok {
		listenerT.EnterDynamicParen(s)
	}
}

func (s *DynamicParenContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kslListener); ok {
		listenerT.ExitDynamicParen(s)
	}
}

func (s *DynamicParenContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case kslVisitor:
		return t.VisitDynamicParen(s)

	default:
		return t.VisitChildren(s)
	}
}

type DynamicSelfContext struct {
	DynamicBodyContext
}

func NewDynamicSelfContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DynamicSelfContext {
	var p = new(DynamicSelfContext)

	InitEmptyDynamicBodyContext(&p.DynamicBodyContext)
	p.parser = parser
	p.CopyAll(ctx.(*DynamicBodyContext))

	return p
}

func (s *DynamicSelfContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DynamicSelfContext) LSQUARE() antlr.TerminalNode {
	return s.GetToken(kslParserLSQUARE, 0)
}

func (s *DynamicSelfContext) AllTypeReference() []ITypeReferenceContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITypeReferenceContext); ok {
			len++
		}
	}

	tst := make([]ITypeReferenceContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITypeReferenceContext); ok {
			tst[i] = t.(ITypeReferenceContext)
			i++
		}
	}

	return tst
}

func (s *DynamicSelfContext) TypeReference(i int) ITypeReferenceContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeReferenceContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeReferenceContext)
}

func (s *DynamicSelfContext) RSQARE() antlr.TerminalNode {
	return s.GetToken(kslParserRSQARE, 0)
}

func (s *DynamicSelfContext) CARDINALITY() antlr.TerminalNode {
	return s.GetToken(kslParserCARDINALITY, 0)
}

func (s *DynamicSelfContext) AllOR() []antlr.TerminalNode {
	return s.GetTokens(kslParserOR)
}

func (s *DynamicSelfContext) OR(i int) antlr.TerminalNode {
	return s.GetToken(kslParserOR, i)
}

func (s *DynamicSelfContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kslListener); ok {
		listenerT.EnterDynamicSelf(s)
	}
}

func (s *DynamicSelfContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kslListener); ok {
		listenerT.ExitDynamicSelf(s)
	}
}

func (s *DynamicSelfContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case kslVisitor:
		return t.VisitDynamicSelf(s)

	default:
		return t.VisitChildren(s)
	}
}

type DynamicAndContext struct {
	DynamicBodyContext
}

func NewDynamicAndContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DynamicAndContext {
	var p = new(DynamicAndContext)

	InitEmptyDynamicBodyContext(&p.DynamicBodyContext)
	p.parser = parser
	p.CopyAll(ctx.(*DynamicBodyContext))

	return p
}

func (s *DynamicAndContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DynamicAndContext) AllDynamicBody() []IDynamicBodyContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDynamicBodyContext); ok {
			len++
		}
	}

	tst := make([]IDynamicBodyContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDynamicBodyContext); ok {
			tst[i] = t.(IDynamicBodyContext)
			i++
		}
	}

	return tst
}

func (s *DynamicAndContext) DynamicBody(i int) IDynamicBodyContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDynamicBodyContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDynamicBodyContext)
}

func (s *DynamicAndContext) AND() antlr.TerminalNode {
	return s.GetToken(kslParserAND, 0)
}

func (s *DynamicAndContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kslListener); ok {
		listenerT.EnterDynamicAnd(s)
	}
}

func (s *DynamicAndContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kslListener); ok {
		listenerT.ExitDynamicAnd(s)
	}
}

func (s *DynamicAndContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case kslVisitor:
		return t.VisitDynamicAnd(s)

	default:
		return t.VisitChildren(s)
	}
}

type DynamicReferenceContext struct {
	DynamicBodyContext
}

func NewDynamicReferenceContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DynamicReferenceContext {
	var p = new(DynamicReferenceContext)

	InitEmptyDynamicBodyContext(&p.DynamicBodyContext)
	p.parser = parser
	p.CopyAll(ctx.(*DynamicBodyContext))

	return p
}

func (s *DynamicReferenceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DynamicReferenceContext) DynamicName() IDynamicNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDynamicNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDynamicNameContext)
}

func (s *DynamicReferenceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kslListener); ok {
		listenerT.EnterDynamicReference(s)
	}
}

func (s *DynamicReferenceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kslListener); ok {
		listenerT.ExitDynamicReference(s)
	}
}

func (s *DynamicReferenceContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case kslVisitor:
		return t.VisitDynamicReference(s)

	default:
		return t.VisitChildren(s)
	}
}

type DynamicSubRelationContext struct {
	DynamicBodyContext
}

func NewDynamicSubRelationContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DynamicSubRelationContext {
	var p = new(DynamicSubRelationContext)

	InitEmptyDynamicBodyContext(&p.DynamicBodyContext)
	p.parser = parser
	p.CopyAll(ctx.(*DynamicBodyContext))

	return p
}

func (s *DynamicSubRelationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DynamicSubRelationContext) AllDynamicName() []IDynamicNameContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDynamicNameContext); ok {
			len++
		}
	}

	tst := make([]IDynamicNameContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDynamicNameContext); ok {
			tst[i] = t.(IDynamicNameContext)
			i++
		}
	}

	return tst
}

func (s *DynamicSubRelationContext) DynamicName(i int) IDynamicNameContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDynamicNameContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDynamicNameContext)
}

func (s *DynamicSubRelationContext) RESOLVE() antlr.TerminalNode {
	return s.GetToken(kslParserRESOLVE, 0)
}

func (s *DynamicSubRelationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kslListener); ok {
		listenerT.EnterDynamicSubRelation(s)
	}
}

func (s *DynamicSubRelationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kslListener); ok {
		listenerT.ExitDynamicSubRelation(s)
	}
}

func (s *DynamicSubRelationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case kslVisitor:
		return t.VisitDynamicSubRelation(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *kslParser) DynamicBody() (localctx IDynamicBodyContext) {
	return p.dynamicBody(0)
}

func (p *kslParser) dynamicBody(_p int) (localctx IDynamicBodyContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewDynamicBodyContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IDynamicBodyContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 36
	p.EnterRecursionRule(localctx, 36, kslParserRULE_dynamicBody, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(270)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 30, p.GetParserRuleContext()) {
	case 1:
		localctx = NewDynamicSelfContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(247)
			p.Match(kslParserLSQUARE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(249)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == kslParserCARDINALITY {
			{
				p.SetState(248)
				p.Match(kslParserCARDINALITY)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(251)
			p.TypeReference()
		}
		p.SetState(256)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == kslParserOR {
			{
				p.SetState(252)
				p.Match(kslParserOR)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(253)
				p.TypeReference()
			}

			p.SetState(258)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(259)
			p.Match(kslParserRSQARE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewDynamicReferenceContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(261)
			p.DynamicName()
		}

	case 3:
		localctx = NewDynamicSubRelationContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(262)
			p.DynamicName()
		}
		{
			p.SetState(263)
			p.Match(kslParserRESOLVE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(264)
			p.DynamicName()
		}

	case 4:
		localctx = NewDynamicParenContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(266)
			p.Match(kslParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(267)
			p.dynamicBody(0)
		}
		{
			p.SetState(268)
			p.Match(kslParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(283)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 32, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(281)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 31, p.GetParserRuleContext()) {
			case 1:
				localctx = NewDynamicAndContext(p, NewDynamicBodyContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, kslParserRULE_dynamicBody)
				p.SetState(272)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
					goto errorExit
				}
				{
					p.SetState(273)
					p.Match(kslParserAND)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(274)
					p.dynamicBody(4)
				}

			case 2:
				localctx = NewDynamicORContext(p, NewDynamicBodyContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, kslParserRULE_dynamicBody)
				p.SetState(275)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
					goto errorExit
				}
				{
					p.SetState(276)
					p.Match(kslParserOR)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(277)
					p.dynamicBody(3)
				}

			case 3:
				localctx = NewDynamicUnlessContext(p, NewDynamicBodyContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, kslParserRULE_dynamicBody)
				p.SetState(278)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
					goto errorExit
				}
				{
					p.SetState(279)
					p.Match(kslParserUNLESS)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(280)
					p.dynamicBody(2)
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(285)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 32, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

func (p *kslParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 12:
		var t *RelationBodyContext = nil
		if localctx != nil {
			t = localctx.(*RelationBodyContext)
		}
		return p.RelationBody_Sempred(t, predIndex)

	case 18:
		var t *DynamicBodyContext = nil
		if localctx != nil {
			t = localctx.(*DynamicBodyContext)
		}
		return p.DynamicBody_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *kslParser) RelationBody_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 2)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *kslParser) DynamicBody_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 3:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 2)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
