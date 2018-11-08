package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

func handleRequest(conn net.Conn) {
	log.Printf("Accepting new connection %v", conn.RemoteAddr())

	close := func() {
		log.Print("Closing connection")
		conn.Close()
	}

	defer close()

	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)

	for {
		str, err := reader.ReadString('\n')

		if err == io.EOF {
			log.Print("Client close connection")
			break
		} else if err != nil {
			log.Panic(err)
		}

		str = strings.TrimSpace(str)

		if str == "STOP" {
			log.Printf("Receive stop signal")
			close()
			break
		} else {
			writer.WriteString(fmt.Sprintf("> %s\n", str))
			writer.Flush()
		}
	}
}

func main() {
	// 1st: create a listener
	l, err := net.Listen("tcp", ":3333")

	if err != nil {
		log.Fatal(err)
	}

	log.Print("Listen on port 3333")

	// close the socket when server is ended
	defer l.Close()

	for {
		conn, err := l.Accept()

		if err != nil {
			log.Fatal(err)
		}

		go handleRequest(conn)
	}
}
