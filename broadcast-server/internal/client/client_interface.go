package client

type Client interface {
	Connect() error
	Start()
	Disconnect()
}