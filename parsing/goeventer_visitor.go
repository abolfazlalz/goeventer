// Code generated from GoEventer.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parsing // GoEventer
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by GoEventerParser.
type GoEventerVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by GoEventerParser#parse.
	VisitParse(ctx *ParseContext) interface{}

	// Visit a parse tree produced by GoEventerParser#block.
	VisitBlock(ctx *BlockContext) interface{}

	// Visit a parse tree produced by GoEventerParser#stat.
	VisitStat(ctx *StatContext) interface{}

	// Visit a parse tree produced by GoEventerParser#assignVariable.
	VisitAssignVariable(ctx *AssignVariableContext) interface{}

	// Visit a parse tree produced by GoEventerParser#forStat.
	VisitForStat(ctx *ForStatContext) interface{}

	// Visit a parse tree produced by GoEventerParser#whileStat.
	VisitWhileStat(ctx *WhileStatContext) interface{}

	// Visit a parse tree produced by GoEventerParser#defineListener.
	VisitDefineListener(ctx *DefineListenerContext) interface{}

	// Visit a parse tree produced by GoEventerParser#defineFunction.
	VisitDefineFunction(ctx *DefineFunctionContext) interface{}

	// Visit a parse tree produced by GoEventerParser#statBlock.
	VisitStatBlock(ctx *StatBlockContext) interface{}

	// Visit a parse tree produced by GoEventerParser#returnStat.
	VisitReturnStat(ctx *ReturnStatContext) interface{}

	// Visit a parse tree produced by GoEventerParser#methodCallStat.
	VisitMethodCallStat(ctx *MethodCallStatContext) interface{}

	// Visit a parse tree produced by GoEventerParser#methodCall.
	VisitMethodCall(ctx *MethodCallContext) interface{}

	// Visit a parse tree produced by GoEventerParser#methodCallArguments.
	VisitMethodCallArguments(ctx *MethodCallArgumentsContext) interface{}

	// Visit a parse tree produced by GoEventerParser#functionDefineArguments.
	VisitFunctionDefineArguments(ctx *FunctionDefineArgumentsContext) interface{}

	// Visit a parse tree produced by GoEventerParser#methodCallExpr.
	VisitMethodCallExpr(ctx *MethodCallExprContext) interface{}

	// Visit a parse tree produced by GoEventerParser#notExpr.
	VisitNotExpr(ctx *NotExprContext) interface{}

	// Visit a parse tree produced by GoEventerParser#unaryMinusExpr.
	VisitUnaryMinusExpr(ctx *UnaryMinusExprContext) interface{}

	// Visit a parse tree produced by GoEventerParser#multiplicationExpr.
	VisitMultiplicationExpr(ctx *MultiplicationExprContext) interface{}

	// Visit a parse tree produced by GoEventerParser#atomExpr.
	VisitAtomExpr(ctx *AtomExprContext) interface{}

	// Visit a parse tree produced by GoEventerParser#orExpr.
	VisitOrExpr(ctx *OrExprContext) interface{}

	// Visit a parse tree produced by GoEventerParser#additiveExpr.
	VisitAdditiveExpr(ctx *AdditiveExprContext) interface{}

	// Visit a parse tree produced by GoEventerParser#powExpr.
	VisitPowExpr(ctx *PowExprContext) interface{}

	// Visit a parse tree produced by GoEventerParser#relationalExpr.
	VisitRelationalExpr(ctx *RelationalExprContext) interface{}

	// Visit a parse tree produced by GoEventerParser#equalityExpr.
	VisitEqualityExpr(ctx *EqualityExprContext) interface{}

	// Visit a parse tree produced by GoEventerParser#andExpr.
	VisitAndExpr(ctx *AndExprContext) interface{}

	// Visit a parse tree produced by GoEventerParser#exprAtom.
	VisitExprAtom(ctx *ExprAtomContext) interface{}

	// Visit a parse tree produced by GoEventerParser#numberAtom.
	VisitNumberAtom(ctx *NumberAtomContext) interface{}

	// Visit a parse tree produced by GoEventerParser#booleanAtom.
	VisitBooleanAtom(ctx *BooleanAtomContext) interface{}

	// Visit a parse tree produced by GoEventerParser#idAtom.
	VisitIdAtom(ctx *IdAtomContext) interface{}

	// Visit a parse tree produced by GoEventerParser#stringAtom.
	VisitStringAtom(ctx *StringAtomContext) interface{}

	// Visit a parse tree produced by GoEventerParser#nilAtom.
	VisitNilAtom(ctx *NilAtomContext) interface{}
}
