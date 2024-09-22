package main

import (
	"github.com/ByteArena/box2d"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"math"
)

const (
	GridCellSize = 16
)

type Vec2F struct {
	X float64
	Y float64
}

func (v Vec2F) AddVec2F(o Vec2F) Vec2F {
	return Vec2F{v.X + o.X, v.Y + o.Y}
}

func (v Vec2F) MulF64(f float64) Vec2F {
	return Vec2F{v.X * f, v.Y * f}
}

func (v Vec2F) Length() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v Vec2F) Normalize() Vec2F {
	length := v.Length()
	return Vec2F{v.X / length, v.Y / length}
}

func (v Vec2F) ToB2D() box2d.B2Vec2 {
	return box2d.B2Vec2{
		X: v.X,
		Y: v.Y,
	}
}

func Vec2FFromB2Vec2(b box2d.B2Vec2) Vec2F {
	return Vec2F{
		X: b.X,
		Y: b.Y,
	}
}

func GridCellOffset(x, y int) Vec2F {
	return Vec2F{
		X: float64(x) * GridCellSize,
		Y: float64(y) * GridCellSize,
	}
}

func ScreenCenter() Vec2F {
	return Vec2F{
		X: gameWidth / 2,
		Y: gameHeight / 2,
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
