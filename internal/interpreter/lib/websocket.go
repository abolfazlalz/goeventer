package lib

import (
	"github.com/abolfazlalz/goeventer/internal/interpreter/miscs"
	"github.com/abolfazlalz/goeventer/internal/websocket"
)

type WebSocketListen struct {
}

func (WebSocketListen) Name() string {
	return "listen_socket"
}

func (WebSocketListen) Call(_ ...interface{}) interface{} {
	socketCh := make(chan *websocket.Chat)
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

type WebSocketSend struct {
}

func (WebSocketSend) Name() string {
	return "send_socket"
}

func (WebSocketSend) Call(i ...interface{}) interface{} {
	userId := miscs.GetVariable(i[0])
	msg := miscs.GetVariable(i[1])
	websocket.Socket().Broadcast(userId.StringValue(), msg.StringValue())
	return nil
}
