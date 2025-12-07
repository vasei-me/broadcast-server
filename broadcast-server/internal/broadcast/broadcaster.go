package broadcast

import (
	"fmt"
	"sync"
)

type ClientInfo struct {
    ID   string
    Send func(message []byte) error
}

type DefaultBroadcaster struct {
    clients map[string]ClientInfo
    mu      sync.RWMutex
}

func NewDefaultBroadcaster() *DefaultBroadcaster {
    return &DefaultBroadcaster{
        clients: make(map[string]ClientInfo),
    }
}

func (b *DefaultBroadcaster) Register(client Client) {
    b.mu.Lock()
    defer b.mu.Unlock()
    
    b.clients[client.ID()] = ClientInfo{
        ID:   client.ID(),
        Send: client.Send,
    }
    fmt.Printf("Client %s connected. Total clients: %d\n", client.ID(), len(b.clients))
}

func (b *DefaultBroadcaster) Unregister(client Client) {
    b.mu.Lock()
    defer b.mu.Unlock()
    
    delete(b.clients, client.ID())
    fmt.Printf("Client %s disconnected. Total clients: %d\n", client.ID(), len(b.clients))
}

func (b *DefaultBroadcaster) Broadcast(message []byte, senderID string) {
    b.mu.RLock()
    defer b.mu.RUnlock()
    
    for id, client := range b.clients {
        if id != senderID {
            go client.Send(message)
        }
    }
}

func (b *DefaultBroadcaster) GetClientsCount() int {
    b.mu.RLock()
    defer b.mu.RUnlock()
    return len(b.clients)
}