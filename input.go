package main

import "github.com/hajimehoshi/ebiten/v2"

func GetPlayerMovementInput() Vec2F {
	movement := Vec2F{}
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		movement.y -= 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		movement.y += 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		movement.x -= 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		movement.x += 1
	}
	if movement.Length() != 0 {
		return movement.Normalize()
	}
	return movement
}
