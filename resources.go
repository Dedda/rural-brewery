package main

import (
	"bytes"
	_ "embed"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"log"
)

var (
	//go:embed assets/JacquardaBastarda9.ttf
	JacquardaBastarda9_ttf       []byte
	JacquardaBastarda9FaceSource *text.GoTextFaceSource
)

func init() {
	s, err := text.NewGoTextFaceSource(bytes.NewReader(JacquardaBastarda9_ttf))
	if err != nil {
		log.Fatal(err)
	}
	JacquardaBastarda9FaceSource = s
}
