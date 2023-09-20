package interpreter

type ReturnStatement struct {
	variable *Variable
}

func NewReturnStatement(variable *Variable) *ReturnStatement {
	return &ReturnStatement{variable: variable}
}
