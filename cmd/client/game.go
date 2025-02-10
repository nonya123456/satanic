package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/nonya123456/satanic/protocol/data"
	"github.com/yohamta/donburi"
)

type Game struct {
	world donburi.World
	sess  *Session
}

func NewGame() *Game {
	sess, err := NewSession("127.0.0.1:12345")
	if err != nil {
		log.Fatal(err)
	}

	go sess.Listen()

	_ = sess.Write(&data.MessageDataT{
		Type:  data.MessageDataPingMessageData,
		Value: &data.PingMessageDataT{},
	})

	world := donburi.NewWorld()
	return &Game{world: world, sess: sess}
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 800, 600
}

func (g *Game) Close() {
	_ = g.sess.Close()
}
