package main

import (
	"flag"
	"log"

	"github.com/nqbao/learn-go/chatserver/client"
	"github.com/nqbao/learn-go/chatserver/client/tui"
	"github.com/nqbao/learn-go/chatserver/server"
)

func main() {
	address := flag.String("server", "", "Which server to connect to")

	flag.Parse()

	if *address == "" {
		server := server.NewServer()
		server.Listen(":3333")

		// start the server
		server.Start()
	} else {
		client := client.NewClient()
		err := client.Dial(*address)

		if err != nil {
			log.Fatal(err)
		}

		// start the client to listen for incoming message
		defer client.Close()

		go client.Start()

		tui.StartUi(client)

		// go func() {
		// 	for msg := range client.Incoming {
		// 		fmt.Printf("> %v\n", msg)
		// 	}
		// }()

		// // send message
		// reader := bufio.NewReader(os.Stdin)
		// for {
		// 	msg, err := reader.ReadString('\n')

		// 	if err != nil {
		// 		panic(err)
		// 	}

		// 	client.Send(msg)
		// }
	}
}
