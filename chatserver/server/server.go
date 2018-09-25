package server

import (
	"io"
	"log"
	"net"
	"sync"

	"github.com/nqbao/learn-go/chatserver/protocol"
)

type client struct {
	conn   net.Conn
	writer *protocol.CommandWriter
}

type ChatServer struct {
	listner net.Listener
	clients []*client
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

func (s *ChatServer) Broadcast(msg interface{}) {
	for _, client := range s.clients {
		// TODO: handle error here?
		client.writer.Write(msg)
	}
}

func (s *ChatServer) accept(conn net.Conn) {
	log.Printf("Accepting connection from %v, total clients: %v", conn.RemoteAddr().String(), len(s.clients)+1)

	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.clients = append(s.clients, &client{
		conn:   conn,
		writer: protocol.NewCommandWriter(conn),
	})

	go func() {
		cmdReader := protocol.NewCommandReader(conn)

		defer func() {
			s.mutex.Lock()
			defer s.mutex.Unlock()

			// remove the connections from clients array
			for i, check := range s.clients {
				if check.conn == conn {
					s.clients = append(s.clients[:i], s.clients[i+1:]...)
				}
			}

			log.Printf("Closing connection from %v", conn.RemoteAddr().String())
			conn.Close()
		}()

		for {
			cmd, err := cmdReader.Read()

			if err != nil && err != io.EOF {
				log.Printf("Read error: %v", err)
			}

			if cmd != nil {
				switch v := cmd.(type) {
				case protocol.SendCommand:
					go s.Broadcast(protocol.MessageCommand{
						Message: v.Message,
					})
				}
			}

			if err == io.EOF {
				break
			}
		}
	}()
}
