package lib

import "log"

type LogMethod struct {
}

func (l LogMethod) Name() string {
	return "log"
}

func (l LogMethod) Call(p ...interface{}) interface{} {
	log.Println(p...)
	return nil
}
