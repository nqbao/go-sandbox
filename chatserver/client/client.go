package client

import (
	"net"
)

type ChatClient struct {
	conn net.Conn
}

func NewClient() *ChatClient {
	return &ChatClient{}
}

func (c *ChatClient) Dial(address string) error {
	conn, err := net.Dial("tcp", address)

	if err == nil {
		c.conn = conn
	}

	return err
}

func (c *ChatClient) Close() {
	c.conn.Close()
}

func (c *ChatClient) Send(message string) {
	c.conn.Write([]byte(message))
}
