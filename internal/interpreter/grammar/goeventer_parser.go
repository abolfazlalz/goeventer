// Code generated from GoEventer.g4 by ANTLR 4.13.0. DO NOT EDIT.

package grammar // GoEventer
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

type GoEventerParser struct {
	*antlr.BaseParser
}

var GoEventerParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func goeventerParserInit() {
	staticData := &GoEventerParserStaticData
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
		"parse", "block", "stat", "ifStat", "conditionBlock", "structStat",
		"structField", "assignVariableStat", "updateVariableStat", "notifyChanStat",
		"forStat", "whileStat", "defineListenerStat", "defineFunctionStat",
		"statBlock", "returnStat", "methodCallStat", "methodCall", "methodCallArguments",
		"functionDefineArguments", "expr", "atom",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 46, 246, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 1, 0, 1, 0, 1, 0, 1, 1, 5, 1, 49, 8, 1, 10, 1, 12, 1, 52, 9,
		1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3,
		2, 65, 8, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 5, 3, 72, 8, 3, 10, 3, 12, 3,
		75, 9, 3, 1, 3, 1, 3, 3, 3, 79, 8, 3, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1,
		5, 1, 5, 1, 5, 1, 5, 5, 5, 90, 8, 5, 10, 5, 12, 5, 93, 9, 5, 1, 5, 1, 5,
		1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8,
		1, 8, 5, 8, 110, 8, 8, 10, 8, 12, 8, 113, 9, 8, 1, 8, 1, 8, 1, 8, 1, 8,
		1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10,
		1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 3, 12, 139,
		8, 12, 1, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1,
		13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 3, 14, 156, 8, 14, 1, 15, 1, 15,
		1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1,
		18, 1, 18, 1, 18, 1, 18, 5, 18, 174, 8, 18, 10, 18, 12, 18, 177, 9, 18,
		3, 18, 179, 8, 18, 1, 19, 1, 19, 1, 19, 1, 19, 5, 19, 185, 8, 19, 10, 19,
		12, 19, 188, 9, 19, 3, 19, 190, 8, 19, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20,
		1, 20, 1, 20, 3, 20, 199, 8, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1,
		20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20,
		1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 5, 20, 222, 8, 20, 10, 20, 12, 20, 225,
		9, 20, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 5,
		21, 236, 8, 21, 10, 21, 12, 21, 239, 9, 21, 1, 21, 1, 21, 1, 21, 3, 21,
		244, 8, 21, 1, 21, 0, 1, 40, 22, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20,
		22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 0, 7, 1, 0, 2, 3, 1, 0, 30,
		32, 1, 0, 28, 29, 1, 0, 24, 27, 1, 0, 22, 23, 1, 0, 43, 44, 1, 0, 35, 36,
		261, 0, 44, 1, 0, 0, 0, 2, 50, 1, 0, 0, 0, 4, 64, 1, 0, 0, 0, 6, 66, 1,
		0, 0, 0, 8, 80, 1, 0, 0, 0, 10, 83, 1, 0, 0, 0, 12, 97, 1, 0, 0, 0, 14,
		101, 1, 0, 0, 0, 16, 106, 1, 0, 0, 0, 18, 118, 1, 0, 0, 0, 20, 123, 1,
		0, 0, 0, 22, 131, 1, 0, 0, 0, 24, 135, 1, 0, 0, 0, 26, 143, 1, 0, 0, 0,
		28, 155, 1, 0, 0, 0, 30, 157, 1, 0, 0, 0, 32, 161, 1, 0, 0, 0, 34, 164,
		1, 0, 0, 0, 36, 178, 1, 0, 0, 0, 38, 189, 1, 0, 0, 0, 40, 198, 1, 0, 0,
		0, 42, 243, 1, 0, 0, 0, 44, 45, 3, 2, 1, 0, 45, 46, 5, 0, 0, 1, 46, 1,
		1, 0, 0, 0, 47, 49, 3, 4, 2, 0, 48, 47, 1, 0, 0, 0, 49, 52, 1, 0, 0, 0,
		50, 48, 1, 0, 0, 0, 50, 51, 1, 0, 0, 0, 51, 3, 1, 0, 0, 0, 52, 50, 1, 0,
		0, 0, 53, 65, 3, 32, 16, 0, 54, 65, 3, 24, 12, 0, 55, 65, 3, 16, 8, 0,
		56, 65, 3, 14, 7, 0, 57, 65, 3, 20, 10, 0, 58, 65, 3, 22, 11, 0, 59, 65,
		3, 26, 13, 0, 60, 65, 3, 30, 15, 0, 61, 65, 3, 18, 9, 0, 62, 65, 3, 10,
		5, 0, 63, 65, 3, 6, 3, 0, 64, 53, 1, 0, 0, 0, 64, 54, 1, 0, 0, 0, 64, 55,
		1, 0, 0, 0, 64, 56, 1, 0, 0, 0, 64, 57, 1, 0, 0, 0, 64, 58, 1, 0, 0, 0,
		64, 59, 1, 0, 0, 0, 64, 60, 1, 0, 0, 0, 64, 61, 1, 0, 0, 0, 64, 62, 1,
		0, 0, 0, 64, 63, 1, 0, 0, 0, 65, 5, 1, 0, 0, 0, 66, 67, 5, 13, 0, 0, 67,
		73, 3, 8, 4, 0, 68, 69, 5, 14, 0, 0, 69, 70, 5, 13, 0, 0, 70, 72, 3, 8,
		4, 0, 71, 68, 1, 0, 0, 0, 72, 75, 1, 0, 0, 0, 73, 71, 1, 0, 0, 0, 73, 74,
		1, 0, 0, 0, 74, 78, 1, 0, 0, 0, 75, 73, 1, 0, 0, 0, 76, 77, 5, 14, 0, 0,
		77, 79, 3, 28, 14, 0, 78, 76, 1, 0, 0, 0, 78, 79, 1, 0, 0, 0, 79, 7, 1,
		0, 0, 0, 80, 81, 3, 40, 20, 0, 81, 82, 3, 28, 14, 0, 82, 9, 1, 0, 0, 0,
		83, 84, 5, 17, 0, 0, 84, 85, 5, 42, 0, 0, 85, 86, 5, 40, 0, 0, 86, 91,
		3, 12, 6, 0, 87, 88, 5, 4, 0, 0, 88, 90, 3, 12, 6, 0, 89, 87, 1, 0, 0,
		0, 90, 93, 1, 0, 0, 0, 91, 89, 1, 0, 0, 0, 91, 92, 1, 0, 0, 0, 92, 94,
		1, 0, 0, 0, 93, 91, 1, 0, 0, 0, 94, 95, 5, 41, 0, 0, 95, 96, 5, 4, 0, 0,
		96, 11, 1, 0, 0, 0, 97, 98, 5, 42, 0, 0, 98, 99, 5, 5, 0, 0, 99, 100, 3,
		40, 20, 0, 100, 13, 1, 0, 0, 0, 101, 102, 5, 42, 0, 0, 102, 103, 7, 0,
		0, 0, 103, 104, 3, 40, 20, 0, 104, 105, 5, 4, 0, 0, 105, 15, 1, 0, 0, 0,
		106, 111, 5, 42, 0, 0, 107, 108, 5, 6, 0, 0, 108, 110, 5, 42, 0, 0, 109,
		107, 1, 0, 0, 0, 110, 113, 1, 0, 0, 0, 111, 109, 1, 0, 0, 0, 111, 112,
		1, 0, 0, 0, 112, 114, 1, 0, 0, 0, 113, 111, 1, 0, 0, 0, 114, 115, 5, 3,
		0, 0, 115, 116, 3, 40, 20, 0, 116, 117, 5, 4, 0, 0, 117, 17, 1, 0, 0, 0,
		118, 119, 5, 42, 0, 0, 119, 120, 5, 16, 0, 0, 120, 121, 3, 40, 20, 0, 121,
		122, 5, 4, 0, 0, 122, 19, 1, 0, 0, 0, 123, 124, 5, 7, 0, 0, 124, 125, 5,
		42, 0, 0, 125, 126, 5, 2, 0, 0, 126, 127, 3, 40, 20, 0, 127, 128, 5, 5,
		0, 0, 128, 129, 3, 40, 20, 0, 129, 130, 3, 28, 14, 0, 130, 21, 1, 0, 0,
		0, 131, 132, 5, 8, 0, 0, 132, 133, 3, 40, 20, 0, 133, 134, 3, 28, 14, 0,
		134, 23, 1, 0, 0, 0, 135, 138, 5, 11, 0, 0, 136, 137, 5, 42, 0, 0, 137,
		139, 5, 2, 0, 0, 138, 136, 1, 0, 0, 0, 138, 139, 1, 0, 0, 0, 139, 140,
		1, 0, 0, 0, 140, 141, 3, 34, 17, 0, 141, 142, 3, 28, 14, 0, 142, 25, 1,
		0, 0, 0, 143, 144, 5, 10, 0, 0, 144, 145, 5, 42, 0, 0, 145, 146, 5, 38,
		0, 0, 146, 147, 3, 38, 19, 0, 147, 148, 5, 39, 0, 0, 148, 149, 3, 28, 14,
		0, 149, 27, 1, 0, 0, 0, 150, 151, 5, 40, 0, 0, 151, 152, 3, 2, 1, 0, 152,
		153, 5, 41, 0, 0, 153, 156, 1, 0, 0, 0, 154, 156, 3, 4, 2, 0, 155, 150,
		1, 0, 0, 0, 155, 154, 1, 0, 0, 0, 156, 29, 1, 0, 0, 0, 157, 158, 5, 12,
		0, 0, 158, 159, 3, 40, 20, 0, 159, 160, 5, 4, 0, 0, 160, 31, 1, 0, 0, 0,
		161, 162, 3, 34, 17, 0, 162, 163, 5, 4, 0, 0, 163, 33, 1, 0, 0, 0, 164,
		165, 5, 42, 0, 0, 165, 166, 5, 38, 0, 0, 166, 167, 3, 36, 18, 0, 167, 168,
		5, 39, 0, 0, 168, 35, 1, 0, 0, 0, 169, 179, 1, 0, 0, 0, 170, 175, 3, 40,
		20, 0, 171, 172, 5, 1, 0, 0, 172, 174, 3, 40, 20, 0, 173, 171, 1, 0, 0,
		0, 174, 177, 1, 0, 0, 0, 175, 173, 1, 0, 0, 0, 175, 176, 1, 0, 0, 0, 176,
		179, 1, 0, 0, 0, 177, 175, 1, 0, 0, 0, 178, 169, 1, 0, 0, 0, 178, 170,
		1, 0, 0, 0, 179, 37, 1, 0, 0, 0, 180, 190, 1, 0, 0, 0, 181, 186, 5, 42,
		0, 0, 182, 183, 5, 1, 0, 0, 183, 185, 5, 42, 0, 0, 184, 182, 1, 0, 0, 0,
		185, 188, 1, 0, 0, 0, 186, 184, 1, 0, 0, 0, 186, 187, 1, 0, 0, 0, 187,
		190, 1, 0, 0, 0, 188, 186, 1, 0, 0, 0, 189, 180, 1, 0, 0, 0, 189, 181,
		1, 0, 0, 0, 190, 39, 1, 0, 0, 0, 191, 192, 6, 20, -1, 0, 192, 199, 3, 34,
		17, 0, 193, 194, 5, 29, 0, 0, 194, 199, 3, 40, 20, 9, 195, 196, 5, 34,
		0, 0, 196, 199, 3, 40, 20, 8, 197, 199, 3, 42, 21, 0, 198, 191, 1, 0, 0,
		0, 198, 193, 1, 0, 0, 0, 198, 195, 1, 0, 0, 0, 198, 197, 1, 0, 0, 0, 199,
		223, 1, 0, 0, 0, 200, 201, 10, 10, 0, 0, 201, 202, 5, 33, 0, 0, 202, 222,
		3, 40, 20, 10, 203, 204, 10, 7, 0, 0, 204, 205, 7, 1, 0, 0, 205, 222, 3,
		40, 20, 8, 206, 207, 10, 6, 0, 0, 207, 208, 7, 2, 0, 0, 208, 222, 3, 40,
		20, 7, 209, 210, 10, 5, 0, 0, 210, 211, 7, 3, 0, 0, 211, 222, 3, 40, 20,
		6, 212, 213, 10, 4, 0, 0, 213, 214, 7, 4, 0, 0, 214, 222, 3, 40, 20, 5,
		215, 216, 10, 3, 0, 0, 216, 217, 5, 21, 0, 0, 217, 222, 3, 40, 20, 4, 218,
		219, 10, 2, 0, 0, 219, 220, 5, 20, 0, 0, 220, 222, 3, 40, 20, 3, 221, 200,
		1, 0, 0, 0, 221, 203, 1, 0, 0, 0, 221, 206, 1, 0, 0, 0, 221, 209, 1, 0,
		0, 0, 221, 212, 1, 0, 0, 0, 221, 215, 1, 0, 0, 0, 221, 218, 1, 0, 0, 0,
		222, 225, 1, 0, 0, 0, 223, 221, 1, 0, 0, 0, 223, 224, 1, 0, 0, 0, 224,
		41, 1, 0, 0, 0, 225, 223, 1, 0, 0, 0, 226, 227, 5, 38, 0, 0, 227, 228,
		3, 40, 20, 0, 228, 229, 5, 39, 0, 0, 229, 244, 1, 0, 0, 0, 230, 244, 7,
		5, 0, 0, 231, 244, 7, 6, 0, 0, 232, 237, 5, 42, 0, 0, 233, 234, 5, 6, 0,
		0, 234, 236, 5, 42, 0, 0, 235, 233, 1, 0, 0, 0, 236, 239, 1, 0, 0, 0, 237,
		235, 1, 0, 0, 0, 237, 238, 1, 0, 0, 0, 238, 244, 1, 0, 0, 0, 239, 237,
		1, 0, 0, 0, 240, 244, 5, 45, 0, 0, 241, 244, 5, 37, 0, 0, 242, 244, 5,
		15, 0, 0, 243, 226, 1, 0, 0, 0, 243, 230, 1, 0, 0, 0, 243, 231, 1, 0, 0,
		0, 243, 232, 1, 0, 0, 0, 243, 240, 1, 0, 0, 0, 243, 241, 1, 0, 0, 0, 243,
		242, 1, 0, 0, 0, 244, 43, 1, 0, 0, 0, 17, 50, 64, 73, 78, 91, 111, 138,
		155, 175, 178, 186, 189, 198, 221, 223, 237, 243,
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

// GoEventerParserInit initializes any static state used to implement GoEventerParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewGoEventerParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func GoEventerParserInit() {
	staticData := &GoEventerParserStaticData
	staticData.once.Do(goeventerParserInit)
}

// NewGoEventerParser produces a new parser instance for the optional input antlr.TokenStream.
func NewGoEventerParser(input antlr.TokenStream) *GoEventerParser {
	GoEventerParserInit()
	this := new(GoEventerParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &GoEventerParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "GoEventer.g4"

	return this
}

// GoEventerParser tokens.
const (
	GoEventerParserEOF          = antlr.TokenEOF
	GoEventerParserT__0         = 1
	GoEventerParserASSIGN       = 2
	GoEventerParserUPDATE       = 3
	GoEventerParserSCOL         = 4
	GoEventerParserCOL          = 5
	GoEventerParserDOT          = 6
	GoEventerParserFOR          = 7
	GoEventerParserWHILE        = 8
	GoEventerParserRANGE        = 9
	GoEventerParserFUNC         = 10
	GoEventerParserON           = 11
	GoEventerParserRETURN       = 12
	GoEventerParserIF           = 13
	GoEventerParserELSE         = 14
	GoEventerParserCHAN         = 15
	GoEventerParserNOTIFY       = 16
	GoEventerParserSTRUCT       = 17
	GoEventerParserSTRING_TYPE  = 18
	GoEventerParserINTEGER_TYPE = 19
	GoEventerParserOR           = 20
	GoEventerParserAND          = 21
	GoEventerParserEQ           = 22
	GoEventerParserNEQ          = 23
	GoEventerParserGT           = 24
	GoEventerParserLT           = 25
	GoEventerParserGTEQ         = 26
	GoEventerParserLTEQ         = 27
	GoEventerParserPLUS         = 28
	GoEventerParserMINUS        = 29
	GoEventerParserMULT         = 30
	GoEventerParserDIV          = 31
	GoEventerParserMOD          = 32
	GoEventerParserPOW          = 33
	GoEventerParserNOT          = 34
	GoEventerParserTRUE         = 35
	GoEventerParserFALSE        = 36
	GoEventerParserNIL          = 37
	GoEventerParserPARAN_OPEN   = 38
	GoEventerParserPARAN_CLOSE  = 39
	GoEventerParserBARACE_OPEN  = 40
	GoEventerParserBARACE_CLOSE = 41
	GoEventerParserID           = 42
	GoEventerParserINT          = 43
	GoEventerParserFLOAT        = 44
	GoEventerParserSTRING       = 45
	GoEventerParserSPACE        = 46
)

// GoEventerParser rules.
const (
	GoEventerParserRULE_parse                   = 0
	GoEventerParserRULE_block                   = 1
	GoEventerParserRULE_stat                    = 2
	GoEventerParserRULE_ifStat                  = 3
	GoEventerParserRULE_conditionBlock          = 4
	GoEventerParserRULE_structStat              = 5
	GoEventerParserRULE_structField             = 6
	GoEventerParserRULE_assignVariableStat      = 7
	GoEventerParserRULE_updateVariableStat      = 8
	GoEventerParserRULE_notifyChanStat          = 9
	GoEventerParserRULE_forStat                 = 10
	GoEventerParserRULE_whileStat               = 11
	GoEventerParserRULE_defineListenerStat      = 12
	GoEventerParserRULE_defineFunctionStat      = 13
	GoEventerParserRULE_statBlock               = 14
	GoEventerParserRULE_returnStat              = 15
	GoEventerParserRULE_methodCallStat          = 16
	GoEventerParserRULE_methodCall              = 17
	GoEventerParserRULE_methodCallArguments     = 18
	GoEventerParserRULE_functionDefineArguments = 19
	GoEventerParserRULE_expr                    = 20
	GoEventerParserRULE_atom                    = 21
)

// IParseContext is an interface to support dynamic dispatch.
type IParseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Block() IBlockContext
	EOF() antlr.TerminalNode

	// IsParseContext differentiates from other interfaces.
	IsParseContext()
}

type ParseContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParseContext() *ParseContext {
	var p = new(ParseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoEventerParserRULE_parse
	return p
}

func InitEmptyParseContext(p *ParseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoEventerParserRULE_parse
}

func (*ParseContext) IsParseContext() {}

func NewParseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParseContext {
	var p = new(ParseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoEventerParserRULE_parse

	return p
}

func (s *ParseContext) GetParser() antlr.Parser { return s.parser }

func (s *ParseContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *ParseContext) EOF() antlr.TerminalNode {
	return s.GetToken(GoEventerParserEOF, 0)
}

func (s *ParseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoEventerListener); ok {
		listenerT.EnterParse(s)
	}
}

func (s *ParseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoEventerListener); ok {
		listenerT.ExitParse(s)
	}
}

func (s *ParseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoEventerVisitor:
		return t.VisitParse(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoEventerParser) Parse() (localctx IParseContext) {
	localctx = NewParseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, GoEventerParserRULE_parse)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(44)
		p.Block()
	}
	{
		p.SetState(45)
		p.Match(GoEventerParserEOF)
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

// IBlockContext is an interface to support dynamic dispatch.
type IBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllStat() []IStatContext
	Stat(i int) IStatContext

	// IsBlockContext differentiates from other interfaces.
	IsBlockContext()
}

type BlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockContext() *BlockContext {
	var p = new(BlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoEventerParserRULE_block
	return p
}

func InitEmptyBlockContext(p *BlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoEventerParserRULE_block
}

func (*BlockContext) IsBlockContext() {}

func NewBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockContext {
	var p = new(BlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoEventerParserRULE_block

	return p
}

func (s *BlockContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockContext) AllStat() []IStatContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatContext); ok {
			len++
		}
	}

	tst := make([]IStatContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatContext); ok {
			tst[i] = t.(IStatContext)
			i++
		}
	}

	return tst
}

func (s *BlockContext) Stat(i int) IStatContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatContext); ok {
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

	return t.(IStatContext)
}

func (s *BlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoEventerListener); ok {
		listenerT.EnterBlock(s)
	}
}

func (s *BlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoEventerListener); ok {
		listenerT.ExitBlock(s)
	}
}

func (s *BlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoEventerVisitor:
		return t.VisitBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoEventerParser) Block() (localctx IBlockContext) {
	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, GoEventerParserRULE_block)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(50)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4398046657920) != 0 {
		{
			p.SetState(47)
			p.Stat()
		}

		p.SetState(52)
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

// IStatContext is an interface to support dynamic dispatch.
type IStatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	MethodCallStat() IMethodCallStatContext
	DefineListenerStat() IDefineListenerStatContext
	UpdateVariableStat() IUpdateVariableStatContext
	AssignVariableStat() IAssignVariableStatContext
	ForStat() IForStatContext
	WhileStat() IWhileStatContext
	DefineFunctionStat() IDefineFunctionStatContext
	ReturnStat() IReturnStatContext
	NotifyChanStat() INotifyChanStatContext
	StructStat() IStructStatContext
	IfStat() IIfStatContext

	// IsStatContext differentiates from other interfaces.
	IsStatContext()
}

type StatContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatContext() *StatContext {
	var p = new(StatContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoEventerParserRULE_stat
	return p
}

func InitEmptyStatContext(p *StatContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoEventerParserRULE_stat
}

func (*StatContext) IsStatContext() {}

func NewStatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatContext {
	var p = new(StatContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoEventerParserRULE_stat

	return p
}

func (s *StatContext) GetParser() antlr.Parser { return s.parser }

func (s *StatContext) MethodCallStat() IMethodCallStatContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMethodCallStatContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMethodCallStatContext)
}

func (s *StatContext) DefineListenerStat() IDefineListenerStatContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDefineListenerStatContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDefineListenerStatContext)
}

func (s *StatContext) UpdateVariableStat() IUpdateVariableStatContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUpdateVariableStatContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUpdateVariableStatContext)
}

func (s *StatContext) AssignVariableStat() IAssignVariableStatContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignVariableStatContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignVariableStatContext)
}

func (s *StatContext) ForStat() IForStatContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IForStatContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IForStatContext)
}

func (s *StatContext) WhileStat() IWhileStatContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWhileStatContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWhileStatContext)
}

func (s *StatContext) DefineFunctionStat() IDefineFunctionStatContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDefineFunctionStatContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDefineFunctionStatContext)
}

func (s *StatContext) ReturnStat() IReturnStatContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReturnStatContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReturnStatContext)
}

func (s *StatContext) NotifyChanStat() INotifyChanStatContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INotifyChanStatContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INotifyChanStatContext)
}

func (s *StatContext) StructStat() IStructStatContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStructStatContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStructStatContext)
}

func (s *StatContext) IfStat() IIfStatContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfStatContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfStatContext)
}

func (s *StatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoEventerListener); ok {
		listenerT.EnterStat(s)
	}
}

func (s *StatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoEventerListener); ok {
		listenerT.ExitStat(s)
	}
}

func (s *StatContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoEventerVisitor:
		return t.VisitStat(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoEventerParser) Stat() (localctx IStatContext) {
	localctx = NewStatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, GoEventerParserRULE_stat)
	p.SetState(64)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(53)
			p.MethodCallStat()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(54)
			p.DefineListenerStat()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(55)
			p.UpdateVariableStat()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(56)
			p.AssignVariableStat()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(57)
			p.ForStat()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(58)
			p.WhileStat()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(59)
			p.DefineFunctionStat()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(60)
			p.ReturnStat()
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(61)
			p.NotifyChanStat()
		}

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(62)
			p.StructStat()
		}

	case 11:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(63)
			p.IfStat()
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

// IIfStatContext is an interface to support dynamic dispatch.
type IIfStatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllIF() []antlr.TerminalNode
	IF(i int) antlr.TerminalNode
	AllConditionBlock() []IConditionBlockContext
	ConditionBlock(i int) IConditionBlockContext
	AllELSE() []antlr.TerminalNode
	ELSE(i int) antlr.TerminalNode
	StatBlock() IStatBlockContext

	// IsIfStatContext differentiates from other interfaces.
	IsIfStatContext()
}

type IfStatContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIfStatContext() *IfStatContext {
	var p = new(IfStatContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoEventerParserRULE_ifStat
	return p
}

func InitEmptyIfStatContext(p *IfStatContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoEventerParserRULE_ifStat
}

func (*IfStatContext) IsIfStatContext() {}

func NewIfStatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfStatContext {
	var p = new(IfStatContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoEventerParserRULE_ifStat

	return p
}

func (s *IfStatContext) GetParser() antlr.Parser { return s.parser }

func (s *IfStatContext) AllIF() []antlr.TerminalNode {
	return s.GetTokens(GoEventerParserIF)
}

func (s *IfStatContext) IF(i int) antlr.TerminalNode {
	return s.GetToken(GoEventerParserIF, i)
}

func (s *IfStatContext) AllConditionBlock() []IConditionBlockContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IConditionBlockContext); ok {
			len++
		}
	}

	tst := make([]IConditionBlockContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IConditionBlockContext); ok {
			tst[i] = t.(IConditionBlockContext)
			i++
		}
	}

	return tst
}

func (s *IfStatContext) ConditionBlock(i int) IConditionBlockContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConditionBlockContext); ok {
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

	return t.(IConditionBlockContext)
}

func (s *IfStatContext) AllELSE() []antlr.TerminalNode {
	return s.GetTokens(GoEventerParserELSE)
}

func (s *IfStatContext) ELSE(i int) antlr.TerminalNode {
	return s.GetToken(GoEventerParserELSE, i)
}

func (s *IfStatContext) StatBlock() IStatBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatBlockContext)
}

func (s *IfStatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfStatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfStatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoEventerListener); ok {
		listenerT.EnterIfStat(s)
	}
}

func (s *IfStatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoEventerListener); ok {
		listenerT.ExitIfStat(s)
	}
}

func (s *IfStatContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoEventerVisitor:
		return t.VisitIfStat(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoEventerParser) IfStat() (localctx IIfStatContext) {
	localctx = NewIfStatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, GoEventerParserRULE_ifStat)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(66)
		p.Match(GoEventerParserIF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(67)
		p.ConditionBlock()
	}
	p.SetState(73)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(68)
				p.Match(GoEventerParserELSE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(69)
				p.Match(GoEventerParserIF)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(70)
				p.ConditionBlock()
			}

		}
		p.SetState(75)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(78)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(76)
			p.Match(GoEventerParserELSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(77)
			p.StatBlock()
		}

	} else if p.HasError() { // JIM
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

// IConditionBlockContext is an interface to support dynamic dispatch.
type IConditionBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expr() IExprContext
	StatBlock() IStatBlockContext

	// IsConditionBlockContext differentiates from other interfaces.
	IsConditionBlockContext()
}

type ConditionBlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConditionBlockContext() *ConditionBlockContext {
	var p = new(ConditionBlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoEventerParserRULE_conditionBlock
	return p
}

func InitEmptyConditionBlockContext(p *ConditionBlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoEventerParserRULE_conditionBlock
}

func (*ConditionBlockContext) IsConditionBlockContext() {}

func NewConditionBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConditionBlockContext {
	var p = new(ConditionBlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoEventerParserRULE_conditionBlock

	return p
}

func (s *ConditionBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *ConditionBlockContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ConditionBlockContext) StatBlock() IStatBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatBlockContext)
}

func (s *ConditionBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConditionBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConditionBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoEventerListener); ok {
		listenerT.EnterConditionBlock(s)
	}
}

func (s *ConditionBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoEventerListener); ok {
		listenerT.ExitConditionBlock(s)
	}
}

func (s *ConditionBlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoEventerVisitor:
		return t.VisitConditionBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoEventerParser) ConditionBlock() (localctx IConditionBlockContext) {
	localctx = NewConditionBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, GoEventerParserRULE_conditionBlock)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(80)
		p.expr(0)
	}
	{
		p.SetState(81)
		p.StatBlock()
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

// IStructStatContext is an interface to support dynamic dispatch.
type IStructStatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	STRUCT() antlr.TerminalNode
	ID() antlr.TerminalNode
	BARACE_OPEN() antlr.TerminalNode
	AllStructField() []IStructFieldContext
	StructField(i int) IStructFieldContext
	BARACE_CLOSE() antlr.TerminalNode
	AllSCOL() []antlr.TerminalNode
	SCOL(i int) antlr.TerminalNode

	// IsStructStatContext differentiates from other interfaces.
	IsStructStatContext()
}

type StructStatContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStructStatContext() *StructStatContext {
	var p = new(StructStatContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoEventerParserRULE_structStat
	return p
}

func InitEmptyStructStatContext(p *StructStatContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoEventerParserRULE_structStat
}

func (*StructStatContext) IsStructStatContext() {}

func NewStructStatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StructStatContext {
	var p = new(StructStatContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoEventerParserRULE_structStat

	return p
}

func (s *StructStatContext) GetParser() antlr.Parser { return s.parser }

func (s *StructStatContext) STRUCT() antlr.TerminalNode {
	return s.GetToken(GoEventerParserSTRUCT, 0)
}

func (s *StructStatContext) ID() antlr.TerminalNode {
	return s.GetToken(GoEventerParserID, 0)
}

func (s *StructStatContext) BARACE_OPEN() antlr.TerminalNode {
	return s.GetToken(GoEventerParserBARACE_OPEN, 0)
}

func (s *StructStatContext) AllStructField() []IStructFieldContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStructFieldContext); ok {
			len++
		}
	}

	tst := make([]IStructFieldContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStructFieldContext); ok {
			tst[i] = t.(IStructFieldContext)
			i++
		}
	}

	return tst
}

func (s *StructStatContext) StructField(i int) IStructFieldContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStructFieldContext); ok {
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

	return t.(IStructFieldContext)
}

func (s *StructStatContext) BARACE_CLOSE() antlr.TerminalNode {
	return s.GetToken(GoEventerParserBARACE_CLOSE, 0)
}

func (s *StructStatContext) AllSCOL() []antlr.TerminalNode {
	return s.GetTokens(GoEventerParserSCOL)
}

func (s *StructStatContext) SCOL(i int) antlr.TerminalNode {
	return s.GetToken(GoEventerParserSCOL, i)
}

func (s *StructStatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructStatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StructStatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoEventerListener); ok {
		listenerT.EnterStructStat(s)
	}
}

func (s *StructStatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoEventerListener); ok {
		listenerT.ExitStructStat(s)
	}
}

func (s *StructStatContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoEventerVisitor:
		return t.VisitStructStat(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoEventerParser) StructStat() (localctx IStructStatContext) {
	localctx = NewStructStatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, GoEventerParserRULE_structStat)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(83)
		p.Match(GoEventerParserSTRUCT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(84)
		p.Match(GoEventerParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(85)
		p.Match(GoEventerParserBARACE_OPEN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(86)
		p.StructField()
	}
	p.SetState(91)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == GoEventerParserSCOL {
		{
			p.SetState(87)
			p.Match(GoEventerParserSCOL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(88)
			p.StructField()
		}

		p.SetState(93)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(94)
		p.Match(GoEventerParserBARACE_CLOSE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(95)
		p.Match(GoEventerParserSCOL)
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

// IStructFieldContext is an interface to support dynamic dispatch.
type IStructFieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	COL() antlr.TerminalNode
	Expr() IExprContext

	// IsStructFieldContext differentiates from other interfaces.
	IsStructFieldContext()
}

type StructFieldContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStructFieldContext() *StructFieldContext {
	var p = new(StructFieldContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoEventerParserRULE_structField
	return p
}

func InitEmptyStructFieldContext(p *StructFieldContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoEventerParserRULE_structField
}

func (*StructFieldContext) IsStructFieldContext() {}

func NewStructFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StructFieldContext {
	var p = new(StructFieldContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoEventerParserRULE_structField

	return p
}

func (s *StructFieldContext) GetParser() antlr.Parser { return s.parser }

func (s *StructFieldContext) ID() antlr.TerminalNode {
	return s.GetToken(GoEventerParserID, 0)
}

func (s *StructFieldContext) COL() antlr.TerminalNode {
	return s.GetToken(GoEventerParserCOL, 0)
}

func (s *StructFieldContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *StructFieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructFieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StructFieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoEventerListener); ok {
		listenerT.EnterStructField(s)
	}
}

func (s *StructFieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoEventerListener); ok {
		listenerT.ExitStructField(s)
	}
}

func (s *StructFieldContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoEventerVisitor:
		return t.VisitStructField(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoEventerParser) StructField() (localctx IStructFieldContext) {
	localctx = NewStructFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, GoEventerParserRULE_structField)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(97)
		p.Match(GoEventerParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(98)
		p.Match(GoEventerParserCOL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(99)
		p.expr(0)
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

// IAssignVariableStatContext is an interface to support dynamic dispatch.
type IAssignVariableStatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op token.
	GetOp() antlr.Token

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// Getter signatures
	ID() antlr.TerminalNode
	Expr() IExprContext
	SCOL() antlr.TerminalNode
	ASSIGN() antlr.TerminalNode
	UPDATE() antlr.TerminalNode

	// IsAssignVariableStatContext differentiates from other interfaces.
	IsAssignVariableStatContext()
}

type AssignVariableStatContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	op     antlr.Token
}

func NewEmptyAssignVariableStatContext() *AssignVariableStatContext {
	var p = new(AssignVariableStatContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoEventerParserRULE_assignVariableStat
	return p
}

func InitEmptyAssignVariableStatContext(p *AssignVariableStatContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoEventerParserRULE_assignVariableStat
}

func (*AssignVariableStatContext) IsAssignVariableStatContext() {}

func NewAssignVariableStatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignVariableStatContext {
	var p = new(AssignVariableStatContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoEventerParserRULE_assignVariableStat

	return p
}

func (s *AssignVariableStatContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignVariableStatContext) GetOp() antlr.Token { return s.op }

func (s *AssignVariableStatContext) SetOp(v antlr.Token) { s.op = v }

func (s *AssignVariableStatContext) ID() antlr.TerminalNode {
	return s.GetToken(GoEventerParserID, 0)
}

func (s *AssignVariableStatContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *AssignVariableStatContext) SCOL() antlr.TerminalNode {
	return s.GetToken(GoEventerParserSCOL, 0)
}

func (s *AssignVariableStatContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(GoEventerParserASSIGN, 0)
}

func (s *AssignVariableStatContext) UPDATE() antlr.TerminalNode {
	return s.GetToken(GoEventerParserUPDATE, 0)
}

func (s *AssignVariableStatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignVariableStatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignVariableStatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoEventerListener); ok {
		listenerT.EnterAssignVariableStat(s)
	}
}

func (s *AssignVariableStatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoEventerListener); ok {
		listenerT.ExitAssignVariableStat(s)
	}
}

func (s *AssignVariableStatContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoEventerVisitor:
		return t.VisitAssignVariableStat(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoEventerParser) AssignVariableStat() (localctx IAssignVariableStatContext) {
	localctx = NewAssignVariableStatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, GoEventerParserRULE_assignVariableStat)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(101)
		p.Match(GoEventerParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(102)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*AssignVariableStatContext).op = _lt

		_la = p.GetTokenStream().LA(1)

		if !(_la == GoEventerParserASSIGN || _la == GoEventerParserUPDATE) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*AssignVariableStatContext).op = _ri
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(103)
		p.expr(0)
	}
	{
		p.SetState(104)
		p.Match(GoEventerParserSCOL)
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

// IUpdateVariableStatContext is an interface to support dynamic dispatch.
type IUpdateVariableStatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllID() []antlr.TerminalNode
	ID(i int) antlr.TerminalNode
	UPDATE() antlr.TerminalNode
	Expr() IExprContext
	SCOL() antlr.TerminalNode
	AllDOT() []antlr.TerminalNode
	DOT(i int) antlr.TerminalNode

	// IsUpdateVariableStatContext differentiates from other interfaces.
	IsUpdateVariableStatContext()
}

type UpdateVariableStatContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUpdateVariableStatContext() *UpdateVariableStatContext {
	var p = new(UpdateVariableStatContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoEventerParserRULE_updateVariableStat
	return p
}

func InitEmptyUpdateVariableStatContext(p *UpdateVariableStatContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoEventerParserRULE_updateVariableStat
}

func (*UpdateVariableStatContext) IsUpdateVariableStatContext() {}

func NewUpdateVariableStatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UpdateVariableStatContext {
	var p = new(UpdateVariableStatContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoEventerParserRULE_updateVariableStat

	return p
}

func (s *UpdateVariableStatContext) GetParser() antlr.Parser { return s.parser }

func (s *UpdateVariableStatContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(GoEventerParserID)
}

func (s *UpdateVariableStatContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(GoEventerParserID, i)
}

func (s *UpdateVariableStatContext) UPDATE() antlr.TerminalNode {
	return s.GetToken(GoEventerParserUPDATE, 0)
}

func (s *UpdateVariableStatContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *UpdateVariableStatContext) SCOL() antlr.TerminalNode {
	return s.GetToken(GoEventerParserSCOL, 0)
}

func (s *UpdateVariableStatContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(GoEventerParserDOT)
}

func (s *UpdateVariableStatContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(GoEventerParserDOT, i)
}

func (s *UpdateVariableStatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UpdateVariableStatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UpdateVariableStatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoEventerListener); ok {
		listenerT.EnterUpdateVariableStat(s)
	}
}

func (s *UpdateVariableStatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoEventerListener); ok {
		listenerT.ExitUpdateVariableStat(s)
	}
}

func (s *UpdateVariableStatContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoEventerVisitor:
		return t.VisitUpdateVariableStat(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoEventerParser) UpdateVariableStat() (localctx IUpdateVariableStatContext) {
	localctx = NewUpdateVariableStatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, GoEventerParserRULE_updateVariableStat)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(106)
		p.Match(GoEventerParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(111)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == GoEventerParserDOT {
		{
			p.SetState(107)
			p.Match(GoEventerParserDOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(108)
			p.Match(GoEventerParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(113)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(114)
		p.Match(GoEventerParserUPDATE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(115)
		p.expr(0)
	}
	{
		p.SetState(116)
		p.Match(GoEventerParserSCOL)
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

// INotifyChanStatContext is an interface to support dynamic dispatch.
type INotifyChanStatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	NOTIFY() antlr.TerminalNode
	Expr() IExprContext
	SCOL() antlr.TerminalNode

	// IsNotifyChanStatContext differentiates from other interfaces.
	IsNotifyChanStatContext()
}

type NotifyChanStatContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNotifyChanStatContext() *NotifyChanStatContext {
	var p = new(NotifyChanStatContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoEventerParserRULE_notifyChanStat
	return p
}

func InitEmptyNotifyChanStatContext(p *NotifyChanStatContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoEventerParserRULE_notifyChanStat
}

func (*NotifyChanStatContext) IsNotifyChanStatContext() {}

func NewNotifyChanStatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NotifyChanStatContext {
	var p = new(NotifyChanStatContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoEventerParserRULE_notifyChanStat

	return p
}

func (s *NotifyChanStatContext) GetParser() antlr.Parser { return s.parser }

func (s *NotifyChanStatContext) ID() antlr.TerminalNode {
	return s.GetToken(GoEventerParserID, 0)
}

func (s *NotifyChanStatContext) NOTIFY() antlr.TerminalNode {
	return s.GetToken(GoEventerParserNOTIFY, 0)
}

func (s *NotifyChanStatContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *NotifyChanStatContext) SCOL() antlr.TerminalNode {
	return s.GetToken(GoEventerParserSCOL, 0)
}

func (s *NotifyChanStatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NotifyChanStatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NotifyChanStatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoEventerListener); ok {
		listenerT.EnterNotifyChanStat(s)
	}
}

func (s *NotifyChanStatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoEventerListener); ok {
		listenerT.ExitNotifyChanStat(s)
	}
}

func (s *NotifyChanStatContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoEventerVisitor:
		return t.VisitNotifyChanStat(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoEventerParser) NotifyChanStat() (localctx INotifyChanStatContext) {
	localctx = NewNotifyChanStatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, GoEventerParserRULE_notifyChanStat)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(118)
		p.Match(GoEventerParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(119)
		p.Match(GoEventerParserNOTIFY)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(120)
		p.expr(0)
	}
	{
		p.SetState(121)
		p.Match(GoEventerParserSCOL)
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

// IForStatContext is an interface to support dynamic dispatch.
type IForStatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FOR() antlr.TerminalNode
	ID() antlr.TerminalNode
	ASSIGN() antlr.TerminalNode
	AllExpr() []IExprContext
	Expr(i int) IExprContext
	COL() antlr.TerminalNode
	StatBlock() IStatBlockContext

	// IsForStatContext differentiates from other interfaces.
	IsForStatContext()
}

type ForStatContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyForStatContext() *ForStatContext {
	var p = new(ForStatContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoEventerParserRULE_forStat
	return p
}

func InitEmptyForStatContext(p *ForStatContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoEventerParserRULE_forStat
}

func (*ForStatContext) IsForStatContext() {}

func NewForStatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ForStatContext {
	var p = new(ForStatContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoEventerParserRULE_forStat

	return p
}

func (s *ForStatContext) GetParser() antlr.Parser { return s.parser }

func (s *ForStatContext) FOR() antlr.TerminalNode {
	return s.GetToken(GoEventerParserFOR, 0)
}

func (s *ForStatContext) ID() antlr.TerminalNode {
	return s.GetToken(GoEventerParserID, 0)
}

func (s *ForStatContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(GoEventerParserASSIGN, 0)
}

func (s *ForStatContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *ForStatContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
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

	return t.(IExprContext)
}

func (s *ForStatContext) COL() antlr.TerminalNode {
	return s.GetToken(GoEventerParserCOL, 0)
}

func (s *ForStatContext) StatBlock() IStatBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatBlockContext)
}

func (s *ForStatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForStatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ForStatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoEventerListener); ok {
		listenerT.EnterForStat(s)
	}
}

func (s *ForStatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoEventerListener); ok {
		listenerT.ExitForStat(s)
	}
}

func (s *ForStatContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoEventerVisitor:
		return t.VisitForStat(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoEventerParser) ForStat() (localctx IForStatContext) {
	localctx = NewForStatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, GoEventerParserRULE_forStat)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(123)
		p.Match(GoEventerParserFOR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(124)
		p.Match(GoEventerParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(125)
		p.Match(GoEventerParserASSIGN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(126)
		p.expr(0)
	}
	{
		p.SetState(127)
		p.Match(GoEventerParserCOL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(128)
		p.expr(0)
	}
	{
		p.SetState(129)
		p.StatBlock()
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

// IWhileStatContext is an interface to support dynamic dispatch.
type IWhileStatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	WHILE() antlr.TerminalNode
	Expr() IExprContext
	StatBlock() IStatBlockContext

	// IsWhileStatContext differentiates from other interfaces.
	IsWhileStatContext()
}

type WhileStatContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWhileStatContext() *WhileStatContext {
	var p = new(WhileStatContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoEventerParserRULE_whileStat
	return p
}

func InitEmptyWhileStatContext(p *WhileStatContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoEventerParserRULE_whileStat
}

func (*WhileStatContext) IsWhileStatContext() {}

func NewWhileStatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WhileStatContext {
	var p = new(WhileStatContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoEventerParserRULE_whileStat

	return p
}

func (s *WhileStatContext) GetParser() antlr.Parser { return s.parser }

func (s *WhileStatContext) WHILE() antlr.TerminalNode {
	return s.GetToken(GoEventerParserWHILE, 0)
}

func (s *WhileStatContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *WhileStatContext) StatBlock() IStatBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatBlockContext)
}

func (s *WhileStatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WhileStatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WhileStatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoEventerListener); ok {
		listenerT.EnterWhileStat(s)
	}
}

func (s *WhileStatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoEventerListener); ok {
		listenerT.ExitWhileStat(s)
	}
}

func (s *WhileStatContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoEventerVisitor:
		return t.VisitWhileStat(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoEventerParser) WhileStat() (localctx IWhileStatContext) {
	localctx = NewWhileStatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, GoEventerParserRULE_whileStat)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(131)
		p.Match(GoEventerParserWHILE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(132)
		p.expr(0)
	}
	{
		p.SetState(133)
		p.StatBlock()
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

// IDefineListenerStatContext is an interface to support dynamic dispatch.
type IDefineListenerStatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ON() antlr.TerminalNode
	MethodCall() IMethodCallContext
	StatBlock() IStatBlockContext
	ID() antlr.TerminalNode
	ASSIGN() antlr.TerminalNode

	// IsDefineListenerStatContext differentiates from other interfaces.
	IsDefineListenerStatContext()
}

type DefineListenerStatContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDefineListenerStatContext() *DefineListenerStatContext {
	var p = new(DefineListenerStatContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoEventerParserRULE_defineListenerStat
	return p
}

func InitEmptyDefineListenerStatContext(p *DefineListenerStatContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoEventerParserRULE_defineListenerStat
}

func (*DefineListenerStatContext) IsDefineListenerStatContext() {}

func NewDefineListenerStatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DefineListenerStatContext {
	var p = new(DefineListenerStatContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoEventerParserRULE_defineListenerStat

	return p
}

func (s *DefineListenerStatContext) GetParser() antlr.Parser { return s.parser }

func (s *DefineListenerStatContext) ON() antlr.TerminalNode {
	return s.GetToken(GoEventerParserON, 0)
}

func (s *DefineListenerStatContext) MethodCall() IMethodCallContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMethodCallContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMethodCallContext)
}

func (s *DefineListenerStatContext) StatBlock() IStatBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatBlockContext)
}

func (s *DefineListenerStatContext) ID() antlr.TerminalNode {
	return s.GetToken(GoEventerParserID, 0)
}

func (s *DefineListenerStatContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(GoEventerParserASSIGN, 0)
}

func (s *DefineListenerStatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DefineListenerStatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DefineListenerStatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoEventerListener); ok {
		listenerT.EnterDefineListenerStat(s)
	}
}

func (s *DefineListenerStatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoEventerListener); ok {
		listenerT.ExitDefineListenerStat(s)
	}
}

func (s *DefineListenerStatContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoEventerVisitor:
		return t.VisitDefineListenerStat(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoEventerParser) DefineListenerStat() (localctx IDefineListenerStatContext) {
	localctx = NewDefineListenerStatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, GoEventerParserRULE_defineListenerStat)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(135)
		p.Match(GoEventerParserON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(138)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(136)
			p.Match(GoEventerParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(137)
			p.Match(GoEventerParserASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	{
		p.SetState(140)
		p.MethodCall()
	}
	{
		p.SetState(141)
		p.StatBlock()
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

// IDefineFunctionStatContext is an interface to support dynamic dispatch.
type IDefineFunctionStatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FUNC() antlr.TerminalNode
	ID() antlr.TerminalNode
	PARAN_OPEN() antlr.TerminalNode
	FunctionDefineArguments() IFunctionDefineArgumentsContext
	PARAN_CLOSE() antlr.TerminalNode
	StatBlock() IStatBlockContext

	// IsDefineFunctionStatContext differentiates from other interfaces.
	IsDefineFunctionStatContext()
}

type DefineFunctionStatContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDefineFunctionStatContext() *DefineFunctionStatContext {
	var p = new(DefineFunctionStatContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoEventerParserRULE_defineFunctionStat
	return p
}

func InitEmptyDefineFunctionStatContext(p *DefineFunctionStatContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoEventerParserRULE_defineFunctionStat
}

func (*DefineFunctionStatContext) IsDefineFunctionStatContext() {}

func NewDefineFunctionStatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DefineFunctionStatContext {
	var p = new(DefineFunctionStatContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoEventerParserRULE_defineFunctionStat

	return p
}

func (s *DefineFunctionStatContext) GetParser() antlr.Parser { return s.parser }

func (s *DefineFunctionStatContext) FUNC() antlr.TerminalNode {
	return s.GetToken(GoEventerParserFUNC, 0)
}

func (s *DefineFunctionStatContext) ID() antlr.TerminalNode {
	return s.GetToken(GoEventerParserID, 0)
}

func (s *DefineFunctionStatContext) PARAN_OPEN() antlr.TerminalNode {
	return s.GetToken(GoEventerParserPARAN_OPEN, 0)
}

func (s *DefineFunctionStatContext) FunctionDefineArguments() IFunctionDefineArgumentsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionDefineArgumentsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionDefineArgumentsContext)
}

func (s *DefineFunctionStatContext) PARAN_CLOSE() antlr.TerminalNode {
	return s.GetToken(GoEventerParserPARAN_CLOSE, 0)
}

func (s *DefineFunctionStatContext) StatBlock() IStatBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatBlockContext)
}

func (s *DefineFunctionStatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DefineFunctionStatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DefineFunctionStatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoEventerListener); ok {
		listenerT.EnterDefineFunctionStat(s)
	}
}

func (s *DefineFunctionStatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoEventerListener); ok {
		listenerT.ExitDefineFunctionStat(s)
	}
}

func (s *DefineFunctionStatContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoEventerVisitor:
		return t.VisitDefineFunctionStat(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoEventerParser) DefineFunctionStat() (localctx IDefineFunctionStatContext) {
	localctx = NewDefineFunctionStatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, GoEventerParserRULE_defineFunctionStat)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(143)
		p.Match(GoEventerParserFUNC)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(144)
		p.Match(GoEventerParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(145)
		p.Match(GoEventerParserPARAN_OPEN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(146)
		p.FunctionDefineArguments()
	}
	{
		p.SetState(147)
		p.Match(GoEventerParserPARAN_CLOSE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(148)
		p.StatBlock()
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

// IStatBlockContext is an interface to support dynamic dispatch.
type IStatBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	BARACE_OPEN() antlr.TerminalNode
	Block() IBlockContext
	BARACE_CLOSE() antlr.TerminalNode
	Stat() IStatContext

	// IsStatBlockContext differentiates from other interfaces.
	IsStatBlockContext()
}

type StatBlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatBlockContext() *StatBlockContext {
	var p = new(StatBlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoEventerParserRULE_statBlock
	return p
}

func InitEmptyStatBlockContext(p *StatBlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoEventerParserRULE_statBlock
}

func (*StatBlockContext) IsStatBlockContext() {}

func NewStatBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatBlockContext {
	var p = new(StatBlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoEventerParserRULE_statBlock

	return p
}

func (s *StatBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *StatBlockContext) BARACE_OPEN() antlr.TerminalNode {
	return s.GetToken(GoEventerParserBARACE_OPEN, 0)
}

func (s *StatBlockContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *StatBlockContext) BARACE_CLOSE() antlr.TerminalNode {
	return s.GetToken(GoEventerParserBARACE_CLOSE, 0)
}

func (s *StatBlockContext) Stat() IStatContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatContext)
}

func (s *StatBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoEventerListener); ok {
		listenerT.EnterStatBlock(s)
	}
}

func (s *StatBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoEventerListener); ok {
		listenerT.ExitStatBlock(s)
	}
}

func (s *StatBlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoEventerVisitor:
		return t.VisitStatBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoEventerParser) StatBlock() (localctx IStatBlockContext) {
	localctx = NewStatBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, GoEventerParserRULE_statBlock)
	p.SetState(155)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case GoEventerParserBARACE_OPEN:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(150)
			p.Match(GoEventerParserBARACE_OPEN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(151)
			p.Block()
		}
		{
			p.SetState(152)
			p.Match(GoEventerParserBARACE_CLOSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case GoEventerParserFOR, GoEventerParserWHILE, GoEventerParserFUNC, GoEventerParserON, GoEventerParserRETURN, GoEventerParserIF, GoEventerParserSTRUCT, GoEventerParserID:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(154)
			p.Stat()
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

// IReturnStatContext is an interface to support dynamic dispatch.
type IReturnStatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	RETURN() antlr.TerminalNode
	Expr() IExprContext
	SCOL() antlr.TerminalNode

	// IsReturnStatContext differentiates from other interfaces.
	IsReturnStatContext()
}

type ReturnStatContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReturnStatContext() *ReturnStatContext {
	var p = new(ReturnStatContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoEventerParserRULE_returnStat
	return p
}

func InitEmptyReturnStatContext(p *ReturnStatContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoEventerParserRULE_returnStat
}

func (*ReturnStatContext) IsReturnStatContext() {}

func NewReturnStatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReturnStatContext {
	var p = new(ReturnStatContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoEventerParserRULE_returnStat

	return p
}

func (s *ReturnStatContext) GetParser() antlr.Parser { return s.parser }

func (s *ReturnStatContext) RETURN() antlr.TerminalNode {
	return s.GetToken(GoEventerParserRETURN, 0)
}

func (s *ReturnStatContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ReturnStatContext) SCOL() antlr.TerminalNode {
	return s.GetToken(GoEventerParserSCOL, 0)
}

func (s *ReturnStatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReturnStatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReturnStatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoEventerListener); ok {
		listenerT.EnterReturnStat(s)
	}
}

func (s *ReturnStatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoEventerListener); ok {
		listenerT.ExitReturnStat(s)
	}
}

func (s *ReturnStatContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoEventerVisitor:
		return t.VisitReturnStat(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoEventerParser) ReturnStat() (localctx IReturnStatContext) {
	localctx = NewReturnStatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, GoEventerParserRULE_returnStat)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(157)
		p.Match(GoEventerParserRETURN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(158)
		p.expr(0)
	}
	{
		p.SetState(159)
		p.Match(GoEventerParserSCOL)
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

// IMethodCallStatContext is an interface to support dynamic dispatch.
type IMethodCallStatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	MethodCall() IMethodCallContext
	SCOL() antlr.TerminalNode

	// IsMethodCallStatContext differentiates from other interfaces.
	IsMethodCallStatContext()
}

type MethodCallStatContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMethodCallStatContext() *MethodCallStatContext {
	var p = new(MethodCallStatContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoEventerParserRULE_methodCallStat
	return p
}

func InitEmptyMethodCallStatContext(p *MethodCallStatContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoEventerParserRULE_methodCallStat
}

func (*MethodCallStatContext) IsMethodCallStatContext() {}

func NewMethodCallStatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MethodCallStatContext {
	var p = new(MethodCallStatContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoEventerParserRULE_methodCallStat

	return p
}

func (s *MethodCallStatContext) GetParser() antlr.Parser { return s.parser }

func (s *MethodCallStatContext) MethodCall() IMethodCallContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMethodCallContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMethodCallContext)
}

func (s *MethodCallStatContext) SCOL() antlr.TerminalNode {
	return s.GetToken(GoEventerParserSCOL, 0)
}

func (s *MethodCallStatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MethodCallStatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MethodCallStatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoEventerListener); ok {
		listenerT.EnterMethodCallStat(s)
	}
}

func (s *MethodCallStatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoEventerListener); ok {
		listenerT.ExitMethodCallStat(s)
	}
}

func (s *MethodCallStatContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoEventerVisitor:
		return t.VisitMethodCallStat(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoEventerParser) MethodCallStat() (localctx IMethodCallStatContext) {
	localctx = NewMethodCallStatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, GoEventerParserRULE_methodCallStat)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(161)
		p.MethodCall()
	}
	{
		p.SetState(162)
		p.Match(GoEventerParserSCOL)
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

// IMethodCallContext is an interface to support dynamic dispatch.
type IMethodCallContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	PARAN_OPEN() antlr.TerminalNode
	MethodCallArguments() IMethodCallArgumentsContext
	PARAN_CLOSE() antlr.TerminalNode

	// IsMethodCallContext differentiates from other interfaces.
	IsMethodCallContext()
}

type MethodCallContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMethodCallContext() *MethodCallContext {
	var p = new(MethodCallContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoEventerParserRULE_methodCall
	return p
}

func InitEmptyMethodCallContext(p *MethodCallContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoEventerParserRULE_methodCall
}

func (*MethodCallContext) IsMethodCallContext() {}

func NewMethodCallContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MethodCallContext {
	var p = new(MethodCallContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoEventerParserRULE_methodCall

	return p
}

func (s *MethodCallContext) GetParser() antlr.Parser { return s.parser }

func (s *MethodCallContext) ID() antlr.TerminalNode {
	return s.GetToken(GoEventerParserID, 0)
}

func (s *MethodCallContext) PARAN_OPEN() antlr.TerminalNode {
	return s.GetToken(GoEventerParserPARAN_OPEN, 0)
}

func (s *MethodCallContext) MethodCallArguments() IMethodCallArgumentsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMethodCallArgumentsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMethodCallArgumentsContext)
}

func (s *MethodCallContext) PARAN_CLOSE() antlr.TerminalNode {
	return s.GetToken(GoEventerParserPARAN_CLOSE, 0)
}

func (s *MethodCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MethodCallContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MethodCallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoEventerListener); ok {
		listenerT.EnterMethodCall(s)
	}
}

func (s *MethodCallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoEventerListener); ok {
		listenerT.ExitMethodCall(s)
	}
}

func (s *MethodCallContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoEventerVisitor:
		return t.VisitMethodCall(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoEventerParser) MethodCall() (localctx IMethodCallContext) {
	localctx = NewMethodCallContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, GoEventerParserRULE_methodCall)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(164)
		p.Match(GoEventerParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(165)
		p.Match(GoEventerParserPARAN_OPEN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(166)
		p.MethodCallArguments()
	}
	{
		p.SetState(167)
		p.Match(GoEventerParserPARAN_CLOSE)
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

// IMethodCallArgumentsContext is an interface to support dynamic dispatch.
type IMethodCallArgumentsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpr() []IExprContext
	Expr(i int) IExprContext

	// IsMethodCallArgumentsContext differentiates from other interfaces.
	IsMethodCallArgumentsContext()
}

type MethodCallArgumentsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMethodCallArgumentsContext() *MethodCallArgumentsContext {
	var p = new(MethodCallArgumentsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoEventerParserRULE_methodCallArguments
	return p
}

func InitEmptyMethodCallArgumentsContext(p *MethodCallArgumentsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoEventerParserRULE_methodCallArguments
}

func (*MethodCallArgumentsContext) IsMethodCallArgumentsContext() {}

func NewMethodCallArgumentsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MethodCallArgumentsContext {
	var p = new(MethodCallArgumentsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoEventerParserRULE_methodCallArguments

	return p
}

func (s *MethodCallArgumentsContext) GetParser() antlr.Parser { return s.parser }

func (s *MethodCallArgumentsContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *MethodCallArgumentsContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
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

	return t.(IExprContext)
}

func (s *MethodCallArgumentsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MethodCallArgumentsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MethodCallArgumentsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoEventerListener); ok {
		listenerT.EnterMethodCallArguments(s)
	}
}

func (s *MethodCallArgumentsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoEventerListener); ok {
		listenerT.ExitMethodCallArguments(s)
	}
}

func (s *MethodCallArgumentsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoEventerVisitor:
		return t.VisitMethodCallArguments(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoEventerParser) MethodCallArguments() (localctx IMethodCallArgumentsContext) {
	localctx = NewMethodCallArgumentsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, GoEventerParserRULE_methodCallArguments)
	var _la int

	p.SetState(178)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case GoEventerParserPARAN_CLOSE:
		p.EnterOuterAlt(localctx, 1)

	case GoEventerParserCHAN, GoEventerParserMINUS, GoEventerParserNOT, GoEventerParserTRUE, GoEventerParserFALSE, GoEventerParserNIL, GoEventerParserPARAN_OPEN, GoEventerParserID, GoEventerParserINT, GoEventerParserFLOAT, GoEventerParserSTRING:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(170)
			p.expr(0)
		}
		p.SetState(175)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == GoEventerParserT__0 {
			{
				p.SetState(171)
				p.Match(GoEventerParserT__0)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(172)
				p.expr(0)
			}

			p.SetState(177)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
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

// IFunctionDefineArgumentsContext is an interface to support dynamic dispatch.
type IFunctionDefineArgumentsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllID() []antlr.TerminalNode
	ID(i int) antlr.TerminalNode

	// IsFunctionDefineArgumentsContext differentiates from other interfaces.
	IsFunctionDefineArgumentsContext()
}

type FunctionDefineArgumentsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionDefineArgumentsContext() *FunctionDefineArgumentsContext {
	var p = new(FunctionDefineArgumentsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoEventerParserRULE_functionDefineArguments
	return p
}

func InitEmptyFunctionDefineArgumentsContext(p *FunctionDefineArgumentsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoEventerParserRULE_functionDefineArguments
}

func (*FunctionDefineArgumentsContext) IsFunctionDefineArgumentsContext() {}

func NewFunctionDefineArgumentsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionDefineArgumentsContext {
	var p = new(FunctionDefineArgumentsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoEventerParserRULE_functionDefineArguments

	return p
}

func (s *FunctionDefineArgumentsContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionDefineArgumentsContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(GoEventerParserID)
}

func (s *FunctionDefineArgumentsContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(GoEventerParserID, i)
}

func (s *FunctionDefineArgumentsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionDefineArgumentsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionDefineArgumentsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoEventerListener); ok {
		listenerT.EnterFunctionDefineArguments(s)
	}
}

func (s *FunctionDefineArgumentsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoEventerListener); ok {
		listenerT.ExitFunctionDefineArguments(s)
	}
}

func (s *FunctionDefineArgumentsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoEventerVisitor:
		return t.VisitFunctionDefineArguments(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoEventerParser) FunctionDefineArguments() (localctx IFunctionDefineArgumentsContext) {
	localctx = NewFunctionDefineArgumentsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, GoEventerParserRULE_functionDefineArguments)
	var _la int

	p.SetState(189)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case GoEventerParserPARAN_CLOSE:
		p.EnterOuterAlt(localctx, 1)

	case GoEventerParserID:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(181)
			p.Match(GoEventerParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(186)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == GoEventerParserT__0 {
			{
				p.SetState(182)
				p.Match(GoEventerParserT__0)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(183)
				p.Match(GoEventerParserID)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(188)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
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

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoEventerParserRULE_expr
	return p
}

func InitEmptyExprContext(p *ExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoEventerParserRULE_expr
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoEventerParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) CopyAll(ctx *ExprContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type MethodCallExprContext struct {
	ExprContext
}

func NewMethodCallExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MethodCallExprContext {
	var p = new(MethodCallExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *MethodCallExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MethodCallExprContext) MethodCall() IMethodCallContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMethodCallContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMethodCallContext)
}

func (s *MethodCallExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoEventerListener); ok {
		listenerT.EnterMethodCallExpr(s)
	}
}

func (s *MethodCallExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoEventerListener); ok {
		listenerT.ExitMethodCallExpr(s)
	}
}

func (s *MethodCallExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoEventerVisitor:
		return t.VisitMethodCallExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type NotExprContext struct {
	ExprContext
}

func NewNotExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NotExprContext {
	var p = new(NotExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *NotExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NotExprContext) NOT() antlr.TerminalNode {
	return s.GetToken(GoEventerParserNOT, 0)
}

func (s *NotExprContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *NotExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoEventerListener); ok {
		listenerT.EnterNotExpr(s)
	}
}

func (s *NotExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoEventerListener); ok {
		listenerT.ExitNotExpr(s)
	}
}

func (s *NotExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoEventerVisitor:
		return t.VisitNotExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type UnaryMinusExprContext struct {
	ExprContext
}

func NewUnaryMinusExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *UnaryMinusExprContext {
	var p = new(UnaryMinusExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *UnaryMinusExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnaryMinusExprContext) MINUS() antlr.TerminalNode {
	return s.GetToken(GoEventerParserMINUS, 0)
}

func (s *UnaryMinusExprContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *UnaryMinusExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoEventerListener); ok {
		listenerT.EnterUnaryMinusExpr(s)
	}
}

func (s *UnaryMinusExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoEventerListener); ok {
		listenerT.ExitUnaryMinusExpr(s)
	}
}

func (s *UnaryMinusExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoEventerVisitor:
		return t.VisitUnaryMinusExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type MultiplicationExprContext struct {
	ExprContext
	op antlr.Token
}

func NewMultiplicationExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MultiplicationExprContext {
	var p = new(MultiplicationExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *MultiplicationExprContext) GetOp() antlr.Token { return s.op }

func (s *MultiplicationExprContext) SetOp(v antlr.Token) { s.op = v }

func (s *MultiplicationExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MultiplicationExprContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *MultiplicationExprContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
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

	return t.(IExprContext)
}

func (s *MultiplicationExprContext) MULT() antlr.TerminalNode {
	return s.GetToken(GoEventerParserMULT, 0)
}

func (s *MultiplicationExprContext) DIV() antlr.TerminalNode {
	return s.GetToken(GoEventerParserDIV, 0)
}

func (s *MultiplicationExprContext) MOD() antlr.TerminalNode {
	return s.GetToken(GoEventerParserMOD, 0)
}

func (s *MultiplicationExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoEventerListener); ok {
		listenerT.EnterMultiplicationExpr(s)
	}
}

func (s *MultiplicationExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoEventerListener); ok {
		listenerT.ExitMultiplicationExpr(s)
	}
}

func (s *MultiplicationExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoEventerVisitor:
		return t.VisitMultiplicationExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type AtomExprContext struct {
	ExprContext
}

func NewAtomExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AtomExprContext {
	var p = new(AtomExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *AtomExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AtomExprContext) Atom() IAtomContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAtomContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAtomContext)
}

func (s *AtomExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoEventerListener); ok {
		listenerT.EnterAtomExpr(s)
	}
}

func (s *AtomExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoEventerListener); ok {
		listenerT.ExitAtomExpr(s)
	}
}

func (s *AtomExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoEventerVisitor:
		return t.VisitAtomExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type OrExprContext struct {
	ExprContext
}

func NewOrExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *OrExprContext {
	var p = new(OrExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *OrExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OrExprContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *OrExprContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
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

	return t.(IExprContext)
}

func (s *OrExprContext) OR() antlr.TerminalNode {
	return s.GetToken(GoEventerParserOR, 0)
}

func (s *OrExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoEventerListener); ok {
		listenerT.EnterOrExpr(s)
	}
}

func (s *OrExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoEventerListener); ok {
		listenerT.ExitOrExpr(s)
	}
}

func (s *OrExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoEventerVisitor:
		return t.VisitOrExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type AdditiveExprContext struct {
	ExprContext
	op antlr.Token
}

func NewAdditiveExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AdditiveExprContext {
	var p = new(AdditiveExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *AdditiveExprContext) GetOp() antlr.Token { return s.op }

func (s *AdditiveExprContext) SetOp(v antlr.Token) { s.op = v }

func (s *AdditiveExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AdditiveExprContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *AdditiveExprContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
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

	return t.(IExprContext)
}

func (s *AdditiveExprContext) PLUS() antlr.TerminalNode {
	return s.GetToken(GoEventerParserPLUS, 0)
}

func (s *AdditiveExprContext) MINUS() antlr.TerminalNode {
	return s.GetToken(GoEventerParserMINUS, 0)
}

func (s *AdditiveExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoEventerListener); ok {
		listenerT.EnterAdditiveExpr(s)
	}
}

func (s *AdditiveExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoEventerListener); ok {
		listenerT.ExitAdditiveExpr(s)
	}
}

func (s *AdditiveExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoEventerVisitor:
		return t.VisitAdditiveExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type PowExprContext struct {
	ExprContext
}

func NewPowExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PowExprContext {
	var p = new(PowExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *PowExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PowExprContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *PowExprContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
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

	return t.(IExprContext)
}

func (s *PowExprContext) POW() antlr.TerminalNode {
	return s.GetToken(GoEventerParserPOW, 0)
}

func (s *PowExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoEventerListener); ok {
		listenerT.EnterPowExpr(s)
	}
}

func (s *PowExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoEventerListener); ok {
		listenerT.ExitPowExpr(s)
	}
}

func (s *PowExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoEventerVisitor:
		return t.VisitPowExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type RelationalExprContext struct {
	ExprContext
	op antlr.Token
}

func NewRelationalExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RelationalExprContext {
	var p = new(RelationalExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *RelationalExprContext) GetOp() antlr.Token { return s.op }

func (s *RelationalExprContext) SetOp(v antlr.Token) { s.op = v }

func (s *RelationalExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RelationalExprContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *RelationalExprContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
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

	return t.(IExprContext)
}

func (s *RelationalExprContext) LTEQ() antlr.TerminalNode {
	return s.GetToken(GoEventerParserLTEQ, 0)
}

func (s *RelationalExprContext) GTEQ() antlr.TerminalNode {
	return s.GetToken(GoEventerParserGTEQ, 0)
}

func (s *RelationalExprContext) LT() antlr.TerminalNode {
	return s.GetToken(GoEventerParserLT, 0)
}

func (s *RelationalExprContext) GT() antlr.TerminalNode {
	return s.GetToken(GoEventerParserGT, 0)
}

func (s *RelationalExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoEventerListener); ok {
		listenerT.EnterRelationalExpr(s)
	}
}

func (s *RelationalExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoEventerListener); ok {
		listenerT.ExitRelationalExpr(s)
	}
}

func (s *RelationalExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoEventerVisitor:
		return t.VisitRelationalExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type EqualityExprContext struct {
	ExprContext
	op antlr.Token
}

func NewEqualityExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EqualityExprContext {
	var p = new(EqualityExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *EqualityExprContext) GetOp() antlr.Token { return s.op }

func (s *EqualityExprContext) SetOp(v antlr.Token) { s.op = v }

func (s *EqualityExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EqualityExprContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *EqualityExprContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
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

	return t.(IExprContext)
}

func (s *EqualityExprContext) EQ() antlr.TerminalNode {
	return s.GetToken(GoEventerParserEQ, 0)
}

func (s *EqualityExprContext) NEQ() antlr.TerminalNode {
	return s.GetToken(GoEventerParserNEQ, 0)
}

func (s *EqualityExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoEventerListener); ok {
		listenerT.EnterEqualityExpr(s)
	}
}

func (s *EqualityExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoEventerListener); ok {
		listenerT.ExitEqualityExpr(s)
	}
}

func (s *EqualityExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoEventerVisitor:
		return t.VisitEqualityExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type AndExprContext struct {
	ExprContext
}

func NewAndExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AndExprContext {
	var p = new(AndExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *AndExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AndExprContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *AndExprContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
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

	return t.(IExprContext)
}

func (s *AndExprContext) AND() antlr.TerminalNode {
	return s.GetToken(GoEventerParserAND, 0)
}

func (s *AndExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoEventerListener); ok {
		listenerT.EnterAndExpr(s)
	}
}

func (s *AndExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoEventerListener); ok {
		listenerT.ExitAndExpr(s)
	}
}

func (s *AndExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoEventerVisitor:
		return t.VisitAndExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoEventerParser) Expr() (localctx IExprContext) {
	return p.expr(0)
}

func (p *GoEventerParser) expr(_p int) (localctx IExprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 40
	p.EnterRecursionRule(localctx, 40, GoEventerParserRULE_expr, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(198)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 12, p.GetParserRuleContext()) {
	case 1:
		localctx = NewMethodCallExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(192)
			p.MethodCall()
		}

	case 2:
		localctx = NewUnaryMinusExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(193)
			p.Match(GoEventerParserMINUS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(194)
			p.expr(9)
		}

	case 3:
		localctx = NewNotExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(195)
			p.Match(GoEventerParserNOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(196)
			p.expr(8)
		}

	case 4:
		localctx = NewAtomExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(197)
			p.Atom()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(223)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 14, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(221)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 13, p.GetParserRuleContext()) {
			case 1:
				localctx = NewPowExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, GoEventerParserRULE_expr)
				p.SetState(200)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
					goto errorExit
				}
				{
					p.SetState(201)
					p.Match(GoEventerParserPOW)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(202)
					p.expr(10)
				}

			case 2:
				localctx = NewMultiplicationExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, GoEventerParserRULE_expr)
				p.SetState(203)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
					goto errorExit
				}
				{
					p.SetState(204)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*MultiplicationExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&7516192768) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*MultiplicationExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(205)
					p.expr(8)
				}

			case 3:
				localctx = NewAdditiveExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, GoEventerParserRULE_expr)
				p.SetState(206)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
					goto errorExit
				}
				{
					p.SetState(207)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*AdditiveExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == GoEventerParserPLUS || _la == GoEventerParserMINUS) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*AdditiveExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(208)
					p.expr(7)
				}

			case 4:
				localctx = NewRelationalExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, GoEventerParserRULE_expr)
				p.SetState(209)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
					goto errorExit
				}
				{
					p.SetState(210)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*RelationalExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&251658240) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*RelationalExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(211)
					p.expr(6)
				}

			case 5:
				localctx = NewEqualityExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, GoEventerParserRULE_expr)
				p.SetState(212)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
					goto errorExit
				}
				{
					p.SetState(213)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*EqualityExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == GoEventerParserEQ || _la == GoEventerParserNEQ) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*EqualityExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(214)
					p.expr(5)
				}

			case 6:
				localctx = NewAndExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, GoEventerParserRULE_expr)
				p.SetState(215)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
					goto errorExit
				}
				{
					p.SetState(216)
					p.Match(GoEventerParserAND)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(217)
					p.expr(4)
				}

			case 7:
				localctx = NewOrExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, GoEventerParserRULE_expr)
				p.SetState(218)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
					goto errorExit
				}
				{
					p.SetState(219)
					p.Match(GoEventerParserOR)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(220)
					p.expr(3)
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(225)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 14, p.GetParserRuleContext())
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

// IAtomContext is an interface to support dynamic dispatch.
type IAtomContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsAtomContext differentiates from other interfaces.
	IsAtomContext()
}

type AtomContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAtomContext() *AtomContext {
	var p = new(AtomContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoEventerParserRULE_atom
	return p
}

func InitEmptyAtomContext(p *AtomContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoEventerParserRULE_atom
}

func (*AtomContext) IsAtomContext() {}

func NewAtomContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AtomContext {
	var p = new(AtomContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoEventerParserRULE_atom

	return p
}

func (s *AtomContext) GetParser() antlr.Parser { return s.parser }

func (s *AtomContext) CopyAll(ctx *AtomContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *AtomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AtomContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ExprAtomContext struct {
	AtomContext
}

func NewExprAtomContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExprAtomContext {
	var p = new(ExprAtomContext)

	InitEmptyAtomContext(&p.AtomContext)
	p.parser = parser
	p.CopyAll(ctx.(*AtomContext))

	return p
}

func (s *ExprAtomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprAtomContext) PARAN_OPEN() antlr.TerminalNode {
	return s.GetToken(GoEventerParserPARAN_OPEN, 0)
}

func (s *ExprAtomContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ExprAtomContext) PARAN_CLOSE() antlr.TerminalNode {
	return s.GetToken(GoEventerParserPARAN_CLOSE, 0)
}

func (s *ExprAtomContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoEventerListener); ok {
		listenerT.EnterExprAtom(s)
	}
}

func (s *ExprAtomContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoEventerListener); ok {
		listenerT.ExitExprAtom(s)
	}
}

func (s *ExprAtomContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoEventerVisitor:
		return t.VisitExprAtom(s)

	default:
		return t.VisitChildren(s)
	}
}

type BooleanAtomContext struct {
	AtomContext
}

func NewBooleanAtomContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BooleanAtomContext {
	var p = new(BooleanAtomContext)

	InitEmptyAtomContext(&p.AtomContext)
	p.parser = parser
	p.CopyAll(ctx.(*AtomContext))

	return p
}

func (s *BooleanAtomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BooleanAtomContext) TRUE() antlr.TerminalNode {
	return s.GetToken(GoEventerParserTRUE, 0)
}

func (s *BooleanAtomContext) FALSE() antlr.TerminalNode {
	return s.GetToken(GoEventerParserFALSE, 0)
}

func (s *BooleanAtomContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoEventerListener); ok {
		listenerT.EnterBooleanAtom(s)
	}
}

func (s *BooleanAtomContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoEventerListener); ok {
		listenerT.ExitBooleanAtom(s)
	}
}

func (s *BooleanAtomContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoEventerVisitor:
		return t.VisitBooleanAtom(s)

	default:
		return t.VisitChildren(s)
	}
}

type IdAtomContext struct {
	AtomContext
}

func NewIdAtomContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IdAtomContext {
	var p = new(IdAtomContext)

	InitEmptyAtomContext(&p.AtomContext)
	p.parser = parser
	p.CopyAll(ctx.(*AtomContext))

	return p
}

func (s *IdAtomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdAtomContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(GoEventerParserID)
}

func (s *IdAtomContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(GoEventerParserID, i)
}

func (s *IdAtomContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(GoEventerParserDOT)
}

func (s *IdAtomContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(GoEventerParserDOT, i)
}

func (s *IdAtomContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoEventerListener); ok {
		listenerT.EnterIdAtom(s)
	}
}

func (s *IdAtomContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoEventerListener); ok {
		listenerT.ExitIdAtom(s)
	}
}

func (s *IdAtomContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoEventerVisitor:
		return t.VisitIdAtom(s)

	default:
		return t.VisitChildren(s)
	}
}

type StringAtomContext struct {
	AtomContext
}

func NewStringAtomContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StringAtomContext {
	var p = new(StringAtomContext)

	InitEmptyAtomContext(&p.AtomContext)
	p.parser = parser
	p.CopyAll(ctx.(*AtomContext))

	return p
}

func (s *StringAtomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringAtomContext) STRING() antlr.TerminalNode {
	return s.GetToken(GoEventerParserSTRING, 0)
}

func (s *StringAtomContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoEventerListener); ok {
		listenerT.EnterStringAtom(s)
	}
}

func (s *StringAtomContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoEventerListener); ok {
		listenerT.ExitStringAtom(s)
	}
}

func (s *StringAtomContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoEventerVisitor:
		return t.VisitStringAtom(s)

	default:
		return t.VisitChildren(s)
	}
}

type MakeChanAtomContext struct {
	AtomContext
}

func NewMakeChanAtomContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MakeChanAtomContext {
	var p = new(MakeChanAtomContext)

	InitEmptyAtomContext(&p.AtomContext)
	p.parser = parser
	p.CopyAll(ctx.(*AtomContext))

	return p
}

func (s *MakeChanAtomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MakeChanAtomContext) CHAN() antlr.TerminalNode {
	return s.GetToken(GoEventerParserCHAN, 0)
}

func (s *MakeChanAtomContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoEventerListener); ok {
		listenerT.EnterMakeChanAtom(s)
	}
}

func (s *MakeChanAtomContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoEventerListener); ok {
		listenerT.ExitMakeChanAtom(s)
	}
}

func (s *MakeChanAtomContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoEventerVisitor:
		return t.VisitMakeChanAtom(s)

	default:
		return t.VisitChildren(s)
	}
}

type NilAtomContext struct {
	AtomContext
}

func NewNilAtomContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NilAtomContext {
	var p = new(NilAtomContext)

	InitEmptyAtomContext(&p.AtomContext)
	p.parser = parser
	p.CopyAll(ctx.(*AtomContext))

	return p
}

func (s *NilAtomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NilAtomContext) NIL() antlr.TerminalNode {
	return s.GetToken(GoEventerParserNIL, 0)
}

func (s *NilAtomContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoEventerListener); ok {
		listenerT.EnterNilAtom(s)
	}
}

func (s *NilAtomContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoEventerListener); ok {
		listenerT.ExitNilAtom(s)
	}
}

func (s *NilAtomContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoEventerVisitor:
		return t.VisitNilAtom(s)

	default:
		return t.VisitChildren(s)
	}
}

type NumberAtomContext struct {
	AtomContext
}

func NewNumberAtomContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NumberAtomContext {
	var p = new(NumberAtomContext)

	InitEmptyAtomContext(&p.AtomContext)
	p.parser = parser
	p.CopyAll(ctx.(*AtomContext))

	return p
}

func (s *NumberAtomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumberAtomContext) INT() antlr.TerminalNode {
	return s.GetToken(GoEventerParserINT, 0)
}

func (s *NumberAtomContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(GoEventerParserFLOAT, 0)
}

func (s *NumberAtomContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoEventerListener); ok {
		listenerT.EnterNumberAtom(s)
	}
}

func (s *NumberAtomContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoEventerListener); ok {
		listenerT.ExitNumberAtom(s)
	}
}

func (s *NumberAtomContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoEventerVisitor:
		return t.VisitNumberAtom(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoEventerParser) Atom() (localctx IAtomContext) {
	localctx = NewAtomContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, GoEventerParserRULE_atom)
	var _la int

	var _alt int

	p.SetState(243)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case GoEventerParserPARAN_OPEN:
		localctx = NewExprAtomContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(226)
			p.Match(GoEventerParserPARAN_OPEN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(227)
			p.expr(0)
		}
		{
			p.SetState(228)
			p.Match(GoEventerParserPARAN_CLOSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case GoEventerParserINT, GoEventerParserFLOAT:
		localctx = NewNumberAtomContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(230)
			_la = p.GetTokenStream().LA(1)

			if !(_la == GoEventerParserINT || _la == GoEventerParserFLOAT) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	case GoEventerParserTRUE, GoEventerParserFALSE:
		localctx = NewBooleanAtomContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(231)
			_la = p.GetTokenStream().LA(1)

			if !(_la == GoEventerParserTRUE || _la == GoEventerParserFALSE) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	case GoEventerParserID:
		localctx = NewIdAtomContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(232)
			p.Match(GoEventerParserID)
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
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 15, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(233)
					p.Match(GoEventerParserDOT)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(234)
					p.Match(GoEventerParserID)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			}
			p.SetState(239)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 15, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}

	case GoEventerParserSTRING:
		localctx = NewStringAtomContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(240)
			p.Match(GoEventerParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case GoEventerParserNIL:
		localctx = NewNilAtomContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(241)
			p.Match(GoEventerParserNIL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case GoEventerParserCHAN:
		localctx = NewMakeChanAtomContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(242)
			p.Match(GoEventerParserCHAN)
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

func (p *GoEventerParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 20:
		var t *ExprContext = nil
		if localctx != nil {
			t = localctx.(*ExprContext)
		}
		return p.Expr_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *GoEventerParser) Expr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 10)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 7)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
