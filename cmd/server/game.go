package main

import (
	"log"

	"github.com/nonya123456/satanic/core"
	"github.com/nonya123456/satanic/protocol/data"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/filter"
)

type Game struct {
	world     donburi.World
	server    *Server
	hasObject bool
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
	if !g.hasObject {
		g.addObject()
		g.hasObject = true
	}

	core.UpdatePosition(g.world)
	core.UpdateBouncing(g.world)
	g.broadcastPosition()
	return nil
}

func (g *Game) Close() {
	_ = g.server.Close()
}

func (g *Game) addObject() {
	entity := g.world.Create(core.Position, core.Velocity)
	entry := g.world.Entry(entity)
	core.Position.SetValue(entry, core.PositionData{X: 10, Y: 20})
	core.Velocity.SetValue(entry, core.VelocityData{X: 2, Y: 2})
}

func (g *Game) broadcastPosition() {
	query := donburi.NewQuery(filter.Contains(core.Position))
	entry, ok := query.First(g.world)
	if !ok {
		return
	}

	pos := core.Position.Get(entry)
	g.server.Broadcast(&data.MessageDataT{
		Type: data.MessageDataGameStateMessageData,
		Value: &data.GameStateMessageDataT{
			Position: &data.PositionT{X: pos.X, Y: pos.Y},
		},
	})
}
