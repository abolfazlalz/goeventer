package websocket

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

type Handler struct {
	clientManager *ClientManager
	upgrader      websocket.Upgrader
	chatCh        []chan *Chat
}

func (handler *Handler) ClientManager() *ClientManager {
	return handler.clientManager
}

func (handler *Handler) AddChatListener(ch chan *Chat) {
	handler.chatCh = append(handler.chatCh, ch)
}

func NewHandler(clientManager *ClientManager) *Handler {
	if clientManager == nil {
		clientManager = NewClientManager()
	}
	return &Handler{
		clientManager: clientManager,
		upgrader: websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
			CheckOrigin: func(r *http.Request) bool {
				// Check origin here for CSRF protection
				return true
			},
		},
	}
}

func (handler *Handler) HandleConnection(w http.ResponseWriter, r *http.Request) {
	client, err := handler.getNewClient(w, r)
	if err != nil {
		log.Print(err)
		return
	}

	handler.handleClientMessages(client)
}

func (handler *Handler) getNewClient(w http.ResponseWriter, r *http.Request) (*Client, error) {
	conn, err := handler.upgrader.Upgrade(w, r, nil)
	if err != nil {
		return nil, err
	}

	clientID := conn.RemoteAddr().String()

	client := NewClient(clientID, conn)
	go func() {
		for {
			client.writer()
		}
	}()
	GetClientListener().Notify(client)

	// Add the client to the Manager
	handler.clientManager.AddClient(client)

	return client, nil
}

func (handler *Handler) handleClientMessages(client *Client) {
	defer func() {
		handler.clientManager.RemoveClient(client.ID)
		client.Conn.Close()
	}()

	for {
		_, message, err := client.Conn.ReadMessage()
		if err != nil {
			log.Print(err)
			return
		}

		chat := NewChat(client.ID, string(message))
		// Put incoming message to client's channel
		for _, listener := range client.listeners {
			listener <- chat
		}
		for _, listener := range handler.chatCh {
			listener <- chat
		}
	}
}

func (handler *Handler) Broadcast(id string, message string) {
	client, ok := handler.clientManager.clients[id]
	if !ok {
		log.Printf("Client %s not found", id)
		return
	}

	// send message to the client's Send channel
	client.send <- []byte(message)
}
