package main

import (
	"github.com/ByteArena/box2d"
	"github.com/Dedda/rural-brewery/assets"
)

const (
	inventorySize   = 32
	playerBaseSpeed = 4
)

type Player struct {
	name      string
	inventory *Inventory
	location  GlobalPosition
	body      *box2d.B2Body
}

func NewPlayer(name string) *Player {
	return &Player{
		name: name,
		inventory: &Inventory{
			items: [32]*InventoryItem{},
		},
		body: nil,
	}
}

func (p *Player) CreateBody(world *box2d.B2World) {
	bd := box2d.NewB2BodyDef()
	bd.Type = box2d.B2BodyType.B2_dynamicBody
	body := world.CreateBody(bd)
	shape := box2d.MakeB2PolygonShape()
	shape.SetAsBox(float64(assets.PlayerImage.Bounds().Dx()), float64(assets.PlayerImage.Bounds().Dy()))
	fixtureDef := box2d.MakeB2FixtureDef()
	fixtureDef.Shape = &shape
	body.CreateFixtureFromDef(&fixtureDef)
	body.SetFixedRotation(true)
	body.SetTransform(box2d.B2Vec2{0, 0}, 0)
	p.body = body
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
