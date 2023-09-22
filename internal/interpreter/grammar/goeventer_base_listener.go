// Code generated from GoEventer.g4 by ANTLR 4.13.0. DO NOT EDIT.

package grammar // GoEventer
import "github.com/antlr4-go/antlr/v4"

// BaseGoEventerListener is a complete listener for a parse tree produced by GoEventerParser.
type BaseGoEventerListener struct{}

var _ GoEventerListener = &BaseGoEventerListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseGoEventerListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseGoEventerListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseGoEventerListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseGoEventerListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterParse is called when production parse is entered.
func (s *BaseGoEventerListener) EnterParse(ctx *ParseContext) {}

// ExitParse is called when production parse is exited.
func (s *BaseGoEventerListener) ExitParse(ctx *ParseContext) {}

// EnterBlock is called when production block is entered.
func (s *BaseGoEventerListener) EnterBlock(ctx *BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *BaseGoEventerListener) ExitBlock(ctx *BlockContext) {}

// EnterStat is called when production stat is entered.
func (s *BaseGoEventerListener) EnterStat(ctx *StatContext) {}

// ExitStat is called when production stat is exited.
func (s *BaseGoEventerListener) ExitStat(ctx *StatContext) {}

// EnterAssignVariable is called when production assignVariable is entered.
func (s *BaseGoEventerListener) EnterAssignVariable(ctx *AssignVariableContext) {}

// ExitAssignVariable is called when production assignVariable is exited.
func (s *BaseGoEventerListener) ExitAssignVariable(ctx *AssignVariableContext) {}

// EnterNotifyChanStat is called when production notifyChanStat is entered.
func (s *BaseGoEventerListener) EnterNotifyChanStat(ctx *NotifyChanStatContext) {}

// ExitNotifyChanStat is called when production notifyChanStat is exited.
func (s *BaseGoEventerListener) ExitNotifyChanStat(ctx *NotifyChanStatContext) {}

// EnterForStat is called when production forStat is entered.
func (s *BaseGoEventerListener) EnterForStat(ctx *ForStatContext) {}

// ExitForStat is called when production forStat is exited.
func (s *BaseGoEventerListener) ExitForStat(ctx *ForStatContext) {}

// EnterWhileStat is called when production whileStat is entered.
func (s *BaseGoEventerListener) EnterWhileStat(ctx *WhileStatContext) {}

// ExitWhileStat is called when production whileStat is exited.
func (s *BaseGoEventerListener) ExitWhileStat(ctx *WhileStatContext) {}

// EnterDefineListener is called when production defineListener is entered.
func (s *BaseGoEventerListener) EnterDefineListener(ctx *DefineListenerContext) {}

// ExitDefineListener is called when production defineListener is exited.
func (s *BaseGoEventerListener) ExitDefineListener(ctx *DefineListenerContext) {}

// EnterDefineFunction is called when production defineFunction is entered.
func (s *BaseGoEventerListener) EnterDefineFunction(ctx *DefineFunctionContext) {}

// ExitDefineFunction is called when production defineFunction is exited.
func (s *BaseGoEventerListener) ExitDefineFunction(ctx *DefineFunctionContext) {}

// EnterStatBlock is called when production statBlock is entered.
func (s *BaseGoEventerListener) EnterStatBlock(ctx *StatBlockContext) {}

// ExitStatBlock is called when production statBlock is exited.
func (s *BaseGoEventerListener) ExitStatBlock(ctx *StatBlockContext) {}

// EnterReturnStat is called when production returnStat is entered.
func (s *BaseGoEventerListener) EnterReturnStat(ctx *ReturnStatContext) {}

// ExitReturnStat is called when production returnStat is exited.
func (s *BaseGoEventerListener) ExitReturnStat(ctx *ReturnStatContext) {}

// EnterMethodCallStat is called when production methodCallStat is entered.
func (s *BaseGoEventerListener) EnterMethodCallStat(ctx *MethodCallStatContext) {}

// ExitMethodCallStat is called when production methodCallStat is exited.
func (s *BaseGoEventerListener) ExitMethodCallStat(ctx *MethodCallStatContext) {}

// EnterMethodCall is called when production methodCall is entered.
func (s *BaseGoEventerListener) EnterMethodCall(ctx *MethodCallContext) {}

// ExitMethodCall is called when production methodCall is exited.
func (s *BaseGoEventerListener) ExitMethodCall(ctx *MethodCallContext) {}

// EnterMethodCallArguments is called when production methodCallArguments is entered.
func (s *BaseGoEventerListener) EnterMethodCallArguments(ctx *MethodCallArgumentsContext) {}

// ExitMethodCallArguments is called when production methodCallArguments is exited.
func (s *BaseGoEventerListener) ExitMethodCallArguments(ctx *MethodCallArgumentsContext) {}

// EnterFunctionDefineArguments is called when production functionDefineArguments is entered.
func (s *BaseGoEventerListener) EnterFunctionDefineArguments(ctx *FunctionDefineArgumentsContext) {}

// ExitFunctionDefineArguments is called when production functionDefineArguments is exited.
func (s *BaseGoEventerListener) ExitFunctionDefineArguments(ctx *FunctionDefineArgumentsContext) {}

// EnterMethodCallExpr is called when production methodCallExpr is entered.
func (s *BaseGoEventerListener) EnterMethodCallExpr(ctx *MethodCallExprContext) {}

// ExitMethodCallExpr is called when production methodCallExpr is exited.
func (s *BaseGoEventerListener) ExitMethodCallExpr(ctx *MethodCallExprContext) {}

// EnterNotExpr is called when production notExpr is entered.
func (s *BaseGoEventerListener) EnterNotExpr(ctx *NotExprContext) {}

// ExitNotExpr is called when production notExpr is exited.
func (s *BaseGoEventerListener) ExitNotExpr(ctx *NotExprContext) {}

// EnterUnaryMinusExpr is called when production unaryMinusExpr is entered.
func (s *BaseGoEventerListener) EnterUnaryMinusExpr(ctx *UnaryMinusExprContext) {}

// ExitUnaryMinusExpr is called when production unaryMinusExpr is exited.
func (s *BaseGoEventerListener) ExitUnaryMinusExpr(ctx *UnaryMinusExprContext) {}

// EnterMultiplicationExpr is called when production multiplicationExpr is entered.
func (s *BaseGoEventerListener) EnterMultiplicationExpr(ctx *MultiplicationExprContext) {}

// ExitMultiplicationExpr is called when production multiplicationExpr is exited.
func (s *BaseGoEventerListener) ExitMultiplicationExpr(ctx *MultiplicationExprContext) {}

// EnterAtomExpr is called when production atomExpr is entered.
func (s *BaseGoEventerListener) EnterAtomExpr(ctx *AtomExprContext) {}

// ExitAtomExpr is called when production atomExpr is exited.
func (s *BaseGoEventerListener) ExitAtomExpr(ctx *AtomExprContext) {}

// EnterOrExpr is called when production orExpr is entered.
func (s *BaseGoEventerListener) EnterOrExpr(ctx *OrExprContext) {}

// ExitOrExpr is called when production orExpr is exited.
func (s *BaseGoEventerListener) ExitOrExpr(ctx *OrExprContext) {}

// EnterAdditiveExpr is called when production additiveExpr is entered.
func (s *BaseGoEventerListener) EnterAdditiveExpr(ctx *AdditiveExprContext) {}

// ExitAdditiveExpr is called when production additiveExpr is exited.
func (s *BaseGoEventerListener) ExitAdditiveExpr(ctx *AdditiveExprContext) {}

// EnterPowExpr is called when production powExpr is entered.
func (s *BaseGoEventerListener) EnterPowExpr(ctx *PowExprContext) {}

// ExitPowExpr is called when production powExpr is exited.
func (s *BaseGoEventerListener) ExitPowExpr(ctx *PowExprContext) {}

// EnterRelationalExpr is called when production relationalExpr is entered.
func (s *BaseGoEventerListener) EnterRelationalExpr(ctx *RelationalExprContext) {}

// ExitRelationalExpr is called when production relationalExpr is exited.
func (s *BaseGoEventerListener) ExitRelationalExpr(ctx *RelationalExprContext) {}

// EnterEqualityExpr is called when production equalityExpr is entered.
func (s *BaseGoEventerListener) EnterEqualityExpr(ctx *EqualityExprContext) {}

// ExitEqualityExpr is called when production equalityExpr is exited.
func (s *BaseGoEventerListener) ExitEqualityExpr(ctx *EqualityExprContext) {}

// EnterAndExpr is called when production andExpr is entered.
func (s *BaseGoEventerListener) EnterAndExpr(ctx *AndExprContext) {}

// ExitAndExpr is called when production andExpr is exited.
func (s *BaseGoEventerListener) ExitAndExpr(ctx *AndExprContext) {}

// EnterExprAtom is called when production exprAtom is entered.
func (s *BaseGoEventerListener) EnterExprAtom(ctx *ExprAtomContext) {}

// ExitExprAtom is called when production exprAtom is exited.
func (s *BaseGoEventerListener) ExitExprAtom(ctx *ExprAtomContext) {}

// EnterNumberAtom is called when production numberAtom is entered.
func (s *BaseGoEventerListener) EnterNumberAtom(ctx *NumberAtomContext) {}

// ExitNumberAtom is called when production numberAtom is exited.
func (s *BaseGoEventerListener) ExitNumberAtom(ctx *NumberAtomContext) {}

// EnterBooleanAtom is called when production booleanAtom is entered.
func (s *BaseGoEventerListener) EnterBooleanAtom(ctx *BooleanAtomContext) {}

// ExitBooleanAtom is called when production booleanAtom is exited.
func (s *BaseGoEventerListener) ExitBooleanAtom(ctx *BooleanAtomContext) {}

// EnterIdAtom is called when production idAtom is entered.
func (s *BaseGoEventerListener) EnterIdAtom(ctx *IdAtomContext) {}

// ExitIdAtom is called when production idAtom is exited.
func (s *BaseGoEventerListener) ExitIdAtom(ctx *IdAtomContext) {}

// EnterStringAtom is called when production stringAtom is entered.
func (s *BaseGoEventerListener) EnterStringAtom(ctx *StringAtomContext) {}

// ExitStringAtom is called when production stringAtom is exited.
func (s *BaseGoEventerListener) ExitStringAtom(ctx *StringAtomContext) {}

// EnterNilAtom is called when production nilAtom is entered.
func (s *BaseGoEventerListener) EnterNilAtom(ctx *NilAtomContext) {}

// ExitNilAtom is called when production nilAtom is exited.
func (s *BaseGoEventerListener) ExitNilAtom(ctx *NilAtomContext) {}

// EnterMakeChanAtom is called when production makeChanAtom is entered.
func (s *BaseGoEventerListener) EnterMakeChanAtom(ctx *MakeChanAtomContext) {}

// ExitMakeChanAtom is called when production makeChanAtom is exited.
func (s *BaseGoEventerListener) ExitMakeChanAtom(ctx *MakeChanAtomContext) {}
