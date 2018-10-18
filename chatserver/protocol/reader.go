package protocol

import (
	"bufio"
	"io"
	"log"
)

type CommandReader struct {
	reader *bufio.Reader
}

func NewCommandReader(reader io.Reader) *CommandReader {
	return &CommandReader{
		reader: bufio.NewReader(reader),
	}
}

func (r *CommandReader) Read() (interface{}, error) {
	// Read the first part
	commandName, err := r.reader.ReadString(' ')

	if err != nil {
		return nil, err
	}

	switch commandName {
	case "MESSAGE ":
		user, err := r.reader.ReadString(' ')

		if err != nil {
			return nil, err
		}

		message, err := r.reader.ReadString('\n')

		if err != nil {
			return nil, err
		}

		return MessageCommand{
			user[:len(user)-1],
			message[:len(message)-1],
		}, nil

	case "SEND ":
		message, err := r.reader.ReadString('\n')

		if err != nil {
			return nil, err
		}

		return SendCommand{message[:len(message)-1]}, nil

	case "NAME ":
		name, err := r.reader.ReadString('\n')

		if err != nil {
			return nil, err
		}

		return NameCommand{name[:len(name)-1]}, nil

	default:
		log.Printf("Unknown command: %v", commandName)
	}

	return nil, UnknownCommand
}

func (r *CommandReader) ReadAll() ([]interface{}, error) {
	commands := []interface{}{}

	for {
		command, err := r.Read()

		if command != nil {
			commands = append(commands, command)
		}

		if err == io.EOF {
			break
		} else if err != nil {
			return commands, err
		}
	}

	return commands, nil
}
