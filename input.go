package main

import "github.com/hajimehoshi/ebiten/v2"

func GetPlayerMovementInput() Vec2F {
	movement := Vec2F{}
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		movement.Y -= 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		movement.Y += 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		movement.X -= 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		movement.X += 1
	}
	if movement.Length() != 0 {
		return movement.Normalize()
	}
	return movement
}
