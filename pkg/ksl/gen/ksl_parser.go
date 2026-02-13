// Code generated from /Users/mmclaugh/go/src/github.com/project-kessel/ksl-schema-language/pkg/ksl/ksl.g4 by ANTLR 4.13.1. DO NOT EDIT.

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
		"'internal'", "'private'", "'type'", "'extend'", "'relation'", "'import'",
		"'extension'", "", "'AtMostOne'", "'ExactlyOne'", "'AtLeastOne'", "'Any'",
		"'as'", "'and'", "'or'", "'unless'", "'allow_duplicates'", "':'", "'{'",
		"'}'", "'@'", "'with'", "'('", "')'", "'['", "']'", "'$'", "'`'", "'''",
		"','", "';'",
	}
	staticData.SymbolicNames = []string{
		"", "KW_ESC", "FORCED_NAME", "VERSION", "VERSIONNUM", "RESOLVE", "NAMESPACE",
		"ACCESS", "PUBLIC", "INTERNAL", "PRIVATE", "TYPE", "EXTEND", "RELATION",
		"IMPORT", "EXTENSION", "CARDINALITY", "ATMOSTONE", "EXACTLYONE", "ATLEASTONE",
		"ANY", "AS", "AND", "OR", "UNLESS", "ALLOW_DUPLICATES", "EXPAND", "LBRACE",
		"RBRACE", "EXTENSION_CALL", "IN_REL_EXTENSION_CALL", "LPAREN", "RPAREN",
		"LSQUARE", "RSQARE", "VARREF", "TEMPLATE_DELIM", "STRING_DELIM", "ARG_DELIM",
		"DECL_END", "NAME", "COMMENT", "WS",
	}
	staticData.RuleNames = []string{
		"name", "file", "version", "namespace", "import_stmt", "declaration",
		"typeExpr", "typeReference", "extendTypeExpr", "extensionParam", "extensionParams",
		"extensionReference", "relation", "relationBody", "paramNames", "extension",
		"dynamicType", "dynamicRelation", "dynamicName", "dynamicBody",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 42, 325, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 1, 0, 1, 0, 1,
		1, 1, 1, 1, 1, 5, 1, 46, 8, 1, 10, 1, 12, 1, 49, 9, 1, 1, 1, 4, 1, 52,
		8, 1, 11, 1, 12, 1, 53, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4,
		1, 4, 1, 5, 1, 5, 3, 5, 67, 8, 5, 1, 5, 1, 5, 3, 5, 71, 8, 5, 1, 5, 1,
		5, 3, 5, 75, 8, 5, 1, 5, 1, 5, 1, 5, 3, 5, 80, 8, 5, 1, 6, 5, 6, 83, 8,
		6, 10, 6, 12, 6, 86, 9, 6, 1, 6, 3, 6, 89, 8, 6, 1, 6, 1, 6, 1, 6, 1, 6,
		5, 6, 95, 8, 6, 10, 6, 12, 6, 98, 9, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 5,
		7, 105, 8, 7, 10, 7, 12, 7, 108, 9, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 5,
		8, 115, 8, 8, 10, 8, 12, 8, 118, 9, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1,
		9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 5, 10, 131, 8, 10, 10, 10, 12, 10,
		134, 9, 10, 1, 11, 1, 11, 1, 11, 1, 11, 3, 11, 140, 8, 11, 1, 11, 1, 11,
		1, 12, 5, 12, 145, 8, 12, 10, 12, 12, 12, 148, 9, 12, 1, 12, 3, 12, 151,
		8, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 3, 13, 161,
		8, 13, 1, 13, 1, 13, 1, 13, 5, 13, 166, 8, 13, 10, 13, 12, 13, 169, 9,
		13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 3, 13, 179,
		8, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1,
		13, 1, 13, 1, 13, 1, 13, 3, 13, 194, 8, 13, 1, 13, 1, 13, 1, 13, 1, 13,
		1, 13, 1, 13, 3, 13, 202, 8, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1,
		13, 1, 13, 1, 13, 1, 13, 5, 13, 213, 8, 13, 10, 13, 12, 13, 216, 9, 13,
		1, 14, 1, 14, 1, 14, 5, 14, 221, 8, 14, 10, 14, 12, 14, 224, 9, 14, 1,
		15, 3, 15, 227, 8, 15, 1, 15, 1, 15, 1, 15, 1, 15, 3, 15, 233, 8, 15, 1,
		15, 1, 15, 1, 15, 4, 15, 238, 8, 15, 11, 15, 12, 15, 239, 1, 15, 1, 15,
		1, 16, 3, 16, 245, 8, 16, 1, 16, 1, 16, 1, 16, 1, 16, 5, 16, 251, 8, 16,
		10, 16, 12, 16, 254, 9, 16, 1, 16, 1, 16, 1, 17, 3, 17, 259, 8, 17, 1,
		17, 3, 17, 262, 8, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 18, 1, 18,
		1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 4, 18, 277, 8, 18, 11, 18, 12,
		18, 278, 1, 18, 1, 18, 3, 18, 283, 8, 18, 1, 19, 1, 19, 1, 19, 3, 19, 288,
		8, 19, 1, 19, 1, 19, 1, 19, 5, 19, 293, 8, 19, 10, 19, 12, 19, 296, 9,
		19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19,
		1, 19, 3, 19, 309, 8, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1,
		19, 1, 19, 1, 19, 5, 19, 320, 8, 19, 10, 19, 12, 19, 323, 9, 19, 1, 19,
		0, 2, 26, 38, 20, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28,
		30, 32, 34, 36, 38, 0, 2, 2, 0, 2, 2, 40, 40, 1, 0, 37, 37, 352, 0, 40,
		1, 0, 0, 0, 2, 42, 1, 0, 0, 0, 4, 55, 1, 0, 0, 0, 6, 58, 1, 0, 0, 0, 8,
		61, 1, 0, 0, 0, 10, 79, 1, 0, 0, 0, 12, 84, 1, 0, 0, 0, 14, 101, 1, 0,
		0, 0, 16, 109, 1, 0, 0, 0, 18, 121, 1, 0, 0, 0, 20, 127, 1, 0, 0, 0, 22,
		135, 1, 0, 0, 0, 24, 146, 1, 0, 0, 0, 26, 201, 1, 0, 0, 0, 28, 217, 1,
		0, 0, 0, 30, 226, 1, 0, 0, 0, 32, 244, 1, 0, 0, 0, 34, 258, 1, 0, 0, 0,
		36, 282, 1, 0, 0, 0, 38, 308, 1, 0, 0, 0, 40, 41, 7, 0, 0, 0, 41, 1, 1,
		0, 0, 0, 42, 43, 3, 4, 2, 0, 43, 47, 3, 6, 3, 0, 44, 46, 3, 8, 4, 0, 45,
		44, 1, 0, 0, 0, 46, 49, 1, 0, 0, 0, 47, 45, 1, 0, 0, 0, 47, 48, 1, 0, 0,
		0, 48, 51, 1, 0, 0, 0, 49, 47, 1, 0, 0, 0, 50, 52, 3, 10, 5, 0, 51, 50,
		1, 0, 0, 0, 52, 53, 1, 0, 0, 0, 53, 51, 1, 0, 0, 0, 53, 54, 1, 0, 0, 0,
		54, 3, 1, 0, 0, 0, 55, 56, 5, 3, 0, 0, 56, 57, 5, 4, 0, 0, 57, 5, 1, 0,
		0, 0, 58, 59, 5, 6, 0, 0, 59, 60, 3, 0, 0, 0, 60, 7, 1, 0, 0, 0, 61, 62,
		5, 14, 0, 0, 62, 63, 3, 0, 0, 0, 63, 9, 1, 0, 0, 0, 64, 66, 3, 12, 6, 0,
		65, 67, 5, 39, 0, 0, 66, 65, 1, 0, 0, 0, 66, 67, 1, 0, 0, 0, 67, 80, 1,
		0, 0, 0, 68, 70, 3, 16, 8, 0, 69, 71, 5, 39, 0, 0, 70, 69, 1, 0, 0, 0,
		70, 71, 1, 0, 0, 0, 71, 80, 1, 0, 0, 0, 72, 74, 3, 30, 15, 0, 73, 75, 5,
		39, 0, 0, 74, 73, 1, 0, 0, 0, 74, 75, 1, 0, 0, 0, 75, 80, 1, 0, 0, 0, 76,
		77, 3, 22, 11, 0, 77, 78, 5, 39, 0, 0, 78, 80, 1, 0, 0, 0, 79, 64, 1, 0,
		0, 0, 79, 68, 1, 0, 0, 0, 79, 72, 1, 0, 0, 0, 79, 76, 1, 0, 0, 0, 80, 11,
		1, 0, 0, 0, 81, 83, 3, 22, 11, 0, 82, 81, 1, 0, 0, 0, 83, 86, 1, 0, 0,
		0, 84, 82, 1, 0, 0, 0, 84, 85, 1, 0, 0, 0, 85, 88, 1, 0, 0, 0, 86, 84,
		1, 0, 0, 0, 87, 89, 5, 7, 0, 0, 88, 87, 1, 0, 0, 0, 88, 89, 1, 0, 0, 0,
		89, 90, 1, 0, 0, 0, 90, 91, 5, 11, 0, 0, 91, 92, 3, 0, 0, 0, 92, 96, 5,
		27, 0, 0, 93, 95, 3, 24, 12, 0, 94, 93, 1, 0, 0, 0, 95, 98, 1, 0, 0, 0,
		96, 94, 1, 0, 0, 0, 96, 97, 1, 0, 0, 0, 97, 99, 1, 0, 0, 0, 98, 96, 1,
		0, 0, 0, 99, 100, 5, 28, 0, 0, 100, 13, 1, 0, 0, 0, 101, 106, 3, 0, 0,
		0, 102, 103, 5, 5, 0, 0, 103, 105, 3, 0, 0, 0, 104, 102, 1, 0, 0, 0, 105,
		108, 1, 0, 0, 0, 106, 104, 1, 0, 0, 0, 106, 107, 1, 0, 0, 0, 107, 15, 1,
		0, 0, 0, 108, 106, 1, 0, 0, 0, 109, 110, 5, 12, 0, 0, 110, 111, 5, 11,
		0, 0, 111, 112, 3, 0, 0, 0, 112, 116, 5, 27, 0, 0, 113, 115, 3, 24, 12,
		0, 114, 113, 1, 0, 0, 0, 115, 118, 1, 0, 0, 0, 116, 114, 1, 0, 0, 0, 116,
		117, 1, 0, 0, 0, 117, 119, 1, 0, 0, 0, 118, 116, 1, 0, 0, 0, 119, 120,
		5, 28, 0, 0, 120, 17, 1, 0, 0, 0, 121, 122, 3, 0, 0, 0, 122, 123, 5, 26,
		0, 0, 123, 124, 5, 37, 0, 0, 124, 125, 8, 1, 0, 0, 125, 126, 5, 37, 0,
		0, 126, 19, 1, 0, 0, 0, 127, 132, 3, 18, 9, 0, 128, 129, 5, 38, 0, 0, 129,
		131, 3, 18, 9, 0, 130, 128, 1, 0, 0, 0, 131, 134, 1, 0, 0, 0, 132, 130,
		1, 0, 0, 0, 132, 133, 1, 0, 0, 0, 133, 21, 1, 0, 0, 0, 134, 132, 1, 0,
		0, 0, 135, 136, 5, 29, 0, 0, 136, 137, 3, 14, 7, 0, 137, 139, 5, 31, 0,
		0, 138, 140, 3, 20, 10, 0, 139, 138, 1, 0, 0, 0, 139, 140, 1, 0, 0, 0,
		140, 141, 1, 0, 0, 0, 141, 142, 5, 32, 0, 0, 142, 23, 1, 0, 0, 0, 143,
		145, 3, 22, 11, 0, 144, 143, 1, 0, 0, 0, 145, 148, 1, 0, 0, 0, 146, 144,
		1, 0, 0, 0, 146, 147, 1, 0, 0, 0, 147, 150, 1, 0, 0, 0, 148, 146, 1, 0,
		0, 0, 149, 151, 5, 7, 0, 0, 150, 149, 1, 0, 0, 0, 150, 151, 1, 0, 0, 0,
		151, 152, 1, 0, 0, 0, 152, 153, 5, 13, 0, 0, 153, 154, 3, 0, 0, 0, 154,
		155, 5, 26, 0, 0, 155, 156, 3, 26, 13, 0, 156, 25, 1, 0, 0, 0, 157, 158,
		6, 13, -1, 0, 158, 160, 5, 33, 0, 0, 159, 161, 5, 16, 0, 0, 160, 159, 1,
		0, 0, 0, 160, 161, 1, 0, 0, 0, 161, 162, 1, 0, 0, 0, 162, 167, 3, 14, 7,
		0, 163, 164, 5, 23, 0, 0, 164, 166, 3, 14, 7, 0, 165, 163, 1, 0, 0, 0,
		166, 169, 1, 0, 0, 0, 167, 165, 1, 0, 0, 0, 167, 168, 1, 0, 0, 0, 168,
		170, 1, 0, 0, 0, 169, 167, 1, 0, 0, 0, 170, 171, 5, 34, 0, 0, 171, 202,
		1, 0, 0, 0, 172, 202, 3, 0, 0, 0, 173, 174, 3, 0, 0, 0, 174, 175, 5, 30,
		0, 0, 175, 176, 3, 14, 7, 0, 176, 178, 5, 31, 0, 0, 177, 179, 3, 20, 10,
		0, 178, 177, 1, 0, 0, 0, 178, 179, 1, 0, 0, 0, 179, 180, 1, 0, 0, 0, 180,
		181, 5, 32, 0, 0, 181, 202, 1, 0, 0, 0, 182, 183, 3, 0, 0, 0, 183, 184,
		5, 5, 0, 0, 184, 185, 3, 0, 0, 0, 185, 202, 1, 0, 0, 0, 186, 187, 3, 0,
		0, 0, 187, 188, 5, 5, 0, 0, 188, 189, 3, 0, 0, 0, 189, 190, 5, 30, 0, 0,
		190, 191, 3, 14, 7, 0, 191, 193, 5, 31, 0, 0, 192, 194, 3, 20, 10, 0, 193,
		192, 1, 0, 0, 0, 193, 194, 1, 0, 0, 0, 194, 195, 1, 0, 0, 0, 195, 196,
		5, 32, 0, 0, 196, 202, 1, 0, 0, 0, 197, 198, 5, 31, 0, 0, 198, 199, 3,
		26, 13, 0, 199, 200, 5, 32, 0, 0, 200, 202, 1, 0, 0, 0, 201, 157, 1, 0,
		0, 0, 201, 172, 1, 0, 0, 0, 201, 173, 1, 0, 0, 0, 201, 182, 1, 0, 0, 0,
		201, 186, 1, 0, 0, 0, 201, 197, 1, 0, 0, 0, 202, 214, 1, 0, 0, 0, 203,
		204, 10, 3, 0, 0, 204, 205, 5, 22, 0, 0, 205, 213, 3, 26, 13, 4, 206, 207,
		10, 2, 0, 0, 207, 208, 5, 23, 0, 0, 208, 213, 3, 26, 13, 3, 209, 210, 10,
		1, 0, 0, 210, 211, 5, 24, 0, 0, 211, 213, 3, 26, 13, 2, 212, 203, 1, 0,
		0, 0, 212, 206, 1, 0, 0, 0, 212, 209, 1, 0, 0, 0, 213, 216, 1, 0, 0, 0,
		214, 212, 1, 0, 0, 0, 214, 215, 1, 0, 0, 0, 215, 27, 1, 0, 0, 0, 216, 214,
		1, 0, 0, 0, 217, 222, 3, 0, 0, 0, 218, 219, 5, 38, 0, 0, 219, 221, 3, 0,
		0, 0, 220, 218, 1, 0, 0, 0, 221, 224, 1, 0, 0, 0, 222, 220, 1, 0, 0, 0,
		222, 223, 1, 0, 0, 0, 223, 29, 1, 0, 0, 0, 224, 222, 1, 0, 0, 0, 225, 227,
		5, 7, 0, 0, 226, 225, 1, 0, 0, 0, 226, 227, 1, 0, 0, 0, 227, 228, 1, 0,
		0, 0, 228, 229, 5, 15, 0, 0, 229, 230, 3, 0, 0, 0, 230, 232, 5, 31, 0,
		0, 231, 233, 3, 28, 14, 0, 232, 231, 1, 0, 0, 0, 232, 233, 1, 0, 0, 0,
		233, 234, 1, 0, 0, 0, 234, 235, 5, 32, 0, 0, 235, 237, 5, 27, 0, 0, 236,
		238, 3, 32, 16, 0, 237, 236, 1, 0, 0, 0, 238, 239, 1, 0, 0, 0, 239, 237,
		1, 0, 0, 0, 239, 240, 1, 0, 0, 0, 240, 241, 1, 0, 0, 0, 241, 242, 5, 28,
		0, 0, 242, 31, 1, 0, 0, 0, 243, 245, 5, 7, 0, 0, 244, 243, 1, 0, 0, 0,
		244, 245, 1, 0, 0, 0, 245, 246, 1, 0, 0, 0, 246, 247, 5, 11, 0, 0, 247,
		248, 3, 36, 18, 0, 248, 252, 5, 27, 0, 0, 249, 251, 3, 34, 17, 0, 250,
		249, 1, 0, 0, 0, 251, 254, 1, 0, 0, 0, 252, 250, 1, 0, 0, 0, 252, 253,
		1, 0, 0, 0, 253, 255, 1, 0, 0, 0, 254, 252, 1, 0, 0, 0, 255, 256, 5, 28,
		0, 0, 256, 33, 1, 0, 0, 0, 257, 259, 5, 25, 0, 0, 258, 257, 1, 0, 0, 0,
		258, 259, 1, 0, 0, 0, 259, 261, 1, 0, 0, 0, 260, 262, 5, 7, 0, 0, 261,
		260, 1, 0, 0, 0, 261, 262, 1, 0, 0, 0, 262, 263, 1, 0, 0, 0, 263, 264,
		5, 13, 0, 0, 264, 265, 3, 36, 18, 0, 265, 266, 5, 26, 0, 0, 266, 267, 3,
		38, 19, 0, 267, 35, 1, 0, 0, 0, 268, 283, 3, 0, 0, 0, 269, 270, 5, 35,
		0, 0, 270, 271, 5, 27, 0, 0, 271, 272, 3, 0, 0, 0, 272, 273, 5, 28, 0,
		0, 273, 283, 1, 0, 0, 0, 274, 276, 5, 36, 0, 0, 275, 277, 3, 36, 18, 0,
		276, 275, 1, 0, 0, 0, 277, 278, 1, 0, 0, 0, 278, 276, 1, 0, 0, 0, 278,
		279, 1, 0, 0, 0, 279, 280, 1, 0, 0, 0, 280, 281, 5, 36, 0, 0, 281, 283,
		1, 0, 0, 0, 282, 268, 1, 0, 0, 0, 282, 269, 1, 0, 0, 0, 282, 274, 1, 0,
		0, 0, 283, 37, 1, 0, 0, 0, 284, 285, 6, 19, -1, 0, 285, 287, 5, 33, 0,
		0, 286, 288, 5, 16, 0, 0, 287, 286, 1, 0, 0, 0, 287, 288, 1, 0, 0, 0, 288,
		289, 1, 0, 0, 0, 289, 294, 3, 14, 7, 0, 290, 291, 5, 23, 0, 0, 291, 293,
		3, 14, 7, 0, 292, 290, 1, 0, 0, 0, 293, 296, 1, 0, 0, 0, 294, 292, 1, 0,
		0, 0, 294, 295, 1, 0, 0, 0, 295, 297, 1, 0, 0, 0, 296, 294, 1, 0, 0, 0,
		297, 298, 5, 34, 0, 0, 298, 309, 1, 0, 0, 0, 299, 309, 3, 36, 18, 0, 300,
		301, 3, 36, 18, 0, 301, 302, 5, 5, 0, 0, 302, 303, 3, 36, 18, 0, 303, 309,
		1, 0, 0, 0, 304, 305, 5, 31, 0, 0, 305, 306, 3, 38, 19, 0, 306, 307, 5,
		32, 0, 0, 307, 309, 1, 0, 0, 0, 308, 284, 1, 0, 0, 0, 308, 299, 1, 0, 0,
		0, 308, 300, 1, 0, 0, 0, 308, 304, 1, 0, 0, 0, 309, 321, 1, 0, 0, 0, 310,
		311, 10, 3, 0, 0, 311, 312, 5, 22, 0, 0, 312, 320, 3, 38, 19, 4, 313, 314,
		10, 2, 0, 0, 314, 315, 5, 23, 0, 0, 315, 320, 3, 38, 19, 3, 316, 317, 10,
		1, 0, 0, 317, 318, 5, 24, 0, 0, 318, 320, 3, 38, 19, 2, 319, 310, 1, 0,
		0, 0, 319, 313, 1, 0, 0, 0, 319, 316, 1, 0, 0, 0, 320, 323, 1, 0, 0, 0,
		321, 319, 1, 0, 0, 0, 321, 322, 1, 0, 0, 0, 322, 39, 1, 0, 0, 0, 323, 321,
		1, 0, 0, 0, 37, 47, 53, 66, 70, 74, 79, 84, 88, 96, 106, 116, 132, 139,
		146, 150, 160, 167, 178, 193, 201, 212, 214, 222, 226, 232, 239, 244, 252,
		258, 261, 278, 282, 287, 294, 308, 319, 321,
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
	kslParserEOF                   = antlr.TokenEOF
	kslParserKW_ESC                = 1
	kslParserFORCED_NAME           = 2
	kslParserVERSION               = 3
	kslParserVERSIONNUM            = 4
	kslParserRESOLVE               = 5
	kslParserNAMESPACE             = 6
	kslParserACCESS                = 7
	kslParserPUBLIC                = 8
	kslParserINTERNAL              = 9
	kslParserPRIVATE               = 10
	kslParserTYPE                  = 11
	kslParserEXTEND                = 12
	kslParserRELATION              = 13
	kslParserIMPORT                = 14
	kslParserEXTENSION             = 15
	kslParserCARDINALITY           = 16
	kslParserATMOSTONE             = 17
	kslParserEXACTLYONE            = 18
	kslParserATLEASTONE            = 19
	kslParserANY                   = 20
	kslParserAS                    = 21
	kslParserAND                   = 22
	kslParserOR                    = 23
	kslParserUNLESS                = 24
	kslParserALLOW_DUPLICATES      = 25
	kslParserEXPAND                = 26
	kslParserLBRACE                = 27
	kslParserRBRACE                = 28
	kslParserEXTENSION_CALL        = 29
	kslParserIN_REL_EXTENSION_CALL = 30
	kslParserLPAREN                = 31
	kslParserRPAREN                = 32
	kslParserLSQUARE               = 33
	kslParserRSQARE                = 34
	kslParserVARREF                = 35
	kslParserTEMPLATE_DELIM        = 36
	kslParserSTRING_DELIM          = 37
	kslParserARG_DELIM             = 38
	kslParserDECL_END              = 39
	kslParserNAME                  = 40
	kslParserCOMMENT               = 41
	kslParserWS                    = 42
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
	kslParserRULE_extendTypeExpr     = 8
	kslParserRULE_extensionParam     = 9
	kslParserRULE_extensionParams    = 10
	kslParserRULE_extensionReference = 11
	kslParserRULE_relation           = 12
	kslParserRULE_relationBody       = 13
	kslParserRULE_paramNames         = 14
	kslParserRULE_extension          = 15
	kslParserRULE_dynamicType        = 16
	kslParserRULE_dynamicRelation    = 17
	kslParserRULE_dynamicName        = 18
	kslParserRULE_dynamicBody        = 19
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
		p.SetState(40)
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
		p.SetState(42)
		p.Version()
	}
	{
		p.SetState(43)
		p.Namespace()
	}
	p.SetState(47)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == kslParserIMPORT {
		{
			p.SetState(44)
			p.Import_stmt()
		}

		p.SetState(49)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(51)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&536909952) != 0) {
		{
			p.SetState(50)
			p.Declaration()
		}

		p.SetState(53)
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
		p.SetState(55)
		p.Match(kslParserVERSION)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(56)
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
		p.SetState(58)
		p.Match(kslParserNAMESPACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(59)
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
		p.SetState(61)
		p.Match(kslParserIMPORT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(62)
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
	ExtendTypeExpr() IExtendTypeExprContext
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

func (s *DeclarationContext) ExtendTypeExpr() IExtendTypeExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExtendTypeExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExtendTypeExprContext)
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

	p.SetState(79)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(64)
			p.TypeExpr()
		}
		p.SetState(66)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == kslParserDECL_END {
			{
				p.SetState(65)
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
			p.SetState(68)
			p.ExtendTypeExpr()
		}
		p.SetState(70)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == kslParserDECL_END {
			{
				p.SetState(69)
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
			p.SetState(72)
			p.Extension()
		}
		p.SetState(74)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == kslParserDECL_END {
			{
				p.SetState(73)
				p.Match(kslParserDECL_END)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(76)
			p.ExtensionReference()
		}
		{
			p.SetState(77)
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
	p.SetState(84)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == kslParserEXTENSION_CALL {
		{
			p.SetState(81)
			p.ExtensionReference()
		}

		p.SetState(86)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(88)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == kslParserACCESS {
		{
			p.SetState(87)
			p.Match(kslParserACCESS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(90)
		p.Match(kslParserTYPE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(91)
		p.Name()
	}
	{
		p.SetState(92)
		p.Match(kslParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(96)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&536879232) != 0 {
		{
			p.SetState(93)
			p.Relation()
		}

		p.SetState(98)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(99)
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
		p.SetState(101)
		p.Name()
	}
	p.SetState(106)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == kslParserRESOLVE {
		{
			p.SetState(102)
			p.Match(kslParserRESOLVE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(103)
			p.Name()
		}

		p.SetState(108)
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

// IExtendTypeExprContext is an interface to support dynamic dispatch.
type IExtendTypeExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EXTEND() antlr.TerminalNode
	TYPE() antlr.TerminalNode
	Name() INameContext
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	AllRelation() []IRelationContext
	Relation(i int) IRelationContext

	// IsExtendTypeExprContext differentiates from other interfaces.
	IsExtendTypeExprContext()
}

type ExtendTypeExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExtendTypeExprContext() *ExtendTypeExprContext {
	var p = new(ExtendTypeExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kslParserRULE_extendTypeExpr
	return p
}

func InitEmptyExtendTypeExprContext(p *ExtendTypeExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kslParserRULE_extendTypeExpr
}

func (*ExtendTypeExprContext) IsExtendTypeExprContext() {}

func NewExtendTypeExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExtendTypeExprContext {
	var p = new(ExtendTypeExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = kslParserRULE_extendTypeExpr

	return p
}

func (s *ExtendTypeExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExtendTypeExprContext) EXTEND() antlr.TerminalNode {
	return s.GetToken(kslParserEXTEND, 0)
}

func (s *ExtendTypeExprContext) TYPE() antlr.TerminalNode {
	return s.GetToken(kslParserTYPE, 0)
}

func (s *ExtendTypeExprContext) Name() INameContext {
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

func (s *ExtendTypeExprContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(kslParserLBRACE, 0)
}

func (s *ExtendTypeExprContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(kslParserRBRACE, 0)
}

func (s *ExtendTypeExprContext) AllRelation() []IRelationContext {
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

func (s *ExtendTypeExprContext) Relation(i int) IRelationContext {
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

func (s *ExtendTypeExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExtendTypeExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExtendTypeExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kslListener); ok {
		listenerT.EnterExtendTypeExpr(s)
	}
}

func (s *ExtendTypeExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kslListener); ok {
		listenerT.ExitExtendTypeExpr(s)
	}
}

func (s *ExtendTypeExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case kslVisitor:
		return t.VisitExtendTypeExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *kslParser) ExtendTypeExpr() (localctx IExtendTypeExprContext) {
	localctx = NewExtendTypeExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, kslParserRULE_extendTypeExpr)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(109)
		p.Match(kslParserEXTEND)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(110)
		p.Match(kslParserTYPE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(111)
		p.Name()
	}
	{
		p.SetState(112)
		p.Match(kslParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(116)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&536879232) != 0 {
		{
			p.SetState(113)
			p.Relation()
		}

		p.SetState(118)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(119)
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
	p.EnterRule(localctx, 18, kslParserRULE_extensionParam)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(121)
		p.Name()
	}
	{
		p.SetState(122)
		p.Match(kslParserEXPAND)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(123)
		p.Match(kslParserSTRING_DELIM)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(124)

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
		p.SetState(125)
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
	p.EnterRule(localctx, 20, kslParserRULE_extensionParams)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(127)
		p.ExtensionParam()
	}
	p.SetState(132)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == kslParserARG_DELIM {
		{
			p.SetState(128)
			p.Match(kslParserARG_DELIM)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(129)
			p.ExtensionParam()
		}

		p.SetState(134)
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
	p.EnterRule(localctx, 22, kslParserRULE_extensionReference)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(135)
		p.Match(kslParserEXTENSION_CALL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(136)
		p.TypeReference()
	}
	{
		p.SetState(137)
		p.Match(kslParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(139)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == kslParserFORCED_NAME || _la == kslParserNAME {
		{
			p.SetState(138)
			p.ExtensionParams()
		}

	}
	{
		p.SetState(141)
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
	p.EnterRule(localctx, 24, kslParserRULE_relation)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(146)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == kslParserEXTENSION_CALL {
		{
			p.SetState(143)
			p.ExtensionReference()
		}

		p.SetState(148)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(150)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == kslParserACCESS {
		{
			p.SetState(149)
			p.Match(kslParserACCESS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(152)
		p.Match(kslParserRELATION)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(153)
		p.Name()
	}
	{
		p.SetState(154)
		p.Match(kslParserEXPAND)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(155)
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

type SubRelationWithExtensionContext struct {
	RelationBodyContext
	relationName    INameContext
	subrelationName INameContext
}

func NewSubRelationWithExtensionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SubRelationWithExtensionContext {
	var p = new(SubRelationWithExtensionContext)

	InitEmptyRelationBodyContext(&p.RelationBodyContext)
	p.parser = parser
	p.CopyAll(ctx.(*RelationBodyContext))

	return p
}

func (s *SubRelationWithExtensionContext) GetRelationName() INameContext { return s.relationName }

func (s *SubRelationWithExtensionContext) GetSubrelationName() INameContext { return s.subrelationName }

func (s *SubRelationWithExtensionContext) SetRelationName(v INameContext) { s.relationName = v }

func (s *SubRelationWithExtensionContext) SetSubrelationName(v INameContext) { s.subrelationName = v }

func (s *SubRelationWithExtensionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SubRelationWithExtensionContext) RESOLVE() antlr.TerminalNode {
	return s.GetToken(kslParserRESOLVE, 0)
}

func (s *SubRelationWithExtensionContext) IN_REL_EXTENSION_CALL() antlr.TerminalNode {
	return s.GetToken(kslParserIN_REL_EXTENSION_CALL, 0)
}

func (s *SubRelationWithExtensionContext) TypeReference() ITypeReferenceContext {
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

func (s *SubRelationWithExtensionContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(kslParserLPAREN, 0)
}

func (s *SubRelationWithExtensionContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(kslParserRPAREN, 0)
}

func (s *SubRelationWithExtensionContext) AllName() []INameContext {
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

func (s *SubRelationWithExtensionContext) Name(i int) INameContext {
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

func (s *SubRelationWithExtensionContext) ExtensionParams() IExtensionParamsContext {
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

func (s *SubRelationWithExtensionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kslListener); ok {
		listenerT.EnterSubRelationWithExtension(s)
	}
}

func (s *SubRelationWithExtensionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kslListener); ok {
		listenerT.ExitSubRelationWithExtension(s)
	}
}

func (s *SubRelationWithExtensionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case kslVisitor:
		return t.VisitSubRelationWithExtension(s)

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

type ReferenceWithExtensionContext struct {
	RelationBodyContext
}

func NewReferenceWithExtensionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ReferenceWithExtensionContext {
	var p = new(ReferenceWithExtensionContext)

	InitEmptyRelationBodyContext(&p.RelationBodyContext)
	p.parser = parser
	p.CopyAll(ctx.(*RelationBodyContext))

	return p
}

func (s *ReferenceWithExtensionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReferenceWithExtensionContext) Name() INameContext {
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

func (s *ReferenceWithExtensionContext) IN_REL_EXTENSION_CALL() antlr.TerminalNode {
	return s.GetToken(kslParserIN_REL_EXTENSION_CALL, 0)
}

func (s *ReferenceWithExtensionContext) TypeReference() ITypeReferenceContext {
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

func (s *ReferenceWithExtensionContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(kslParserLPAREN, 0)
}

func (s *ReferenceWithExtensionContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(kslParserRPAREN, 0)
}

func (s *ReferenceWithExtensionContext) ExtensionParams() IExtensionParamsContext {
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

func (s *ReferenceWithExtensionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kslListener); ok {
		listenerT.EnterReferenceWithExtension(s)
	}
}

func (s *ReferenceWithExtensionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kslListener); ok {
		listenerT.ExitReferenceWithExtension(s)
	}
}

func (s *ReferenceWithExtensionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case kslVisitor:
		return t.VisitReferenceWithExtension(s)

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
	_startState := 26
	p.EnterRecursionRule(localctx, 26, kslParserRULE_relationBody, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(201)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 19, p.GetParserRuleContext()) {
	case 1:
		localctx = NewSelfContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(158)
			p.Match(kslParserLSQUARE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(160)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == kslParserCARDINALITY {
			{
				p.SetState(159)
				p.Match(kslParserCARDINALITY)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(162)
			p.TypeReference()
		}
		p.SetState(167)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == kslParserOR {
			{
				p.SetState(163)
				p.Match(kslParserOR)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(164)
				p.TypeReference()
			}

			p.SetState(169)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(170)
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
			p.SetState(172)
			p.Name()
		}

	case 3:
		localctx = NewReferenceWithExtensionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(173)
			p.Name()
		}
		{
			p.SetState(174)
			p.Match(kslParserIN_REL_EXTENSION_CALL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(175)
			p.TypeReference()
		}
		{
			p.SetState(176)
			p.Match(kslParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(178)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == kslParserFORCED_NAME || _la == kslParserNAME {
			{
				p.SetState(177)
				p.ExtensionParams()
			}

		}
		{
			p.SetState(180)
			p.Match(kslParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		localctx = NewSubRelationContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(182)

			var _x = p.Name()

			localctx.(*SubRelationContext).relationName = _x
		}
		{
			p.SetState(183)
			p.Match(kslParserRESOLVE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(184)

			var _x = p.Name()

			localctx.(*SubRelationContext).subrelationName = _x
		}

	case 5:
		localctx = NewSubRelationWithExtensionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(186)

			var _x = p.Name()

			localctx.(*SubRelationWithExtensionContext).relationName = _x
		}
		{
			p.SetState(187)
			p.Match(kslParserRESOLVE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(188)

			var _x = p.Name()

			localctx.(*SubRelationWithExtensionContext).subrelationName = _x
		}
		{
			p.SetState(189)
			p.Match(kslParserIN_REL_EXTENSION_CALL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(190)
			p.TypeReference()
		}
		{
			p.SetState(191)
			p.Match(kslParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(193)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == kslParserFORCED_NAME || _la == kslParserNAME {
			{
				p.SetState(192)
				p.ExtensionParams()
			}

		}
		{
			p.SetState(195)
			p.Match(kslParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 6:
		localctx = NewParenContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(197)
			p.Match(kslParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(198)
			p.relationBody(0)
		}
		{
			p.SetState(199)
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
	p.SetState(214)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 21, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(212)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 20, p.GetParserRuleContext()) {
			case 1:
				localctx = NewAndContext(p, NewRelationBodyContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, kslParserRULE_relationBody)
				p.SetState(203)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
					goto errorExit
				}
				{
					p.SetState(204)
					p.Match(kslParserAND)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(205)
					p.relationBody(4)
				}

			case 2:
				localctx = NewORContext(p, NewRelationBodyContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, kslParserRULE_relationBody)
				p.SetState(206)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
					goto errorExit
				}
				{
					p.SetState(207)
					p.Match(kslParserOR)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(208)
					p.relationBody(3)
				}

			case 3:
				localctx = NewUnlessContext(p, NewRelationBodyContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, kslParserRULE_relationBody)
				p.SetState(209)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
					goto errorExit
				}
				{
					p.SetState(210)
					p.Match(kslParserUNLESS)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(211)
					p.relationBody(2)
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(216)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 21, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 28, kslParserRULE_paramNames)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(217)
		p.Name()
	}
	p.SetState(222)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == kslParserARG_DELIM {
		{
			p.SetState(218)
			p.Match(kslParserARG_DELIM)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(219)
			p.Name()
		}

		p.SetState(224)
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
	p.EnterRule(localctx, 30, kslParserRULE_extension)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(226)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == kslParserACCESS {
		{
			p.SetState(225)
			p.Match(kslParserACCESS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(228)
		p.Match(kslParserEXTENSION)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(229)
		p.Name()
	}
	{
		p.SetState(230)
		p.Match(kslParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(232)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == kslParserFORCED_NAME || _la == kslParserNAME {
		{
			p.SetState(231)
			p.ParamNames()
		}

	}
	{
		p.SetState(234)
		p.Match(kslParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(235)
		p.Match(kslParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(237)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == kslParserACCESS || _la == kslParserTYPE {
		{
			p.SetState(236)
			p.DynamicType()
		}

		p.SetState(239)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(241)
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
	p.EnterRule(localctx, 32, kslParserRULE_dynamicType)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(244)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == kslParserACCESS {
		{
			p.SetState(243)
			p.Match(kslParserACCESS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(246)
		p.Match(kslParserTYPE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(247)
		p.DynamicName()
	}
	{
		p.SetState(248)
		p.Match(kslParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(252)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&33562752) != 0 {
		{
			p.SetState(249)
			p.DynamicRelation()
		}

		p.SetState(254)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(255)
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
	p.EnterRule(localctx, 34, kslParserRULE_dynamicRelation)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(258)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == kslParserALLOW_DUPLICATES {
		{
			p.SetState(257)
			p.Match(kslParserALLOW_DUPLICATES)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	p.SetState(261)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == kslParserACCESS {
		{
			p.SetState(260)
			p.Match(kslParserACCESS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(263)
		p.Match(kslParserRELATION)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(264)
		p.DynamicName()
	}
	{
		p.SetState(265)
		p.Match(kslParserEXPAND)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(266)
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
	p.EnterRule(localctx, 36, kslParserRULE_dynamicName)
	var _alt int

	p.SetState(282)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case kslParserFORCED_NAME, kslParserNAME:
		localctx = NewLiteralContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(268)
			p.Name()
		}

	case kslParserVARREF:
		localctx = NewVariableContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(269)
			p.Match(kslParserVARREF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(270)
			p.Match(kslParserLBRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(271)
			p.Name()
		}
		{
			p.SetState(272)
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
			p.SetState(274)
			p.Match(kslParserTEMPLATE_DELIM)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(276)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = 1
		for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			switch _alt {
			case 1:
				{
					p.SetState(275)
					p.DynamicName()
				}

			default:
				p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
				goto errorExit
			}

			p.SetState(278)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 30, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}
		{
			p.SetState(280)
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
	_startState := 38
	p.EnterRecursionRule(localctx, 38, kslParserRULE_dynamicBody, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(308)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 34, p.GetParserRuleContext()) {
	case 1:
		localctx = NewDynamicSelfContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(285)
			p.Match(kslParserLSQUARE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(287)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == kslParserCARDINALITY {
			{
				p.SetState(286)
				p.Match(kslParserCARDINALITY)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(289)
			p.TypeReference()
		}
		p.SetState(294)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == kslParserOR {
			{
				p.SetState(290)
				p.Match(kslParserOR)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(291)
				p.TypeReference()
			}

			p.SetState(296)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(297)
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
			p.SetState(299)
			p.DynamicName()
		}

	case 3:
		localctx = NewDynamicSubRelationContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(300)
			p.DynamicName()
		}
		{
			p.SetState(301)
			p.Match(kslParserRESOLVE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(302)
			p.DynamicName()
		}

	case 4:
		localctx = NewDynamicParenContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(304)
			p.Match(kslParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(305)
			p.dynamicBody(0)
		}
		{
			p.SetState(306)
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
	p.SetState(321)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 36, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(319)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 35, p.GetParserRuleContext()) {
			case 1:
				localctx = NewDynamicAndContext(p, NewDynamicBodyContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, kslParserRULE_dynamicBody)
				p.SetState(310)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
					goto errorExit
				}
				{
					p.SetState(311)
					p.Match(kslParserAND)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(312)
					p.dynamicBody(4)
				}

			case 2:
				localctx = NewDynamicORContext(p, NewDynamicBodyContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, kslParserRULE_dynamicBody)
				p.SetState(313)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
					goto errorExit
				}
				{
					p.SetState(314)
					p.Match(kslParserOR)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(315)
					p.dynamicBody(3)
				}

			case 3:
				localctx = NewDynamicUnlessContext(p, NewDynamicBodyContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, kslParserRULE_dynamicBody)
				p.SetState(316)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
					goto errorExit
				}
				{
					p.SetState(317)
					p.Match(kslParserUNLESS)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(318)
					p.dynamicBody(2)
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(323)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 36, p.GetParserRuleContext())
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
	case 13:
		var t *RelationBodyContext = nil
		if localctx != nil {
			t = localctx.(*RelationBodyContext)
		}
		return p.RelationBody_Sempred(t, predIndex)

	case 19:
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
