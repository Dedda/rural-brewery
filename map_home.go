package main

import (
	"github.com/ByteArena/box2d"
	"github.com/Dedda/rural-brewery/assets"
	"github.com/hajimehoshi/ebiten/v2"
)

var (
	homeData = WorldMapBaseData{
		id:   "home",
		name: "Home",
		size: Vec2F{
			X: 48,
			Y: 32,
		},
		teleporters: []Teleporter{},
	}
)

type MapHome struct {
}

func NewMapHome() *MapHome {
	return &MapHome{}
}

func (m *MapHome) Enter(g *Game, b2dWorld *box2d.B2World) {

}

func (m *MapHome) Leave(g *Game, b2dWorld *box2d.B2World) {

}

func (m MapHome) BaseInfo() *WorldMapBaseData {
	return &homeData
}

func (m MapHome) Update(g *Game, current bool) error {
	return nil
}

func (m MapHome) Draw(screen *ebiten.Image, g *Game) {
	op := &ebiten.DrawImageOptions{}
	screen.DrawImage(assets.BackgroundHome, op)
}
