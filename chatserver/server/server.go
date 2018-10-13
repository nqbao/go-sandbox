package server

type ChatServer interface {
	Listen(address string) error
	Broadcast(msg interface{})
	Start()
	Close()
}
