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
		"", "','", "':='", "'='", "';'", "':'", "'for'", "'range'", "'func'",
		"'on'", "'return'", "'||'", "'&&'", "'=='", "'!='", "'>'", "'<'", "'>='",
		"'<='", "'+'", "'-'", "'*'", "'/'", "'%'", "'^'", "'!'", "'true'", "'false'",
		"'nil'", "'('", "')'", "'{'", "'}'",
	}
	staticData.SymbolicNames = []string{
		"", "", "ASSIGN", "UPDATE", "SCOL", "COL", "FOR", "RANGE", "FUNC", "ON",
		"RETURN", "OR", "AND", "EQ", "NEQ", "GT", "LT", "GTEQ", "LTEQ", "PLUS",
		"MINUS", "MULT", "DIV", "MOD", "POW", "NOT", "TRUE", "FALSE", "NIL",
		"PARAN_OPEN", "PARAN_CLOSE", "BARACE_OPEN", "BARACE_CLOSE", "ID", "INT",
		"FLOAT", "STRING", "SPACE",
	}
	staticData.RuleNames = []string{
		"T__0", "ASSIGN", "UPDATE", "SCOL", "COL", "FOR", "RANGE", "FUNC", "ON",
		"RETURN", "OR", "AND", "EQ", "NEQ", "GT", "LT", "GTEQ", "LTEQ", "PLUS",
		"MINUS", "MULT", "DIV", "MOD", "POW", "NOT", "TRUE", "FALSE", "NIL",
		"PARAN_OPEN", "PARAN_CLOSE", "BARACE_OPEN", "BARACE_CLOSE", "ID", "INT",
		"FLOAT", "STRING", "SPACE",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 37, 217, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36,
		7, 36, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4,
		1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7,
		1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9,
		1, 9, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1,
		13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 1, 17,
		1, 17, 1, 17, 1, 18, 1, 18, 1, 19, 1, 19, 1, 20, 1, 20, 1, 21, 1, 21, 1,
		22, 1, 22, 1, 23, 1, 23, 1, 24, 1, 24, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25,
		1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 27, 1, 27, 1, 27, 1, 27, 1,
		28, 1, 28, 1, 29, 1, 29, 1, 30, 1, 30, 1, 31, 1, 31, 1, 32, 1, 32, 5, 32,
		173, 8, 32, 10, 32, 12, 32, 176, 9, 32, 1, 33, 4, 33, 179, 8, 33, 11, 33,
		12, 33, 180, 1, 34, 4, 34, 184, 8, 34, 11, 34, 12, 34, 185, 1, 34, 1, 34,
		5, 34, 190, 8, 34, 10, 34, 12, 34, 193, 9, 34, 1, 34, 1, 34, 4, 34, 197,
		8, 34, 11, 34, 12, 34, 198, 3, 34, 201, 8, 34, 1, 35, 1, 35, 1, 35, 1,
		35, 5, 35, 207, 8, 35, 10, 35, 12, 35, 210, 9, 35, 1, 35, 1, 35, 1, 36,
		1, 36, 1, 36, 1, 36, 0, 0, 37, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13,
		7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16,
		33, 17, 35, 18, 37, 19, 39, 20, 41, 21, 43, 22, 45, 23, 47, 24, 49, 25,
		51, 26, 53, 27, 55, 28, 57, 29, 59, 30, 61, 31, 63, 32, 65, 33, 67, 34,
		69, 35, 71, 36, 73, 37, 1, 0, 5, 3, 0, 65, 90, 95, 95, 97, 122, 4, 0, 48,
		57, 65, 90, 95, 95, 97, 122, 1, 0, 48, 57, 3, 0, 10, 10, 13, 13, 34, 34,
		3, 0, 9, 10, 13, 13, 32, 32, 224, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0,
		5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0,
		13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0,
		0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0,
		0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0,
		0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1,
		0, 0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1, 0, 0, 0, 0, 49, 1, 0, 0, 0, 0, 51,
		1, 0, 0, 0, 0, 53, 1, 0, 0, 0, 0, 55, 1, 0, 0, 0, 0, 57, 1, 0, 0, 0, 0,
		59, 1, 0, 0, 0, 0, 61, 1, 0, 0, 0, 0, 63, 1, 0, 0, 0, 0, 65, 1, 0, 0, 0,
		0, 67, 1, 0, 0, 0, 0, 69, 1, 0, 0, 0, 0, 71, 1, 0, 0, 0, 0, 73, 1, 0, 0,
		0, 1, 75, 1, 0, 0, 0, 3, 77, 1, 0, 0, 0, 5, 80, 1, 0, 0, 0, 7, 82, 1, 0,
		0, 0, 9, 84, 1, 0, 0, 0, 11, 86, 1, 0, 0, 0, 13, 90, 1, 0, 0, 0, 15, 96,
		1, 0, 0, 0, 17, 101, 1, 0, 0, 0, 19, 104, 1, 0, 0, 0, 21, 111, 1, 0, 0,
		0, 23, 114, 1, 0, 0, 0, 25, 117, 1, 0, 0, 0, 27, 120, 1, 0, 0, 0, 29, 123,
		1, 0, 0, 0, 31, 125, 1, 0, 0, 0, 33, 127, 1, 0, 0, 0, 35, 130, 1, 0, 0,
		0, 37, 133, 1, 0, 0, 0, 39, 135, 1, 0, 0, 0, 41, 137, 1, 0, 0, 0, 43, 139,
		1, 0, 0, 0, 45, 141, 1, 0, 0, 0, 47, 143, 1, 0, 0, 0, 49, 145, 1, 0, 0,
		0, 51, 147, 1, 0, 0, 0, 53, 152, 1, 0, 0, 0, 55, 158, 1, 0, 0, 0, 57, 162,
		1, 0, 0, 0, 59, 164, 1, 0, 0, 0, 61, 166, 1, 0, 0, 0, 63, 168, 1, 0, 0,
		0, 65, 170, 1, 0, 0, 0, 67, 178, 1, 0, 0, 0, 69, 200, 1, 0, 0, 0, 71, 202,
		1, 0, 0, 0, 73, 213, 1, 0, 0, 0, 75, 76, 5, 44, 0, 0, 76, 2, 1, 0, 0, 0,
		77, 78, 5, 58, 0, 0, 78, 79, 5, 61, 0, 0, 79, 4, 1, 0, 0, 0, 80, 81, 5,
		61, 0, 0, 81, 6, 1, 0, 0, 0, 82, 83, 5, 59, 0, 0, 83, 8, 1, 0, 0, 0, 84,
		85, 5, 58, 0, 0, 85, 10, 1, 0, 0, 0, 86, 87, 5, 102, 0, 0, 87, 88, 5, 111,
		0, 0, 88, 89, 5, 114, 0, 0, 89, 12, 1, 0, 0, 0, 90, 91, 5, 114, 0, 0, 91,
		92, 5, 97, 0, 0, 92, 93, 5, 110, 0, 0, 93, 94, 5, 103, 0, 0, 94, 95, 5,
		101, 0, 0, 95, 14, 1, 0, 0, 0, 96, 97, 5, 102, 0, 0, 97, 98, 5, 117, 0,
		0, 98, 99, 5, 110, 0, 0, 99, 100, 5, 99, 0, 0, 100, 16, 1, 0, 0, 0, 101,
		102, 5, 111, 0, 0, 102, 103, 5, 110, 0, 0, 103, 18, 1, 0, 0, 0, 104, 105,
		5, 114, 0, 0, 105, 106, 5, 101, 0, 0, 106, 107, 5, 116, 0, 0, 107, 108,
		5, 117, 0, 0, 108, 109, 5, 114, 0, 0, 109, 110, 5, 110, 0, 0, 110, 20,
		1, 0, 0, 0, 111, 112, 5, 124, 0, 0, 112, 113, 5, 124, 0, 0, 113, 22, 1,
		0, 0, 0, 114, 115, 5, 38, 0, 0, 115, 116, 5, 38, 0, 0, 116, 24, 1, 0, 0,
		0, 117, 118, 5, 61, 0, 0, 118, 119, 5, 61, 0, 0, 119, 26, 1, 0, 0, 0, 120,
		121, 5, 33, 0, 0, 121, 122, 5, 61, 0, 0, 122, 28, 1, 0, 0, 0, 123, 124,
		5, 62, 0, 0, 124, 30, 1, 0, 0, 0, 125, 126, 5, 60, 0, 0, 126, 32, 1, 0,
		0, 0, 127, 128, 5, 62, 0, 0, 128, 129, 5, 61, 0, 0, 129, 34, 1, 0, 0, 0,
		130, 131, 5, 60, 0, 0, 131, 132, 5, 61, 0, 0, 132, 36, 1, 0, 0, 0, 133,
		134, 5, 43, 0, 0, 134, 38, 1, 0, 0, 0, 135, 136, 5, 45, 0, 0, 136, 40,
		1, 0, 0, 0, 137, 138, 5, 42, 0, 0, 138, 42, 1, 0, 0, 0, 139, 140, 5, 47,
		0, 0, 140, 44, 1, 0, 0, 0, 141, 142, 5, 37, 0, 0, 142, 46, 1, 0, 0, 0,
		143, 144, 5, 94, 0, 0, 144, 48, 1, 0, 0, 0, 145, 146, 5, 33, 0, 0, 146,
		50, 1, 0, 0, 0, 147, 148, 5, 116, 0, 0, 148, 149, 5, 114, 0, 0, 149, 150,
		5, 117, 0, 0, 150, 151, 5, 101, 0, 0, 151, 52, 1, 0, 0, 0, 152, 153, 5,
		102, 0, 0, 153, 154, 5, 97, 0, 0, 154, 155, 5, 108, 0, 0, 155, 156, 5,
		115, 0, 0, 156, 157, 5, 101, 0, 0, 157, 54, 1, 0, 0, 0, 158, 159, 5, 110,
		0, 0, 159, 160, 5, 105, 0, 0, 160, 161, 5, 108, 0, 0, 161, 56, 1, 0, 0,
		0, 162, 163, 5, 40, 0, 0, 163, 58, 1, 0, 0, 0, 164, 165, 5, 41, 0, 0, 165,
		60, 1, 0, 0, 0, 166, 167, 5, 123, 0, 0, 167, 62, 1, 0, 0, 0, 168, 169,
		5, 125, 0, 0, 169, 64, 1, 0, 0, 0, 170, 174, 7, 0, 0, 0, 171, 173, 7, 1,
		0, 0, 172, 171, 1, 0, 0, 0, 173, 176, 1, 0, 0, 0, 174, 172, 1, 0, 0, 0,
		174, 175, 1, 0, 0, 0, 175, 66, 1, 0, 0, 0, 176, 174, 1, 0, 0, 0, 177, 179,
		7, 2, 0, 0, 178, 177, 1, 0, 0, 0, 179, 180, 1, 0, 0, 0, 180, 178, 1, 0,
		0, 0, 180, 181, 1, 0, 0, 0, 181, 68, 1, 0, 0, 0, 182, 184, 7, 2, 0, 0,
		183, 182, 1, 0, 0, 0, 184, 185, 1, 0, 0, 0, 185, 183, 1, 0, 0, 0, 185,
		186, 1, 0, 0, 0, 186, 187, 1, 0, 0, 0, 187, 191, 5, 46, 0, 0, 188, 190,
		7, 2, 0, 0, 189, 188, 1, 0, 0, 0, 190, 193, 1, 0, 0, 0, 191, 189, 1, 0,
		0, 0, 191, 192, 1, 0, 0, 0, 192, 201, 1, 0, 0, 0, 193, 191, 1, 0, 0, 0,
		194, 196, 5, 46, 0, 0, 195, 197, 7, 2, 0, 0, 196, 195, 1, 0, 0, 0, 197,
		198, 1, 0, 0, 0, 198, 196, 1, 0, 0, 0, 198, 199, 1, 0, 0, 0, 199, 201,
		1, 0, 0, 0, 200, 183, 1, 0, 0, 0, 200, 194, 1, 0, 0, 0, 201, 70, 1, 0,
		0, 0, 202, 208, 5, 34, 0, 0, 203, 207, 8, 3, 0, 0, 204, 205, 5, 34, 0,
		0, 205, 207, 5, 34, 0, 0, 206, 203, 1, 0, 0, 0, 206, 204, 1, 0, 0, 0, 207,
		210, 1, 0, 0, 0, 208, 206, 1, 0, 0, 0, 208, 209, 1, 0, 0, 0, 209, 211,
		1, 0, 0, 0, 210, 208, 1, 0, 0, 0, 211, 212, 5, 34, 0, 0, 212, 72, 1, 0,
		0, 0, 213, 214, 7, 4, 0, 0, 214, 215, 1, 0, 0, 0, 215, 216, 6, 36, 0, 0,
		216, 74, 1, 0, 0, 0, 9, 0, 174, 180, 185, 191, 198, 200, 206, 208, 1, 6,
		0, 0,
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
	GoEventerLexerRANGE        = 7
	GoEventerLexerFUNC         = 8
	GoEventerLexerON           = 9
	GoEventerLexerRETURN       = 10
	GoEventerLexerOR           = 11
	GoEventerLexerAND          = 12
	GoEventerLexerEQ           = 13
	GoEventerLexerNEQ          = 14
	GoEventerLexerGT           = 15
	GoEventerLexerLT           = 16
	GoEventerLexerGTEQ         = 17
	GoEventerLexerLTEQ         = 18
	GoEventerLexerPLUS         = 19
	GoEventerLexerMINUS        = 20
	GoEventerLexerMULT         = 21
	GoEventerLexerDIV          = 22
	GoEventerLexerMOD          = 23
	GoEventerLexerPOW          = 24
	GoEventerLexerNOT          = 25
	GoEventerLexerTRUE         = 26
	GoEventerLexerFALSE        = 27
	GoEventerLexerNIL          = 28
	GoEventerLexerPARAN_OPEN   = 29
	GoEventerLexerPARAN_CLOSE  = 30
	GoEventerLexerBARACE_OPEN  = 31
	GoEventerLexerBARACE_CLOSE = 32
	GoEventerLexerID           = 33
	GoEventerLexerINT          = 34
	GoEventerLexerFLOAT        = 35
	GoEventerLexerSTRING       = 36
	GoEventerLexerSPACE        = 37
)
