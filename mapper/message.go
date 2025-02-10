package mapper

import (
	"errors"

	"github.com/nonya123456/satanic/core"
	"github.com/nonya123456/satanic/protocol/data"
)

func FromGameStateData(s core.GameStateData) *data.MessageDataT {
	return &data.MessageDataT{
		Type: data.MessageDataGameStateMessageData,
		Value: &data.GameStateMessageDataT{
			Position: &data.PositionT{
				X: s.Position.X,
				Y: s.Position.Y,
			},
		},
	}
}

func ToGameStateData(m *data.MessageDataT) (core.GameStateData, error) {
	gameStateT, ok := m.Value.(*data.GameStateMessageDataT)
	if !ok {
		return core.GameStateData{}, errors.New("invalid message data type")
	}

	return core.GameStateData{
		Position: core.PositionData{
			X: gameStateT.Position.X,
			Y: gameStateT.Position.Y,
		},
	}, nil
}
