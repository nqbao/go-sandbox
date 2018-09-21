package main

import (
	"flag"
	"log"

	"github.com/nqbao/learn-go/chatserver/client"
	"github.com/nqbao/learn-go/chatserver/server"
)

func main() {
	address := flag.String("server", "", "Which server to connect to")

	flag.Parse()

	if *address == "" {
		server := server.NewServer()
		server.Listen(":3333")
		server.Start()
	} else {
		client := client.NewClient()
		err := client.Dial(*address)

		if err != nil {
			log.Fatal(err)
		}

		client.Send("Hello")

		defer client.Close()
	}
}
