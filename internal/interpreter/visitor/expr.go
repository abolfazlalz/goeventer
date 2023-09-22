package visitor

import (
	"github.com/abolfazlalz/goeventer/internal/interpreter/grammar"
	"github.com/abolfazlalz/goeventer/internal/interpreter/miscs"
)

func (v *Visitor) VisitNotExpr(ctx *grammar.NotExprContext) interface{} {
	result := v.Visit(ctx.Expr())
	val := v.variableFromContext(result)
	return miscs.NewVariable(!val.Boolean())
}

func (v *Visitor) VisitUnaryMinusExpr(ctx *grammar.UnaryMinusExprContext) interface{} {
	value := v.variableFromContext(v.Visit(ctx.Expr()))
	return miscs.NewVariable(value.Multiple(miscs.NewVariable(-1)))
}

func (v *Visitor) VisitMultiplicationExpr(ctx *grammar.MultiplicationExprContext) interface{} {
	left := v.variableFromContext(v.Visit(ctx.Expr(0)))
	right := v.variableFromContext(v.Visit(ctx.Expr(1)))
	switch ctx.GetOp().GetTokenType() {
	case grammar.GoEventerLexerMULT:
		return left.Multiple(right)
	case grammar.GoEventerLexerDIV:
		return left.Divide(right)
	case grammar.GoEventerLexerMOD:
		return left.Mod(right)
	default:
		panic("Undefined operator")
	}
}

func (v *Visitor) VisitOrExpr(ctx *grammar.OrExprContext) interface{} {
	v1 := v.variableFromContext(v.Visit(ctx.Expr(0)))
	v2 := v.variableFromContext(v.Visit(ctx.Expr(1)))
	return miscs.NewVariable(v1.Boolean() || v2.Boolean())
}

func (v *Visitor) VisitAdditiveExpr(ctx *grammar.AdditiveExprContext) interface{} {
	v1 := v.variableFromContext(v.Visit(ctx.Expr(0)))
	v2 := v.variableFromContext(v.Visit(ctx.Expr(1)))
	return miscs.NewVariable(v1.Sum(v2))
}

func (v *Visitor) VisitPowExpr(ctx *grammar.PowExprContext) interface{} {
	v1 := v.variableFromContext(v.Visit(ctx.Expr(0)))
	v2 := v.variableFromContext(v.Visit(ctx.Expr(1)))
	return miscs.NewVariable(v1.Pow(v2))
}

func (v *Visitor) VisitRelationalExpr(ctx *grammar.RelationalExprContext) interface{} {
	v1 := v.variableFromContext(v.Visit(ctx.Expr(0)))
	v2 := v.variableFromContext(v.Visit(ctx.Expr(1)))
	val := false
	if ctx.LT() != nil {
		val = v1.Integer() < v2.Integer()
	} else if ctx.GT() != nil {
		val = v1.Integer() > v2.Integer()
	} else if ctx.GTEQ() != nil {
		val = v1.Integer() >= v2.Integer()
	} else if ctx.LTEQ() != nil {
		val = v1.Integer() <= v2.Integer()
	}

	return miscs.NewVariable(val)
}

func (v *Visitor) VisitEqualityExpr(ctx *grammar.EqualityExprContext) interface{} {
	v1 := v.variableFromContext(v.Visit(ctx.Expr(0)))
	v2 := v.variableFromContext(v.Visit(ctx.Expr(1)))
	return miscs.NewVariable(v1.Equal(v2))
}

func (v *Visitor) VisitAndExpr(ctx *grammar.AndExprContext) interface{} {
	return miscs.NewVariable(v.Visit(ctx.Expr(0)).(bool) && v.Visit(ctx.Expr(1)).(bool))
}

func (v *Visitor) VisitAtomExpr(ctx *grammar.AtomExprContext) interface{} {
	return v.Visit(ctx.Atom())
}
