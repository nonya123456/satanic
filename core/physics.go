package core

import (
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/filter"
)

type PositionData struct {
	X, Y float32
}

type VelocityData struct {
	X, Y float32
}

var Position = donburi.NewComponentType[PositionData]()
var Velocity = donburi.NewComponentType[VelocityData]()

func UpdatePosition(world donburi.World) {
	query := donburi.NewQuery(filter.Contains(Position, Velocity))
	for entry := range query.Iter(world) {
		position := Position.Get(entry)
		velocity := Velocity.Get(entry)
		position.X += velocity.X
		position.Y += velocity.Y
	}
}

func UpdateBouncing(world donburi.World) {
	query := donburi.NewQuery(filter.Contains(Position, Velocity))
	for entry := range query.Iter(world) {
		position := Position.Get(entry)
		velocity := Velocity.Get(entry)

		newVelocity := *velocity
		if position.X < 0 || position.X > 800 {
			newVelocity.X = -velocity.X
		}
		if position.Y < 0 || position.Y > 600 {
			newVelocity.Y = -velocity.Y
		}
		Velocity.SetValue(entry, newVelocity)
	}
}
