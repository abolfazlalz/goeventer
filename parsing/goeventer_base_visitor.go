// Code generated from GoEventer.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parsing // GoEventer
import "github.com/antlr4-go/antlr/v4"

type BaseGoEventerVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseGoEventerVisitor) VisitParse(ctx *ParseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoEventerVisitor) VisitBlock(ctx *BlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoEventerVisitor) VisitStat(ctx *StatContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoEventerVisitor) VisitAssignVariable(ctx *AssignVariableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoEventerVisitor) VisitForStat(ctx *ForStatContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoEventerVisitor) VisitDefineListener(ctx *DefineListenerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoEventerVisitor) VisitDefineFunction(ctx *DefineFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoEventerVisitor) VisitStatBlock(ctx *StatBlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoEventerVisitor) VisitReturnStat(ctx *ReturnStatContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoEventerVisitor) VisitMethodCallStat(ctx *MethodCallStatContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoEventerVisitor) VisitMethodCall(ctx *MethodCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoEventerVisitor) VisitMethodCallArguments(ctx *MethodCallArgumentsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoEventerVisitor) VisitFunctionDefineArguments(ctx *FunctionDefineArgumentsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoEventerVisitor) VisitMethodCallExpr(ctx *MethodCallExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoEventerVisitor) VisitNotExpr(ctx *NotExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoEventerVisitor) VisitUnaryMinusExpr(ctx *UnaryMinusExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoEventerVisitor) VisitMultiplicationExpr(ctx *MultiplicationExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoEventerVisitor) VisitAtomExpr(ctx *AtomExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoEventerVisitor) VisitOrExpr(ctx *OrExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoEventerVisitor) VisitAdditiveExpr(ctx *AdditiveExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoEventerVisitor) VisitPowExpr(ctx *PowExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoEventerVisitor) VisitRelationalExpr(ctx *RelationalExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoEventerVisitor) VisitEqualityExpr(ctx *EqualityExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoEventerVisitor) VisitAndExpr(ctx *AndExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoEventerVisitor) VisitExprAtom(ctx *ExprAtomContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoEventerVisitor) VisitNumberAtom(ctx *NumberAtomContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoEventerVisitor) VisitBooleanAtom(ctx *BooleanAtomContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoEventerVisitor) VisitIdAtom(ctx *IdAtomContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoEventerVisitor) VisitStringAtom(ctx *StringAtomContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoEventerVisitor) VisitNilAtom(ctx *NilAtomContext) interface{} {
	return v.VisitChildren(ctx)
}
