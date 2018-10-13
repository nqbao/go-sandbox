package protocol_test

import (
	"bytes"
	"testing"

	"github.com/nqbao/learn-go/chatserver/protocol"
)

func TestWriteCommand(t *testing.T) {
	tests := []struct {
		commands []interface{}
		result   string
	}{
		{
			[]interface{}{
				protocol.SendCommand{"Hello"},
			},
			"SEND Hello\n",
		},
	}

	buf := new(bytes.Buffer)
	for _, test := range tests {
		buf.Reset()
		cmdWriter := protocol.NewCommandWriter(buf)

		for _, cmd := range test.commands {
			if cmdWriter.Write(cmd) != nil {
				t.Errorf("Unable to write command %v", cmd)
			}
		}

		if buf.String() != test.result {
			t.Errorf("Command output is not the same: %v %v", buf.String(), test.result)
		}
	}
}
