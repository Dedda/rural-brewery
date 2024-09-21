package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"log"
)

const (
	screenWidth  = 1280
	screenHeight = 960
)

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("🍺 Rural Brewery")
	game, _ := NewGameWrapper()
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
