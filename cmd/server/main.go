package main

import (
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	game := NewGame()
	defer game.Close()

	if os.Getenv("HEADLESS") != "" {
		log.Println("Running headless server")
		if err := RunHeadlessGame(game, 60); err != nil {
			log.Fatal(err)
		}
	} else {
		log.Println("Running server")
		ebiten.SetWindowSize(800, 600)
		ebiten.SetWindowTitle("Server Game")
		if err := ebiten.RunGame(game); err != nil {
			log.Fatal(err)
		}
	}
}
