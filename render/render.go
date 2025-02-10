package render

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/nonya123456/satanic/core"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/filter"
)

type RendererData struct {
}

var Renderer = donburi.NewComponentType[RendererData]()

func Render(w donburi.World, image *ebiten.Image) {
	query := donburi.NewQuery(filter.Contains(Renderer, core.Position))
	for entry := range query.Iter(w) {
		_ = Renderer.Get(entry)
		pos := core.Position.Get(entry)
		vector.DrawFilledCircle(image, pos.X, pos.Y, 10, color.Black, false)
	}
}
