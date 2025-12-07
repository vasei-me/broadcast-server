package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"broadcast-server/internal/broadcast"
	"broadcast-server/internal/websocket"
)

type WebSocketServer struct {
    addr        string
    broadcaster broadcast.Broadcaster
    upgrader    websocket.Upgrader
    server      *http.Server
    stopChan    chan struct{}
}

func NewWebSocketServer(addr string, broadcaster broadcast.Broadcaster, upgrader websocket.Upgrader) *WebSocketServer {
    return &WebSocketServer{
        addr:        addr,
        broadcaster: broadcaster,
        upgrader:    upgrader,
        stopChan:    make(chan struct{}),
    }
}

func (s *WebSocketServer) Start() error {
    mux := http.NewServeMux()
    mux.HandleFunc("/ws", s.handleWebSocket)
    mux.HandleFunc("/health", s.handleHealthCheck)
    
    s.server = &http.Server{
        Addr:    s.addr,
        Handler: mux,
    }
    
    go s.gracefulShutdown()
    
    log.Printf("Server starting on %s", s.addr)
    return s.server.ListenAndServe()
}

func (s *WebSocketServer) handleWebSocket(w http.ResponseWriter, r *http.Request) {
    conn, err := s.upgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Printf("Failed to upgrade connection: %v", err)
        return
    }
    defer conn.Close()
    
    client := NewWebSocketClient(conn, s.broadcaster)
    s.broadcaster.Register(client)
    defer s.broadcaster.Unregister(client)
    
    client.HandleMessages()
}

func (s *WebSocketServer) handleHealthCheck(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, "Clients connected: %d", s.broadcaster.GetClientsCount())
}

func (s *WebSocketServer) Stop() {
    if s.server != nil {
        ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
        defer cancel()
        s.server.Shutdown(ctx)
    }
    close(s.stopChan)
}

func (s *WebSocketServer) gracefulShutdown() {
    sigChan := make(chan os.Signal, 1)
    signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
    
    select {
    case <-sigChan:
        log.Println("Received shutdown signal")
    case <-s.stopChan:
        log.Println("Received stop signal")
    }
    
    s.Stop()
}

func (s *WebSocketServer) GetClientsCount() int {
    return s.broadcaster.GetClientsCount()
}