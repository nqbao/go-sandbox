package client

type messageHandler func(string)

type ChatClient interface {
	Dial(address string) error
	Start()
	Close()
	Send(message string) error
}
