package visitor

import (
	"github.com/abolfazlalz/goeventer/internal/interpreter/grammar"
	"github.com/abolfazlalz/goeventer/internal/interpreter/miscs"
)

func (v *Visitor) VisitWhileStat(ctx *grammar.WhileStatContext) interface{} {
	defer v.exitLastState()
	v.enterNewState()

	for {
		result := v.Visit(ctx.Expr())
		val := v.variableFromContext(result)
		if val.Type != miscs.BoolVariable || !val.Boolean() {
			break
		}
		v.Visit(ctx.StatBlock())
	}
	return nil
}

func (v *Visitor) VisitForStat(ctx *grammar.ForStatContext) interface{} {
	defer v.exitLastState()
	v.enterNewState()
	name := ctx.ID().GetText()
	start, okStart := v.Visit(ctx.Expr(0)).(*miscs.Variable)
	end, okEnd := v.Visit(ctx.Expr(1)).(*miscs.Variable)
	if !okStart || start.Type != miscs.IntegerVariable {
		panic("invalid start value")
	} else if !okEnd || end.Type != miscs.IntegerVariable {
		panic("invalid end value")
	}
	v.variables[v.state][name] = start
	for {
		v.Visit(ctx.StatBlock())
		val := v.variables[v.state][name].Sum(miscs.NewVariable(1)).(int)
		v.variables[v.state][name] = miscs.NewVariable(val)
		if val >= end.Integer() {
			break
		}
	}
	return nil
}
