package websocket

type Chat struct {
	SocketID string
	Text     string
}

func NewChat(socketID string, text string) *Chat {
	return &Chat{SocketID: socketID, Text: text}
}
