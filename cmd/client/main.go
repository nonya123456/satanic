package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	game := NewGame()
	defer game.Close()

	ebiten.SetWindowSize(800, 600)
	ebiten.SetWindowTitle("Client Game")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
