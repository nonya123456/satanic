package protocol_test

import (
	"bytes"
	"testing"

	"github.com/nonya123456/satanic/protocol"
	"github.com/nonya123456/satanic/protocol/data"
	"github.com/stretchr/testify/assert"
)

func TestReadWriteMessage(t *testing.T) {
	tests := []struct {
		name        string
		messageData *data.MessageDataT
	}{
		{
			name:        "GameStateUpdated message",
			messageData: createMessageDataT(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buffer := &bytes.Buffer{}
			err := protocol.WriteMessage(buffer, tt.messageData)
			assert.NoError(t, err, "Error writing message")
			readMessage, err := protocol.ReadMessage(buffer)
			assert.NoError(t, err, "Error reading message")
			assert.Equal(t, tt.messageData, readMessage)
		})
	}
}

func createMessageDataT() *data.MessageDataT {
	return &data.MessageDataT{
		Type:  data.MessageDataPingMessageData,
		Value: &data.PingMessageDataT{},
	}
}
