package main

import (
	"bytes"
	_ "embed"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"image"
	"log"
)

var (
	//go:embed assets/JacquardaBastarda9.ttf
	JacquardaBastarda9_ttf       []byte
	JacquardaBastarda9FaceSource *text.GoTextFaceSource
	//go:embed assets/player.png
	Player_png  []byte
	PlayerImage *ebiten.Image
)

func init() {
	s, err := text.NewGoTextFaceSource(bytes.NewReader(JacquardaBastarda9_ttf))
	if err != nil {
		log.Fatal(err)
	}
	JacquardaBastarda9FaceSource = s
	img, _, err := image.Decode(bytes.NewReader(Player_png))
	if err != nil {
		log.Fatal(err)
	}
	PlayerImage = ebiten.NewImageFromImage(img)
}
