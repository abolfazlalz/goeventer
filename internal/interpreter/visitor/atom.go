package visitor

import (
	"github.com/abolfazlalz/goeventer/internal/interpreter/grammar"
	"github.com/abolfazlalz/goeventer/internal/interpreter/miscs"
	"strconv"
	"strings"
)

func (v *Visitor) VisitBooleanAtom(ctx *grammar.BooleanAtomContext) interface{} {
	return miscs.NewVariable(ctx.TRUE() != nil)
}

func (v *Visitor) VisitIdAtom(ctx *grammar.IdAtomContext) interface{} {
	return v.getVariable(ctx.ID().GetText())
}

func (v *Visitor) VisitStringAtom(ctx *grammar.StringAtomContext) interface{} {
	str := ctx.GetText()
	str = str[1 : len(str)-1]
	str = strings.ReplaceAll(str, "\"\"", "\"")
	return miscs.NewVariable(str)
}

func (v *Visitor) VisitNilAtom(_ *grammar.NilAtomContext) interface{} {
	return nil
}

func (v *Visitor) VisitMakeChanAtom(_ *grammar.MakeChanAtomContext) interface{} {
	return miscs.NewChanVariable()
}

func (v *Visitor) VisitNumberAtom(ctx *grammar.NumberAtomContext) interface{} {
	num, err := strconv.Atoi(ctx.GetText())
	if err != nil {
		panic(err.Error())
	}
	return miscs.NewVariable(num)
}

func (v *Visitor) VisitExprAtom(ctx *grammar.ExprAtomContext) interface{} {
	return v.Visit(ctx.Expr())
}
