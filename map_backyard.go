package main

import (
	"github.com/ByteArena/box2d"
	"github.com/hajimehoshi/ebiten/v2"
)

var (
	backyardData = WorldMapBaseData{
		id:   "backyard",
		name: "Backyard",
		size: Vec2F{
			X: 128,
			Y: 128,
		},
		teleporters: []Teleporter{},
	}
)

type MapBackyard struct {
}

func NewMapBackyard() *MapBackyard {
	return &MapBackyard{}
}

func (m *MapBackyard) Enter(g *Game, b2dWorld *box2d.B2World) {

}

func (m *MapBackyard) Leave(g *Game, b2dWorld *box2d.B2World) {

}

func (m *MapBackyard) BaseInfo() *WorldMapBaseData {
	return &backyardData
}

func (m *MapBackyard) Update(g *Game, current bool) error {
	return nil
}

func (m *MapBackyard) Draw(screen *ebiten.Image, g *Game) {}
