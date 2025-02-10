package main

import (
	"log"
	"net"

	"github.com/nonya123456/satanic/protocol"
	"github.com/nonya123456/satanic/protocol/data"
	"github.com/xtaci/kcp-go/v5"
)

func handleConnection(conn net.Conn) {
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

func main() {
	listener, err := kcp.ListenWithOptions("127.0.0.1:12345", nil, 10, 3)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	for {
		conn, err := listener.AcceptKCP()
		if err != nil {
			log.Fatal(err)
			continue
		}
		go handleConnection(conn)
	}
}
