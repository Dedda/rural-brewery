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
	mapCache []*WorldMap
}

func (w *World) LoadMap(id MapIdentifier) (*WorldMap, error) {
	if len(w.mapCache) > 5 {
		w.mapCache = w.mapCache[len(w.mapCache):]
	}
	loaded, err := w.LoadMap(id)
	w.mapCache = append(w.mapCache, loaded)
	return loaded, err
}

func (w *World) GetMap(id MapIdentifier) (*WorldMap, error) {
	for _, worldMap := range w.mapCache {
		if worldMap.identifier == id {
			return worldMap, nil
		}
	}
	return w.LoadMap(id)
}

type WorldMap struct {
	identifier MapIdentifier
	name       string
}
