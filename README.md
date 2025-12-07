\# Broadcast Server - WebSocket Broadcast System



A SOLID architecture WebSocket broadcast server implemented in Go, allowing real-time message broadcasting to multiple connected clients.



!\[Go Version](https://img.shields.io/badge/Go-1.21%2B-blue)

!\[WebSocket](https://img.shields.io/badge/WebSocket-Gorilla-orange)

!\[SOLID](https://img.shields.io/badge/Architecture-SOLID-green)

!\[License](https://img.shields.io/badge/License-MIT-brightgreen)



\## 📋 Project Overview



This project implements a broadcast server that enables real-time communication between multiple clients using WebSockets. When a client sends a message, the server broadcasts it to all other connected clients. The implementation follows SOLID principles for clean, maintainable, and extensible code.



\## ✨ Features



\- \*\*Real-time broadcasting\*\* - Instant message delivery to all connected clients

\- \*\*Multi-client support\*\* - Handle multiple simultaneous connections

\- \*\*Graceful connection management\*\* - Automatic client registration and cleanup

\- \*\*SOLID architecture\*\* - Clean separation of concerns and extensible design

\- \*\*CLI interface\*\* - Easy-to-use command line interface for both server and client

\- \*\*Health monitoring\*\* - Built-in health check endpoint

\- \*\*Error handling\*\* - Robust error handling and graceful shutdown



\## 🏗️ Architecture



The project follows a clean architecture with SOLID principles:



```

broadcast-server/

├── cmd/                    # Entry points

│   ├── server/            # Server CLI

│   └── client/            # Client CLI

├── internal/              # Internal packages

│   ├── websocket/         # WebSocket abstraction layer

│   ├── broadcast/         # Broadcast logic and interfaces

│   ├── server/            # Server implementation

│   └── client/            # Client implementation

└── pkg/cli/               # CLI command parsing

```



\### SOLID Principles Applied:



\- \*\*Single Responsibility\*\*: Each class has a single, well-defined responsibility

\- \*\*Open/Closed\*\*: Components are open for extension but closed for modification

\- \*\*Liskov Substitution\*\*: Interfaces can be substituted with their implementations

\- \*\*Interface Segregation\*\*: Small, focused interfaces instead of large ones

\- \*\*Dependency Inversion\*\*: Depend on abstractions, not concretions



\## 🚀 Getting Started



\### Prerequisites



\- Go 1.21 or higher

\- Git



\### Installation



1\. Clone the repository:

```bash

git clone https://github.com/yourusername/broadcast-server.git

cd broadcast-server

```



2\. Build the executables:

```bash

\# Build server

go build -o broadcast-server ./cmd/server



\# Build client

go build -o broadcast-client ./cmd/client



\# Or build both at once

go build ./cmd/...

```



\### Usage



\#### Starting the Server



```bash

\# Default port (8080)

./broadcast-server start



\# Custom port

./broadcast-server start --addr=:9090

```



\#### Connecting as a Client



```bash

\# Connect to default server

./broadcast-client --addr=localhost:8080



\# Connect to custom server

./broadcast-client --addr=localhost:9090

```



\#### Using the CLI Commands



```bash

\# Show usage

./broadcast-server



\# Start server on port 3000

./broadcast-server start --addr=:3000



\# Show connection instructions

./broadcast-server connect --addr=localhost:3000

```



\## 📖 How It Works



1\. \*\*Server Initialization\*\*: The server starts and listens for WebSocket connections

2\. \*\*Client Connection\*\*: Clients connect to the server via WebSocket

3\. \*\*Message Broadcasting\*\*: When a client sends a message, the server broadcasts it to all other connected clients

4\. \*\*Connection Management\*\*: The server tracks all active connections and cleans up disconnected clients



\### Message Flow

```

Client A → Server → Broadcast → Client B

&nbsp;                        ↘ Client C

&nbsp;                        ↘ Client D

```



\## 🧪 Testing



\### Manual Testing



1\. Start the server in one terminal:

```bash

./broadcast-server start

```



2\. Open multiple client terminals:

```bash

\# Terminal 2

./broadcast-client



\# Terminal 3

./broadcast-client



\# Terminal 4

./broadcast-client

```



3\. Send messages from any client and watch them appear in all other clients.



\### Health Check

The server provides a health endpoint:

```bash

curl http://localhost:8080/health

\# Response: Clients connected: X

```



\## 🔧 Technical Details



\### Dependencies



\- \*\*\[gorilla/websocket](https://github.com/gorilla/websocket)\*\* - Robust WebSocket implementation

\- \*\*Standard Library\*\* - All other dependencies are from Go's standard library



\### Key Components



1\. \*\*WebSocket Layer\*\* (`internal/websocket/`)

&nbsp;  - Abstract interface for WebSocket connections

&nbsp;  - Implementation using gorilla/websocket



2\. \*\*Broadcast Manager\*\* (`internal/broadcast/`)

&nbsp;  - Manages client registrations

&nbsp;  - Handles message broadcasting

&nbsp;  - Thread-safe client management



3\. \*\*Server Core\*\* (`internal/server/`)

&nbsp;  - HTTP server with WebSocket upgrade

&nbsp;  - Client connection handling

&nbsp;  - Graceful shutdown implementation



4\. \*\*CLI Interface\*\* (`pkg/cli/`)

&nbsp;  - Command parsing and routing

&nbsp;  - Clean command pattern implementation



\## 📁 Project Structure



```go

// Example of interface definitions

type Broadcaster interface {

&nbsp;   Register(client Client)

&nbsp;   Unregister(client Client)

&nbsp;   Broadcast(message \[]byte, senderID string)

&nbsp;   GetClientsCount() int

}



type Connection interface {

&nbsp;   ReadMessage() (messageType int, p \[]byte, err error)

&nbsp;   WriteMessage(messageType int, data \[]byte) error

&nbsp;   Close() error

}

```



\## 🛠️ Development



\### Running Tests

```bash

\# Run all tests

go test ./...



\# Run tests with coverage

go test ./... -cover

```



\### Code Quality

```bash

\# Format code

go fmt ./...



\# Vet code

go vet ./...



\# Static analysis

go install honnef.co/go/tools/cmd/staticcheck@latest

staticcheck ./...

```



\## 🔮 Future Enhancements



Potential features for future development:



1\. \*\*Authentication\*\* - User authentication and authorization

2\. \*\*Chat Rooms\*\* - Support for multiple chat rooms/channels

3\. \*\*Message History\*\* - Store and retrieve message history

4\. \*\*TLS/SSL Support\*\* - Secure WebSocket connections (wss://)

5\. \*\*JSON Messages\*\* - Structured message format with metadata

6\. \*\*Client Commands\*\* - Special commands like `/join`, `/users`, `/help`

7\. \*\*Monitoring\*\* - Metrics and monitoring endpoints

8\. \*\*Docker Support\*\* - Containerization for easy deployment



\## 🤝 Contributing



Contributions are welcome! Please feel free to submit a Pull Request.



1\. Fork the repository

2\. Create your feature branch (`git checkout -b feature/AmazingFeature`)

3\. Commit your changes (`git commit -m 'Add some AmazingFeature'`)

4\. Push to the branch (`git push origin feature/AmazingFeature`)

5\. Open a Pull Request



\### Development Guidelines

\- Follow Go coding standards

\- Write tests for new functionality

\- Update documentation as needed

\- Follow SOLID principles



\## 📄 License



This project is licensed under the MIT License - see the \[LICENSE](LICENSE) file for details.



\## 🙏 Acknowledgments



\- \[Gorilla WebSocket](https://github.com/gorilla/websocket) for the excellent WebSocket library

\- SOLID principles for guiding the architecture

\- The Go community for excellent tools and libraries



\## 📞 Support



For support, please:

1\. Check the existing issues

2\. Create a new issue with detailed information

3\. Include steps to reproduce, expected behavior, and actual behavior



---



\*\*Happy Broadcasting!\*\* 🚀



\## 📊 Project Stats



!\[Lines of Code](https://img.shields.io/tokei/lines/github/yourusername/broadcast-server)

!\[GitHub stars](https://img.shields.io/github/stars/yourusername/broadcast-server?style=social)

!\[GitHub forks](https://img.shields.io/github/forks/yourusername/broadcast-server?style=social)



---



\*Built with ❤️ using Go and SOLID principles\*

