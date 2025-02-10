package main

import (
	"log"
)

func main() {
	game := NewGame()
	defer game.Close()

	log.Println("Running headless server")
	if err := RunHeadlessGame(game, 60); err != nil {
		log.Fatal(err)
	}
}
