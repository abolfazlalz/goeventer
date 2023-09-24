package websocket

import (
	"github.com/gorilla/websocket"
	"log"
)

type Client struct {
	ID        string
	Conn      *websocket.Conn
	Channel   chan []byte
	send      chan []byte
	listeners []chan *Chat
}

func NewClient(id string, conn *websocket.Conn) *Client {
	client := &Client{
		ID:      id,
		Conn:    conn,
		Channel: make(chan []byte),
		send:    make(chan []byte),
	}
	return client
}

func (client *Client) Send(msg []byte) {
	client.send <- msg
}

func (client *Client) AddListener(ch chan *Chat) {
	client.listeners = append(client.listeners, ch)
}

func (client *Client) writer() {
	for message := range client.send {
		err := client.Conn.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			log.Println(err)
			return
		}
	}
	err := client.Conn.Close()
	if err != nil {
		return
	}
}
