package main

const (
	inventorySize   = 32
	playerBaseSpeed = 2
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
	item   Item
	amount int
}

func (i *InventoryItem) AddInventoryItem(o *InventoryItem) *InventoryItem {
	if i.item.Id() != o.item.Id() {
		return o
	}
	total := i.amount + o.amount
	if total <= i.item.MaxStackSize() {
		i.amount = total
		return nil
	}
	i.amount = i.item.MaxStackSize()
	remainder := total - i.item.MaxStackSize()
	return &InventoryItem{
		item:   i.item,
		amount: remainder,
	}
}
