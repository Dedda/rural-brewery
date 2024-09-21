package main

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	splashScreenTime = 5 * 60
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
		fmt.Println("Entering main menu...")
		return NewMainMenu()
	}
	return nil, nil
}

func (s *SplashScreen) Draw(screen *ebiten.Image) {

}
