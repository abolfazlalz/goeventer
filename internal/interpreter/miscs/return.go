package miscs

type ReturnStatement struct {
	variable *Variable
}

func NewReturnStatement(variable *Variable) *ReturnStatement {
	return &ReturnStatement{variable: variable}
}

func (rs ReturnStatement) Variable() *Variable {
	return rs.variable
}
