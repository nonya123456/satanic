package protocol

import (
	"encoding/binary"
	"io"

	flatbuffers "github.com/google/flatbuffers/go"
	"github.com/nonya123456/satanic/protocol/data"
)

func ReadMessage(r io.Reader) (*data.MessageDataT, error) {
	lengthBuf := make([]byte, flatbuffers.SizeUint32)
	if _, err := io.ReadFull(r, lengthBuf); err != nil {
		return nil, err
	}

	length := binary.LittleEndian.Uint32(lengthBuf)
	messageBuf := make([]byte, length)
	if _, err := io.ReadFull(r, messageBuf); err != nil {
		return nil, err
	}

	message := data.GetRootAsMessage(messageBuf, 0)
	unpacked := message.UnPack()
	return unpacked.Data, nil
}

func WriteMessage(w io.Writer, msg *data.MessageDataT) error {
	wrapped := &data.MessageT{Data: msg}

	builder := flatbuffers.NewBuilder(0)
	msgOffset := wrapped.Pack(builder)
	builder.FinishSizePrefixed(msgOffset)

	message := builder.FinishedBytes()
	if _, err := w.Write(message); err != nil {
		return err
	}

	return nil
}
