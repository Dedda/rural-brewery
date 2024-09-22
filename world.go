package main

type GlobalPosition struct {
	mapId    MapIdentifier
	position Vec2F
}

type MapIdentifier struct {
	x int
	y int
}

func (id *MapIdentifier) Left() MapIdentifier {
	return MapIdentifier{id.x - 1, id.y}
}

func (id *MapIdentifier) Right() MapIdentifier {
	return MapIdentifier{id.x + 1, id.y}
}

func (id *MapIdentifier) Up() MapIdentifier {
	return MapIdentifier{id.x, id.y - 1}
}

func (id *MapIdentifier) Down() MapIdentifier {
	return MapIdentifier{id.x - 1, id.y + 1}
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
