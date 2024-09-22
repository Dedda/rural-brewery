package assets

import (
	"bytes"
	"embed"
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	"log"
)

var (
	//go:embed world/background/*
	backgrounds_png   embed.FS
	BackgroundHome    *ebiten.Image
	BackgroundBrewery *ebiten.Image
)

func init() {
	BackgroundHome = loadImage("home.png")
	BackgroundBrewery = loadImage("brewery.png")
}

func loadImage(name string) *ebiten.Image {
	data, err := backgrounds_png.ReadFile("world/background/" + name)
	if err != nil {
		log.Fatal(err)
	}
	img, _, err := image.Decode(bytes.NewReader(data))
	if err != nil {
		log.Fatal(err)
	}
	return ebiten.NewImageFromImage(img)
}
