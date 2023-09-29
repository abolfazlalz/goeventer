package visitor

import (
	"github.com/abolfazlalz/goeventer/internal/interpreter/lib"
	"github.com/abolfazlalz/goeventer/internal/interpreter/miscs"
	"github.com/abolfazlalz/goeventer/internal/websocket"
)

func (v *Visitor) loadBuiltinMethods() {
	methods := lib.Methods()
	for _, method := range methods {
		v.functions[method.Name()] = method.Call
	}
	v.onSocketMessage()
}

func (v *Visitor) onSocketMessage() {
	defer func() {
		v.mu.Unlock()
	}()
	v.mu.Lock()

	v.functions["socket"] = func(i ...interface{}) interface{} {
		socketCh := make(chan *websocket.Chat, 1)
		ch := miscs.NewChanVariable()
		websocket.Socket().AddChatListener(socketCh)
		go func() {
			for {
				chat := <-socketCh
				value := miscs.StructType{
					"user_id": miscs.NewVariable(chat.SocketID),
					"text":    miscs.NewVariable(chat.Text),
				}
				ch.NotifyChan(*miscs.NewVariable(value))
			}
		}()
		return miscs.NewReturnStatement(ch)
	}
}

func (v *Visitor) sendSocketMessage() {

}
