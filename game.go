package main

import "github.com/hajimehoshi/ebiten/v2"

type Game struct {
	player *Player
	world  *World
}

func NewGame(p *Player) (*Game, error) {
	world := &World{}
	return &Game{
		player: p,
		world:  world,
	}, nil
}

func (g *Game) Update() (GameState, error) {
	return nil, nil
}

func (g *Game) Draw(screen *ebiten.Image) {

}
