// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package data

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type PongMessageDataT struct {
}

func (t *PongMessageDataT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil {
		return 0
	}
	PongMessageDataStart(builder)
	return PongMessageDataEnd(builder)
}

func (rcv *PongMessageData) UnPackTo(t *PongMessageDataT) {
}

func (rcv *PongMessageData) UnPack() *PongMessageDataT {
	if rcv == nil {
		return nil
	}
	t := &PongMessageDataT{}
	rcv.UnPackTo(t)
	return t
}

type PongMessageData struct {
	_tab flatbuffers.Table
}

func GetRootAsPongMessageData(buf []byte, offset flatbuffers.UOffsetT) *PongMessageData {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &PongMessageData{}
	x.Init(buf, n+offset)
	return x
}

func FinishPongMessageDataBuffer(builder *flatbuffers.Builder, offset flatbuffers.UOffsetT) {
	builder.Finish(offset)
}

func GetSizePrefixedRootAsPongMessageData(buf []byte, offset flatbuffers.UOffsetT) *PongMessageData {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &PongMessageData{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func FinishSizePrefixedPongMessageDataBuffer(builder *flatbuffers.Builder, offset flatbuffers.UOffsetT) {
	builder.FinishSizePrefixed(offset)
}

func (rcv *PongMessageData) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *PongMessageData) Table() flatbuffers.Table {
	return rcv._tab
}

func PongMessageDataStart(builder *flatbuffers.Builder) {
	builder.StartObject(0)
}
func PongMessageDataEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
