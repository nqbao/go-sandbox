package protocol_test

import (
	"reflect"
	"strings"
	"testing"

	"github.com/nqbao/learn-go/chatserver/protocol"
)

func TestCommandReader(t *testing.T) {
	tests := []struct {
		input   string
		results []interface{}
	}{
		{
			"SEND test\n",
			[]interface{}{
				protocol.SendCommand{"test"},
			},
		},
		{
			"MESSAGE user1 hello\nMESSAGE user2 world\n",
			[]interface{}{
				protocol.MessageCommand{"user1", "hello"},
				protocol.MessageCommand{"user2", "world"},
			},
		},
	}

	for _, test := range tests {
		reader := protocol.NewCommandReader(strings.NewReader(test.input))
		results, err := reader.ReadAll()

		t.Log(results)

		if err != nil {
			t.Errorf("Unable to read command, error %v", err)
		} else if !reflect.DeepEqual(results, test.results) {
			t.Errorf("Command output is not the same: %v %v", results, test.results)
		}
	}
}
