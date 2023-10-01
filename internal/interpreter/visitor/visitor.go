package visitor

import (
	"github.com/abolfazlalz/goeventer/internal/interpreter/grammar"
	"github.com/abolfazlalz/goeventer/internal/interpreter/miscs"
	"github.com/abolfazlalz/goeventer/internal/websocket"
	"github.com/antlr4-go/antlr/v4"
	"github.com/samber/lo"
	"sync"
)

type Visitor struct {
	*grammar.BaseGoEventerVisitor
	mu        sync.Mutex
	variables map[int]map[string]*miscs.Variable
	functions map[string]func(...interface{}) interface{}
	state     int
	ch        chan *websocket.Chat
}

func (v *Visitor) VisitStructStat(ctx *grammar.StructStatContext) interface{} {
	name := ctx.ID().GetText()
	values := make(miscs.StructType)
	for _, structValue := range ctx.AllStructField() {
		fieldName := structValue.ID().GetText()
		values[fieldName] = v.Visit(structValue.Expr()).(*miscs.Variable)
	}
	v.defineVariable(name, miscs.NewVariable(values))
	return nil
}

func NewVisitor() *Visitor {
	ch := make(chan *websocket.Chat, 1)

	v := &Visitor{
		BaseGoEventerVisitor: &grammar.BaseGoEventerVisitor{},
		variables:            map[int]map[string]*miscs.Variable{0: make(map[string]*miscs.Variable)},
		state:                0,
		functions:            map[string]func(...interface{}) interface{}{},
		ch:                   ch,
	}

	v.loadBuiltinMethods()
	return v
}

func (v *Visitor) VisitNotifyChanStat(ctx *grammar.NotifyChanStatContext) interface{} {
	defer func() {
		v.mu.Unlock()
	}()
	v.mu.Lock()
	val := v.Visit(ctx.Expr()).(*miscs.Variable)
	name := ctx.ID().GetText()
	variable := v.getVariable(name)
	variable.NotifyChan(*val)
	return nil
}

func (v *Visitor) VisitReturnStat(ctx *grammar.ReturnStatContext) interface{} {
	value := v.Visit(ctx.Expr()).(*miscs.Variable)
	return miscs.NewReturnStatement(value)
}

func (v *Visitor) Visit(tree antlr.ParseTree) interface{} {
	return tree.Accept(v)
}

func (v *Visitor) VisitStatBlock(ctx *grammar.StatBlockContext) interface{} {
	if ctx.Stat() != nil {
		return v.Visit(ctx.Stat())
	} else if ctx.Block() != nil {
		return v.Visit(ctx.Block())
	}
	return nil
}

func (v *Visitor) VisitAssignVariableStat(ctx *grammar.AssignVariableStatContext) interface{} {
	name := ctx.ID().GetText()
	value := v.variableFromContext(v.Visit(ctx.Expr()))

	if ctx.GetOp().GetTokenType() == grammar.GoEventerLexerASSIGN {
		v.variables[v.state][name] = value
	} else {
		v.setVariable(name, value)
	}
	return value
}

func (v *Visitor) VisitUpdateVariableStat(ctx *grammar.UpdateVariableStatContext) interface{} {
	ids := lo.Map(ctx.AllID(), func(item antlr.TerminalNode, _ int) string {
		return item.GetText()
	})
	name := ids[0]
	value := v.Visit(ctx.Expr()).(*miscs.Variable)
	if len(ids) > 1 {
		variable := v.getVariable(name)
		for i := 1; i < len(ids)-1; i++ {
			variable = variable.FieldStruct(ids[i])
		}
		variable.SetFieldStruct(ids[len(ids)-1], value)
	} else {
		v.setVariable(name, value)
	}
	return nil
}

func (v *Visitor) VisitDefineFunctionStat(ctx *grammar.DefineFunctionStatContext) interface{} {
	methodName := ctx.ID().GetText()

	v.functions[methodName] = func(args ...interface{}) interface{} {
		defer v.exitLastState()
		v.enterNewState()

		argNames := v.Visit(ctx.FunctionDefineArguments()).([]string)
		for i, name := range argNames {
			v.defineVariable(name, args[i].(*miscs.Variable))
		}

		r := v.Visit(ctx.StatBlock())
		return r
	}
	return nil
}

func (v *Visitor) VisitFunctionDefineArguments(ctx *grammar.FunctionDefineArgumentsContext) interface{} {
	ids := make([]string, len(ctx.AllID()))
	for i, id := range ctx.AllID() {
		ids[i] = id.GetText()
	}
	return ids
}

func (v *Visitor) enterNewState() {
	defer func() {
		v.mu.Unlock()
	}()
	v.mu.Lock()
	v.state++
	v.variables[v.state] = map[string]*miscs.Variable{}
}

func (v *Visitor) exitLastState() {
	defer func() {
		v.mu.Unlock()
	}()
	v.mu.Lock()
	delete(v.variables, v.state)
	v.state--
}

func (v *Visitor) VisitMethodCallStat(ctx *grammar.MethodCallStatContext) interface{} {
	return v.Visit(ctx.MethodCall())
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

func (v *Visitor) VisitParse(ctx *grammar.ParseContext) interface{} {
	return v.Visit(ctx.Block())
}

func (v *Visitor) VisitBlock(ctx *grammar.BlockContext) interface{} {
	for _, stat := range ctx.AllStat() {
		if result, ok := v.Visit(stat).(*miscs.ReturnStatement); ok {
			return result
		}
	}
	return nil
}

func (v *Visitor) VisitStat(ctx *grammar.StatContext) interface{} {
	for _, child := range ctx.GetChildren() {
		return v.Visit(child.(antlr.ParseTree))
	}
	return nil
}

func (v *Visitor) VisitIfStat(ctx *grammar.IfStatContext) interface{} {
	conditions := ctx.AllConditionBlock()
	evaluatedBlock := false
	for _, condition := range conditions {
		evaluated := v.Visit(condition.Expr()).(*miscs.Variable)
		if evaluated.Boolean() {
			evaluatedBlock = true
			v.Visit(condition.StatBlock())
			break
		}
	}

	if !evaluatedBlock && ctx.StatBlock() != nil {
		v.Visit(ctx.StatBlock())
	}

	return nil
}

func (v *Visitor) VisitDefineListenerStat(ctx *grammar.DefineListenerStatContext) interface{} {
	variable := v.Visit(ctx.MethodCall())
	var ch chan miscs.Variable
	if variableStatement, ok := variable.(*miscs.Variable); ok {
		ch = variableStatement.Chan()
	} else if returnStatement, ok := variable.(*miscs.ReturnStatement); ok {
		ch = returnStatement.Variable().Chan()
	} else {
		return nil
	}
	go func() {
		for {
			value := <-ch
			if ctx.ID() != nil {
				idName := ctx.ID().GetText()
				v.defineVariable(idName, &value)
			}
			v.Visit(ctx.StatBlock())
		}
	}()
	return nil
}
