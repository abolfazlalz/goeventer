// Code generated from GoEventer.g4 by ANTLR 4.13.0. DO NOT EDIT.

package grammar

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

type GoEventerLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var GoEventerLexerLexerStaticData struct {
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

func goeventerlexerLexerInit() {
	staticData := &GoEventerLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "','", "':='", "'='", "';'", "':'", "'.'", "'for'", "'while'", "'range'",
		"'func'", "'on'", "'return'", "'chan'", "'notify'", "'struct'", "'string'",
		"'int'", "'||'", "'&&'", "'=='", "'!='", "'>'", "'<'", "'>='", "'<='",
		"'+'", "'-'", "'*'", "'/'", "'%'", "'^'", "'!'", "'true'", "'false'",
		"'nil'", "'('", "')'", "'{'", "'}'",
	}
	staticData.SymbolicNames = []string{
		"", "", "ASSIGN", "UPDATE", "SCOL", "COL", "DOT", "FOR", "WHILE", "RANGE",
		"FUNC", "ON", "RETURN", "CHAN", "NOTIFY", "STRUCT", "STRING_TYPE", "INTEGER_TYPE",
		"OR", "AND", "EQ", "NEQ", "GT", "LT", "GTEQ", "LTEQ", "PLUS", "MINUS",
		"MULT", "DIV", "MOD", "POW", "NOT", "TRUE", "FALSE", "NIL", "PARAN_OPEN",
		"PARAN_CLOSE", "BARACE_OPEN", "BARACE_CLOSE", "ID", "INT", "FLOAT",
		"STRING", "SPACE",
	}
	staticData.RuleNames = []string{
		"T__0", "ASSIGN", "UPDATE", "SCOL", "COL", "DOT", "FOR", "WHILE", "RANGE",
		"FUNC", "ON", "RETURN", "CHAN", "NOTIFY", "STRUCT", "STRING_TYPE", "INTEGER_TYPE",
		"OR", "AND", "EQ", "NEQ", "GT", "LT", "GTEQ", "LTEQ", "PLUS", "MINUS",
		"MULT", "DIV", "MOD", "POW", "NOT", "TRUE", "FALSE", "NIL", "PARAN_OPEN",
		"PARAN_CLOSE", "BARACE_OPEN", "BARACE_CLOSE", "ID", "INT", "FLOAT",
		"STRING", "SPACE",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 44, 269, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36,
		7, 36, 2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7,
		41, 2, 42, 7, 42, 2, 43, 7, 43, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 2, 1,
		2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1,
		7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1,
		9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1,
		11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13,
		1, 13, 1, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1,
		14, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16,
		1, 16, 1, 17, 1, 17, 1, 17, 1, 18, 1, 18, 1, 18, 1, 19, 1, 19, 1, 19, 1,
		20, 1, 20, 1, 20, 1, 21, 1, 21, 1, 22, 1, 22, 1, 23, 1, 23, 1, 23, 1, 24,
		1, 24, 1, 24, 1, 25, 1, 25, 1, 26, 1, 26, 1, 27, 1, 27, 1, 28, 1, 28, 1,
		29, 1, 29, 1, 30, 1, 30, 1, 31, 1, 31, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32,
		1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 34, 1, 34, 1, 34, 1, 34, 1,
		35, 1, 35, 1, 36, 1, 36, 1, 37, 1, 37, 1, 38, 1, 38, 1, 39, 1, 39, 5, 39,
		225, 8, 39, 10, 39, 12, 39, 228, 9, 39, 1, 40, 4, 40, 231, 8, 40, 11, 40,
		12, 40, 232, 1, 41, 4, 41, 236, 8, 41, 11, 41, 12, 41, 237, 1, 41, 1, 41,
		5, 41, 242, 8, 41, 10, 41, 12, 41, 245, 9, 41, 1, 41, 1, 41, 4, 41, 249,
		8, 41, 11, 41, 12, 41, 250, 3, 41, 253, 8, 41, 1, 42, 1, 42, 1, 42, 1,
		42, 5, 42, 259, 8, 42, 10, 42, 12, 42, 262, 9, 42, 1, 42, 1, 42, 1, 43,
		1, 43, 1, 43, 1, 43, 0, 0, 44, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13,
		7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16,
		33, 17, 35, 18, 37, 19, 39, 20, 41, 21, 43, 22, 45, 23, 47, 24, 49, 25,
		51, 26, 53, 27, 55, 28, 57, 29, 59, 30, 61, 31, 63, 32, 65, 33, 67, 34,
		69, 35, 71, 36, 73, 37, 75, 38, 77, 39, 79, 40, 81, 41, 83, 42, 85, 43,
		87, 44, 1, 0, 5, 3, 0, 65, 90, 95, 95, 97, 122, 4, 0, 48, 57, 65, 90, 95,
		95, 97, 122, 1, 0, 48, 57, 3, 0, 10, 10, 13, 13, 34, 34, 3, 0, 9, 10, 13,
		13, 32, 32, 276, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0,
		0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0,
		0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0,
		0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1,
		0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37,
		1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0,
		45, 1, 0, 0, 0, 0, 47, 1, 0, 0, 0, 0, 49, 1, 0, 0, 0, 0, 51, 1, 0, 0, 0,
		0, 53, 1, 0, 0, 0, 0, 55, 1, 0, 0, 0, 0, 57, 1, 0, 0, 0, 0, 59, 1, 0, 0,
		0, 0, 61, 1, 0, 0, 0, 0, 63, 1, 0, 0, 0, 0, 65, 1, 0, 0, 0, 0, 67, 1, 0,
		0, 0, 0, 69, 1, 0, 0, 0, 0, 71, 1, 0, 0, 0, 0, 73, 1, 0, 0, 0, 0, 75, 1,
		0, 0, 0, 0, 77, 1, 0, 0, 0, 0, 79, 1, 0, 0, 0, 0, 81, 1, 0, 0, 0, 0, 83,
		1, 0, 0, 0, 0, 85, 1, 0, 0, 0, 0, 87, 1, 0, 0, 0, 1, 89, 1, 0, 0, 0, 3,
		91, 1, 0, 0, 0, 5, 94, 1, 0, 0, 0, 7, 96, 1, 0, 0, 0, 9, 98, 1, 0, 0, 0,
		11, 100, 1, 0, 0, 0, 13, 102, 1, 0, 0, 0, 15, 106, 1, 0, 0, 0, 17, 112,
		1, 0, 0, 0, 19, 118, 1, 0, 0, 0, 21, 123, 1, 0, 0, 0, 23, 126, 1, 0, 0,
		0, 25, 133, 1, 0, 0, 0, 27, 138, 1, 0, 0, 0, 29, 145, 1, 0, 0, 0, 31, 152,
		1, 0, 0, 0, 33, 159, 1, 0, 0, 0, 35, 163, 1, 0, 0, 0, 37, 166, 1, 0, 0,
		0, 39, 169, 1, 0, 0, 0, 41, 172, 1, 0, 0, 0, 43, 175, 1, 0, 0, 0, 45, 177,
		1, 0, 0, 0, 47, 179, 1, 0, 0, 0, 49, 182, 1, 0, 0, 0, 51, 185, 1, 0, 0,
		0, 53, 187, 1, 0, 0, 0, 55, 189, 1, 0, 0, 0, 57, 191, 1, 0, 0, 0, 59, 193,
		1, 0, 0, 0, 61, 195, 1, 0, 0, 0, 63, 197, 1, 0, 0, 0, 65, 199, 1, 0, 0,
		0, 67, 204, 1, 0, 0, 0, 69, 210, 1, 0, 0, 0, 71, 214, 1, 0, 0, 0, 73, 216,
		1, 0, 0, 0, 75, 218, 1, 0, 0, 0, 77, 220, 1, 0, 0, 0, 79, 222, 1, 0, 0,
		0, 81, 230, 1, 0, 0, 0, 83, 252, 1, 0, 0, 0, 85, 254, 1, 0, 0, 0, 87, 265,
		1, 0, 0, 0, 89, 90, 5, 44, 0, 0, 90, 2, 1, 0, 0, 0, 91, 92, 5, 58, 0, 0,
		92, 93, 5, 61, 0, 0, 93, 4, 1, 0, 0, 0, 94, 95, 5, 61, 0, 0, 95, 6, 1,
		0, 0, 0, 96, 97, 5, 59, 0, 0, 97, 8, 1, 0, 0, 0, 98, 99, 5, 58, 0, 0, 99,
		10, 1, 0, 0, 0, 100, 101, 5, 46, 0, 0, 101, 12, 1, 0, 0, 0, 102, 103, 5,
		102, 0, 0, 103, 104, 5, 111, 0, 0, 104, 105, 5, 114, 0, 0, 105, 14, 1,
		0, 0, 0, 106, 107, 5, 119, 0, 0, 107, 108, 5, 104, 0, 0, 108, 109, 5, 105,
		0, 0, 109, 110, 5, 108, 0, 0, 110, 111, 5, 101, 0, 0, 111, 16, 1, 0, 0,
		0, 112, 113, 5, 114, 0, 0, 113, 114, 5, 97, 0, 0, 114, 115, 5, 110, 0,
		0, 115, 116, 5, 103, 0, 0, 116, 117, 5, 101, 0, 0, 117, 18, 1, 0, 0, 0,
		118, 119, 5, 102, 0, 0, 119, 120, 5, 117, 0, 0, 120, 121, 5, 110, 0, 0,
		121, 122, 5, 99, 0, 0, 122, 20, 1, 0, 0, 0, 123, 124, 5, 111, 0, 0, 124,
		125, 5, 110, 0, 0, 125, 22, 1, 0, 0, 0, 126, 127, 5, 114, 0, 0, 127, 128,
		5, 101, 0, 0, 128, 129, 5, 116, 0, 0, 129, 130, 5, 117, 0, 0, 130, 131,
		5, 114, 0, 0, 131, 132, 5, 110, 0, 0, 132, 24, 1, 0, 0, 0, 133, 134, 5,
		99, 0, 0, 134, 135, 5, 104, 0, 0, 135, 136, 5, 97, 0, 0, 136, 137, 5, 110,
		0, 0, 137, 26, 1, 0, 0, 0, 138, 139, 5, 110, 0, 0, 139, 140, 5, 111, 0,
		0, 140, 141, 5, 116, 0, 0, 141, 142, 5, 105, 0, 0, 142, 143, 5, 102, 0,
		0, 143, 144, 5, 121, 0, 0, 144, 28, 1, 0, 0, 0, 145, 146, 5, 115, 0, 0,
		146, 147, 5, 116, 0, 0, 147, 148, 5, 114, 0, 0, 148, 149, 5, 117, 0, 0,
		149, 150, 5, 99, 0, 0, 150, 151, 5, 116, 0, 0, 151, 30, 1, 0, 0, 0, 152,
		153, 5, 115, 0, 0, 153, 154, 5, 116, 0, 0, 154, 155, 5, 114, 0, 0, 155,
		156, 5, 105, 0, 0, 156, 157, 5, 110, 0, 0, 157, 158, 5, 103, 0, 0, 158,
		32, 1, 0, 0, 0, 159, 160, 5, 105, 0, 0, 160, 161, 5, 110, 0, 0, 161, 162,
		5, 116, 0, 0, 162, 34, 1, 0, 0, 0, 163, 164, 5, 124, 0, 0, 164, 165, 5,
		124, 0, 0, 165, 36, 1, 0, 0, 0, 166, 167, 5, 38, 0, 0, 167, 168, 5, 38,
		0, 0, 168, 38, 1, 0, 0, 0, 169, 170, 5, 61, 0, 0, 170, 171, 5, 61, 0, 0,
		171, 40, 1, 0, 0, 0, 172, 173, 5, 33, 0, 0, 173, 174, 5, 61, 0, 0, 174,
		42, 1, 0, 0, 0, 175, 176, 5, 62, 0, 0, 176, 44, 1, 0, 0, 0, 177, 178, 5,
		60, 0, 0, 178, 46, 1, 0, 0, 0, 179, 180, 5, 62, 0, 0, 180, 181, 5, 61,
		0, 0, 181, 48, 1, 0, 0, 0, 182, 183, 5, 60, 0, 0, 183, 184, 5, 61, 0, 0,
		184, 50, 1, 0, 0, 0, 185, 186, 5, 43, 0, 0, 186, 52, 1, 0, 0, 0, 187, 188,
		5, 45, 0, 0, 188, 54, 1, 0, 0, 0, 189, 190, 5, 42, 0, 0, 190, 56, 1, 0,
		0, 0, 191, 192, 5, 47, 0, 0, 192, 58, 1, 0, 0, 0, 193, 194, 5, 37, 0, 0,
		194, 60, 1, 0, 0, 0, 195, 196, 5, 94, 0, 0, 196, 62, 1, 0, 0, 0, 197, 198,
		5, 33, 0, 0, 198, 64, 1, 0, 0, 0, 199, 200, 5, 116, 0, 0, 200, 201, 5,
		114, 0, 0, 201, 202, 5, 117, 0, 0, 202, 203, 5, 101, 0, 0, 203, 66, 1,
		0, 0, 0, 204, 205, 5, 102, 0, 0, 205, 206, 5, 97, 0, 0, 206, 207, 5, 108,
		0, 0, 207, 208, 5, 115, 0, 0, 208, 209, 5, 101, 0, 0, 209, 68, 1, 0, 0,
		0, 210, 211, 5, 110, 0, 0, 211, 212, 5, 105, 0, 0, 212, 213, 5, 108, 0,
		0, 213, 70, 1, 0, 0, 0, 214, 215, 5, 40, 0, 0, 215, 72, 1, 0, 0, 0, 216,
		217, 5, 41, 0, 0, 217, 74, 1, 0, 0, 0, 218, 219, 5, 123, 0, 0, 219, 76,
		1, 0, 0, 0, 220, 221, 5, 125, 0, 0, 221, 78, 1, 0, 0, 0, 222, 226, 7, 0,
		0, 0, 223, 225, 7, 1, 0, 0, 224, 223, 1, 0, 0, 0, 225, 228, 1, 0, 0, 0,
		226, 224, 1, 0, 0, 0, 226, 227, 1, 0, 0, 0, 227, 80, 1, 0, 0, 0, 228, 226,
		1, 0, 0, 0, 229, 231, 7, 2, 0, 0, 230, 229, 1, 0, 0, 0, 231, 232, 1, 0,
		0, 0, 232, 230, 1, 0, 0, 0, 232, 233, 1, 0, 0, 0, 233, 82, 1, 0, 0, 0,
		234, 236, 7, 2, 0, 0, 235, 234, 1, 0, 0, 0, 236, 237, 1, 0, 0, 0, 237,
		235, 1, 0, 0, 0, 237, 238, 1, 0, 0, 0, 238, 239, 1, 0, 0, 0, 239, 243,
		5, 46, 0, 0, 240, 242, 7, 2, 0, 0, 241, 240, 1, 0, 0, 0, 242, 245, 1, 0,
		0, 0, 243, 241, 1, 0, 0, 0, 243, 244, 1, 0, 0, 0, 244, 253, 1, 0, 0, 0,
		245, 243, 1, 0, 0, 0, 246, 248, 5, 46, 0, 0, 247, 249, 7, 2, 0, 0, 248,
		247, 1, 0, 0, 0, 249, 250, 1, 0, 0, 0, 250, 248, 1, 0, 0, 0, 250, 251,
		1, 0, 0, 0, 251, 253, 1, 0, 0, 0, 252, 235, 1, 0, 0, 0, 252, 246, 1, 0,
		0, 0, 253, 84, 1, 0, 0, 0, 254, 260, 5, 34, 0, 0, 255, 259, 8, 3, 0, 0,
		256, 257, 5, 34, 0, 0, 257, 259, 5, 34, 0, 0, 258, 255, 1, 0, 0, 0, 258,
		256, 1, 0, 0, 0, 259, 262, 1, 0, 0, 0, 260, 258, 1, 0, 0, 0, 260, 261,
		1, 0, 0, 0, 261, 263, 1, 0, 0, 0, 262, 260, 1, 0, 0, 0, 263, 264, 5, 34,
		0, 0, 264, 86, 1, 0, 0, 0, 265, 266, 7, 4, 0, 0, 266, 267, 1, 0, 0, 0,
		267, 268, 6, 43, 0, 0, 268, 88, 1, 0, 0, 0, 9, 0, 226, 232, 237, 243, 250,
		252, 258, 260, 1, 6, 0, 0,
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

// GoEventerLexerInit initializes any static state used to implement GoEventerLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewGoEventerLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func GoEventerLexerInit() {
	staticData := &GoEventerLexerLexerStaticData
	staticData.once.Do(goeventerlexerLexerInit)
}

// NewGoEventerLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewGoEventerLexer(input antlr.CharStream) *GoEventerLexer {
	GoEventerLexerInit()
	l := new(GoEventerLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &GoEventerLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "GoEventer.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// GoEventerLexer tokens.
const (
	GoEventerLexerT__0         = 1
	GoEventerLexerASSIGN       = 2
	GoEventerLexerUPDATE       = 3
	GoEventerLexerSCOL         = 4
	GoEventerLexerCOL          = 5
	GoEventerLexerDOT          = 6
	GoEventerLexerFOR          = 7
	GoEventerLexerWHILE        = 8
	GoEventerLexerRANGE        = 9
	GoEventerLexerFUNC         = 10
	GoEventerLexerON           = 11
	GoEventerLexerRETURN       = 12
	GoEventerLexerCHAN         = 13
	GoEventerLexerNOTIFY       = 14
	GoEventerLexerSTRUCT       = 15
	GoEventerLexerSTRING_TYPE  = 16
	GoEventerLexerINTEGER_TYPE = 17
	GoEventerLexerOR           = 18
	GoEventerLexerAND          = 19
	GoEventerLexerEQ           = 20
	GoEventerLexerNEQ          = 21
	GoEventerLexerGT           = 22
	GoEventerLexerLT           = 23
	GoEventerLexerGTEQ         = 24
	GoEventerLexerLTEQ         = 25
	GoEventerLexerPLUS         = 26
	GoEventerLexerMINUS        = 27
	GoEventerLexerMULT         = 28
	GoEventerLexerDIV          = 29
	GoEventerLexerMOD          = 30
	GoEventerLexerPOW          = 31
	GoEventerLexerNOT          = 32
	GoEventerLexerTRUE         = 33
	GoEventerLexerFALSE        = 34
	GoEventerLexerNIL          = 35
	GoEventerLexerPARAN_OPEN   = 36
	GoEventerLexerPARAN_CLOSE  = 37
	GoEventerLexerBARACE_OPEN  = 38
	GoEventerLexerBARACE_CLOSE = 39
	GoEventerLexerID           = 40
	GoEventerLexerINT          = 41
	GoEventerLexerFLOAT        = 42
	GoEventerLexerSTRING       = 43
	GoEventerLexerSPACE        = 44
)
