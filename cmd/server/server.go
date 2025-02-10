package main

import (
	"log"

	"github.com/nonya123456/satanic/protocol"
	"github.com/nonya123456/satanic/protocol/data"
	"github.com/xtaci/kcp-go/v5"
)

type Server struct {
	listener *kcp.Listener
}

func NewServer(addr string) (*Server, error) {
	listener, err := kcp.ListenWithOptions("127.0.0.1:12345", nil, 10, 3)
	if err != nil {
		return nil, err
	}

	return &Server{listener: listener}, nil
}

func (s *Server) Listen() {
	for {
		conn, err := s.listener.AcceptKCP()
		if err != nil {
			log.Fatal(err)
			continue
		}
		go handleConnection(conn)
	}
}

func (s *Server) Close() error {
	return s.listener.Close()
}

func handleConnection(conn *kcp.UDPSession) {
	defer conn.Close()
	for {
		msgT, err := protocol.ReadMessage(conn)
		if err != nil {
			log.Println(err)
			return
		}
		if msgT == nil {
			log.Println("got nil message")
			continue
		}

		switch msgT.Type {
		case data.MessageDataPingMessageData:
			err := protocol.WriteMessage(conn, &data.MessageDataT{
				Type:  data.MessageDataPongMessageData,
				Value: &data.PongMessageDataT{},
			})
			if err != nil {
				log.Println(err)
				return
			}
		}
	}
}
