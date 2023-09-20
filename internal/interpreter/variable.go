package interpreter

import (
	"fmt"
	"math"
	"strconv"
)

type VariableType int

const (
	StringVariable VariableType = iota
	IntegerVariable
	BoolVariable
	ChanVariable
	UndefinedVariable
)

func GetVariableType(value interface{}) VariableType {
	if _, ok := value.(int); ok {
		return IntegerVariable
	} else if _, ok := value.(string); ok {
		return StringVariable
	} else if _, ok := value.(bool); ok {
		return BoolVariable
	} else if _, ok := value.(chan Variable); ok {
		return ChanVariable
	}
	return UndefinedVariable
}

func (v VariableType) String() string {
	if v == StringVariable {
		return "string"
	} else if v == IntegerVariable {
		return "integer"
	}
	panic("Unknown variable type")
}

type Variable struct {
	Value interface{}
	Type  VariableType
}

func NewVariable(value interface{}) *Variable {
	return &Variable{
		Value: value,
		Type:  GetVariableType(value),
	}
}

func (v Variable) TypeName() string {
	return v.Type.String()
}

func (v Variable) String() string {
	return fmt.Sprintf("%v", v.Value)
}

func (v Variable) Boolean() bool {
	if v.Type != BoolVariable {
		panic("value is not a bool value.")
	}

	if val, ok := v.Value.(bool); ok {
		return val
	}
	return false
}

func (v Variable) StringValue() string {
	if v.Type != StringVariable {
		panic("value is not a string value.")
	}
	return fmt.Sprintf("%v", v.Value)
}

func (v Variable) Integer() int {
	if v.Type != IntegerVariable {
		panic("Value is not a integer value.")
	}
	if val, ok := v.Value.(int); ok {
		return val
	} else if val, err := strconv.Atoi(v.StringValue()); err != nil {
		return val
	}
	return 0
}

func (v Variable) panicForUndefinedVariable(sec *Variable) {
	if v.Type != sec.Type {
		panic(fmt.Sprintf("Expected same types but received %T and %T", v.TypeName(), sec.TypeName()))
	}
}

func (v Variable) Sum(sec *Variable) interface{} {
	v.panicForUndefinedVariable(sec)
	if v.Type == IntegerVariable {
		return v.Integer() + sec.Integer()
	}
	return v.StringValue() + sec.StringValue()
}

func (v Variable) Equal(sec *Variable) interface{} {
	v.panicForUndefinedVariable(sec)
	if v.Type == IntegerVariable {
		return v.Integer() == sec.Integer()
	} else if v.Type == BoolVariable {
		return v.Boolean() == sec.Boolean()
	}
	return v.StringValue() == sec.StringValue()
}

func (v Variable) Minus(sec *Variable) interface{} {
	v.panicForUndefinedVariable(sec)
	if v.Type != IntegerVariable {
		panic("Invalid operation: Only Integer Variables are allowed for minus")
	}
	return v.Integer() - sec.Integer()
}

func (v Variable) Multiple(sec *Variable) interface{} {
	v.panicForUndefinedVariable(sec)
	if v.Type != IntegerVariable {
		panic("Invalid operation: Only Integer Variables are allowed for multiplication")
	}
	return v.Integer() * sec.Integer()
}

func (v Variable) Divide(sec *Variable) interface{} {
	v.panicForUndefinedVariable(sec)
	if v.Type != IntegerVariable {
		panic("Invalid operation: Only Integer Variables are allowed for dividing")
	}
	return v.Integer() / sec.Integer()
}

func (v Variable) Mod(sec *Variable) interface{} {
	v.panicForUndefinedVariable(sec)
	if v.Type != IntegerVariable {
		panic("Invalid operation: Only Integer Variables are allowed for dividing")
	}
	return int(math.Mod(float64(v.Integer()), float64(sec.Integer())))
}

func (v Variable) Pow(sec *Variable) interface{} {
	v.panicForUndefinedVariable(sec)
	if v.Type != IntegerVariable {
		panic("Invalid operation: Only Integer Variables are allowed ")
	}
	return math.Pow(float64(v.Integer()), float64(sec.Integer()))
}
