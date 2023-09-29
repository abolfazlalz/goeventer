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
	v.socketMethods()
}

func (v *Visitor) socketMethods() {
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

	v.functions["send_socket"] = func(i ...interface{}) interface{} {
		userId := i[0].(*miscs.Variable)
		msg := i[1].(*miscs.Variable)
		websocket.Socket().Broadcast(userId.StringValue(), msg.StringValue())
		return nil
	}
}
