package main

import (
	"github.com/nqbao/learn-go/chatserver/server"
)

func main() {
	server := server.NewServer()
	server.Listen(":3333")

	// start the server
	server.Start()
}
