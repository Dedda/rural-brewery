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
	backgrounds_png embed.FS
	BackgroundHome  *ebiten.Image
)

func init() {
	data, err := backgrounds_png.ReadFile("world/background/home.png")
	if err != nil {
		log.Fatal(err)
	}
	img, _, err := image.Decode(bytes.NewReader(data))
	if err != nil {
		log.Fatal(err)
	}
	BackgroundHome = ebiten.NewImageFromImage(img)
}
