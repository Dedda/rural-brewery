package main

import "github.com/hajimehoshi/ebiten/v2"

type MainMenu struct {
}

func NewMainMenu() (*MainMenu, error) {
	return &MainMenu{}, nil
}

func (m *MainMenu) Update() (GameState, error) {
	return nil, nil
}

func (m *MainMenu) Draw(screen *ebiten.Image) {

}
