package main

import (
	"github.com/ByteArena/box2d"
	"github.com/Dedda/rural-brewery/assets"
	"github.com/hajimehoshi/ebiten/v2"
)

var (
	breweryData = WorldMapBaseData{
		id:   "brewery",
		name: "Brewery",
		size: Vec2F{
			X: 64,
			Y: 64,
		},
		teleporters: []Teleporter{},
	}
)

type MapBrewery struct {
}

func NewMapBrewery() *MapBrewery {
	return &MapBrewery{}
}

func (m *MapBrewery) Enter(g *Game, b2dWorld *box2d.B2World) {

}

func (m *MapBrewery) Leave(g *Game, b2dWorld *box2d.B2World) {

}

func (m *MapBrewery) BaseInfo() *WorldMapBaseData {
	return &breweryData
}

func (m *MapBrewery) Update(g *Game, current bool) error {
	return nil
}

func (m *MapBrewery) Draw(screen *ebiten.Image, g *Game) {
	op := &ebiten.DrawImageOptions{}
	screen.DrawImage(assets.BackgroundBrewery, op)
}
