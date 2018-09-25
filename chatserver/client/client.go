package client

import (
	"io"
	"log"
	"net"

	"github.com/nqbao/learn-go/chatserver/protocol"
)

type ChatClient struct {
	conn      net.Conn
	cmdReader *protocol.CommandReader
	cmdWriter *protocol.CommandWriter
	Incoming  chan string
}

func NewClient() *ChatClient {
	return &ChatClient{
		Incoming: make(chan string),
	}
}

func (c *ChatClient) Dial(address string) error {
	conn, err := net.Dial("tcp", address)

	if err == nil {
		c.conn = conn
	}

	c.cmdReader = protocol.NewCommandReader(conn)
	c.cmdWriter = protocol.NewCommandWriter(conn)

	return err
}

func (c *ChatClient) Start() {
	for {
		cmd, err := c.cmdReader.Read()

		if err == io.EOF {
			break
		} else if err != nil {
			log.Printf("Read error %v", err)
		}

		if cmd != nil {
			switch v := cmd.(type) {
			case protocol.MessageCommand:
				c.Incoming <- v.Message
			default:
				log.Printf("Unknown command: %v", v)
			}
		}
	}
}

func (c *ChatClient) Close() {
	c.conn.Close()
}

func (c *ChatClient) Send(message string) error {
	return c.cmdWriter.Write(protocol.SendCommand{
		Message: message,
	})
}
