package websocket

import (
	"net/http"
)

type Connection interface {
    ReadMessage() (messageType int, p []byte, err error)
    WriteMessage(messageType int, data []byte) error
    Close() error
}

type Upgrader interface {
    Upgrade(w http.ResponseWriter, r *http.Request, responseHeader http.Header) (Connection, error)
}