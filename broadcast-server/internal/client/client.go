package client

import (
	"bufio"
	"fmt"
	"log"
	"net/url"
	"os"
	"os/signal"

	"github.com/gorilla/websocket"
)

type WebSocketClient struct {
    conn   *websocket.Conn
    done   chan struct{}
    server string
}

func NewWebSocketClient(serverAddr string) *WebSocketClient {
    return &WebSocketClient{
        server: serverAddr,
        done:   make(chan struct{}),
    }
}

func (c *WebSocketClient) Connect() error {
    u := url.URL{Scheme: "ws", Host: c.server, Path: "/ws"}
    
    conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
    if err != nil {
        return fmt.Errorf("failed to connect: %w", err)
    }
    
    c.conn = conn
    log.Printf("Connected to %s", u.String())
    
    return nil
}

func (c *WebSocketClient) Start() {
    // Handle interrupt signal
    interrupt := make(chan os.Signal, 1)
    signal.Notify(interrupt, os.Interrupt)
    
    // Read messages from server
    go func() {
        defer close(c.done)
        for {
            _, message, err := c.conn.ReadMessage()
            if err != nil {
                log.Println("Read error:", err)
                return
            }
            fmt.Printf("\n[Broadcast] %s\n", string(message))
            fmt.Print("> ")
        }
    }()
    
    // Read from stdin and send
    go func() {
        scanner := bufio.NewScanner(os.Stdin)
        fmt.Print("> ")
        for scanner.Scan() {
            text := scanner.Text()
            if text == "" {
                fmt.Print("> ")
                continue
            }
            
            err := c.conn.WriteMessage(websocket.TextMessage, []byte(text))
            if err != nil {
                log.Println("Write error:", err)
                return
            }
            fmt.Print("> ")
        }
    }()
    
    // Wait for interrupt or done
    select {
    case <-interrupt:
        log.Println("Interrupt received")
    case <-c.done:
    }
    
    c.Disconnect()
}

func (c *WebSocketClient) Disconnect() {
    if c.conn != nil {
        c.conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
        c.conn.Close()
    }
}