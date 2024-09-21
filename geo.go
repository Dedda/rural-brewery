package main

import (
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type Point2F struct {
	x, y float64
}

func ScreenCenter() Point2F {
	return Point2F{
		x: gameWidth / 2,
		y: gameHeight / 2,
	}
}

func MaxWidthOfTexts(texts []string, face *text.GoTextFace) float64 {
	max := 0.0
	for _, t := range texts {
		w, _ := text.Measure(t, face, 0)
		if w > max {
			max = w
		}
	}
	return max
}

func CenterWidthHorizontally(width float64, containerWidth float64) float64 {
	return containerWidth/2 - width/2
}

func CenterTextHorizontally(msg *string, face *text.GoTextFace, containerWidth float64) float64 {
	w, _ := text.Measure(*msg, face, 0)
	return containerWidth/2 - w/2
}
