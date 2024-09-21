package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	gameWidth  = 1280
	gameHeight = 960
)

type GameState interface {
	Update() (GameState, error)
	Draw(screen *ebiten.Image)
}

type GameWrapper struct {
	state GameState
}

func NewGameWrapper() (*GameWrapper, error) {
	splash, err := NewSplashScreen()
	return &GameWrapper{
		state: splash,
	}, err
}

func (g *GameWrapper) Update() error {
	newState, err := g.state.Update()
	if newState != nil {
		g.state = newState
	}
	return err
}

func (g *GameWrapper) Draw(screen *ebiten.Image) {
	g.state.Draw(screen)
}

func (g *GameWrapper) Layout(_, _ int) (screenWidth, screenHeight int) {
	return gameWidth, gameHeight
}
