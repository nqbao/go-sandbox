package server

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"sync"
)

type ChatServer struct {
	listner net.Listener
	clients []net.Conn
	mutex   *sync.Mutex
}

func NewServer() *ChatServer {
	return &ChatServer{
		mutex: &sync.Mutex{},
	}
}

func (s *ChatServer) Listen(address string) error {
	l, err := net.Listen("tcp", address)

	if err == nil {
		s.listner = l
	}

	log.Printf("Listening on %v", address)

	return err
}

func (s *ChatServer) Close() {
	s.listner.Close()
}

func (s *ChatServer) Start() {
	for {
		// XXX: need a way to break the loop
		conn, err := s.listner.Accept()

		if err != nil {
			log.Print(err)
		} else {
			// handle connection
			go s.accept(conn)
		}
	}
}

func (s *ChatServer) Broadcast(msg []byte) {
	for _, conn := range s.clients {
		// TODO: handle error here?
		conn.Write(msg)
	}
}

func (s *ChatServer) accept(conn net.Conn) {
	log.Printf("Accepting connection from %v", conn.RemoteAddr().String())

	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.clients = append(s.clients, conn)

	go func() {
		reader := bufio.NewReader(conn)

		defer conn.Close()

		for {
			msg, err := reader.ReadString('\n')

			fmt.Printf("%v\n", msg)

			if err == io.EOF {
				// TODO: close the connection
				log.Printf("Closing connection from %v", conn.RemoteAddr().String())
				break
			} else {
				log.Printf("Read error %v", err)
			}
		}
	}()
}
