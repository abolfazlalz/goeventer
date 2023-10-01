package miscs

import (
	"fmt"
	"math"
	"strconv"
)

type VariableType int

type StructType map[string]*Variable

const (
	StringVariable VariableType = iota
	IntegerVariable
	BoolVariable
	ChanVariable
	StructVariable
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
	} else if _, ok := value.(StructType); ok {
		return StructVariable
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

func GetVariable(value interface{}) *Variable {
	if variable, ok := value.(*Variable); ok {
		return variable
	}
	if returnStat, ok := value.(*ReturnStatement); ok {
		return returnStat.Variable()
	}
	return NewVariable(nil)
}

func NewChanVariable() *Variable {
	return NewVariable(make(chan Variable))
}

func (v *Variable) TypeName() string {
	return v.Type.String()
}

func (v *Variable) String() string {
	return fmt.Sprintf("%v", v.Value)
}

func (v *Variable) Boolean() bool {
	if v.Type != BoolVariable {
		panic("value is not a bool value.")
	}

	if val, ok := v.Value.(bool); ok {
		return val
	}
	return false
}

func (v *Variable) StringValue() string {
	if v.Type != StringVariable {
		panic("value is not a string value.")
	}
	return fmt.Sprintf("%v", v.Value)
}

func (v *Variable) StructValue() StructType {
	if v.Type != StructVariable {
		panic("value is not a struct type")
	}
	return v.Value.(StructType)
}

func (v *Variable) FieldStruct(field string) *Variable {
	structValue := v.StructValue()
	variable := structValue[field]
	return variable
}

func (v *Variable) SetFieldStruct(field string, value *Variable) {
	structValue := v.StructValue()
	structValue[field] = value
	v.Value = structValue
}

func (v *Variable) Integer() int {
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

func (v *Variable) Chan() chan Variable {
	if v.Type != ChanVariable {
		panic("value is not channel")
	}
	if val, ok := v.Value.(chan Variable); ok {
		return val
	}
	panic("not defined channel")
}

func (v *Variable) panicForUndefinedVariable(sec *Variable) {
	if v.Type != sec.Type {
		panic(fmt.Sprintf("Expected same types but received %T and %T", v.TypeName(), sec.TypeName()))
	}
}

func (v *Variable) Sum(sec *Variable) interface{} {
	v.panicForUndefinedVariable(sec)
	if v.Type == IntegerVariable {
		return v.Integer() + sec.Integer()
	}
	return v.StringValue() + sec.StringValue()
}

func (v *Variable) Equal(sec *Variable) interface{} {
	v.panicForUndefinedVariable(sec)
	if v.Type == IntegerVariable {
		return v.Integer() == sec.Integer()
	} else if v.Type == BoolVariable {
		return v.Boolean() == sec.Boolean()
	}
	return v.StringValue() == sec.StringValue()
}

func (v *Variable) Minus(sec *Variable) interface{} {
	v.panicForUndefinedVariable(sec)
	if v.Type != IntegerVariable {
		panic("Invalid operation: Only Integer Variables are allowed for minus")
	}
	return v.Integer() - sec.Integer()
}

func (v *Variable) Multiple(sec *Variable) interface{} {
	v.panicForUndefinedVariable(sec)
	if v.Type != IntegerVariable {
		panic("Invalid operation: Only Integer Variables are allowed for multiplication")
	}
	return v.Integer() * sec.Integer()
}

func (v *Variable) Divide(sec *Variable) interface{} {
	v.panicForUndefinedVariable(sec)
	if v.Type != IntegerVariable {
		panic("Invalid operation: Only Integer Variables are allowed for dividing")
	}
	return v.Integer() / sec.Integer()
}

func (v *Variable) Mod(sec *Variable) interface{} {
	v.panicForUndefinedVariable(sec)
	if v.Type != IntegerVariable {
		panic("Invalid operation: Only Integer Variables are allowed for dividing")
	}
	return int(math.Mod(float64(v.Integer()), float64(sec.Integer())))
}

func (v *Variable) Pow(sec *Variable) interface{} {
	v.panicForUndefinedVariable(sec)
	if v.Type != IntegerVariable {
		panic("Invalid operation: Only Integer Variables are allowed ")
	}
	return math.Pow(float64(v.Integer()), float64(sec.Integer()))
}

func (v *Variable) NotifyChan(vb Variable) {
	ch := v.Chan()
	ch <- vb
}
