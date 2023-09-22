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
		"", "','", "':='", "'='", "';'", "':'", "'for'", "'while'", "'range'",
		"'func'", "'on'", "'return'", "'chan'", "'notify'", "'string'", "'int'",
		"'||'", "'&&'", "'=='", "'!='", "'>'", "'<'", "'>='", "'<='", "'+'",
		"'-'", "'*'", "'/'", "'%'", "'^'", "'!'", "'true'", "'false'", "'nil'",
		"'('", "')'", "'{'", "'}'",
	}
	staticData.SymbolicNames = []string{
		"", "", "ASSIGN", "UPDATE", "SCOL", "COL", "FOR", "WHILE", "RANGE",
		"FUNC", "ON", "RETURN", "CHAN", "NOTIFY", "STRING_TYPE", "INTEGER_TYPE",
		"OR", "AND", "EQ", "NEQ", "GT", "LT", "GTEQ", "LTEQ", "PLUS", "MINUS",
		"MULT", "DIV", "MOD", "POW", "NOT", "TRUE", "FALSE", "NIL", "PARAN_OPEN",
		"PARAN_CLOSE", "BARACE_OPEN", "BARACE_CLOSE", "ID", "INT", "FLOAT",
		"STRING", "SPACE",
	}
	staticData.RuleNames = []string{
		"T__0", "ASSIGN", "UPDATE", "SCOL", "COL", "FOR", "WHILE", "RANGE",
		"FUNC", "ON", "RETURN", "CHAN", "NOTIFY", "STRING_TYPE", "INTEGER_TYPE",
		"OR", "AND", "EQ", "NEQ", "GT", "LT", "GTEQ", "LTEQ", "PLUS", "MINUS",
		"MULT", "DIV", "MOD", "POW", "NOT", "TRUE", "FALSE", "NIL", "PARAN_OPEN",
		"PARAN_CLOSE", "BARACE_OPEN", "BARACE_CLOSE", "ID", "INT", "FLOAT",
		"STRING", "SPACE",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 42, 256, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36,
		7, 36, 2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7,
		41, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1,
		5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1,
		7, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1,
		10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 11,
		1, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1,
		13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 15, 1, 15,
		1, 15, 1, 16, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 1, 18, 1, 18, 1, 18, 1,
		19, 1, 19, 1, 20, 1, 20, 1, 21, 1, 21, 1, 21, 1, 22, 1, 22, 1, 22, 1, 23,
		1, 23, 1, 24, 1, 24, 1, 25, 1, 25, 1, 26, 1, 26, 1, 27, 1, 27, 1, 28, 1,
		28, 1, 29, 1, 29, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 31, 1, 31, 1, 31,
		1, 31, 1, 31, 1, 31, 1, 32, 1, 32, 1, 32, 1, 32, 1, 33, 1, 33, 1, 34, 1,
		34, 1, 35, 1, 35, 1, 36, 1, 36, 1, 37, 1, 37, 5, 37, 212, 8, 37, 10, 37,
		12, 37, 215, 9, 37, 1, 38, 4, 38, 218, 8, 38, 11, 38, 12, 38, 219, 1, 39,
		4, 39, 223, 8, 39, 11, 39, 12, 39, 224, 1, 39, 1, 39, 5, 39, 229, 8, 39,
		10, 39, 12, 39, 232, 9, 39, 1, 39, 1, 39, 4, 39, 236, 8, 39, 11, 39, 12,
		39, 237, 3, 39, 240, 8, 39, 1, 40, 1, 40, 1, 40, 1, 40, 5, 40, 246, 8,
		40, 10, 40, 12, 40, 249, 9, 40, 1, 40, 1, 40, 1, 41, 1, 41, 1, 41, 1, 41,
		0, 0, 42, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19,
		10, 21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17, 35, 18, 37,
		19, 39, 20, 41, 21, 43, 22, 45, 23, 47, 24, 49, 25, 51, 26, 53, 27, 55,
		28, 57, 29, 59, 30, 61, 31, 63, 32, 65, 33, 67, 34, 69, 35, 71, 36, 73,
		37, 75, 38, 77, 39, 79, 40, 81, 41, 83, 42, 1, 0, 5, 3, 0, 65, 90, 95,
		95, 97, 122, 4, 0, 48, 57, 65, 90, 95, 95, 97, 122, 1, 0, 48, 57, 3, 0,
		10, 10, 13, 13, 34, 34, 3, 0, 9, 10, 13, 13, 32, 32, 263, 0, 1, 1, 0, 0,
		0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0,
		0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0,
		0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1,
		0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33,
		1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0,
		41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1, 0, 0, 0,
		0, 49, 1, 0, 0, 0, 0, 51, 1, 0, 0, 0, 0, 53, 1, 0, 0, 0, 0, 55, 1, 0, 0,
		0, 0, 57, 1, 0, 0, 0, 0, 59, 1, 0, 0, 0, 0, 61, 1, 0, 0, 0, 0, 63, 1, 0,
		0, 0, 0, 65, 1, 0, 0, 0, 0, 67, 1, 0, 0, 0, 0, 69, 1, 0, 0, 0, 0, 71, 1,
		0, 0, 0, 0, 73, 1, 0, 0, 0, 0, 75, 1, 0, 0, 0, 0, 77, 1, 0, 0, 0, 0, 79,
		1, 0, 0, 0, 0, 81, 1, 0, 0, 0, 0, 83, 1, 0, 0, 0, 1, 85, 1, 0, 0, 0, 3,
		87, 1, 0, 0, 0, 5, 90, 1, 0, 0, 0, 7, 92, 1, 0, 0, 0, 9, 94, 1, 0, 0, 0,
		11, 96, 1, 0, 0, 0, 13, 100, 1, 0, 0, 0, 15, 106, 1, 0, 0, 0, 17, 112,
		1, 0, 0, 0, 19, 117, 1, 0, 0, 0, 21, 120, 1, 0, 0, 0, 23, 127, 1, 0, 0,
		0, 25, 132, 1, 0, 0, 0, 27, 139, 1, 0, 0, 0, 29, 146, 1, 0, 0, 0, 31, 150,
		1, 0, 0, 0, 33, 153, 1, 0, 0, 0, 35, 156, 1, 0, 0, 0, 37, 159, 1, 0, 0,
		0, 39, 162, 1, 0, 0, 0, 41, 164, 1, 0, 0, 0, 43, 166, 1, 0, 0, 0, 45, 169,
		1, 0, 0, 0, 47, 172, 1, 0, 0, 0, 49, 174, 1, 0, 0, 0, 51, 176, 1, 0, 0,
		0, 53, 178, 1, 0, 0, 0, 55, 180, 1, 0, 0, 0, 57, 182, 1, 0, 0, 0, 59, 184,
		1, 0, 0, 0, 61, 186, 1, 0, 0, 0, 63, 191, 1, 0, 0, 0, 65, 197, 1, 0, 0,
		0, 67, 201, 1, 0, 0, 0, 69, 203, 1, 0, 0, 0, 71, 205, 1, 0, 0, 0, 73, 207,
		1, 0, 0, 0, 75, 209, 1, 0, 0, 0, 77, 217, 1, 0, 0, 0, 79, 239, 1, 0, 0,
		0, 81, 241, 1, 0, 0, 0, 83, 252, 1, 0, 0, 0, 85, 86, 5, 44, 0, 0, 86, 2,
		1, 0, 0, 0, 87, 88, 5, 58, 0, 0, 88, 89, 5, 61, 0, 0, 89, 4, 1, 0, 0, 0,
		90, 91, 5, 61, 0, 0, 91, 6, 1, 0, 0, 0, 92, 93, 5, 59, 0, 0, 93, 8, 1,
		0, 0, 0, 94, 95, 5, 58, 0, 0, 95, 10, 1, 0, 0, 0, 96, 97, 5, 102, 0, 0,
		97, 98, 5, 111, 0, 0, 98, 99, 5, 114, 0, 0, 99, 12, 1, 0, 0, 0, 100, 101,
		5, 119, 0, 0, 101, 102, 5, 104, 0, 0, 102, 103, 5, 105, 0, 0, 103, 104,
		5, 108, 0, 0, 104, 105, 5, 101, 0, 0, 105, 14, 1, 0, 0, 0, 106, 107, 5,
		114, 0, 0, 107, 108, 5, 97, 0, 0, 108, 109, 5, 110, 0, 0, 109, 110, 5,
		103, 0, 0, 110, 111, 5, 101, 0, 0, 111, 16, 1, 0, 0, 0, 112, 113, 5, 102,
		0, 0, 113, 114, 5, 117, 0, 0, 114, 115, 5, 110, 0, 0, 115, 116, 5, 99,
		0, 0, 116, 18, 1, 0, 0, 0, 117, 118, 5, 111, 0, 0, 118, 119, 5, 110, 0,
		0, 119, 20, 1, 0, 0, 0, 120, 121, 5, 114, 0, 0, 121, 122, 5, 101, 0, 0,
		122, 123, 5, 116, 0, 0, 123, 124, 5, 117, 0, 0, 124, 125, 5, 114, 0, 0,
		125, 126, 5, 110, 0, 0, 126, 22, 1, 0, 0, 0, 127, 128, 5, 99, 0, 0, 128,
		129, 5, 104, 0, 0, 129, 130, 5, 97, 0, 0, 130, 131, 5, 110, 0, 0, 131,
		24, 1, 0, 0, 0, 132, 133, 5, 110, 0, 0, 133, 134, 5, 111, 0, 0, 134, 135,
		5, 116, 0, 0, 135, 136, 5, 105, 0, 0, 136, 137, 5, 102, 0, 0, 137, 138,
		5, 121, 0, 0, 138, 26, 1, 0, 0, 0, 139, 140, 5, 115, 0, 0, 140, 141, 5,
		116, 0, 0, 141, 142, 5, 114, 0, 0, 142, 143, 5, 105, 0, 0, 143, 144, 5,
		110, 0, 0, 144, 145, 5, 103, 0, 0, 145, 28, 1, 0, 0, 0, 146, 147, 5, 105,
		0, 0, 147, 148, 5, 110, 0, 0, 148, 149, 5, 116, 0, 0, 149, 30, 1, 0, 0,
		0, 150, 151, 5, 124, 0, 0, 151, 152, 5, 124, 0, 0, 152, 32, 1, 0, 0, 0,
		153, 154, 5, 38, 0, 0, 154, 155, 5, 38, 0, 0, 155, 34, 1, 0, 0, 0, 156,
		157, 5, 61, 0, 0, 157, 158, 5, 61, 0, 0, 158, 36, 1, 0, 0, 0, 159, 160,
		5, 33, 0, 0, 160, 161, 5, 61, 0, 0, 161, 38, 1, 0, 0, 0, 162, 163, 5, 62,
		0, 0, 163, 40, 1, 0, 0, 0, 164, 165, 5, 60, 0, 0, 165, 42, 1, 0, 0, 0,
		166, 167, 5, 62, 0, 0, 167, 168, 5, 61, 0, 0, 168, 44, 1, 0, 0, 0, 169,
		170, 5, 60, 0, 0, 170, 171, 5, 61, 0, 0, 171, 46, 1, 0, 0, 0, 172, 173,
		5, 43, 0, 0, 173, 48, 1, 0, 0, 0, 174, 175, 5, 45, 0, 0, 175, 50, 1, 0,
		0, 0, 176, 177, 5, 42, 0, 0, 177, 52, 1, 0, 0, 0, 178, 179, 5, 47, 0, 0,
		179, 54, 1, 0, 0, 0, 180, 181, 5, 37, 0, 0, 181, 56, 1, 0, 0, 0, 182, 183,
		5, 94, 0, 0, 183, 58, 1, 0, 0, 0, 184, 185, 5, 33, 0, 0, 185, 60, 1, 0,
		0, 0, 186, 187, 5, 116, 0, 0, 187, 188, 5, 114, 0, 0, 188, 189, 5, 117,
		0, 0, 189, 190, 5, 101, 0, 0, 190, 62, 1, 0, 0, 0, 191, 192, 5, 102, 0,
		0, 192, 193, 5, 97, 0, 0, 193, 194, 5, 108, 0, 0, 194, 195, 5, 115, 0,
		0, 195, 196, 5, 101, 0, 0, 196, 64, 1, 0, 0, 0, 197, 198, 5, 110, 0, 0,
		198, 199, 5, 105, 0, 0, 199, 200, 5, 108, 0, 0, 200, 66, 1, 0, 0, 0, 201,
		202, 5, 40, 0, 0, 202, 68, 1, 0, 0, 0, 203, 204, 5, 41, 0, 0, 204, 70,
		1, 0, 0, 0, 205, 206, 5, 123, 0, 0, 206, 72, 1, 0, 0, 0, 207, 208, 5, 125,
		0, 0, 208, 74, 1, 0, 0, 0, 209, 213, 7, 0, 0, 0, 210, 212, 7, 1, 0, 0,
		211, 210, 1, 0, 0, 0, 212, 215, 1, 0, 0, 0, 213, 211, 1, 0, 0, 0, 213,
		214, 1, 0, 0, 0, 214, 76, 1, 0, 0, 0, 215, 213, 1, 0, 0, 0, 216, 218, 7,
		2, 0, 0, 217, 216, 1, 0, 0, 0, 218, 219, 1, 0, 0, 0, 219, 217, 1, 0, 0,
		0, 219, 220, 1, 0, 0, 0, 220, 78, 1, 0, 0, 0, 221, 223, 7, 2, 0, 0, 222,
		221, 1, 0, 0, 0, 223, 224, 1, 0, 0, 0, 224, 222, 1, 0, 0, 0, 224, 225,
		1, 0, 0, 0, 225, 226, 1, 0, 0, 0, 226, 230, 5, 46, 0, 0, 227, 229, 7, 2,
		0, 0, 228, 227, 1, 0, 0, 0, 229, 232, 1, 0, 0, 0, 230, 228, 1, 0, 0, 0,
		230, 231, 1, 0, 0, 0, 231, 240, 1, 0, 0, 0, 232, 230, 1, 0, 0, 0, 233,
		235, 5, 46, 0, 0, 234, 236, 7, 2, 0, 0, 235, 234, 1, 0, 0, 0, 236, 237,
		1, 0, 0, 0, 237, 235, 1, 0, 0, 0, 237, 238, 1, 0, 0, 0, 238, 240, 1, 0,
		0, 0, 239, 222, 1, 0, 0, 0, 239, 233, 1, 0, 0, 0, 240, 80, 1, 0, 0, 0,
		241, 247, 5, 34, 0, 0, 242, 246, 8, 3, 0, 0, 243, 244, 5, 34, 0, 0, 244,
		246, 5, 34, 0, 0, 245, 242, 1, 0, 0, 0, 245, 243, 1, 0, 0, 0, 246, 249,
		1, 0, 0, 0, 247, 245, 1, 0, 0, 0, 247, 248, 1, 0, 0, 0, 248, 250, 1, 0,
		0, 0, 249, 247, 1, 0, 0, 0, 250, 251, 5, 34, 0, 0, 251, 82, 1, 0, 0, 0,
		252, 253, 7, 4, 0, 0, 253, 254, 1, 0, 0, 0, 254, 255, 6, 41, 0, 0, 255,
		84, 1, 0, 0, 0, 9, 0, 213, 219, 224, 230, 237, 239, 245, 247, 1, 6, 0,
		0,
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
	GoEventerLexerFOR          = 6
	GoEventerLexerWHILE        = 7
	GoEventerLexerRANGE        = 8
	GoEventerLexerFUNC         = 9
	GoEventerLexerON           = 10
	GoEventerLexerRETURN       = 11
	GoEventerLexerCHAN         = 12
	GoEventerLexerNOTIFY       = 13
	GoEventerLexerSTRING_TYPE  = 14
	GoEventerLexerINTEGER_TYPE = 15
	GoEventerLexerOR           = 16
	GoEventerLexerAND          = 17
	GoEventerLexerEQ           = 18
	GoEventerLexerNEQ          = 19
	GoEventerLexerGT           = 20
	GoEventerLexerLT           = 21
	GoEventerLexerGTEQ         = 22
	GoEventerLexerLTEQ         = 23
	GoEventerLexerPLUS         = 24
	GoEventerLexerMINUS        = 25
	GoEventerLexerMULT         = 26
	GoEventerLexerDIV          = 27
	GoEventerLexerMOD          = 28
	GoEventerLexerPOW          = 29
	GoEventerLexerNOT          = 30
	GoEventerLexerTRUE         = 31
	GoEventerLexerFALSE        = 32
	GoEventerLexerNIL          = 33
	GoEventerLexerPARAN_OPEN   = 34
	GoEventerLexerPARAN_CLOSE  = 35
	GoEventerLexerBARACE_OPEN  = 36
	GoEventerLexerBARACE_CLOSE = 37
	GoEventerLexerID           = 38
	GoEventerLexerINT          = 39
	GoEventerLexerFLOAT        = 40
	GoEventerLexerSTRING       = 41
	GoEventerLexerSPACE        = 42
)
