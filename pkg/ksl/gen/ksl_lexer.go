// Code generated from /Users/jmarcant/Documents/ksl-schema-language/pkg/ksl/ksl.g4 by ANTLR 4.13.1. DO NOT EDIT.

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
		"", "", "'.'", "", "'public'", "'internal'", "'private'", "'type'",
		"'relation'", "'import'", "'extension'", "", "'AtMostOne'", "'ExactlyOne'",
		"'AtLeastOne'", "'Any'", "'as'", "'and'", "'or'", "'unless'", "'allow_duplicates'",
		"':'", "'{'", "'}'", "'@'", "'('", "')'", "'['", "']'", "'$'", "'`'",
		"'''", "','", "';'",
	}
	staticData.SymbolicNames = []string{
		"", "VERSIONNUM", "RESOLVE", "ACCESS", "PUBLIC", "INTERNAL", "PRIVATE",
		"TYPE", "RELATION", "IMPORT", "EXTENSION", "CARDINALITY", "ATMOSTONE",
		"EXACTLYONE", "ATLEASTONE", "ANY", "AS", "AND", "OR", "UNLESS", "ALLOW_DUPLICATES",
		"EXPAND", "LBRACE", "RBRACE", "EXTENSION_CALL", "LPAREN", "RPAREN",
		"LSQUARE", "RSQARE", "VARREF", "TEMPLATE_DELIM", "STRING_DELIM", "ARG_DELIM",
		"DECL_END", "NAME", "COMMENT", "WS",
	}
	staticData.RuleNames = []string{
		"VERSIONNUM", "RESOLVE", "ACCESS", "PUBLIC", "INTERNAL", "PRIVATE",
		"TYPE", "RELATION", "IMPORT", "EXTENSION", "CARDINALITY", "ATMOSTONE",
		"EXACTLYONE", "ATLEASTONE", "ANY", "AS", "AND", "OR", "UNLESS", "ALLOW_DUPLICATES",
		"EXPAND", "LBRACE", "RBRACE", "EXTENSION_CALL", "LPAREN", "RPAREN",
		"LSQUARE", "RSQARE", "VARREF", "TEMPLATE_DELIM", "STRING_DELIM", "ARG_DELIM",
		"DECL_END", "NAME", "COMMENT", "WS",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 36, 269, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 1, 0,
		1, 0, 1, 0, 4, 0, 77, 8, 0, 11, 0, 12, 0, 78, 1, 1, 1, 1, 1, 2, 1, 2, 1,
		2, 3, 2, 86, 8, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4,
		1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5,
		1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7,
		1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8,
		1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10,
		1, 10, 1, 10, 3, 10, 147, 8, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1,
		11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12,
		1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1,
		13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 15,
		1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 1, 18, 1,
		18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19,
		1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1,
		19, 1, 19, 1, 20, 1, 20, 1, 21, 1, 21, 1, 22, 1, 22, 1, 23, 1, 23, 1, 24,
		1, 24, 1, 25, 1, 25, 1, 26, 1, 26, 1, 27, 1, 27, 1, 28, 1, 28, 1, 29, 1,
		29, 1, 30, 1, 30, 1, 31, 1, 31, 1, 32, 1, 32, 1, 33, 1, 33, 5, 33, 247,
		8, 33, 10, 33, 12, 33, 250, 9, 33, 1, 34, 1, 34, 1, 34, 1, 34, 5, 34, 256,
		8, 34, 10, 34, 12, 34, 259, 9, 34, 1, 34, 1, 34, 1, 35, 4, 35, 264, 8,
		35, 11, 35, 12, 35, 265, 1, 35, 1, 35, 0, 0, 36, 1, 1, 3, 2, 5, 3, 7, 4,
		9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14,
		29, 15, 31, 16, 33, 17, 35, 18, 37, 19, 39, 20, 41, 21, 43, 22, 45, 23,
		47, 24, 49, 25, 51, 26, 53, 27, 55, 28, 57, 29, 59, 30, 61, 31, 63, 32,
		65, 33, 67, 34, 69, 35, 71, 36, 1, 0, 5, 1, 0, 48, 57, 3, 0, 65, 90, 95,
		95, 97, 122, 4, 0, 48, 57, 65, 90, 95, 95, 97, 122, 2, 0, 10, 10, 13, 13,
		3, 0, 9, 10, 13, 13, 32, 32, 277, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0,
		5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0,
		13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0,
		0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0,
		0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0,
		0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1,
		0, 0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1, 0, 0, 0, 0, 49, 1, 0, 0, 0, 0, 51,
		1, 0, 0, 0, 0, 53, 1, 0, 0, 0, 0, 55, 1, 0, 0, 0, 0, 57, 1, 0, 0, 0, 0,
		59, 1, 0, 0, 0, 0, 61, 1, 0, 0, 0, 0, 63, 1, 0, 0, 0, 0, 65, 1, 0, 0, 0,
		0, 67, 1, 0, 0, 0, 0, 69, 1, 0, 0, 0, 0, 71, 1, 0, 0, 0, 1, 73, 1, 0, 0,
		0, 3, 80, 1, 0, 0, 0, 5, 85, 1, 0, 0, 0, 7, 87, 1, 0, 0, 0, 9, 94, 1, 0,
		0, 0, 11, 103, 1, 0, 0, 0, 13, 111, 1, 0, 0, 0, 15, 116, 1, 0, 0, 0, 17,
		125, 1, 0, 0, 0, 19, 132, 1, 0, 0, 0, 21, 146, 1, 0, 0, 0, 23, 148, 1,
		0, 0, 0, 25, 158, 1, 0, 0, 0, 27, 169, 1, 0, 0, 0, 29, 180, 1, 0, 0, 0,
		31, 184, 1, 0, 0, 0, 33, 187, 1, 0, 0, 0, 35, 191, 1, 0, 0, 0, 37, 194,
		1, 0, 0, 0, 39, 201, 1, 0, 0, 0, 41, 218, 1, 0, 0, 0, 43, 220, 1, 0, 0,
		0, 45, 222, 1, 0, 0, 0, 47, 224, 1, 0, 0, 0, 49, 226, 1, 0, 0, 0, 51, 228,
		1, 0, 0, 0, 53, 230, 1, 0, 0, 0, 55, 232, 1, 0, 0, 0, 57, 234, 1, 0, 0,
		0, 59, 236, 1, 0, 0, 0, 61, 238, 1, 0, 0, 0, 63, 240, 1, 0, 0, 0, 65, 242,
		1, 0, 0, 0, 67, 244, 1, 0, 0, 0, 69, 251, 1, 0, 0, 0, 71, 263, 1, 0, 0,
		0, 73, 74, 7, 0, 0, 0, 74, 76, 3, 3, 1, 0, 75, 77, 7, 0, 0, 0, 76, 75,
		1, 0, 0, 0, 77, 78, 1, 0, 0, 0, 78, 76, 1, 0, 0, 0, 78, 79, 1, 0, 0, 0,
		79, 2, 1, 0, 0, 0, 80, 81, 5, 46, 0, 0, 81, 4, 1, 0, 0, 0, 82, 86, 3, 7,
		3, 0, 83, 86, 3, 9, 4, 0, 84, 86, 3, 11, 5, 0, 85, 82, 1, 0, 0, 0, 85,
		83, 1, 0, 0, 0, 85, 84, 1, 0, 0, 0, 86, 6, 1, 0, 0, 0, 87, 88, 5, 112,
		0, 0, 88, 89, 5, 117, 0, 0, 89, 90, 5, 98, 0, 0, 90, 91, 5, 108, 0, 0,
		91, 92, 5, 105, 0, 0, 92, 93, 5, 99, 0, 0, 93, 8, 1, 0, 0, 0, 94, 95, 5,
		105, 0, 0, 95, 96, 5, 110, 0, 0, 96, 97, 5, 116, 0, 0, 97, 98, 5, 101,
		0, 0, 98, 99, 5, 114, 0, 0, 99, 100, 5, 110, 0, 0, 100, 101, 5, 97, 0,
		0, 101, 102, 5, 108, 0, 0, 102, 10, 1, 0, 0, 0, 103, 104, 5, 112, 0, 0,
		104, 105, 5, 114, 0, 0, 105, 106, 5, 105, 0, 0, 106, 107, 5, 118, 0, 0,
		107, 108, 5, 97, 0, 0, 108, 109, 5, 116, 0, 0, 109, 110, 5, 101, 0, 0,
		110, 12, 1, 0, 0, 0, 111, 112, 5, 116, 0, 0, 112, 113, 5, 121, 0, 0, 113,
		114, 5, 112, 0, 0, 114, 115, 5, 101, 0, 0, 115, 14, 1, 0, 0, 0, 116, 117,
		5, 114, 0, 0, 117, 118, 5, 101, 0, 0, 118, 119, 5, 108, 0, 0, 119, 120,
		5, 97, 0, 0, 120, 121, 5, 116, 0, 0, 121, 122, 5, 105, 0, 0, 122, 123,
		5, 111, 0, 0, 123, 124, 5, 110, 0, 0, 124, 16, 1, 0, 0, 0, 125, 126, 5,
		105, 0, 0, 126, 127, 5, 109, 0, 0, 127, 128, 5, 112, 0, 0, 128, 129, 5,
		111, 0, 0, 129, 130, 5, 114, 0, 0, 130, 131, 5, 116, 0, 0, 131, 18, 1,
		0, 0, 0, 132, 133, 5, 101, 0, 0, 133, 134, 5, 120, 0, 0, 134, 135, 5, 116,
		0, 0, 135, 136, 5, 101, 0, 0, 136, 137, 5, 110, 0, 0, 137, 138, 5, 115,
		0, 0, 138, 139, 5, 105, 0, 0, 139, 140, 5, 111, 0, 0, 140, 141, 5, 110,
		0, 0, 141, 20, 1, 0, 0, 0, 142, 147, 3, 23, 11, 0, 143, 147, 3, 25, 12,
		0, 144, 147, 3, 27, 13, 0, 145, 147, 3, 29, 14, 0, 146, 142, 1, 0, 0, 0,
		146, 143, 1, 0, 0, 0, 146, 144, 1, 0, 0, 0, 146, 145, 1, 0, 0, 0, 147,
		22, 1, 0, 0, 0, 148, 149, 5, 65, 0, 0, 149, 150, 5, 116, 0, 0, 150, 151,
		5, 77, 0, 0, 151, 152, 5, 111, 0, 0, 152, 153, 5, 115, 0, 0, 153, 154,
		5, 116, 0, 0, 154, 155, 5, 79, 0, 0, 155, 156, 5, 110, 0, 0, 156, 157,
		5, 101, 0, 0, 157, 24, 1, 0, 0, 0, 158, 159, 5, 69, 0, 0, 159, 160, 5,
		120, 0, 0, 160, 161, 5, 97, 0, 0, 161, 162, 5, 99, 0, 0, 162, 163, 5, 116,
		0, 0, 163, 164, 5, 108, 0, 0, 164, 165, 5, 121, 0, 0, 165, 166, 5, 79,
		0, 0, 166, 167, 5, 110, 0, 0, 167, 168, 5, 101, 0, 0, 168, 26, 1, 0, 0,
		0, 169, 170, 5, 65, 0, 0, 170, 171, 5, 116, 0, 0, 171, 172, 5, 76, 0, 0,
		172, 173, 5, 101, 0, 0, 173, 174, 5, 97, 0, 0, 174, 175, 5, 115, 0, 0,
		175, 176, 5, 116, 0, 0, 176, 177, 5, 79, 0, 0, 177, 178, 5, 110, 0, 0,
		178, 179, 5, 101, 0, 0, 179, 28, 1, 0, 0, 0, 180, 181, 5, 65, 0, 0, 181,
		182, 5, 110, 0, 0, 182, 183, 5, 121, 0, 0, 183, 30, 1, 0, 0, 0, 184, 185,
		5, 97, 0, 0, 185, 186, 5, 115, 0, 0, 186, 32, 1, 0, 0, 0, 187, 188, 5,
		97, 0, 0, 188, 189, 5, 110, 0, 0, 189, 190, 5, 100, 0, 0, 190, 34, 1, 0,
		0, 0, 191, 192, 5, 111, 0, 0, 192, 193, 5, 114, 0, 0, 193, 36, 1, 0, 0,
		0, 194, 195, 5, 117, 0, 0, 195, 196, 5, 110, 0, 0, 196, 197, 5, 108, 0,
		0, 197, 198, 5, 101, 0, 0, 198, 199, 5, 115, 0, 0, 199, 200, 5, 115, 0,
		0, 200, 38, 1, 0, 0, 0, 201, 202, 5, 97, 0, 0, 202, 203, 5, 108, 0, 0,
		203, 204, 5, 108, 0, 0, 204, 205, 5, 111, 0, 0, 205, 206, 5, 119, 0, 0,
		206, 207, 5, 95, 0, 0, 207, 208, 5, 100, 0, 0, 208, 209, 5, 117, 0, 0,
		209, 210, 5, 112, 0, 0, 210, 211, 5, 108, 0, 0, 211, 212, 5, 105, 0, 0,
		212, 213, 5, 99, 0, 0, 213, 214, 5, 97, 0, 0, 214, 215, 5, 116, 0, 0, 215,
		216, 5, 101, 0, 0, 216, 217, 5, 115, 0, 0, 217, 40, 1, 0, 0, 0, 218, 219,
		5, 58, 0, 0, 219, 42, 1, 0, 0, 0, 220, 221, 5, 123, 0, 0, 221, 44, 1, 0,
		0, 0, 222, 223, 5, 125, 0, 0, 223, 46, 1, 0, 0, 0, 224, 225, 5, 64, 0,
		0, 225, 48, 1, 0, 0, 0, 226, 227, 5, 40, 0, 0, 227, 50, 1, 0, 0, 0, 228,
		229, 5, 41, 0, 0, 229, 52, 1, 0, 0, 0, 230, 231, 5, 91, 0, 0, 231, 54,
		1, 0, 0, 0, 232, 233, 5, 93, 0, 0, 233, 56, 1, 0, 0, 0, 234, 235, 5, 36,
		0, 0, 235, 58, 1, 0, 0, 0, 236, 237, 5, 96, 0, 0, 237, 60, 1, 0, 0, 0,
		238, 239, 5, 39, 0, 0, 239, 62, 1, 0, 0, 0, 240, 241, 5, 44, 0, 0, 241,
		64, 1, 0, 0, 0, 242, 243, 5, 59, 0, 0, 243, 66, 1, 0, 0, 0, 244, 248, 7,
		1, 0, 0, 245, 247, 7, 2, 0, 0, 246, 245, 1, 0, 0, 0, 247, 250, 1, 0, 0,
		0, 248, 246, 1, 0, 0, 0, 248, 249, 1, 0, 0, 0, 249, 68, 1, 0, 0, 0, 250,
		248, 1, 0, 0, 0, 251, 252, 5, 47, 0, 0, 252, 253, 5, 47, 0, 0, 253, 257,
		1, 0, 0, 0, 254, 256, 8, 3, 0, 0, 255, 254, 1, 0, 0, 0, 256, 259, 1, 0,
		0, 0, 257, 255, 1, 0, 0, 0, 257, 258, 1, 0, 0, 0, 258, 260, 1, 0, 0, 0,
		259, 257, 1, 0, 0, 0, 260, 261, 6, 34, 0, 0, 261, 70, 1, 0, 0, 0, 262,
		264, 7, 4, 0, 0, 263, 262, 1, 0, 0, 0, 264, 265, 1, 0, 0, 0, 265, 263,
		1, 0, 0, 0, 265, 266, 1, 0, 0, 0, 266, 267, 1, 0, 0, 0, 267, 268, 6, 35,
		0, 0, 268, 72, 1, 0, 0, 0, 7, 0, 78, 85, 146, 248, 257, 265, 1, 6, 0, 0,
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
	kslLexerVERSIONNUM       = 1
	kslLexerRESOLVE          = 2
	kslLexerACCESS           = 3
	kslLexerPUBLIC           = 4
	kslLexerINTERNAL         = 5
	kslLexerPRIVATE          = 6
	kslLexerTYPE             = 7
	kslLexerRELATION         = 8
	kslLexerIMPORT           = 9
	kslLexerEXTENSION        = 10
	kslLexerCARDINALITY      = 11
	kslLexerATMOSTONE        = 12
	kslLexerEXACTLYONE       = 13
	kslLexerATLEASTONE       = 14
	kslLexerANY              = 15
	kslLexerAS               = 16
	kslLexerAND              = 17
	kslLexerOR               = 18
	kslLexerUNLESS           = 19
	kslLexerALLOW_DUPLICATES = 20
	kslLexerEXPAND           = 21
	kslLexerLBRACE           = 22
	kslLexerRBRACE           = 23
	kslLexerEXTENSION_CALL   = 24
	kslLexerLPAREN           = 25
	kslLexerRPAREN           = 26
	kslLexerLSQUARE          = 27
	kslLexerRSQARE           = 28
	kslLexerVARREF           = 29
	kslLexerTEMPLATE_DELIM   = 30
	kslLexerSTRING_DELIM     = 31
	kslLexerARG_DELIM        = 32
	kslLexerDECL_END         = 33
	kslLexerNAME             = 34
	kslLexerCOMMENT          = 35
	kslLexerWS               = 36
)
