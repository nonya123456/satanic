// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package data

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type MessageT struct {
	Data *MessageDataT `json:"data"`
}

func (t *MessageT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil {
		return 0
	}
	dataOffset := t.Data.Pack(builder)

	MessageStart(builder)
	if t.Data != nil {
		MessageAddDataType(builder, t.Data.Type)
	}
	MessageAddData(builder, dataOffset)
	return MessageEnd(builder)
}

func (rcv *Message) UnPackTo(t *MessageT) {
	dataTable := flatbuffers.Table{}
	if rcv.Data(&dataTable) {
		t.Data = rcv.DataType().UnPack(dataTable)
	}
}

func (rcv *Message) UnPack() *MessageT {
	if rcv == nil {
		return nil
	}
	t := &MessageT{}
	rcv.UnPackTo(t)
	return t
}

type Message struct {
	_tab flatbuffers.Table
}

func GetRootAsMessage(buf []byte, offset flatbuffers.UOffsetT) *Message {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Message{}
	x.Init(buf, n+offset)
	return x
}

func FinishMessageBuffer(builder *flatbuffers.Builder, offset flatbuffers.UOffsetT) {
	builder.Finish(offset)
}

func GetSizePrefixedRootAsMessage(buf []byte, offset flatbuffers.UOffsetT) *Message {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &Message{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func FinishSizePrefixedMessageBuffer(builder *flatbuffers.Builder, offset flatbuffers.UOffsetT) {
	builder.FinishSizePrefixed(offset)
}

func (rcv *Message) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Message) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Message) DataType() MessageData {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return MessageData(rcv._tab.GetByte(o + rcv._tab.Pos))
	}
	return 0
}

func (rcv *Message) MutateDataType(n MessageData) bool {
	return rcv._tab.MutateByteSlot(4, byte(n))
}

func (rcv *Message) Data(obj *flatbuffers.Table) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		rcv._tab.Union(obj, o)
		return true
	}
	return false
}

func MessageStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func MessageAddDataType(builder *flatbuffers.Builder, dataType MessageData) {
	builder.PrependByteSlot(0, byte(dataType), 0)
}
func MessageAddData(builder *flatbuffers.Builder, data flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(data), 0)
}
func MessageEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
