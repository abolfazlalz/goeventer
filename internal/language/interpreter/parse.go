package interpreter

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/abolfazlalz/goeventer/parsing"
	"github.com/antlr4-go/antlr/v4"
)

type Visitor struct {
	*parsing.BaseGoEventerVisitor
	mu        sync.Mutex
	variables map[int]map[string]*Variable
	functions map[string]func(...interface{}) interface{}
	state     int
}

func (v *Visitor) VisitReturnStat(ctx *parsing.ReturnStatContext) interface{} {
	value := v.Visit(ctx.Expr()).(*Variable)
	return NewReturnStatement(value)
}

func NewVisitor() *Visitor {
	return &Visitor{
		BaseGoEventerVisitor: &parsing.BaseGoEventerVisitor{},
		variables:            map[int]map[string]*Variable{0: make(map[string]*Variable)},
		state:                0,
		functions: map[string]func(...interface{}) interface{}{
			"print": func(i ...interface{}) interface{} {
				fmt.Println(i...)
				return nil
			},
		},
	}
}

func (v *Visitor) getVariable(name string) *Variable {
	for i := v.state; 0 <= i; i-- {
		if values, ok := v.variables[i]; ok {
			if value, ok := values[name]; ok {
				return value
			}
		}
	}
	panic("undefined variable")
}

func (v *Visitor) setVariable(name string, value *Variable) {
	for i := v.state; 0 <= i; i-- {
		if values, ok := v.variables[i]; ok {
			if _, ok := values[name]; ok {
				v.variables[i][name] = value
				return
			}
		}
	}
	panic("undefined variable")
}

func (v *Visitor) defineVariable(name string, value *Variable) {
	v.variables[v.state][name] = value
}

func (v *Visitor) VisitStatBlock(ctx *parsing.StatBlockContext) interface{} {
	if ctx.Stat() != nil {
		return v.Visit(ctx.Stat())
	} else if ctx.Block() != nil {
		return v.Visit(ctx.Block())
	}
	return nil
}

func (v *Visitor) VisitAssignVariable(ctx *parsing.AssignVariableContext) interface{} {
	name := ctx.ID().GetText()
	value := v.Visit(ctx.Expr()).(*Variable)

	if ctx.GetOp().GetTokenType() == parsing.GoEventerLexerASSIGN {
		v.variables[v.state][name] = value
	} else {
		v.setVariable(name, value)
	}
	return value
}

func (v *Visitor) VisitDefineFunction(ctx *parsing.DefineFunctionContext) interface{} {
	methodName := ctx.ID().GetText()

	v.functions[methodName] = func(args ...interface{}) interface{} {
		defer v.exitLastState()
		v.enterNewState()

		argNames := v.Visit(ctx.FunctionDefineArguments()).([]string)
		for i, name := range argNames {
			v.defineVariable(name, args[i].(*Variable))
		}

		r := v.Visit(ctx.StatBlock())
		return r
	}
	return nil
}

func (v *Visitor) VisitFunctionDefineArguments(ctx *parsing.FunctionDefineArgumentsContext) interface{} {
	ids := make([]string, len(ctx.AllID()))
	for i, id := range ctx.AllID() {
		ids[i] = id.GetText()
	}
	return ids
}

func (v *Visitor) VisitForStat(ctx *parsing.ForStatContext) interface{} {
	defer v.exitLastState()
	v.enterNewState()
	name := ctx.ID().GetText()
	start, okStart := v.Visit(ctx.Expr(0)).(*Variable)
	end, okEnd := v.Visit(ctx.Expr(1)).(*Variable)
	if !okStart || start.Type != IntegerVariable {
		panic("invalid start value")
	} else if !okEnd || end.Type != IntegerVariable {
		panic("invalid end value")
	}
	v.variables[v.state][name] = start
	for {
		v.Visit(ctx.StatBlock())
		val := v.variables[v.state][name].Sum(NewVariable(1)).(int)
		v.variables[v.state][name] = NewVariable(val)
		if val >= end.Integer() {
			break
		}
	}
	return nil
}

func (v *Visitor) enterNewState() {
	defer func() {
		v.mu.Unlock()
	}()
	v.mu.Lock()
	v.state++
	v.variables[v.state] = map[string]*Variable{}
}

func (v *Visitor) exitLastState() {
	defer func() {
		v.mu.Unlock()
	}()
	v.mu.Lock()
	delete(v.variables, v.state)
	v.state--
}

func (v *Visitor) VisitMethodCallStat(ctx *parsing.MethodCallStatContext) interface{} {
	return v.Visit(ctx.MethodCall())
}

func (v *Visitor) Visit(tree antlr.ParseTree) interface{} {
	return tree.Accept(v)
}

func (v *Visitor) VisitChildren(node antlr.RuleNode) interface{} {
	return v.BaseGoEventerVisitor.VisitChildren(node)
}

func (v *Visitor) VisitTerminal(node antlr.TerminalNode) interface{} {
	return v.BaseGoEventerVisitor.VisitTerminal(node)
}

func (v *Visitor) VisitErrorNode(node antlr.ErrorNode) interface{} {
	return v.BaseGoEventerVisitor.VisitErrorNode(node)
}

func (v *Visitor) VisitParse(ctx *parsing.ParseContext) interface{} {
	return v.Visit(ctx.Block())
}

func (v *Visitor) VisitBlock(ctx *parsing.BlockContext) interface{} {
	for _, stat := range ctx.AllStat() {
		if result, ok := v.Visit(stat).(*ReturnStatement); ok {
			return result
		}
	}
	return nil
}

func (v *Visitor) VisitStat(ctx *parsing.StatContext) interface{} {
	for _, child := range ctx.GetChildren() {
		return v.Visit(child.(antlr.ParseTree))
	}
	return nil
}

func (v *Visitor) VisitDefineListener(ctx *parsing.DefineListenerContext) interface{} {
	ticker := time.NewTicker(1 * time.Second)
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				result := v.Visit(ctx.MethodCall()).(*ReturnStatement)
				if result != nil && result.variable.Type == BoolVariable && result.variable.Boolean() {
					v.Visit(ctx.StatBlock())
				}
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()
	return nil
}

func (v *Visitor) VisitMethodCall(ctx *parsing.MethodCallContext) interface{} {
	name := ctx.ID().GetText()
	argsAny := v.Visit(ctx.MethodCallArguments())
	args, ok := argsAny.([]interface{})
	if !ok {
		panic(fmt.Sprintf("undefined parameter for %s method.", name))
	}
	if callback, ok := v.functions[name]; ok {
		if re, ok := callback(args...).(*ReturnStatement); ok {
			return re
		}
	}
	return nil
}

func (v *Visitor) VisitMethodCallArguments(ctx *parsing.MethodCallArgumentsContext) interface{} {
	arr := make([]interface{}, len(ctx.AllExpr()))
	for i, expr := range ctx.AllExpr() {
		arr[i] = v.Visit(expr)
	}
	return arr
}

func (v *Visitor) VisitMethodCallExpr(ctx *parsing.MethodCallExprContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v *Visitor) VisitNotExpr(ctx *parsing.NotExprContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v *Visitor) VisitUnaryMinusExpr(ctx *parsing.UnaryMinusExprContext) interface{} {
	value := v.Visit(ctx.Expr()).(*Variable)
	return NewVariable(value.Multiple(NewVariable(-1)))
}

func (v *Visitor) VisitMultiplicationExpr(ctx *parsing.MultiplicationExprContext) interface{} {
	left := v.Visit(ctx.Expr(0)).(*Variable)
	right := v.Visit(ctx.Expr(1)).(*Variable)
	switch ctx.GetOp().GetTokenType() {
	case parsing.GoEventerLexerMULT:
		return left.Multiple(right)
	case parsing.GoEventerLexerDIV:
		return left.Divide(right)
	case parsing.GoEventerLexerMOD:
		return left.Mod(right)
	default:
		panic("Undefined operator")
	}
}

func (v *Visitor) VisitAtomExpr(ctx *parsing.AtomExprContext) interface{} {
	return v.Visit(ctx.Atom())
}

func (v *Visitor) VisitOrExpr(ctx *parsing.OrExprContext) interface{} {
	v1 := v.Visit(ctx.Expr(0)).(*Variable)
	v2 := v.Visit(ctx.Expr(1)).(*Variable)
	return v1.Boolean() || v2.Boolean()
}

func (v *Visitor) VisitAdditiveExpr(ctx *parsing.AdditiveExprContext) interface{} {
	v1 := v.Visit(ctx.Expr(0)).(*Variable)
	v2 := v.Visit(ctx.Expr(1)).(*Variable)
	return NewVariable(v1.Sum(v2))
}

func (v *Visitor) VisitPowExpr(ctx *parsing.PowExprContext) interface{} {
	v1 := v.Visit(ctx.Expr(0)).(*Variable)
	v2 := v.Visit(ctx.Expr(1)).(*Variable)
	return NewVariable(v1.Pow(v2))
}

func (v *Visitor) VisitRelationalExpr(ctx *parsing.RelationalExprContext) interface{} {
	v1 := v.Visit(ctx.Expr(0)).(*Variable)
	v2 := v.Visit(ctx.Expr(1)).(*Variable)
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

	return NewVariable(val)
}

func (v *Visitor) VisitEqualityExpr(ctx *parsing.EqualityExprContext) interface{} {
	v1 := v.Visit(ctx.Expr(0)).(*Variable)
	v2 := v.Visit(ctx.Expr(1)).(*Variable)
	return NewVariable(v1.Equal(v2))
}

func (v *Visitor) VisitAndExpr(ctx *parsing.AndExprContext) interface{} {
	return v.Visit(ctx.Expr(0)).(bool) && v.Visit(ctx.Expr(1)).(bool)
}

func (v *Visitor) VisitExprAtom(ctx *parsing.ExprAtomContext) interface{} {
	return v.Visit(ctx.Expr())
}

func (v *Visitor) VisitNumberAtom(ctx *parsing.NumberAtomContext) interface{} {
	num, err := strconv.Atoi(ctx.GetText())
	if err != nil {
		panic(err.Error())
	}
	return NewVariable(num)
}

func (v *Visitor) VisitBooleanAtom(ctx *parsing.BooleanAtomContext) interface{} {
	return NewVariable(ctx.TRUE() != nil)
}

func (v *Visitor) VisitIdAtom(ctx *parsing.IdAtomContext) interface{} {
	return v.getVariable(ctx.ID().GetText())
}

func (v *Visitor) VisitStringAtom(ctx *parsing.StringAtomContext) interface{} {
	str := ctx.GetText()
	str = str[1 : len(str)-1]
	str = strings.ReplaceAll(str, "\"\"", "\"")
	return NewVariable(str)
}

func (v *Visitor) VisitNilAtom(ctx *parsing.NilAtomContext) interface{} {
	return nil
}
