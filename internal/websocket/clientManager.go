package websocket

import (
	"sync"
)

type ClientManager struct {
	clientsMutex *sync.Mutex
	clients      map[string]*Client
	clientCh     chan *Client
	quitCh       chan bool
}

func NewClientManager() *ClientManager {
	return &ClientManager{
		clientsMutex: &sync.Mutex{},
		clients:      make(map[string]*Client),
	}
}

func (manager *ClientManager) AddClient(client *Client) {
	manager.clientsMutex.Lock()
	manager.clients[client.ID] = client
	manager.clientsMutex.Unlock()
}

func (manager *ClientManager) RemoveClient(id string) {
	manager.clientsMutex.Lock()
	delete(manager.clients, id)
	manager.clientsMutex.Unlock()
	manager.quitCh <- true
}
