// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package data

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type PingMessageDataT struct {
}

func (t *PingMessageDataT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil {
		return 0
	}
	PingMessageDataStart(builder)
	return PingMessageDataEnd(builder)
}

func (rcv *PingMessageData) UnPackTo(t *PingMessageDataT) {
}

func (rcv *PingMessageData) UnPack() *PingMessageDataT {
	if rcv == nil {
		return nil
	}
	t := &PingMessageDataT{}
	rcv.UnPackTo(t)
	return t
}

type PingMessageData struct {
	_tab flatbuffers.Table
}

func GetRootAsPingMessageData(buf []byte, offset flatbuffers.UOffsetT) *PingMessageData {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &PingMessageData{}
	x.Init(buf, n+offset)
	return x
}

func FinishPingMessageDataBuffer(builder *flatbuffers.Builder, offset flatbuffers.UOffsetT) {
	builder.Finish(offset)
}

func GetSizePrefixedRootAsPingMessageData(buf []byte, offset flatbuffers.UOffsetT) *PingMessageData {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &PingMessageData{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func FinishSizePrefixedPingMessageDataBuffer(builder *flatbuffers.Builder, offset flatbuffers.UOffsetT) {
	builder.FinishSizePrefixed(offset)
}

func (rcv *PingMessageData) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *PingMessageData) Table() flatbuffers.Table {
	return rcv._tab
}

func PingMessageDataStart(builder *flatbuffers.Builder) {
	builder.StartObject(0)
}
func PingMessageDataEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
