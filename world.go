package main

type GlobalPosition struct {
	mapId    MapIdentifier
	position LocalPosition
}

type MapIdentifier struct {
	x int
	y int
}

type LocalPosition struct {
	x float64
	y float64
}

func (p *LocalPosition) Left() LocalPosition {
	return LocalPosition{p.x - 1, p.y}
}

func (p *LocalPosition) Right() LocalPosition {
	return LocalPosition{p.x + 1, p.y}
}

func (p *LocalPosition) Up() LocalPosition {
	return LocalPosition{p.x, p.y - 1}
}

func (p *LocalPosition) Down() LocalPosition {
	return LocalPosition{p.x - 1, p.y + 1}
}

type World struct {
}

func (w *World) LoadMap(id MapIdentifier) *WorldMap {
	return nil
}

type WorldMap struct {
	identifier MapIdentifier
}
