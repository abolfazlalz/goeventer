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
		"'func'", "'on'", "'return'", "'if'", "'else'", "'chan'", "'notify'",
		"'struct'", "'string'", "'int'", "'||'", "'&&'", "'=='", "'!='", "'>'",
		"'<'", "'>='", "'<='", "'+'", "'-'", "'*'", "'/'", "'%'", "'^'", "'!'",
		"'true'", "'false'", "'nil'", "'('", "')'", "'{'", "'}'",
	}
	staticData.SymbolicNames = []string{
		"", "", "ASSIGN", "UPDATE", "SCOL", "COL", "DOT", "FOR", "WHILE", "RANGE",
		"FUNC", "ON", "RETURN", "IF", "ELSE", "CHAN", "NOTIFY", "STRUCT", "STRING_TYPE",
		"INTEGER_TYPE", "OR", "AND", "EQ", "NEQ", "GT", "LT", "GTEQ", "LTEQ",
		"PLUS", "MINUS", "MULT", "DIV", "MOD", "POW", "NOT", "TRUE", "FALSE",
		"NIL", "PARAN_OPEN", "PARAN_CLOSE", "BARACE_OPEN", "BARACE_CLOSE", "ID",
		"INT", "FLOAT", "STRING", "SPACE",
	}
	staticData.RuleNames = []string{
		"T__0", "ASSIGN", "UPDATE", "SCOL", "COL", "DOT", "FOR", "WHILE", "RANGE",
		"FUNC", "ON", "RETURN", "IF", "ELSE", "CHAN", "NOTIFY", "STRUCT", "STRING_TYPE",
		"INTEGER_TYPE", "OR", "AND", "EQ", "NEQ", "GT", "LT", "GTEQ", "LTEQ",
		"PLUS", "MINUS", "MULT", "DIV", "MOD", "POW", "NOT", "TRUE", "FALSE",
		"NIL", "PARAN_OPEN", "PARAN_CLOSE", "BARACE_OPEN", "BARACE_CLOSE", "ID",
		"INT", "FLOAT", "STRING", "SPACE",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 46, 281, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36,
		7, 36, 2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7,
		41, 2, 42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 2, 45, 7, 45, 1, 0, 1, 0,
		1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 6,
		1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8,
		1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1,
		11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 13,
		1, 13, 1, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 15, 1,
		15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16,
		1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 18, 1,
		18, 1, 18, 1, 18, 1, 19, 1, 19, 1, 19, 1, 20, 1, 20, 1, 20, 1, 21, 1, 21,
		1, 21, 1, 22, 1, 22, 1, 22, 1, 23, 1, 23, 1, 24, 1, 24, 1, 25, 1, 25, 1,
		25, 1, 26, 1, 26, 1, 26, 1, 27, 1, 27, 1, 28, 1, 28, 1, 29, 1, 29, 1, 30,
		1, 30, 1, 31, 1, 31, 1, 32, 1, 32, 1, 33, 1, 33, 1, 34, 1, 34, 1, 34, 1,
		34, 1, 34, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 36, 1, 36, 1, 36,
		1, 36, 1, 37, 1, 37, 1, 38, 1, 38, 1, 39, 1, 39, 1, 40, 1, 40, 1, 41, 1,
		41, 5, 41, 237, 8, 41, 10, 41, 12, 41, 240, 9, 41, 1, 42, 4, 42, 243, 8,
		42, 11, 42, 12, 42, 244, 1, 43, 4, 43, 248, 8, 43, 11, 43, 12, 43, 249,
		1, 43, 1, 43, 5, 43, 254, 8, 43, 10, 43, 12, 43, 257, 9, 43, 1, 43, 1,
		43, 4, 43, 261, 8, 43, 11, 43, 12, 43, 262, 3, 43, 265, 8, 43, 1, 44, 1,
		44, 1, 44, 1, 44, 5, 44, 271, 8, 44, 10, 44, 12, 44, 274, 9, 44, 1, 44,
		1, 44, 1, 45, 1, 45, 1, 45, 1, 45, 0, 0, 46, 1, 1, 3, 2, 5, 3, 7, 4, 9,
		5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14,
		29, 15, 31, 16, 33, 17, 35, 18, 37, 19, 39, 20, 41, 21, 43, 22, 45, 23,
		47, 24, 49, 25, 51, 26, 53, 27, 55, 28, 57, 29, 59, 30, 61, 31, 63, 32,
		65, 33, 67, 34, 69, 35, 71, 36, 73, 37, 75, 38, 77, 39, 79, 40, 81, 41,
		83, 42, 85, 43, 87, 44, 89, 45, 91, 46, 1, 0, 5, 3, 0, 65, 90, 95, 95,
		97, 122, 4, 0, 48, 57, 65, 90, 95, 95, 97, 122, 1, 0, 48, 57, 3, 0, 10,
		10, 13, 13, 34, 34, 3, 0, 9, 10, 13, 13, 32, 32, 288, 0, 1, 1, 0, 0, 0,
		0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0,
		0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0,
		0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0,
		0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1,
		0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41,
		1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1, 0, 0, 0, 0,
		49, 1, 0, 0, 0, 0, 51, 1, 0, 0, 0, 0, 53, 1, 0, 0, 0, 0, 55, 1, 0, 0, 0,
		0, 57, 1, 0, 0, 0, 0, 59, 1, 0, 0, 0, 0, 61, 1, 0, 0, 0, 0, 63, 1, 0, 0,
		0, 0, 65, 1, 0, 0, 0, 0, 67, 1, 0, 0, 0, 0, 69, 1, 0, 0, 0, 0, 71, 1, 0,
		0, 0, 0, 73, 1, 0, 0, 0, 0, 75, 1, 0, 0, 0, 0, 77, 1, 0, 0, 0, 0, 79, 1,
		0, 0, 0, 0, 81, 1, 0, 0, 0, 0, 83, 1, 0, 0, 0, 0, 85, 1, 0, 0, 0, 0, 87,
		1, 0, 0, 0, 0, 89, 1, 0, 0, 0, 0, 91, 1, 0, 0, 0, 1, 93, 1, 0, 0, 0, 3,
		95, 1, 0, 0, 0, 5, 98, 1, 0, 0, 0, 7, 100, 1, 0, 0, 0, 9, 102, 1, 0, 0,
		0, 11, 104, 1, 0, 0, 0, 13, 106, 1, 0, 0, 0, 15, 110, 1, 0, 0, 0, 17, 116,
		1, 0, 0, 0, 19, 122, 1, 0, 0, 0, 21, 127, 1, 0, 0, 0, 23, 130, 1, 0, 0,
		0, 25, 137, 1, 0, 0, 0, 27, 140, 1, 0, 0, 0, 29, 145, 1, 0, 0, 0, 31, 150,
		1, 0, 0, 0, 33, 157, 1, 0, 0, 0, 35, 164, 1, 0, 0, 0, 37, 171, 1, 0, 0,
		0, 39, 175, 1, 0, 0, 0, 41, 178, 1, 0, 0, 0, 43, 181, 1, 0, 0, 0, 45, 184,
		1, 0, 0, 0, 47, 187, 1, 0, 0, 0, 49, 189, 1, 0, 0, 0, 51, 191, 1, 0, 0,
		0, 53, 194, 1, 0, 0, 0, 55, 197, 1, 0, 0, 0, 57, 199, 1, 0, 0, 0, 59, 201,
		1, 0, 0, 0, 61, 203, 1, 0, 0, 0, 63, 205, 1, 0, 0, 0, 65, 207, 1, 0, 0,
		0, 67, 209, 1, 0, 0, 0, 69, 211, 1, 0, 0, 0, 71, 216, 1, 0, 0, 0, 73, 222,
		1, 0, 0, 0, 75, 226, 1, 0, 0, 0, 77, 228, 1, 0, 0, 0, 79, 230, 1, 0, 0,
		0, 81, 232, 1, 0, 0, 0, 83, 234, 1, 0, 0, 0, 85, 242, 1, 0, 0, 0, 87, 264,
		1, 0, 0, 0, 89, 266, 1, 0, 0, 0, 91, 277, 1, 0, 0, 0, 93, 94, 5, 44, 0,
		0, 94, 2, 1, 0, 0, 0, 95, 96, 5, 58, 0, 0, 96, 97, 5, 61, 0, 0, 97, 4,
		1, 0, 0, 0, 98, 99, 5, 61, 0, 0, 99, 6, 1, 0, 0, 0, 100, 101, 5, 59, 0,
		0, 101, 8, 1, 0, 0, 0, 102, 103, 5, 58, 0, 0, 103, 10, 1, 0, 0, 0, 104,
		105, 5, 46, 0, 0, 105, 12, 1, 0, 0, 0, 106, 107, 5, 102, 0, 0, 107, 108,
		5, 111, 0, 0, 108, 109, 5, 114, 0, 0, 109, 14, 1, 0, 0, 0, 110, 111, 5,
		119, 0, 0, 111, 112, 5, 104, 0, 0, 112, 113, 5, 105, 0, 0, 113, 114, 5,
		108, 0, 0, 114, 115, 5, 101, 0, 0, 115, 16, 1, 0, 0, 0, 116, 117, 5, 114,
		0, 0, 117, 118, 5, 97, 0, 0, 118, 119, 5, 110, 0, 0, 119, 120, 5, 103,
		0, 0, 120, 121, 5, 101, 0, 0, 121, 18, 1, 0, 0, 0, 122, 123, 5, 102, 0,
		0, 123, 124, 5, 117, 0, 0, 124, 125, 5, 110, 0, 0, 125, 126, 5, 99, 0,
		0, 126, 20, 1, 0, 0, 0, 127, 128, 5, 111, 0, 0, 128, 129, 5, 110, 0, 0,
		129, 22, 1, 0, 0, 0, 130, 131, 5, 114, 0, 0, 131, 132, 5, 101, 0, 0, 132,
		133, 5, 116, 0, 0, 133, 134, 5, 117, 0, 0, 134, 135, 5, 114, 0, 0, 135,
		136, 5, 110, 0, 0, 136, 24, 1, 0, 0, 0, 137, 138, 5, 105, 0, 0, 138, 139,
		5, 102, 0, 0, 139, 26, 1, 0, 0, 0, 140, 141, 5, 101, 0, 0, 141, 142, 5,
		108, 0, 0, 142, 143, 5, 115, 0, 0, 143, 144, 5, 101, 0, 0, 144, 28, 1,
		0, 0, 0, 145, 146, 5, 99, 0, 0, 146, 147, 5, 104, 0, 0, 147, 148, 5, 97,
		0, 0, 148, 149, 5, 110, 0, 0, 149, 30, 1, 0, 0, 0, 150, 151, 5, 110, 0,
		0, 151, 152, 5, 111, 0, 0, 152, 153, 5, 116, 0, 0, 153, 154, 5, 105, 0,
		0, 154, 155, 5, 102, 0, 0, 155, 156, 5, 121, 0, 0, 156, 32, 1, 0, 0, 0,
		157, 158, 5, 115, 0, 0, 158, 159, 5, 116, 0, 0, 159, 160, 5, 114, 0, 0,
		160, 161, 5, 117, 0, 0, 161, 162, 5, 99, 0, 0, 162, 163, 5, 116, 0, 0,
		163, 34, 1, 0, 0, 0, 164, 165, 5, 115, 0, 0, 165, 166, 5, 116, 0, 0, 166,
		167, 5, 114, 0, 0, 167, 168, 5, 105, 0, 0, 168, 169, 5, 110, 0, 0, 169,
		170, 5, 103, 0, 0, 170, 36, 1, 0, 0, 0, 171, 172, 5, 105, 0, 0, 172, 173,
		5, 110, 0, 0, 173, 174, 5, 116, 0, 0, 174, 38, 1, 0, 0, 0, 175, 176, 5,
		124, 0, 0, 176, 177, 5, 124, 0, 0, 177, 40, 1, 0, 0, 0, 178, 179, 5, 38,
		0, 0, 179, 180, 5, 38, 0, 0, 180, 42, 1, 0, 0, 0, 181, 182, 5, 61, 0, 0,
		182, 183, 5, 61, 0, 0, 183, 44, 1, 0, 0, 0, 184, 185, 5, 33, 0, 0, 185,
		186, 5, 61, 0, 0, 186, 46, 1, 0, 0, 0, 187, 188, 5, 62, 0, 0, 188, 48,
		1, 0, 0, 0, 189, 190, 5, 60, 0, 0, 190, 50, 1, 0, 0, 0, 191, 192, 5, 62,
		0, 0, 192, 193, 5, 61, 0, 0, 193, 52, 1, 0, 0, 0, 194, 195, 5, 60, 0, 0,
		195, 196, 5, 61, 0, 0, 196, 54, 1, 0, 0, 0, 197, 198, 5, 43, 0, 0, 198,
		56, 1, 0, 0, 0, 199, 200, 5, 45, 0, 0, 200, 58, 1, 0, 0, 0, 201, 202, 5,
		42, 0, 0, 202, 60, 1, 0, 0, 0, 203, 204, 5, 47, 0, 0, 204, 62, 1, 0, 0,
		0, 205, 206, 5, 37, 0, 0, 206, 64, 1, 0, 0, 0, 207, 208, 5, 94, 0, 0, 208,
		66, 1, 0, 0, 0, 209, 210, 5, 33, 0, 0, 210, 68, 1, 0, 0, 0, 211, 212, 5,
		116, 0, 0, 212, 213, 5, 114, 0, 0, 213, 214, 5, 117, 0, 0, 214, 215, 5,
		101, 0, 0, 215, 70, 1, 0, 0, 0, 216, 217, 5, 102, 0, 0, 217, 218, 5, 97,
		0, 0, 218, 219, 5, 108, 0, 0, 219, 220, 5, 115, 0, 0, 220, 221, 5, 101,
		0, 0, 221, 72, 1, 0, 0, 0, 222, 223, 5, 110, 0, 0, 223, 224, 5, 105, 0,
		0, 224, 225, 5, 108, 0, 0, 225, 74, 1, 0, 0, 0, 226, 227, 5, 40, 0, 0,
		227, 76, 1, 0, 0, 0, 228, 229, 5, 41, 0, 0, 229, 78, 1, 0, 0, 0, 230, 231,
		5, 123, 0, 0, 231, 80, 1, 0, 0, 0, 232, 233, 5, 125, 0, 0, 233, 82, 1,
		0, 0, 0, 234, 238, 7, 0, 0, 0, 235, 237, 7, 1, 0, 0, 236, 235, 1, 0, 0,
		0, 237, 240, 1, 0, 0, 0, 238, 236, 1, 0, 0, 0, 238, 239, 1, 0, 0, 0, 239,
		84, 1, 0, 0, 0, 240, 238, 1, 0, 0, 0, 241, 243, 7, 2, 0, 0, 242, 241, 1,
		0, 0, 0, 243, 244, 1, 0, 0, 0, 244, 242, 1, 0, 0, 0, 244, 245, 1, 0, 0,
		0, 245, 86, 1, 0, 0, 0, 246, 248, 7, 2, 0, 0, 247, 246, 1, 0, 0, 0, 248,
		249, 1, 0, 0, 0, 249, 247, 1, 0, 0, 0, 249, 250, 1, 0, 0, 0, 250, 251,
		1, 0, 0, 0, 251, 255, 5, 46, 0, 0, 252, 254, 7, 2, 0, 0, 253, 252, 1, 0,
		0, 0, 254, 257, 1, 0, 0, 0, 255, 253, 1, 0, 0, 0, 255, 256, 1, 0, 0, 0,
		256, 265, 1, 0, 0, 0, 257, 255, 1, 0, 0, 0, 258, 260, 5, 46, 0, 0, 259,
		261, 7, 2, 0, 0, 260, 259, 1, 0, 0, 0, 261, 262, 1, 0, 0, 0, 262, 260,
		1, 0, 0, 0, 262, 263, 1, 0, 0, 0, 263, 265, 1, 0, 0, 0, 264, 247, 1, 0,
		0, 0, 264, 258, 1, 0, 0, 0, 265, 88, 1, 0, 0, 0, 266, 272, 5, 34, 0, 0,
		267, 271, 8, 3, 0, 0, 268, 269, 5, 34, 0, 0, 269, 271, 5, 34, 0, 0, 270,
		267, 1, 0, 0, 0, 270, 268, 1, 0, 0, 0, 271, 274, 1, 0, 0, 0, 272, 270,
		1, 0, 0, 0, 272, 273, 1, 0, 0, 0, 273, 275, 1, 0, 0, 0, 274, 272, 1, 0,
		0, 0, 275, 276, 5, 34, 0, 0, 276, 90, 1, 0, 0, 0, 277, 278, 7, 4, 0, 0,
		278, 279, 1, 0, 0, 0, 279, 280, 6, 45, 0, 0, 280, 92, 1, 0, 0, 0, 9, 0,
		238, 244, 249, 255, 262, 264, 270, 272, 1, 6, 0, 0,
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
	GoEventerLexerIF           = 13
	GoEventerLexerELSE         = 14
	GoEventerLexerCHAN         = 15
	GoEventerLexerNOTIFY       = 16
	GoEventerLexerSTRUCT       = 17
	GoEventerLexerSTRING_TYPE  = 18
	GoEventerLexerINTEGER_TYPE = 19
	GoEventerLexerOR           = 20
	GoEventerLexerAND          = 21
	GoEventerLexerEQ           = 22
	GoEventerLexerNEQ          = 23
	GoEventerLexerGT           = 24
	GoEventerLexerLT           = 25
	GoEventerLexerGTEQ         = 26
	GoEventerLexerLTEQ         = 27
	GoEventerLexerPLUS         = 28
	GoEventerLexerMINUS        = 29
	GoEventerLexerMULT         = 30
	GoEventerLexerDIV          = 31
	GoEventerLexerMOD          = 32
	GoEventerLexerPOW          = 33
	GoEventerLexerNOT          = 34
	GoEventerLexerTRUE         = 35
	GoEventerLexerFALSE        = 36
	GoEventerLexerNIL          = 37
	GoEventerLexerPARAN_OPEN   = 38
	GoEventerLexerPARAN_CLOSE  = 39
	GoEventerLexerBARACE_OPEN  = 40
	GoEventerLexerBARACE_CLOSE = 41
	GoEventerLexerID           = 42
	GoEventerLexerINT          = 43
	GoEventerLexerFLOAT        = 44
	GoEventerLexerSTRING       = 45
	GoEventerLexerSPACE        = 46
)
