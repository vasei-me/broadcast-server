package server

import (
	"broadcast-server/internal/broadcast"
	"broadcast-server/internal/websocket"
	"fmt"
	"log"
	"time"
)

// WebSocketClient (این برای سرور است، برای مدیریت connectionهای ورودی)
type WebSocketClient struct {
    id          string
    conn        websocket.Connection
    broadcaster broadcast.Broadcaster
    sendChan    chan []byte
    doneChan    chan struct{}
}

func NewWebSocketClient(conn websocket.Connection, broadcaster broadcast.Broadcaster) *WebSocketClient {
    return &WebSocketClient{
        id:          fmt.Sprintf("client-%d", time.Now().UnixNano()),
        conn:        conn,
        broadcaster: broadcaster,
        sendChan:    make(chan []byte, 256),
        doneChan:    make(chan struct{}),
    }
}

func (c *WebSocketClient) ID() string {
    return c.id
}

func (c *WebSocketClient) Send(message []byte) error {
    select {
    case c.sendChan <- message:
        return nil
    default:
        return fmt.Errorf("send channel full")
    }
}

func (c *WebSocketClient) HandleMessages() {
    go c.writePump()
    c.readPump()
}

func (c *WebSocketClient) readPump() {
    defer close(c.doneChan)
    
    for {
        _, message, err := c.conn.ReadMessage()
        if err != nil {
            log.Printf("Error reading message: %v", err)
            break
        }
        
        log.Printf("Received message from %s: %s", c.id, string(message))
        c.broadcaster.Broadcast(message, c.id)
    }
}

func (c *WebSocketClient) writePump() {
    for {
        select {
        case message := <-c.sendChan:
            err := c.conn.WriteMessage(1, message)
            if err != nil {
                log.Printf("Error writing message: %v", err)
                return
            }
        case <-c.doneChan:
            return
        }
    }
}