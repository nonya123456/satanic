package main

import (
	"time"
)

func RunHeadlessGame(game *Game, tickRate int) error {
	ticker := time.NewTicker(time.Second / time.Duration(tickRate))
	defer ticker.Stop()
	for range ticker.C {
		if err := game.Update(); err != nil {
			return err
		}
	}
	return nil
}
