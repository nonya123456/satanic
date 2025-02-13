// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package data

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type GameStateMessageDataT struct {
	Position *PositionT `json:"position"`
}

func (t *GameStateMessageDataT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil {
		return 0
	}
	positionOffset := t.Position.Pack(builder)
	GameStateMessageDataStart(builder)
	GameStateMessageDataAddPosition(builder, positionOffset)
	return GameStateMessageDataEnd(builder)
}

func (rcv *GameStateMessageData) UnPackTo(t *GameStateMessageDataT) {
	t.Position = rcv.Position(nil).UnPack()
}

func (rcv *GameStateMessageData) UnPack() *GameStateMessageDataT {
	if rcv == nil {
		return nil
	}
	t := &GameStateMessageDataT{}
	rcv.UnPackTo(t)
	return t
}

type GameStateMessageData struct {
	_tab flatbuffers.Table
}

func GetRootAsGameStateMessageData(buf []byte, offset flatbuffers.UOffsetT) *GameStateMessageData {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &GameStateMessageData{}
	x.Init(buf, n+offset)
	return x
}

func FinishGameStateMessageDataBuffer(builder *flatbuffers.Builder, offset flatbuffers.UOffsetT) {
	builder.Finish(offset)
}

func GetSizePrefixedRootAsGameStateMessageData(buf []byte, offset flatbuffers.UOffsetT) *GameStateMessageData {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &GameStateMessageData{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func FinishSizePrefixedGameStateMessageDataBuffer(builder *flatbuffers.Builder, offset flatbuffers.UOffsetT) {
	builder.FinishSizePrefixed(offset)
}

func (rcv *GameStateMessageData) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *GameStateMessageData) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *GameStateMessageData) Position(obj *Position) *Position {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(Position)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func GameStateMessageDataStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func GameStateMessageDataAddPosition(builder *flatbuffers.Builder, position flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(position), 0)
}
func GameStateMessageDataEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
