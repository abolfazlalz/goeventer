package lib

type MethodInterface interface {
	Name() string
	Call(params ...interface{}) interface{}
}

func Methods() []MethodInterface {
	return []MethodInterface{
		LogMethod{},
		Print{},
		WebSocketListen{},
		WebSocketSend{},
		JsonDecode{},
		JsonEncode{},
	}
}
