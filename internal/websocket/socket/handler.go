package socket

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

type Handler struct {
	ClientManager *ClientManager
	upgrader      websocket.Upgrader
}

func NewHandler(clientManager *ClientManager) *Handler {
	return &Handler{
		ClientManager: clientManager,
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

	// Add the client to the Manager
	handler.ClientManager.AddClient(client)

	return client, nil
}

func (handler *Handler) handleClientMessages(client *Client) {
	defer func() {
		handler.ClientManager.RemoveClient(client.ID)
		client.Conn.Close()
	}()

	for {
		_, message, err := client.Conn.ReadMessage()
		if err != nil {
			log.Print(err)
			return
		}

		// Put incoming message to client's channel
		client.Channel <- message
	}
}

func (handler *Handler) Broadcast(id string, message string) {
	client, ok := handler.ClientManager.Clients[id]
	if !ok {
		log.Printf("Client %s not found", id)
		return
	}
	err := client.Conn.WriteMessage(websocket.TextMessage, []byte(message))
	if err != nil {
		log.Println(err)
		return
	}
}
