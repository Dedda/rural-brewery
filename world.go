package main

import "github.com/hajimehoshi/ebiten/v2"

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
	return &World{
		maps: maps,
	}, nil
}

func (w *World) GetMap(id string) (WorldMap, error) {
	worldMap := w.maps[id]
	return worldMap, nil
}

type WorldMap interface {
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
