package socket

import (
	"sync"
)

type ClientManager struct {
	ClientsMutex *sync.Mutex
	Clients      map[string]*Client
}

func NewClientManager() *ClientManager {
	return &ClientManager{
		ClientsMutex: &sync.Mutex{},
		Clients:      make(map[string]*Client),
	}
}

func (manager *ClientManager) AddClient(client *Client) {
	manager.ClientsMutex.Lock()
	manager.Clients[client.ID] = client
	manager.ClientsMutex.Unlock()
}

func (manager *ClientManager) RemoveClient(id string) {
	manager.ClientsMutex.Lock()
	delete(manager.Clients, id)
	manager.ClientsMutex.Unlock()
}
