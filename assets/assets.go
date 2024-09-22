package assets

import (
	"bytes"
	_ "embed"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"image"
	"log"
)

var (
	//go:embed JacquardaBastarda9.ttf
	jacquardaBastarda9_ttf       []byte
	JacquardaBastarda9FaceSource *text.GoTextFaceSource
	//go:embed player.png
	player_png  []byte
	PlayerImage *ebiten.Image
)

func init() {
	s, err := text.NewGoTextFaceSource(bytes.NewReader(jacquardaBastarda9_ttf))
	if err != nil {
		log.Fatal(err)
	}
	JacquardaBastarda9FaceSource = s
	img, _, err := image.Decode(bytes.NewReader(player_png))
	if err != nil {
		log.Fatal(err)
	}
	PlayerImage = ebiten.NewImageFromImage(img)
}
