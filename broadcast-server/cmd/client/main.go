package main

import (
	"flag"
	"log"

	"broadcast-server/internal/client"
)

func main() {
    addr := flag.String("addr", "localhost:8080", "server address")
    flag.Parse()
    
    wsClient := client.NewWebSocketClient(*addr)
    
    if err := wsClient.Connect(); err != nil {
        log.Fatal("Connection error:", err)
    }
    
    wsClient.Start()
}