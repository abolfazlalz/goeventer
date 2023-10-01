package visitor

import (
	"github.com/abolfazlalz/goeventer/internal/interpreter/lib"
)

func (v *Visitor) loadBuiltinMethods() {
	defer func() { v.mu.Unlock() }()
	v.mu.Lock()
	methods := lib.Methods()
	for _, method := range methods {
		v.functions[method.Name()] = method.Call
	}
}
