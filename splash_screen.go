package main

import (
	"github.com/Dedda/rural-brewery/assets"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"image/color"
)

const (
	splashScreenTime = 2 * 60
)

type SplashScreen struct {
	remainingTime int
}

func NewSplashScreen() (*SplashScreen, error) {
	return &SplashScreen{
		remainingTime: splashScreenTime,
	}, nil
}

func (s *SplashScreen) Update() (GameState, error) {
	s.remainingTime--
	if s.remainingTime == 0 {
		return NewMainMenu()
	}
	return nil, nil
}

func (s *SplashScreen) Draw(screen *ebiten.Image) {
	yStep := float64(gameHeight-100) / splashScreenTime
	y := 100 + yStep*float64(s.remainingTime)
	headerText := "Rural Brewery"
	headerFace := &text.GoTextFace{
		Source: assets.JacquardaBastarda9FaceSource,
		Size:   mainMenuHeaderSize,
	}
	op := &text.DrawOptions{}
	op.GeoM.Translate(CenterTextHorizontally(&headerText, headerFace, gameWidth), y)
	op.ColorScale.ScaleWithColor(color.White)
	text.Draw(screen, headerText, headerFace, op)
}
