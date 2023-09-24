package websocket

import "net/http"

var socket *Handler

func Socket() *Handler {
	if socket == nil {
		client := NewClientManager()
		socket = NewHandler(client)
	}
	return socket
}

func Broadcast(id string, message string) {
	socket.Broadcast(id, message)
}

func Clients() map[string]*Client {
	return Socket().ClientManager().clients
}

func HandleConnection(w http.ResponseWriter, r *http.Request) {
	Socket().HandleConnection(w, r)
}
