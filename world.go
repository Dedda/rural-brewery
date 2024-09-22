package main

import "encoding/json"

type GlobalPosition struct {
	mapId    string
	position Vec2F
}

type World struct {
	maps map[string]WorldMap
}

func LoadWorld() (*World, error) {
	var maps map[string]WorldMap
	err := json.Unmarshal(Maps_json, &maps)
	return &World{
		maps: maps,
	}, err
}

func (w *World) GetMap(id string) (*WorldMap, error) {
	return nil, nil
}

type WorldMap struct {
	Identifier  string        `json:"id"`
	Name        string        `json:"name"`
	Size        Vec2F         `json:"size"`
	Teleporters []*Teleporter `json:"teleporters"`
}

type Teleporter struct {
	TopLeft        Vec2F  `json:"topLeft"`
	BottomRight    Vec2F  `json:"bottomRight"`
	DestinationMap string `json:"destinationMap"`
	Destination    Vec2F  `json:"destination"`
}
