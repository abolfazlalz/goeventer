package socket

import (
	"github.com/gorilla/websocket"
)

type Client struct {
	ID      string
	Conn    *websocket.Conn
	Channel chan []byte
}

func NewClient(id string, conn *websocket.Conn) *Client {
	return &Client{
		ID:      id,
		Conn:    conn,
		Channel: make(chan []byte),
	}
}
