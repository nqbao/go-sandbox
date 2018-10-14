package client

import "github.com/nqbao/learn-go/chatserver/protocol"

type messageHandler func(string)

type ChatClient interface {
	Dial(address string) error
	Start()
	Close()
	SetName(name string) error
	Send(message string) error
	Incoming() chan protocol.MessageCommand
}
