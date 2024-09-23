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
	boundaries []*box2d.B2Body
}

func NewMapHome() *MapHome {
	return &MapHome{
		boundaries: []*box2d.B2Body{},
	}
}

func (m *MapHome) Enter(g *Game, b2dWorld *box2d.B2World) {
	m.createBoundaries(b2dWorld)
}

func (m *MapHome) Leave(g *Game, b2dWorld *box2d.B2World) {
	for _, boundary := range m.boundaries {
		b2dWorld.DestroyBody(boundary)
	}
	m.boundaries = []*box2d.B2Body{}
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

func (m *MapHome) createBoundaries(world *box2d.B2World) {
	m.boundaries = createMapBoundaries(world, homeData.size.MulF64(GridCellSize))
}
