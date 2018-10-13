package main

import (
	"github.com/nqbao/learn-go/chatserver/server"
)

func main() {
	var s server.ChatServer
	s = server.NewServer()
	s.Listen(":3333")

	// start the server
	s.Start()
}
