package visitor

import (
	"fmt"
	"github.com/abolfazlalz/goeventer/internal/interpreter/grammar"
	"github.com/abolfazlalz/goeventer/internal/interpreter/miscs"
)

func (v *Visitor) VisitMethodCall(ctx *grammar.MethodCallContext) interface{} {
	name := ctx.ID().GetText()
	argsAny := v.Visit(ctx.MethodCallArguments())
	args, ok := argsAny.([]interface{})
	if !ok {
		panic(fmt.Sprintf("undefined parameter for %s method.", name))
	}
	if callback, ok := v.functions[name]; ok {
		if re, ok := callback(args...).(*miscs.ReturnStatement); ok {
			return re
		}
		return nil
	}
	panic(fmt.Sprintf("undefined %s method", name))
}

func (v *Visitor) VisitMethodCallArguments(ctx *grammar.MethodCallArgumentsContext) interface{} {
	arr := make([]interface{}, len(ctx.AllExpr()))
	for i, expr := range ctx.AllExpr() {
		arr[i] = v.Visit(expr)
	}
	return arr
}

func (v *Visitor) VisitMethodCallExpr(ctx *grammar.MethodCallExprContext) interface{} {
	return v.Visit(ctx.MethodCall())
}
