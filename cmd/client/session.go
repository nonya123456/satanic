package main

import (
	"io"
	"log"

	"github.com/nonya123456/satanic/core"
	"github.com/nonya123456/satanic/mapper"
	"github.com/nonya123456/satanic/protocol"
	"github.com/nonya123456/satanic/protocol/data"
	"github.com/xtaci/kcp-go/v5"
)

type Session struct {
	conn *kcp.UDPSession
	ch   chan core.GameStateData
}

func NewSession(addr string) (*Session, error) {
	conn, err := kcp.DialWithOptions(addr, nil, 10, 3)
	if err != nil {
		return nil, err
	}

	return &Session{
		conn: conn,
		ch:   make(chan core.GameStateData, 100),
	}, nil
}

func (s *Session) Listen() {
	for {
		msgT, err := protocol.ReadMessage(s.conn)
		if err != nil {
			if err == io.EOF {
				log.Println("connection closed")
				return
			}
			log.Println("error reading from server:", err)
			continue
		}
		if msgT == nil {
			log.Println("got nil message")
			continue
		}

		switch msgT.Type {
		case data.MessageDataPongMessageData:
			log.Println("got pong message")
		case data.MessageDataGameStateMessageData:
			gameStateT, err := mapper.ToGameStateData(msgT)
			if err != nil {
				log.Println("error mapping game state data:", err)
			}
			select {
			case s.ch <- gameStateT:
			default:
			}
		default:
			log.Println("got unknown message")
		}
	}
}

func (s *Session) Write(msg *data.MessageDataT) error {
	return protocol.WriteMessage(s.conn, msg)
}

func (s *Session) Close() error {
	return s.conn.Close()
}

func (s *Session) Channel() <-chan core.GameStateData {
	return s.ch
}
