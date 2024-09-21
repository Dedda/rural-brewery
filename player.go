package main

const (
	inventorySize = 32
)

type Player struct {
	name      string
	inventory *Inventory
	location  GlobalPosition
}

func NewPlayer(name string) *Player {
	return &Player{
		name: name,
		inventory: &Inventory{
			items: [32]*InventoryItem{},
		},
	}
}

type Inventory struct {
	items [inventorySize]*InventoryItem
}

type InventoryItem struct {
	item   *Item
	amount int
}
