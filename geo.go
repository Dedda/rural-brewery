package main

import (
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"math"
)

type Vec2F struct {
	x, y float64
}

func (v Vec2F) AddVec2F(o Vec2F) Vec2F {
	return Vec2F{v.x + o.x, v.y + o.y}
}

func (v Vec2F) MulF64(f float64) Vec2F {
	return Vec2F{v.x * f, v.y * f}
}

func (v Vec2F) Length() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}

func (v Vec2F) Normalize() Vec2F {
	length := v.Length()
	return Vec2F{v.x / length, v.y / length}
}

func ScreenCenter() Vec2F {
	return Vec2F{
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
