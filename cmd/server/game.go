package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
)

type Game struct {
	world  donburi.World
	server *Server
}

func NewGame() *Game {
	server, err := NewServer("127.0.0.1:12345")
	if err != nil {
		log.Fatal(err)
	}

	go server.Listen()

	world := donburi.NewWorld()
	return &Game{world: world, server: server}
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
	_ = g.server.Close()
}
