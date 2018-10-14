package client

import (
	"io"
	"log"
	"net"

	"github.com/nqbao/learn-go/chatserver/protocol"
)

type TcpChatClient struct {
	conn      net.Conn
	cmdReader *protocol.CommandReader
	cmdWriter *protocol.CommandWriter
	name      string
	incoming  chan protocol.MessageCommand
}

func NewClient() *TcpChatClient {
	return &TcpChatClient{
		incoming: make(chan protocol.MessageCommand),
	}
}

func (c *TcpChatClient) Dial(address string) error {
	conn, err := net.Dial("tcp", address)

	if err == nil {
		c.conn = conn
	}

	c.cmdReader = protocol.NewCommandReader(conn)
	c.cmdWriter = protocol.NewCommandWriter(conn)

	return err
}

func (c *TcpChatClient) Start() {
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
				c.incoming <- v
			default:
				log.Printf("Unknown command: %v", v)
			}
		}
	}
}

func (c *TcpChatClient) Close() {
	c.conn.Close()
}

func (c *TcpChatClient) Incoming() chan protocol.MessageCommand {
	return c.incoming
}

func (c *TcpChatClient) SetName(name string) error {
	return c.cmdWriter.Write(protocol.NameCommand{name})
}

func (c *TcpChatClient) Send(message string) error {
	return c.cmdWriter.Write(protocol.SendCommand{
		Message: message,
	})
}
