package visitor

import "github.com/abolfazlalz/goeventer/internal/interpreter/miscs"

func (v *Visitor) getVariable(name string) *miscs.Variable {
	for i := v.state; 0 <= i; i-- {
		if values, ok := v.variables[i]; ok {
			if value, ok := values[name]; ok {
				return value
			}
		}
	}
	panic("undefined variable")
}

func (v *Visitor) setVariable(name string, value *miscs.Variable) {
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

func (v *Visitor) variableFromContext(value interface{}) *miscs.Variable {
	if variable, ok := value.(*miscs.Variable); ok {
		return variable
	} else if returnValue, ok := value.(*miscs.ReturnStatement); ok {
		return returnValue.Variable()
	}
	panic("undefined variable")
}

func (v *Visitor) defineVariable(name string, value *miscs.Variable) {
	v.variables[v.state][name] = value
}
