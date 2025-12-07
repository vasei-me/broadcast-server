package broadcast

type Client interface {
	Send(message []byte) error
	ID() string
}

type Broadcaster interface {
	Register(client Client)
	Unregister(client Client)
	Broadcast(message []byte, senderID string)
	GetClientsCount() int
}