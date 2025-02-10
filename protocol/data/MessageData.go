// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package data

import (
	flatbuffers "github.com/google/flatbuffers/go"
	"strconv"
)

type MessageData byte

const (
	MessageDataNONE                 MessageData = 0
	MessageDataPingMessageData      MessageData = 1
	MessageDataPongMessageData      MessageData = 2
	MessageDataGameStateMessageData MessageData = 3
)

var EnumNamesMessageData = map[MessageData]string{
	MessageDataNONE:                 "NONE",
	MessageDataPingMessageData:      "PingMessageData",
	MessageDataPongMessageData:      "PongMessageData",
	MessageDataGameStateMessageData: "GameStateMessageData",
}

var EnumValuesMessageData = map[string]MessageData{
	"NONE":                 MessageDataNONE,
	"PingMessageData":      MessageDataPingMessageData,
	"PongMessageData":      MessageDataPongMessageData,
	"GameStateMessageData": MessageDataGameStateMessageData,
}

func (v MessageData) String() string {
	if s, ok := EnumNamesMessageData[v]; ok {
		return s
	}
	return "MessageData(" + strconv.FormatInt(int64(v), 10) + ")"
}

type MessageDataT struct {
	Type MessageData
	Value interface{}
}

func (t *MessageDataT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil {
		return 0
	}
	switch t.Type {
	case MessageDataPingMessageData:
		return t.Value.(*PingMessageDataT).Pack(builder)
	case MessageDataPongMessageData:
		return t.Value.(*PongMessageDataT).Pack(builder)
	case MessageDataGameStateMessageData:
		return t.Value.(*GameStateMessageDataT).Pack(builder)
	}
	return 0
}

func (rcv MessageData) UnPack(table flatbuffers.Table) *MessageDataT {
	switch rcv {
	case MessageDataPingMessageData:
		var x PingMessageData
		x.Init(table.Bytes, table.Pos)
		return &MessageDataT{Type: MessageDataPingMessageData, Value: x.UnPack()}
	case MessageDataPongMessageData:
		var x PongMessageData
		x.Init(table.Bytes, table.Pos)
		return &MessageDataT{Type: MessageDataPongMessageData, Value: x.UnPack()}
	case MessageDataGameStateMessageData:
		var x GameStateMessageData
		x.Init(table.Bytes, table.Pos)
		return &MessageDataT{Type: MessageDataGameStateMessageData, Value: x.UnPack()}
	}
	return nil
}
