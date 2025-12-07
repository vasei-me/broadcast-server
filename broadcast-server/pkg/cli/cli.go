package cli

import (
	"broadcast-server/internal/broadcast"
	"broadcast-server/internal/server"
	"broadcast-server/internal/websocket"
	"flag"
	"fmt"
	"log"
	"os"
)

type Command interface {
    Execute() error
    Name() string
}

type StartCommand struct {
    addr string
}

func NewStartCommand(addr string) *StartCommand {
    return &StartCommand{addr: addr}
}

func (c *StartCommand) Execute() error {
    broadcaster := broadcast.NewDefaultBroadcaster()
    upgrader := websocket.NewWebSocketUpgrader()
    srv := server.NewWebSocketServer(c.addr, broadcaster, upgrader)
    
    log.Printf("Starting broadcast server on %s", c.addr)
    return srv.Start()
}

func (c *StartCommand) Name() string {
    return "start"
}

type ConnectCommand struct {
    addr string
}

func NewConnectCommand(addr string) *ConnectCommand {
    return &ConnectCommand{addr: addr}
}

func (c *ConnectCommand) Execute() error {
    log.Printf("Use the broadcast-client executable to connect")
    log.Printf("Or run: go run cmd/client/main.go --addr=%s", c.addr)
    return nil
}

func (c *ConnectCommand) Name() string {
    return "connect"
}

func ParseCommands() {
    if len(os.Args) < 2 {
        printUsage()
        os.Exit(1)
    }
    
    switch os.Args[1] {
    case "start":
        addr := flag.String("addr", ":8080", "server address")
        flag.CommandLine.Parse(os.Args[2:])
        
        cmd := NewStartCommand(*addr)
        if err := cmd.Execute(); err != nil {
            log.Fatal(err)
        }
        
    case "connect":
        addr := flag.String("addr", "localhost:8080", "server address")
        flag.CommandLine.Parse(os.Args[2:])
        
        cmd := NewConnectCommand(*addr)
        if err := cmd.Execute(); err != nil {
            log.Fatal(err)
        }
        
    default:
        printUsage()
        os.Exit(1)
    }
}

func printUsage() {
    fmt.Println("Broadcast Server Commands:")
    fmt.Println("  broadcast-server start [--addr=:8080]")
    fmt.Println("  broadcast-server connect [--addr=localhost:8080]")
    fmt.Println("\nExamples:")
    fmt.Println("  broadcast-server start --addr=:9090")
    fmt.Println("  broadcast-client --addr=localhost:9090")
}