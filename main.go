package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct{}

func NewGame() *Game {
	return &Game{}
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 800, 600
}

func main() {
	game := NewGame()
	ebiten.SetWindowSize(800, 600)
	ebiten.SetWindowTitle("Multiplayer Dungeon Game")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
