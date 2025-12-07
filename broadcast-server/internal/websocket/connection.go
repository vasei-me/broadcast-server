package websocket

import (
	"net/http"

	"github.com/gorilla/websocket"
)

type WebSocketConnection struct {
    conn *websocket.Conn
}

func NewWebSocketConnection(conn *websocket.Conn) *WebSocketConnection {
    return &WebSocketConnection{conn: conn}
}

func (w *WebSocketConnection) ReadMessage() (int, []byte, error) {
    return w.conn.ReadMessage()
}

func (w *WebSocketConnection) WriteMessage(messageType int, data []byte) error {
    return w.conn.WriteMessage(messageType, data)
}

func (w *WebSocketConnection) Close() error {
    return w.conn.Close()
}

type WebSocketUpgrader struct {
    upgrader *websocket.Upgrader
}

func NewWebSocketUpgrader() *WebSocketUpgrader {
    return &WebSocketUpgrader{
        upgrader: &websocket.Upgrader{
            CheckOrigin: func(r *http.Request) bool { return true },
        },
    }
}

func (w *WebSocketUpgrader) Upgrade(writer http.ResponseWriter, request *http.Request, responseHeader http.Header) (Connection, error) {
    conn, err := w.upgrader.Upgrade(writer, request, responseHeader)
    if err != nil {
        return nil, err
    }
    return NewWebSocketConnection(conn), nil
}