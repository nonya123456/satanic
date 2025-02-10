package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/nonya123456/satanic/core"
	"github.com/nonya123456/satanic/protocol/data"
	"github.com/nonya123456/satanic/render"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/filter"
)

type Game struct {
	world     donburi.World
	sess      *Session
	hasObject bool
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

func (g *Game) addObject() {
	entity := g.world.Create(core.Position, core.Velocity, render.Renderer)
	entry := g.world.Entry(entity)
	core.Position.SetValue(entry, core.PositionData{X: 10, Y: 20})
	core.Velocity.SetValue(entry, core.VelocityData{X: 1, Y: 2})
	g.hasObject = true
}

func (g *Game) Update() error {
	if !g.hasObject {
		g.addObject()
	}
	g.updateServerPosition()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{245, 245, 245, 255})
	render.Render(g.world, screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 800, 600
}

func (g *Game) Close() {
	_ = g.sess.Close()
}

func (g *Game) updateServerPosition() {
	state, ok := getLatestServerState(g.sess.Channel())
	if !ok {
		return
	}

	query := donburi.NewQuery(filter.Contains(core.Position))
	entry, ok := query.First(g.world)
	if !ok {
		return
	}

	core.Position.SetValue(entry, state.Position)
}

func getLatestServerState(ch <-chan core.GameStateData) (core.GameStateData, bool) {
	var state core.GameStateData
	var ok = false
	for {
		select {
		case state = <-ch:
			ok = true
		default:
			return state, ok
		}
	}
}
