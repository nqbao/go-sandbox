package client

import (
	"bufio"
	"io"
	"log"
	"net"
	"strings"
)

type ChatClient struct {
	conn     net.Conn
	Incoming chan string
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

	return err
}

func (c *ChatClient) Start() {
	reader := bufio.NewReader(c.conn)

	for {
		msg, err := reader.ReadString('\n')

		if err == io.EOF {
			break
		} else if err != nil {
			log.Printf("Read error %v", err)
		}

		c.Incoming <- strings.TrimSpace(msg)
	}
}

func (c *ChatClient) Close() {
	c.conn.Close()
}

func (c *ChatClient) Send(message string) {
	c.conn.Write([]byte(message))
}
