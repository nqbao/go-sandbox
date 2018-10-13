# Simple Chat Server

A simple chat server written in Golang, with very basic features:

  * There is only a single chat room for now
  * User can connect to the chat server
  * User can set their name
  * User can send message to the chat room

## Protocol

For this excersie , I will use simple text-based message over TCP:

  * All messages are terminated with `\n`
  * To send a chat message, client will send: 
    * `SEND chat message`
    * For now, chat message can not contain new line.
  * To set client name, client will send:
    * `NAME username`
    * For now, username can not contain space
  * Server will send the following command to all clients when there are new message:
    * `MESSAGE username the actual message`

Later on I will try protobuf, or Gob to write protocol.

## References

  * https://gist.github.com/drewolson/3950226
  * https://scotch.io/bar-talk/build-a-realtime-chat-server-with-go-and-websockets
  * [tui-go](https://github.com/marcusolsson/tui-go) for chat client. tview may be a better option, but textare input is not available yet.
