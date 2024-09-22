package main

import "github.com/hajimehoshi/ebiten/v2"

type MapHome struct {
}

func NewMapHome() *MapHome {
	return &MapHome{}
}

func (m MapHome) BaseInfo() *WorldMapBaseData {
	return &WorldMapBaseData{
		id:   "home",
		name: "Home",
		size: Vec2F{
			X: 48,
			Y: 32,
		},
		teleporters: []Teleporter{},
	}
}

func (m MapHome) Update(g *Game, current bool) error {
	return nil
}

func (m MapHome) Draw(screen *ebiten.Image, g *Game) {

}
