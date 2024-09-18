// Code generated from /home/wscalf/Projects/project-kessel/ksl-schema-language/pkg/ksl/ksl.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type kslLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var KslLexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func ksllexerLexerInit() {
	staticData := &KslLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "'version'", "", "'.'", "'namespace'", "", "'public'", "'internal'",
		"'private'", "'type'", "'relation'", "'import'", "'extension'", "",
		"'AtMostOne'", "'ExactlyOne'", "'AtLeastOne'", "'Any'", "'as'", "'and'",
		"'or'", "'unless'", "'allow_duplicates'", "':'", "'{'", "'}'", "'@'",
		"'('", "')'", "'['", "']'", "'$'", "'`'", "'''", "','", "';'",
	}
	staticData.SymbolicNames = []string{
		"", "VERSION", "VERSIONNUM", "RESOLVE", "NAMESPACE", "ACCESS", "PUBLIC",
		"INTERNAL", "PRIVATE", "TYPE", "RELATION", "IMPORT", "EXTENSION", "CARDINALITY",
		"ATMOSTONE", "EXACTLYONE", "ATLEASTONE", "ANY", "AS", "AND", "OR", "UNLESS",
		"ALLOW_DUPLICATES", "EXPAND", "LBRACE", "RBRACE", "EXTENSION_CALL",
		"LPAREN", "RPAREN", "LSQUARE", "RSQARE", "VARREF", "TEMPLATE_DELIM",
		"STRING_DELIM", "ARG_DELIM", "DECL_END", "NAME", "COMMENT", "WS",
	}
	staticData.RuleNames = []string{
		"VERSION", "VERSIONNUM", "RESOLVE", "NAMESPACE", "ACCESS", "PUBLIC",
		"INTERNAL", "PRIVATE", "TYPE", "RELATION", "IMPORT", "EXTENSION", "CARDINALITY",
		"ATMOSTONE", "EXACTLYONE", "ATLEASTONE", "ANY", "AS", "AND", "OR", "UNLESS",
		"ALLOW_DUPLICATES", "EXPAND", "LBRACE", "RBRACE", "EXTENSION_CALL",
		"LPAREN", "RPAREN", "LSQUARE", "RSQARE", "VARREF", "TEMPLATE_DELIM",
		"STRING_DELIM", "ARG_DELIM", "DECL_END", "NAME", "COMMENT", "WS",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 38, 291, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36,
		7, 36, 2, 37, 7, 37, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1,
		1, 1, 1, 1, 1, 4, 1, 89, 8, 1, 11, 1, 12, 1, 90, 1, 2, 1, 2, 1, 3, 1, 3,
		1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 3, 4,
		108, 8, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6,
		1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7,
		1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9,
		1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10,
		1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1,
		12, 1, 12, 1, 12, 1, 12, 3, 12, 169, 8, 12, 1, 13, 1, 13, 1, 13, 1, 13,
		1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1,
		14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 1, 15,
		1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 1,
		16, 1, 17, 1, 17, 1, 17, 1, 18, 1, 18, 1, 18, 1, 18, 1, 19, 1, 19, 1, 19,
		1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 21, 1, 21, 1, 21, 1,
		21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21,
		1, 21, 1, 21, 1, 21, 1, 22, 1, 22, 1, 23, 1, 23, 1, 24, 1, 24, 1, 25, 1,
		25, 1, 26, 1, 26, 1, 27, 1, 27, 1, 28, 1, 28, 1, 29, 1, 29, 1, 30, 1, 30,
		1, 31, 1, 31, 1, 32, 1, 32, 1, 33, 1, 33, 1, 34, 1, 34, 1, 35, 1, 35, 5,
		35, 269, 8, 35, 10, 35, 12, 35, 272, 9, 35, 1, 36, 1, 36, 1, 36, 1, 36,
		5, 36, 278, 8, 36, 10, 36, 12, 36, 281, 9, 36, 1, 36, 1, 36, 1, 37, 4,
		37, 286, 8, 37, 11, 37, 12, 37, 287, 1, 37, 1, 37, 0, 0, 38, 1, 1, 3, 2,
		5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25,
		13, 27, 14, 29, 15, 31, 16, 33, 17, 35, 18, 37, 19, 39, 20, 41, 21, 43,
		22, 45, 23, 47, 24, 49, 25, 51, 26, 53, 27, 55, 28, 57, 29, 59, 30, 61,
		31, 63, 32, 65, 33, 67, 34, 69, 35, 71, 36, 73, 37, 75, 38, 1, 0, 5, 1,
		0, 48, 57, 3, 0, 65, 90, 95, 95, 97, 122, 4, 0, 48, 57, 65, 90, 95, 95,
		97, 122, 2, 0, 10, 10, 13, 13, 3, 0, 9, 10, 13, 13, 32, 32, 299, 0, 1,
		1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9,
		1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0,
		17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0,
		0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0,
		0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0,
		0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1,
		0, 0, 0, 0, 49, 1, 0, 0, 0, 0, 51, 1, 0, 0, 0, 0, 53, 1, 0, 0, 0, 0, 55,
		1, 0, 0, 0, 0, 57, 1, 0, 0, 0, 0, 59, 1, 0, 0, 0, 0, 61, 1, 0, 0, 0, 0,
		63, 1, 0, 0, 0, 0, 65, 1, 0, 0, 0, 0, 67, 1, 0, 0, 0, 0, 69, 1, 0, 0, 0,
		0, 71, 1, 0, 0, 0, 0, 73, 1, 0, 0, 0, 0, 75, 1, 0, 0, 0, 1, 77, 1, 0, 0,
		0, 3, 85, 1, 0, 0, 0, 5, 92, 1, 0, 0, 0, 7, 94, 1, 0, 0, 0, 9, 107, 1,
		0, 0, 0, 11, 109, 1, 0, 0, 0, 13, 116, 1, 0, 0, 0, 15, 125, 1, 0, 0, 0,
		17, 133, 1, 0, 0, 0, 19, 138, 1, 0, 0, 0, 21, 147, 1, 0, 0, 0, 23, 154,
		1, 0, 0, 0, 25, 168, 1, 0, 0, 0, 27, 170, 1, 0, 0, 0, 29, 180, 1, 0, 0,
		0, 31, 191, 1, 0, 0, 0, 33, 202, 1, 0, 0, 0, 35, 206, 1, 0, 0, 0, 37, 209,
		1, 0, 0, 0, 39, 213, 1, 0, 0, 0, 41, 216, 1, 0, 0, 0, 43, 223, 1, 0, 0,
		0, 45, 240, 1, 0, 0, 0, 47, 242, 1, 0, 0, 0, 49, 244, 1, 0, 0, 0, 51, 246,
		1, 0, 0, 0, 53, 248, 1, 0, 0, 0, 55, 250, 1, 0, 0, 0, 57, 252, 1, 0, 0,
		0, 59, 254, 1, 0, 0, 0, 61, 256, 1, 0, 0, 0, 63, 258, 1, 0, 0, 0, 65, 260,
		1, 0, 0, 0, 67, 262, 1, 0, 0, 0, 69, 264, 1, 0, 0, 0, 71, 266, 1, 0, 0,
		0, 73, 273, 1, 0, 0, 0, 75, 285, 1, 0, 0, 0, 77, 78, 5, 118, 0, 0, 78,
		79, 5, 101, 0, 0, 79, 80, 5, 114, 0, 0, 80, 81, 5, 115, 0, 0, 81, 82, 5,
		105, 0, 0, 82, 83, 5, 111, 0, 0, 83, 84, 5, 110, 0, 0, 84, 2, 1, 0, 0,
		0, 85, 86, 7, 0, 0, 0, 86, 88, 3, 5, 2, 0, 87, 89, 7, 0, 0, 0, 88, 87,
		1, 0, 0, 0, 89, 90, 1, 0, 0, 0, 90, 88, 1, 0, 0, 0, 90, 91, 1, 0, 0, 0,
		91, 4, 1, 0, 0, 0, 92, 93, 5, 46, 0, 0, 93, 6, 1, 0, 0, 0, 94, 95, 5, 110,
		0, 0, 95, 96, 5, 97, 0, 0, 96, 97, 5, 109, 0, 0, 97, 98, 5, 101, 0, 0,
		98, 99, 5, 115, 0, 0, 99, 100, 5, 112, 0, 0, 100, 101, 5, 97, 0, 0, 101,
		102, 5, 99, 0, 0, 102, 103, 5, 101, 0, 0, 103, 8, 1, 0, 0, 0, 104, 108,
		3, 11, 5, 0, 105, 108, 3, 13, 6, 0, 106, 108, 3, 15, 7, 0, 107, 104, 1,
		0, 0, 0, 107, 105, 1, 0, 0, 0, 107, 106, 1, 0, 0, 0, 108, 10, 1, 0, 0,
		0, 109, 110, 5, 112, 0, 0, 110, 111, 5, 117, 0, 0, 111, 112, 5, 98, 0,
		0, 112, 113, 5, 108, 0, 0, 113, 114, 5, 105, 0, 0, 114, 115, 5, 99, 0,
		0, 115, 12, 1, 0, 0, 0, 116, 117, 5, 105, 0, 0, 117, 118, 5, 110, 0, 0,
		118, 119, 5, 116, 0, 0, 119, 120, 5, 101, 0, 0, 120, 121, 5, 114, 0, 0,
		121, 122, 5, 110, 0, 0, 122, 123, 5, 97, 0, 0, 123, 124, 5, 108, 0, 0,
		124, 14, 1, 0, 0, 0, 125, 126, 5, 112, 0, 0, 126, 127, 5, 114, 0, 0, 127,
		128, 5, 105, 0, 0, 128, 129, 5, 118, 0, 0, 129, 130, 5, 97, 0, 0, 130,
		131, 5, 116, 0, 0, 131, 132, 5, 101, 0, 0, 132, 16, 1, 0, 0, 0, 133, 134,
		5, 116, 0, 0, 134, 135, 5, 121, 0, 0, 135, 136, 5, 112, 0, 0, 136, 137,
		5, 101, 0, 0, 137, 18, 1, 0, 0, 0, 138, 139, 5, 114, 0, 0, 139, 140, 5,
		101, 0, 0, 140, 141, 5, 108, 0, 0, 141, 142, 5, 97, 0, 0, 142, 143, 5,
		116, 0, 0, 143, 144, 5, 105, 0, 0, 144, 145, 5, 111, 0, 0, 145, 146, 5,
		110, 0, 0, 146, 20, 1, 0, 0, 0, 147, 148, 5, 105, 0, 0, 148, 149, 5, 109,
		0, 0, 149, 150, 5, 112, 0, 0, 150, 151, 5, 111, 0, 0, 151, 152, 5, 114,
		0, 0, 152, 153, 5, 116, 0, 0, 153, 22, 1, 0, 0, 0, 154, 155, 5, 101, 0,
		0, 155, 156, 5, 120, 0, 0, 156, 157, 5, 116, 0, 0, 157, 158, 5, 101, 0,
		0, 158, 159, 5, 110, 0, 0, 159, 160, 5, 115, 0, 0, 160, 161, 5, 105, 0,
		0, 161, 162, 5, 111, 0, 0, 162, 163, 5, 110, 0, 0, 163, 24, 1, 0, 0, 0,
		164, 169, 3, 27, 13, 0, 165, 169, 3, 29, 14, 0, 166, 169, 3, 31, 15, 0,
		167, 169, 3, 33, 16, 0, 168, 164, 1, 0, 0, 0, 168, 165, 1, 0, 0, 0, 168,
		166, 1, 0, 0, 0, 168, 167, 1, 0, 0, 0, 169, 26, 1, 0, 0, 0, 170, 171, 5,
		65, 0, 0, 171, 172, 5, 116, 0, 0, 172, 173, 5, 77, 0, 0, 173, 174, 5, 111,
		0, 0, 174, 175, 5, 115, 0, 0, 175, 176, 5, 116, 0, 0, 176, 177, 5, 79,
		0, 0, 177, 178, 5, 110, 0, 0, 178, 179, 5, 101, 0, 0, 179, 28, 1, 0, 0,
		0, 180, 181, 5, 69, 0, 0, 181, 182, 5, 120, 0, 0, 182, 183, 5, 97, 0, 0,
		183, 184, 5, 99, 0, 0, 184, 185, 5, 116, 0, 0, 185, 186, 5, 108, 0, 0,
		186, 187, 5, 121, 0, 0, 187, 188, 5, 79, 0, 0, 188, 189, 5, 110, 0, 0,
		189, 190, 5, 101, 0, 0, 190, 30, 1, 0, 0, 0, 191, 192, 5, 65, 0, 0, 192,
		193, 5, 116, 0, 0, 193, 194, 5, 76, 0, 0, 194, 195, 5, 101, 0, 0, 195,
		196, 5, 97, 0, 0, 196, 197, 5, 115, 0, 0, 197, 198, 5, 116, 0, 0, 198,
		199, 5, 79, 0, 0, 199, 200, 5, 110, 0, 0, 200, 201, 5, 101, 0, 0, 201,
		32, 1, 0, 0, 0, 202, 203, 5, 65, 0, 0, 203, 204, 5, 110, 0, 0, 204, 205,
		5, 121, 0, 0, 205, 34, 1, 0, 0, 0, 206, 207, 5, 97, 0, 0, 207, 208, 5,
		115, 0, 0, 208, 36, 1, 0, 0, 0, 209, 210, 5, 97, 0, 0, 210, 211, 5, 110,
		0, 0, 211, 212, 5, 100, 0, 0, 212, 38, 1, 0, 0, 0, 213, 214, 5, 111, 0,
		0, 214, 215, 5, 114, 0, 0, 215, 40, 1, 0, 0, 0, 216, 217, 5, 117, 0, 0,
		217, 218, 5, 110, 0, 0, 218, 219, 5, 108, 0, 0, 219, 220, 5, 101, 0, 0,
		220, 221, 5, 115, 0, 0, 221, 222, 5, 115, 0, 0, 222, 42, 1, 0, 0, 0, 223,
		224, 5, 97, 0, 0, 224, 225, 5, 108, 0, 0, 225, 226, 5, 108, 0, 0, 226,
		227, 5, 111, 0, 0, 227, 228, 5, 119, 0, 0, 228, 229, 5, 95, 0, 0, 229,
		230, 5, 100, 0, 0, 230, 231, 5, 117, 0, 0, 231, 232, 5, 112, 0, 0, 232,
		233, 5, 108, 0, 0, 233, 234, 5, 105, 0, 0, 234, 235, 5, 99, 0, 0, 235,
		236, 5, 97, 0, 0, 236, 237, 5, 116, 0, 0, 237, 238, 5, 101, 0, 0, 238,
		239, 5, 115, 0, 0, 239, 44, 1, 0, 0, 0, 240, 241, 5, 58, 0, 0, 241, 46,
		1, 0, 0, 0, 242, 243, 5, 123, 0, 0, 243, 48, 1, 0, 0, 0, 244, 245, 5, 125,
		0, 0, 245, 50, 1, 0, 0, 0, 246, 247, 5, 64, 0, 0, 247, 52, 1, 0, 0, 0,
		248, 249, 5, 40, 0, 0, 249, 54, 1, 0, 0, 0, 250, 251, 5, 41, 0, 0, 251,
		56, 1, 0, 0, 0, 252, 253, 5, 91, 0, 0, 253, 58, 1, 0, 0, 0, 254, 255, 5,
		93, 0, 0, 255, 60, 1, 0, 0, 0, 256, 257, 5, 36, 0, 0, 257, 62, 1, 0, 0,
		0, 258, 259, 5, 96, 0, 0, 259, 64, 1, 0, 0, 0, 260, 261, 5, 39, 0, 0, 261,
		66, 1, 0, 0, 0, 262, 263, 5, 44, 0, 0, 263, 68, 1, 0, 0, 0, 264, 265, 5,
		59, 0, 0, 265, 70, 1, 0, 0, 0, 266, 270, 7, 1, 0, 0, 267, 269, 7, 2, 0,
		0, 268, 267, 1, 0, 0, 0, 269, 272, 1, 0, 0, 0, 270, 268, 1, 0, 0, 0, 270,
		271, 1, 0, 0, 0, 271, 72, 1, 0, 0, 0, 272, 270, 1, 0, 0, 0, 273, 274, 5,
		47, 0, 0, 274, 275, 5, 47, 0, 0, 275, 279, 1, 0, 0, 0, 276, 278, 8, 3,
		0, 0, 277, 276, 1, 0, 0, 0, 278, 281, 1, 0, 0, 0, 279, 277, 1, 0, 0, 0,
		279, 280, 1, 0, 0, 0, 280, 282, 1, 0, 0, 0, 281, 279, 1, 0, 0, 0, 282,
		283, 6, 36, 0, 0, 283, 74, 1, 0, 0, 0, 284, 286, 7, 4, 0, 0, 285, 284,
		1, 0, 0, 0, 286, 287, 1, 0, 0, 0, 287, 285, 1, 0, 0, 0, 287, 288, 1, 0,
		0, 0, 288, 289, 1, 0, 0, 0, 289, 290, 6, 37, 0, 0, 290, 76, 1, 0, 0, 0,
		7, 0, 90, 107, 168, 270, 279, 287, 1, 6, 0, 0,
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

// kslLexerInit initializes any static state used to implement kslLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewkslLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func KslLexerInit() {
	staticData := &KslLexerLexerStaticData
	staticData.once.Do(ksllexerLexerInit)
}

// NewkslLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewkslLexer(input antlr.CharStream) *kslLexer {
	KslLexerInit()
	l := new(kslLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &KslLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "ksl.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// kslLexer tokens.
const (
	kslLexerVERSION          = 1
	kslLexerVERSIONNUM       = 2
	kslLexerRESOLVE          = 3
	kslLexerNAMESPACE        = 4
	kslLexerACCESS           = 5
	kslLexerPUBLIC           = 6
	kslLexerINTERNAL         = 7
	kslLexerPRIVATE          = 8
	kslLexerTYPE             = 9
	kslLexerRELATION         = 10
	kslLexerIMPORT           = 11
	kslLexerEXTENSION        = 12
	kslLexerCARDINALITY      = 13
	kslLexerATMOSTONE        = 14
	kslLexerEXACTLYONE       = 15
	kslLexerATLEASTONE       = 16
	kslLexerANY              = 17
	kslLexerAS               = 18
	kslLexerAND              = 19
	kslLexerOR               = 20
	kslLexerUNLESS           = 21
	kslLexerALLOW_DUPLICATES = 22
	kslLexerEXPAND           = 23
	kslLexerLBRACE           = 24
	kslLexerRBRACE           = 25
	kslLexerEXTENSION_CALL   = 26
	kslLexerLPAREN           = 27
	kslLexerRPAREN           = 28
	kslLexerLSQUARE          = 29
	kslLexerRSQARE           = 30
	kslLexerVARREF           = 31
	kslLexerTEMPLATE_DELIM   = 32
	kslLexerSTRING_DELIM     = 33
	kslLexerARG_DELIM        = 34
	kslLexerDECL_END         = 35
	kslLexerNAME             = 36
	kslLexerCOMMENT          = 37
	kslLexerWS               = 38
)
