// Code generated from GoEventer.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parsing

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
		"'func'", "'on'", "'return'", "'||'", "'&&'", "'=='", "'!='", "'>'",
		"'<'", "'>='", "'<='", "'+'", "'-'", "'*'", "'/'", "'%'", "'^'", "'!'",
		"'true'", "'false'", "'nil'", "'('", "')'", "'{'", "'}'",
	}
	staticData.SymbolicNames = []string{
		"", "", "ASSIGN", "UPDATE", "SCOL", "COL", "FOR", "WHILE", "RANGE",
		"FUNC", "ON", "RETURN", "OR", "AND", "EQ", "NEQ", "GT", "LT", "GTEQ",
		"LTEQ", "PLUS", "MINUS", "MULT", "DIV", "MOD", "POW", "NOT", "TRUE",
		"FALSE", "NIL", "PARAN_OPEN", "PARAN_CLOSE", "BARACE_OPEN", "BARACE_CLOSE",
		"ID", "INT", "FLOAT", "STRING", "SPACE",
	}
	staticData.RuleNames = []string{
		"T__0", "ASSIGN", "UPDATE", "SCOL", "COL", "FOR", "WHILE", "RANGE",
		"FUNC", "ON", "RETURN", "OR", "AND", "EQ", "NEQ", "GT", "LT", "GTEQ",
		"LTEQ", "PLUS", "MINUS", "MULT", "DIV", "MOD", "POW", "NOT", "TRUE",
		"FALSE", "NIL", "PARAN_OPEN", "PARAN_CLOSE", "BARACE_OPEN", "BARACE_CLOSE",
		"ID", "INT", "FLOAT", "STRING", "SPACE",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 38, 225, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36,
		7, 36, 2, 37, 7, 37, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1,
		3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1,
		6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1,
		9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 11,
		1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1,
		14, 1, 15, 1, 15, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 1, 18, 1, 18, 1, 18,
		1, 19, 1, 19, 1, 20, 1, 20, 1, 21, 1, 21, 1, 22, 1, 22, 1, 23, 1, 23, 1,
		24, 1, 24, 1, 25, 1, 25, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 27, 1, 27,
		1, 27, 1, 27, 1, 27, 1, 27, 1, 28, 1, 28, 1, 28, 1, 28, 1, 29, 1, 29, 1,
		30, 1, 30, 1, 31, 1, 31, 1, 32, 1, 32, 1, 33, 1, 33, 5, 33, 181, 8, 33,
		10, 33, 12, 33, 184, 9, 33, 1, 34, 4, 34, 187, 8, 34, 11, 34, 12, 34, 188,
		1, 35, 4, 35, 192, 8, 35, 11, 35, 12, 35, 193, 1, 35, 1, 35, 5, 35, 198,
		8, 35, 10, 35, 12, 35, 201, 9, 35, 1, 35, 1, 35, 4, 35, 205, 8, 35, 11,
		35, 12, 35, 206, 3, 35, 209, 8, 35, 1, 36, 1, 36, 1, 36, 1, 36, 5, 36,
		215, 8, 36, 10, 36, 12, 36, 218, 9, 36, 1, 36, 1, 36, 1, 37, 1, 37, 1,
		37, 1, 37, 0, 0, 38, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8,
		17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17,
		35, 18, 37, 19, 39, 20, 41, 21, 43, 22, 45, 23, 47, 24, 49, 25, 51, 26,
		53, 27, 55, 28, 57, 29, 59, 30, 61, 31, 63, 32, 65, 33, 67, 34, 69, 35,
		71, 36, 73, 37, 75, 38, 1, 0, 5, 3, 0, 65, 90, 95, 95, 97, 122, 4, 0, 48,
		57, 65, 90, 95, 95, 97, 122, 1, 0, 48, 57, 3, 0, 10, 10, 13, 13, 34, 34,
		3, 0, 9, 10, 13, 13, 32, 32, 232, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0,
		5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0,
		13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0,
		0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0,
		0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0,
		0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1,
		0, 0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1, 0, 0, 0, 0, 49, 1, 0, 0, 0, 0, 51,
		1, 0, 0, 0, 0, 53, 1, 0, 0, 0, 0, 55, 1, 0, 0, 0, 0, 57, 1, 0, 0, 0, 0,
		59, 1, 0, 0, 0, 0, 61, 1, 0, 0, 0, 0, 63, 1, 0, 0, 0, 0, 65, 1, 0, 0, 0,
		0, 67, 1, 0, 0, 0, 0, 69, 1, 0, 0, 0, 0, 71, 1, 0, 0, 0, 0, 73, 1, 0, 0,
		0, 0, 75, 1, 0, 0, 0, 1, 77, 1, 0, 0, 0, 3, 79, 1, 0, 0, 0, 5, 82, 1, 0,
		0, 0, 7, 84, 1, 0, 0, 0, 9, 86, 1, 0, 0, 0, 11, 88, 1, 0, 0, 0, 13, 92,
		1, 0, 0, 0, 15, 98, 1, 0, 0, 0, 17, 104, 1, 0, 0, 0, 19, 109, 1, 0, 0,
		0, 21, 112, 1, 0, 0, 0, 23, 119, 1, 0, 0, 0, 25, 122, 1, 0, 0, 0, 27, 125,
		1, 0, 0, 0, 29, 128, 1, 0, 0, 0, 31, 131, 1, 0, 0, 0, 33, 133, 1, 0, 0,
		0, 35, 135, 1, 0, 0, 0, 37, 138, 1, 0, 0, 0, 39, 141, 1, 0, 0, 0, 41, 143,
		1, 0, 0, 0, 43, 145, 1, 0, 0, 0, 45, 147, 1, 0, 0, 0, 47, 149, 1, 0, 0,
		0, 49, 151, 1, 0, 0, 0, 51, 153, 1, 0, 0, 0, 53, 155, 1, 0, 0, 0, 55, 160,
		1, 0, 0, 0, 57, 166, 1, 0, 0, 0, 59, 170, 1, 0, 0, 0, 61, 172, 1, 0, 0,
		0, 63, 174, 1, 0, 0, 0, 65, 176, 1, 0, 0, 0, 67, 178, 1, 0, 0, 0, 69, 186,
		1, 0, 0, 0, 71, 208, 1, 0, 0, 0, 73, 210, 1, 0, 0, 0, 75, 221, 1, 0, 0,
		0, 77, 78, 5, 44, 0, 0, 78, 2, 1, 0, 0, 0, 79, 80, 5, 58, 0, 0, 80, 81,
		5, 61, 0, 0, 81, 4, 1, 0, 0, 0, 82, 83, 5, 61, 0, 0, 83, 6, 1, 0, 0, 0,
		84, 85, 5, 59, 0, 0, 85, 8, 1, 0, 0, 0, 86, 87, 5, 58, 0, 0, 87, 10, 1,
		0, 0, 0, 88, 89, 5, 102, 0, 0, 89, 90, 5, 111, 0, 0, 90, 91, 5, 114, 0,
		0, 91, 12, 1, 0, 0, 0, 92, 93, 5, 119, 0, 0, 93, 94, 5, 104, 0, 0, 94,
		95, 5, 105, 0, 0, 95, 96, 5, 108, 0, 0, 96, 97, 5, 101, 0, 0, 97, 14, 1,
		0, 0, 0, 98, 99, 5, 114, 0, 0, 99, 100, 5, 97, 0, 0, 100, 101, 5, 110,
		0, 0, 101, 102, 5, 103, 0, 0, 102, 103, 5, 101, 0, 0, 103, 16, 1, 0, 0,
		0, 104, 105, 5, 102, 0, 0, 105, 106, 5, 117, 0, 0, 106, 107, 5, 110, 0,
		0, 107, 108, 5, 99, 0, 0, 108, 18, 1, 0, 0, 0, 109, 110, 5, 111, 0, 0,
		110, 111, 5, 110, 0, 0, 111, 20, 1, 0, 0, 0, 112, 113, 5, 114, 0, 0, 113,
		114, 5, 101, 0, 0, 114, 115, 5, 116, 0, 0, 115, 116, 5, 117, 0, 0, 116,
		117, 5, 114, 0, 0, 117, 118, 5, 110, 0, 0, 118, 22, 1, 0, 0, 0, 119, 120,
		5, 124, 0, 0, 120, 121, 5, 124, 0, 0, 121, 24, 1, 0, 0, 0, 122, 123, 5,
		38, 0, 0, 123, 124, 5, 38, 0, 0, 124, 26, 1, 0, 0, 0, 125, 126, 5, 61,
		0, 0, 126, 127, 5, 61, 0, 0, 127, 28, 1, 0, 0, 0, 128, 129, 5, 33, 0, 0,
		129, 130, 5, 61, 0, 0, 130, 30, 1, 0, 0, 0, 131, 132, 5, 62, 0, 0, 132,
		32, 1, 0, 0, 0, 133, 134, 5, 60, 0, 0, 134, 34, 1, 0, 0, 0, 135, 136, 5,
		62, 0, 0, 136, 137, 5, 61, 0, 0, 137, 36, 1, 0, 0, 0, 138, 139, 5, 60,
		0, 0, 139, 140, 5, 61, 0, 0, 140, 38, 1, 0, 0, 0, 141, 142, 5, 43, 0, 0,
		142, 40, 1, 0, 0, 0, 143, 144, 5, 45, 0, 0, 144, 42, 1, 0, 0, 0, 145, 146,
		5, 42, 0, 0, 146, 44, 1, 0, 0, 0, 147, 148, 5, 47, 0, 0, 148, 46, 1, 0,
		0, 0, 149, 150, 5, 37, 0, 0, 150, 48, 1, 0, 0, 0, 151, 152, 5, 94, 0, 0,
		152, 50, 1, 0, 0, 0, 153, 154, 5, 33, 0, 0, 154, 52, 1, 0, 0, 0, 155, 156,
		5, 116, 0, 0, 156, 157, 5, 114, 0, 0, 157, 158, 5, 117, 0, 0, 158, 159,
		5, 101, 0, 0, 159, 54, 1, 0, 0, 0, 160, 161, 5, 102, 0, 0, 161, 162, 5,
		97, 0, 0, 162, 163, 5, 108, 0, 0, 163, 164, 5, 115, 0, 0, 164, 165, 5,
		101, 0, 0, 165, 56, 1, 0, 0, 0, 166, 167, 5, 110, 0, 0, 167, 168, 5, 105,
		0, 0, 168, 169, 5, 108, 0, 0, 169, 58, 1, 0, 0, 0, 170, 171, 5, 40, 0,
		0, 171, 60, 1, 0, 0, 0, 172, 173, 5, 41, 0, 0, 173, 62, 1, 0, 0, 0, 174,
		175, 5, 123, 0, 0, 175, 64, 1, 0, 0, 0, 176, 177, 5, 125, 0, 0, 177, 66,
		1, 0, 0, 0, 178, 182, 7, 0, 0, 0, 179, 181, 7, 1, 0, 0, 180, 179, 1, 0,
		0, 0, 181, 184, 1, 0, 0, 0, 182, 180, 1, 0, 0, 0, 182, 183, 1, 0, 0, 0,
		183, 68, 1, 0, 0, 0, 184, 182, 1, 0, 0, 0, 185, 187, 7, 2, 0, 0, 186, 185,
		1, 0, 0, 0, 187, 188, 1, 0, 0, 0, 188, 186, 1, 0, 0, 0, 188, 189, 1, 0,
		0, 0, 189, 70, 1, 0, 0, 0, 190, 192, 7, 2, 0, 0, 191, 190, 1, 0, 0, 0,
		192, 193, 1, 0, 0, 0, 193, 191, 1, 0, 0, 0, 193, 194, 1, 0, 0, 0, 194,
		195, 1, 0, 0, 0, 195, 199, 5, 46, 0, 0, 196, 198, 7, 2, 0, 0, 197, 196,
		1, 0, 0, 0, 198, 201, 1, 0, 0, 0, 199, 197, 1, 0, 0, 0, 199, 200, 1, 0,
		0, 0, 200, 209, 1, 0, 0, 0, 201, 199, 1, 0, 0, 0, 202, 204, 5, 46, 0, 0,
		203, 205, 7, 2, 0, 0, 204, 203, 1, 0, 0, 0, 205, 206, 1, 0, 0, 0, 206,
		204, 1, 0, 0, 0, 206, 207, 1, 0, 0, 0, 207, 209, 1, 0, 0, 0, 208, 191,
		1, 0, 0, 0, 208, 202, 1, 0, 0, 0, 209, 72, 1, 0, 0, 0, 210, 216, 5, 34,
		0, 0, 211, 215, 8, 3, 0, 0, 212, 213, 5, 34, 0, 0, 213, 215, 5, 34, 0,
		0, 214, 211, 1, 0, 0, 0, 214, 212, 1, 0, 0, 0, 215, 218, 1, 0, 0, 0, 216,
		214, 1, 0, 0, 0, 216, 217, 1, 0, 0, 0, 217, 219, 1, 0, 0, 0, 218, 216,
		1, 0, 0, 0, 219, 220, 5, 34, 0, 0, 220, 74, 1, 0, 0, 0, 221, 222, 7, 4,
		0, 0, 222, 223, 1, 0, 0, 0, 223, 224, 6, 37, 0, 0, 224, 76, 1, 0, 0, 0,
		9, 0, 182, 188, 193, 199, 206, 208, 214, 216, 1, 6, 0, 0,
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
	GoEventerLexerOR           = 12
	GoEventerLexerAND          = 13
	GoEventerLexerEQ           = 14
	GoEventerLexerNEQ          = 15
	GoEventerLexerGT           = 16
	GoEventerLexerLT           = 17
	GoEventerLexerGTEQ         = 18
	GoEventerLexerLTEQ         = 19
	GoEventerLexerPLUS         = 20
	GoEventerLexerMINUS        = 21
	GoEventerLexerMULT         = 22
	GoEventerLexerDIV          = 23
	GoEventerLexerMOD          = 24
	GoEventerLexerPOW          = 25
	GoEventerLexerNOT          = 26
	GoEventerLexerTRUE         = 27
	GoEventerLexerFALSE        = 28
	GoEventerLexerNIL          = 29
	GoEventerLexerPARAN_OPEN   = 30
	GoEventerLexerPARAN_CLOSE  = 31
	GoEventerLexerBARACE_OPEN  = 32
	GoEventerLexerBARACE_CLOSE = 33
	GoEventerLexerID           = 34
	GoEventerLexerINT          = 35
	GoEventerLexerFLOAT        = 36
	GoEventerLexerSTRING       = 37
	GoEventerLexerSPACE        = 38
)
