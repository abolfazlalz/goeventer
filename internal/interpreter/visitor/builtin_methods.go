package visitor

import "github.com/abolfazlalz/goeventer/internal/interpreter/lib"

func (v *Visitor) loadBuiltinMethods() {
	methods := lib.Methods()
	for _, method := range methods {
		v.functions[method.Name()] = method.Call
	}
}
