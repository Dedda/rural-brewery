package main

import (
	"github.com/ByteArena/box2d"
	"github.com/hajimehoshi/ebiten/v2"
)

type GlobalPosition struct {
	mapId    string
	position Vec2F
}

type World struct {
	maps map[string]WorldMap
}

func LoadWorld() (*World, error) {
	maps := make(map[string]WorldMap)
	maps["home"] = NewMapHome()
	maps["backyard"] = NewMapBackyard()
	maps["brewery"] = NewMapBrewery()
	return &World{
		maps: maps,
	}, nil
}

func (w *World) GetMap(id string) (WorldMap, error) {
	worldMap := w.maps[id]
	return worldMap, nil
}

type WorldMap interface {
	Enter(g *Game, b2dWorld *box2d.B2World)
	Leave(g *Game, b2dWorld *box2d.B2World)
	BaseInfo() *WorldMapBaseData
	Update(g *Game, current bool) error
	Draw(screen *ebiten.Image, g *Game)
}

type Teleporter struct {
	TopLeft        Vec2F
	BottomRight    Vec2F
	DestinationMap string
	Destination    Vec2F
}

type WorldMapBaseData struct {
	id          string
	name        string
	size        Vec2F
	teleporters []Teleporter
}
