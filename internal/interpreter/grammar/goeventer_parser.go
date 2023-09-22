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
		"parse", "block", "stat", "assignVariable", "notifyChanStat", "forStat",
		"whileStat", "defineListener", "defineFunction", "statBlock", "returnStat",
		"methodCallStat", "methodCall", "methodCallArguments", "functionDefineArguments",
		"expr", "atom",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 42, 175, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 1, 0, 1, 0, 1, 0, 1, 1, 5, 1, 39, 8, 1, 10, 1, 12, 1, 42,
		9, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 52, 8, 2, 1,
		3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1,
		5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1,
		7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1,
		9, 1, 9, 3, 9, 92, 8, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11,
		1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 1, 13, 5, 13, 110,
		8, 13, 10, 13, 12, 13, 113, 9, 13, 3, 13, 115, 8, 13, 1, 14, 1, 14, 1,
		14, 1, 14, 5, 14, 121, 8, 14, 10, 14, 12, 14, 124, 9, 14, 3, 14, 126, 8,
		14, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 3, 15, 135, 8, 15,
		1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1,
		15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15,
		5, 15, 158, 8, 15, 10, 15, 12, 15, 161, 9, 15, 1, 16, 1, 16, 1, 16, 1,
		16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 3, 16, 173, 8, 16, 1, 16,
		0, 1, 30, 17, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30,
		32, 0, 7, 1, 0, 2, 3, 1, 0, 26, 28, 1, 0, 24, 25, 1, 0, 20, 23, 1, 0, 18,
		19, 1, 0, 39, 40, 1, 0, 31, 32, 186, 0, 34, 1, 0, 0, 0, 2, 40, 1, 0, 0,
		0, 4, 51, 1, 0, 0, 0, 6, 53, 1, 0, 0, 0, 8, 58, 1, 0, 0, 0, 10, 63, 1,
		0, 0, 0, 12, 71, 1, 0, 0, 0, 14, 75, 1, 0, 0, 0, 16, 79, 1, 0, 0, 0, 18,
		91, 1, 0, 0, 0, 20, 93, 1, 0, 0, 0, 22, 97, 1, 0, 0, 0, 24, 100, 1, 0,
		0, 0, 26, 114, 1, 0, 0, 0, 28, 125, 1, 0, 0, 0, 30, 134, 1, 0, 0, 0, 32,
		172, 1, 0, 0, 0, 34, 35, 3, 2, 1, 0, 35, 36, 5, 0, 0, 1, 36, 1, 1, 0, 0,
		0, 37, 39, 3, 4, 2, 0, 38, 37, 1, 0, 0, 0, 39, 42, 1, 0, 0, 0, 40, 38,
		1, 0, 0, 0, 40, 41, 1, 0, 0, 0, 41, 3, 1, 0, 0, 0, 42, 40, 1, 0, 0, 0,
		43, 52, 3, 22, 11, 0, 44, 52, 3, 14, 7, 0, 45, 52, 3, 6, 3, 0, 46, 52,
		3, 10, 5, 0, 47, 52, 3, 12, 6, 0, 48, 52, 3, 16, 8, 0, 49, 52, 3, 20, 10,
		0, 50, 52, 3, 8, 4, 0, 51, 43, 1, 0, 0, 0, 51, 44, 1, 0, 0, 0, 51, 45,
		1, 0, 0, 0, 51, 46, 1, 0, 0, 0, 51, 47, 1, 0, 0, 0, 51, 48, 1, 0, 0, 0,
		51, 49, 1, 0, 0, 0, 51, 50, 1, 0, 0, 0, 52, 5, 1, 0, 0, 0, 53, 54, 5, 38,
		0, 0, 54, 55, 7, 0, 0, 0, 55, 56, 3, 30, 15, 0, 56, 57, 5, 4, 0, 0, 57,
		7, 1, 0, 0, 0, 58, 59, 5, 38, 0, 0, 59, 60, 5, 13, 0, 0, 60, 61, 3, 30,
		15, 0, 61, 62, 5, 4, 0, 0, 62, 9, 1, 0, 0, 0, 63, 64, 5, 6, 0, 0, 64, 65,
		5, 38, 0, 0, 65, 66, 5, 2, 0, 0, 66, 67, 3, 30, 15, 0, 67, 68, 5, 5, 0,
		0, 68, 69, 3, 30, 15, 0, 69, 70, 3, 18, 9, 0, 70, 11, 1, 0, 0, 0, 71, 72,
		5, 7, 0, 0, 72, 73, 3, 30, 15, 0, 73, 74, 3, 18, 9, 0, 74, 13, 1, 0, 0,
		0, 75, 76, 5, 10, 0, 0, 76, 77, 3, 24, 12, 0, 77, 78, 3, 18, 9, 0, 78,
		15, 1, 0, 0, 0, 79, 80, 5, 9, 0, 0, 80, 81, 5, 38, 0, 0, 81, 82, 5, 34,
		0, 0, 82, 83, 3, 28, 14, 0, 83, 84, 5, 35, 0, 0, 84, 85, 3, 18, 9, 0, 85,
		17, 1, 0, 0, 0, 86, 87, 5, 36, 0, 0, 87, 88, 3, 2, 1, 0, 88, 89, 5, 37,
		0, 0, 89, 92, 1, 0, 0, 0, 90, 92, 3, 4, 2, 0, 91, 86, 1, 0, 0, 0, 91, 90,
		1, 0, 0, 0, 92, 19, 1, 0, 0, 0, 93, 94, 5, 11, 0, 0, 94, 95, 3, 30, 15,
		0, 95, 96, 5, 4, 0, 0, 96, 21, 1, 0, 0, 0, 97, 98, 3, 24, 12, 0, 98, 99,
		5, 4, 0, 0, 99, 23, 1, 0, 0, 0, 100, 101, 5, 38, 0, 0, 101, 102, 5, 34,
		0, 0, 102, 103, 3, 26, 13, 0, 103, 104, 5, 35, 0, 0, 104, 25, 1, 0, 0,
		0, 105, 115, 1, 0, 0, 0, 106, 111, 3, 30, 15, 0, 107, 108, 5, 1, 0, 0,
		108, 110, 3, 30, 15, 0, 109, 107, 1, 0, 0, 0, 110, 113, 1, 0, 0, 0, 111,
		109, 1, 0, 0, 0, 111, 112, 1, 0, 0, 0, 112, 115, 1, 0, 0, 0, 113, 111,
		1, 0, 0, 0, 114, 105, 1, 0, 0, 0, 114, 106, 1, 0, 0, 0, 115, 27, 1, 0,
		0, 0, 116, 126, 1, 0, 0, 0, 117, 122, 5, 38, 0, 0, 118, 119, 5, 1, 0, 0,
		119, 121, 5, 38, 0, 0, 120, 118, 1, 0, 0, 0, 121, 124, 1, 0, 0, 0, 122,
		120, 1, 0, 0, 0, 122, 123, 1, 0, 0, 0, 123, 126, 1, 0, 0, 0, 124, 122,
		1, 0, 0, 0, 125, 116, 1, 0, 0, 0, 125, 117, 1, 0, 0, 0, 126, 29, 1, 0,
		0, 0, 127, 128, 6, 15, -1, 0, 128, 135, 3, 24, 12, 0, 129, 130, 5, 25,
		0, 0, 130, 135, 3, 30, 15, 9, 131, 132, 5, 30, 0, 0, 132, 135, 3, 30, 15,
		8, 133, 135, 3, 32, 16, 0, 134, 127, 1, 0, 0, 0, 134, 129, 1, 0, 0, 0,
		134, 131, 1, 0, 0, 0, 134, 133, 1, 0, 0, 0, 135, 159, 1, 0, 0, 0, 136,
		137, 10, 10, 0, 0, 137, 138, 5, 29, 0, 0, 138, 158, 3, 30, 15, 10, 139,
		140, 10, 7, 0, 0, 140, 141, 7, 1, 0, 0, 141, 158, 3, 30, 15, 8, 142, 143,
		10, 6, 0, 0, 143, 144, 7, 2, 0, 0, 144, 158, 3, 30, 15, 7, 145, 146, 10,
		5, 0, 0, 146, 147, 7, 3, 0, 0, 147, 158, 3, 30, 15, 6, 148, 149, 10, 4,
		0, 0, 149, 150, 7, 4, 0, 0, 150, 158, 3, 30, 15, 5, 151, 152, 10, 3, 0,
		0, 152, 153, 5, 17, 0, 0, 153, 158, 3, 30, 15, 4, 154, 155, 10, 2, 0, 0,
		155, 156, 5, 16, 0, 0, 156, 158, 3, 30, 15, 3, 157, 136, 1, 0, 0, 0, 157,
		139, 1, 0, 0, 0, 157, 142, 1, 0, 0, 0, 157, 145, 1, 0, 0, 0, 157, 148,
		1, 0, 0, 0, 157, 151, 1, 0, 0, 0, 157, 154, 1, 0, 0, 0, 158, 161, 1, 0,
		0, 0, 159, 157, 1, 0, 0, 0, 159, 160, 1, 0, 0, 0, 160, 31, 1, 0, 0, 0,
		161, 159, 1, 0, 0, 0, 162, 163, 5, 34, 0, 0, 163, 164, 3, 30, 15, 0, 164,
		165, 5, 35, 0, 0, 165, 173, 1, 0, 0, 0, 166, 173, 7, 5, 0, 0, 167, 173,
		7, 6, 0, 0, 168, 173, 5, 38, 0, 0, 169, 173, 5, 41, 0, 0, 170, 173, 5,
		33, 0, 0, 171, 173, 5, 12, 0, 0, 172, 162, 1, 0, 0, 0, 172, 166, 1, 0,
		0, 0, 172, 167, 1, 0, 0, 0, 172, 168, 1, 0, 0, 0, 172, 169, 1, 0, 0, 0,
		172, 170, 1, 0, 0, 0, 172, 171, 1, 0, 0, 0, 173, 33, 1, 0, 0, 0, 11, 40,
		51, 91, 111, 114, 122, 125, 134, 157, 159, 172,
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
// static state used to implement the grammar is lazily initialized during the first call to
// NewGoEventerParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func GoEventerParserInit() {
	staticData := &GoEventerParserStaticData
	staticData.once.Do(goeventerParserInit)
}

// NewGoEventerParser produces a new grammar instance for the optional input antlr.TokenStream.
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
	GoEventerParserFOR          = 6
	GoEventerParserWHILE        = 7
	GoEventerParserRANGE        = 8
	GoEventerParserFUNC         = 9
	GoEventerParserON           = 10
	GoEventerParserRETURN       = 11
	GoEventerParserCHAN         = 12
	GoEventerParserNOTIFY       = 13
	GoEventerParserSTRING_TYPE  = 14
	GoEventerParserINTEGER_TYPE = 15
	GoEventerParserOR           = 16
	GoEventerParserAND          = 17
	GoEventerParserEQ           = 18
	GoEventerParserNEQ          = 19
	GoEventerParserGT           = 20
	GoEventerParserLT           = 21
	GoEventerParserGTEQ         = 22
	GoEventerParserLTEQ         = 23
	GoEventerParserPLUS         = 24
	GoEventerParserMINUS        = 25
	GoEventerParserMULT         = 26
	GoEventerParserDIV          = 27
	GoEventerParserMOD          = 28
	GoEventerParserPOW          = 29
	GoEventerParserNOT          = 30
	GoEventerParserTRUE         = 31
	GoEventerParserFALSE        = 32
	GoEventerParserNIL          = 33
	GoEventerParserPARAN_OPEN   = 34
	GoEventerParserPARAN_CLOSE  = 35
	GoEventerParserBARACE_OPEN  = 36
	GoEventerParserBARACE_CLOSE = 37
	GoEventerParserID           = 38
	GoEventerParserINT          = 39
	GoEventerParserFLOAT        = 40
	GoEventerParserSTRING       = 41
	GoEventerParserSPACE        = 42
)

// GoEventerParser rules.
const (
	GoEventerParserRULE_parse                   = 0
	GoEventerParserRULE_block                   = 1
	GoEventerParserRULE_stat                    = 2
	GoEventerParserRULE_assignVariable          = 3
	GoEventerParserRULE_notifyChanStat          = 4
	GoEventerParserRULE_forStat                 = 5
	GoEventerParserRULE_whileStat               = 6
	GoEventerParserRULE_defineListener          = 7
	GoEventerParserRULE_defineFunction          = 8
	GoEventerParserRULE_statBlock               = 9
	GoEventerParserRULE_returnStat              = 10
	GoEventerParserRULE_methodCallStat          = 11
	GoEventerParserRULE_methodCall              = 12
	GoEventerParserRULE_methodCallArguments     = 13
	GoEventerParserRULE_functionDefineArguments = 14
	GoEventerParserRULE_expr                    = 15
	GoEventerParserRULE_atom                    = 16
)

// IParseContext is an interface to support dynamic dispatch.
type IParseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the grammar.
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
		p.SetState(34)
		p.Block()
	}
	{
		p.SetState(35)
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

	// GetParser returns the grammar.
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
	p.SetState(40)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&274877910720) != 0 {
		{
			p.SetState(37)
			p.Stat()
		}

		p.SetState(42)
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

	// GetParser returns the grammar.
	GetParser() antlr.Parser

	// Getter signatures
	MethodCallStat() IMethodCallStatContext
	DefineListener() IDefineListenerContext
	AssignVariable() IAssignVariableContext
	ForStat() IForStatContext
	WhileStat() IWhileStatContext
	DefineFunction() IDefineFunctionContext
	ReturnStat() IReturnStatContext
	NotifyChanStat() INotifyChanStatContext

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

func (s *StatContext) DefineListener() IDefineListenerContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDefineListenerContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDefineListenerContext)
}

func (s *StatContext) AssignVariable() IAssignVariableContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignVariableContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignVariableContext)
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

func (s *StatContext) DefineFunction() IDefineFunctionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDefineFunctionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDefineFunctionContext)
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
	p.SetState(51)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(43)
			p.MethodCallStat()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(44)
			p.DefineListener()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(45)
			p.AssignVariable()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(46)
			p.ForStat()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(47)
			p.WhileStat()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(48)
			p.DefineFunction()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(49)
			p.ReturnStat()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(50)
			p.NotifyChanStat()
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

// IAssignVariableContext is an interface to support dynamic dispatch.
type IAssignVariableContext interface {
	antlr.ParserRuleContext

	// GetParser returns the grammar.
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

	// IsAssignVariableContext differentiates from other interfaces.
	IsAssignVariableContext()
}

type AssignVariableContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	op     antlr.Token
}

func NewEmptyAssignVariableContext() *AssignVariableContext {
	var p = new(AssignVariableContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoEventerParserRULE_assignVariable
	return p
}

func InitEmptyAssignVariableContext(p *AssignVariableContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoEventerParserRULE_assignVariable
}

func (*AssignVariableContext) IsAssignVariableContext() {}

func NewAssignVariableContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignVariableContext {
	var p = new(AssignVariableContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoEventerParserRULE_assignVariable

	return p
}

func (s *AssignVariableContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignVariableContext) GetOp() antlr.Token { return s.op }

func (s *AssignVariableContext) SetOp(v antlr.Token) { s.op = v }

func (s *AssignVariableContext) ID() antlr.TerminalNode {
	return s.GetToken(GoEventerParserID, 0)
}

func (s *AssignVariableContext) Expr() IExprContext {
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

func (s *AssignVariableContext) SCOL() antlr.TerminalNode {
	return s.GetToken(GoEventerParserSCOL, 0)
}

func (s *AssignVariableContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(GoEventerParserASSIGN, 0)
}

func (s *AssignVariableContext) UPDATE() antlr.TerminalNode {
	return s.GetToken(GoEventerParserUPDATE, 0)
}

func (s *AssignVariableContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignVariableContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignVariableContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoEventerListener); ok {
		listenerT.EnterAssignVariable(s)
	}
}

func (s *AssignVariableContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoEventerListener); ok {
		listenerT.ExitAssignVariable(s)
	}
}

func (s *AssignVariableContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoEventerVisitor:
		return t.VisitAssignVariable(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoEventerParser) AssignVariable() (localctx IAssignVariableContext) {
	localctx = NewAssignVariableContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, GoEventerParserRULE_assignVariable)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(53)
		p.Match(GoEventerParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(54)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*AssignVariableContext).op = _lt

		_la = p.GetTokenStream().LA(1)

		if !(_la == GoEventerParserASSIGN || _la == GoEventerParserUPDATE) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*AssignVariableContext).op = _ri
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(55)
		p.expr(0)
	}
	{
		p.SetState(56)
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

	// GetParser returns the grammar.
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
	p.EnterRule(localctx, 8, GoEventerParserRULE_notifyChanStat)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(58)
		p.Match(GoEventerParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(59)
		p.Match(GoEventerParserNOTIFY)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(60)
		p.expr(0)
	}
	{
		p.SetState(61)
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

	// GetParser returns the grammar.
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
	p.EnterRule(localctx, 10, GoEventerParserRULE_forStat)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(63)
		p.Match(GoEventerParserFOR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(64)
		p.Match(GoEventerParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(65)
		p.Match(GoEventerParserASSIGN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(66)
		p.expr(0)
	}
	{
		p.SetState(67)
		p.Match(GoEventerParserCOL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(68)
		p.expr(0)
	}
	{
		p.SetState(69)
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

	// GetParser returns the grammar.
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
	p.EnterRule(localctx, 12, GoEventerParserRULE_whileStat)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(71)
		p.Match(GoEventerParserWHILE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(72)
		p.expr(0)
	}
	{
		p.SetState(73)
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

// IDefineListenerContext is an interface to support dynamic dispatch.
type IDefineListenerContext interface {
	antlr.ParserRuleContext

	// GetParser returns the grammar.
	GetParser() antlr.Parser

	// Getter signatures
	ON() antlr.TerminalNode
	MethodCall() IMethodCallContext
	StatBlock() IStatBlockContext

	// IsDefineListenerContext differentiates from other interfaces.
	IsDefineListenerContext()
}

type DefineListenerContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDefineListenerContext() *DefineListenerContext {
	var p = new(DefineListenerContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoEventerParserRULE_defineListener
	return p
}

func InitEmptyDefineListenerContext(p *DefineListenerContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoEventerParserRULE_defineListener
}

func (*DefineListenerContext) IsDefineListenerContext() {}

func NewDefineListenerContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DefineListenerContext {
	var p = new(DefineListenerContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoEventerParserRULE_defineListener

	return p
}

func (s *DefineListenerContext) GetParser() antlr.Parser { return s.parser }

func (s *DefineListenerContext) ON() antlr.TerminalNode {
	return s.GetToken(GoEventerParserON, 0)
}

func (s *DefineListenerContext) MethodCall() IMethodCallContext {
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

func (s *DefineListenerContext) StatBlock() IStatBlockContext {
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

func (s *DefineListenerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DefineListenerContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DefineListenerContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoEventerListener); ok {
		listenerT.EnterDefineListener(s)
	}
}

func (s *DefineListenerContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoEventerListener); ok {
		listenerT.ExitDefineListener(s)
	}
}

func (s *DefineListenerContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoEventerVisitor:
		return t.VisitDefineListener(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoEventerParser) DefineListener() (localctx IDefineListenerContext) {
	localctx = NewDefineListenerContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, GoEventerParserRULE_defineListener)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(75)
		p.Match(GoEventerParserON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(76)
		p.MethodCall()
	}
	{
		p.SetState(77)
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

// IDefineFunctionContext is an interface to support dynamic dispatch.
type IDefineFunctionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the grammar.
	GetParser() antlr.Parser

	// Getter signatures
	FUNC() antlr.TerminalNode
	ID() antlr.TerminalNode
	PARAN_OPEN() antlr.TerminalNode
	FunctionDefineArguments() IFunctionDefineArgumentsContext
	PARAN_CLOSE() antlr.TerminalNode
	StatBlock() IStatBlockContext

	// IsDefineFunctionContext differentiates from other interfaces.
	IsDefineFunctionContext()
}

type DefineFunctionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDefineFunctionContext() *DefineFunctionContext {
	var p = new(DefineFunctionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoEventerParserRULE_defineFunction
	return p
}

func InitEmptyDefineFunctionContext(p *DefineFunctionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoEventerParserRULE_defineFunction
}

func (*DefineFunctionContext) IsDefineFunctionContext() {}

func NewDefineFunctionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DefineFunctionContext {
	var p = new(DefineFunctionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoEventerParserRULE_defineFunction

	return p
}

func (s *DefineFunctionContext) GetParser() antlr.Parser { return s.parser }

func (s *DefineFunctionContext) FUNC() antlr.TerminalNode {
	return s.GetToken(GoEventerParserFUNC, 0)
}

func (s *DefineFunctionContext) ID() antlr.TerminalNode {
	return s.GetToken(GoEventerParserID, 0)
}

func (s *DefineFunctionContext) PARAN_OPEN() antlr.TerminalNode {
	return s.GetToken(GoEventerParserPARAN_OPEN, 0)
}

func (s *DefineFunctionContext) FunctionDefineArguments() IFunctionDefineArgumentsContext {
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

func (s *DefineFunctionContext) PARAN_CLOSE() antlr.TerminalNode {
	return s.GetToken(GoEventerParserPARAN_CLOSE, 0)
}

func (s *DefineFunctionContext) StatBlock() IStatBlockContext {
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

func (s *DefineFunctionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DefineFunctionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DefineFunctionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoEventerListener); ok {
		listenerT.EnterDefineFunction(s)
	}
}

func (s *DefineFunctionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoEventerListener); ok {
		listenerT.ExitDefineFunction(s)
	}
}

func (s *DefineFunctionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoEventerVisitor:
		return t.VisitDefineFunction(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoEventerParser) DefineFunction() (localctx IDefineFunctionContext) {
	localctx = NewDefineFunctionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, GoEventerParserRULE_defineFunction)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(79)
		p.Match(GoEventerParserFUNC)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(80)
		p.Match(GoEventerParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(81)
		p.Match(GoEventerParserPARAN_OPEN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(82)
		p.FunctionDefineArguments()
	}
	{
		p.SetState(83)
		p.Match(GoEventerParserPARAN_CLOSE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(84)
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

	// GetParser returns the grammar.
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
	p.EnterRule(localctx, 18, GoEventerParserRULE_statBlock)
	p.SetState(91)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case GoEventerParserBARACE_OPEN:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(86)
			p.Match(GoEventerParserBARACE_OPEN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(87)
			p.Block()
		}
		{
			p.SetState(88)
			p.Match(GoEventerParserBARACE_CLOSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case GoEventerParserFOR, GoEventerParserWHILE, GoEventerParserFUNC, GoEventerParserON, GoEventerParserRETURN, GoEventerParserID:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(90)
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

	// GetParser returns the grammar.
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
	p.EnterRule(localctx, 20, GoEventerParserRULE_returnStat)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(93)
		p.Match(GoEventerParserRETURN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(94)
		p.expr(0)
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

// IMethodCallStatContext is an interface to support dynamic dispatch.
type IMethodCallStatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the grammar.
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
	p.EnterRule(localctx, 22, GoEventerParserRULE_methodCallStat)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(97)
		p.MethodCall()
	}
	{
		p.SetState(98)
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

	// GetParser returns the grammar.
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
	p.EnterRule(localctx, 24, GoEventerParserRULE_methodCall)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(100)
		p.Match(GoEventerParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(101)
		p.Match(GoEventerParserPARAN_OPEN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(102)
		p.MethodCallArguments()
	}
	{
		p.SetState(103)
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

	// GetParser returns the grammar.
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
	p.EnterRule(localctx, 26, GoEventerParserRULE_methodCallArguments)
	var _la int

	p.SetState(114)
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
			p.SetState(106)
			p.expr(0)
		}
		p.SetState(111)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == GoEventerParserT__0 {
			{
				p.SetState(107)
				p.Match(GoEventerParserT__0)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(108)
				p.expr(0)
			}

			p.SetState(113)
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

	// GetParser returns the grammar.
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
	p.EnterRule(localctx, 28, GoEventerParserRULE_functionDefineArguments)
	var _la int

	p.SetState(125)
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
			p.SetState(117)
			p.Match(GoEventerParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(122)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == GoEventerParserT__0 {
			{
				p.SetState(118)
				p.Match(GoEventerParserT__0)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(119)
				p.Match(GoEventerParserID)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(124)
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

	// GetParser returns the grammar.
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
	_startState := 30
	p.EnterRecursionRule(localctx, 30, GoEventerParserRULE_expr, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(134)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 7, p.GetParserRuleContext()) {
	case 1:
		localctx = NewMethodCallExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(128)
			p.MethodCall()
		}

	case 2:
		localctx = NewUnaryMinusExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(129)
			p.Match(GoEventerParserMINUS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(130)
			p.expr(9)
		}

	case 3:
		localctx = NewNotExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(131)
			p.Match(GoEventerParserNOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(132)
			p.expr(8)
		}

	case 4:
		localctx = NewAtomExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(133)
			p.Atom()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(159)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(157)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext()) {
			case 1:
				localctx = NewPowExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, GoEventerParserRULE_expr)
				p.SetState(136)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
					goto errorExit
				}
				{
					p.SetState(137)
					p.Match(GoEventerParserPOW)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(138)
					p.expr(10)
				}

			case 2:
				localctx = NewMultiplicationExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, GoEventerParserRULE_expr)
				p.SetState(139)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
					goto errorExit
				}
				{
					p.SetState(140)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*MultiplicationExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&469762048) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*MultiplicationExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(141)
					p.expr(8)
				}

			case 3:
				localctx = NewAdditiveExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, GoEventerParserRULE_expr)
				p.SetState(142)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
					goto errorExit
				}
				{
					p.SetState(143)

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
					p.SetState(144)
					p.expr(7)
				}

			case 4:
				localctx = NewRelationalExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, GoEventerParserRULE_expr)
				p.SetState(145)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
					goto errorExit
				}
				{
					p.SetState(146)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*RelationalExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&15728640) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*RelationalExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(147)
					p.expr(6)
				}

			case 5:
				localctx = NewEqualityExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, GoEventerParserRULE_expr)
				p.SetState(148)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
					goto errorExit
				}
				{
					p.SetState(149)

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
					p.SetState(150)
					p.expr(5)
				}

			case 6:
				localctx = NewAndExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, GoEventerParserRULE_expr)
				p.SetState(151)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
					goto errorExit
				}
				{
					p.SetState(152)
					p.Match(GoEventerParserAND)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(153)
					p.expr(4)
				}

			case 7:
				localctx = NewOrExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, GoEventerParserRULE_expr)
				p.SetState(154)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
					goto errorExit
				}
				{
					p.SetState(155)
					p.Match(GoEventerParserOR)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(156)
					p.expr(3)
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(161)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext())
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

	// GetParser returns the grammar.
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

func (s *IdAtomContext) ID() antlr.TerminalNode {
	return s.GetToken(GoEventerParserID, 0)
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
	p.EnterRule(localctx, 32, GoEventerParserRULE_atom)
	var _la int

	p.SetState(172)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case GoEventerParserPARAN_OPEN:
		localctx = NewExprAtomContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(162)
			p.Match(GoEventerParserPARAN_OPEN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(163)
			p.expr(0)
		}
		{
			p.SetState(164)
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
			p.SetState(166)
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
			p.SetState(167)
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
			p.SetState(168)
			p.Match(GoEventerParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case GoEventerParserSTRING:
		localctx = NewStringAtomContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(169)
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
			p.SetState(170)
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
			p.SetState(171)
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
	case 15:
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
