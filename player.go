package main

import (
	"fmt"
	"github.com/ByteArena/box2d"
	"github.com/Dedda/rural-brewery/assets"
	"github.com/Dedda/rural-brewery/items"
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
	fmt.Println("Creating player body...")
	bodyDef := box2d.NewB2BodyDef()
	bodyDef.Type = box2d.B2BodyType.B2_dynamicBody
	bodyDef.Position = box2d.MakeB2Vec2(p.location.position.X, p.location.position.Y)
	body := world.CreateBody(bodyDef)
	box := box2d.NewB2PolygonShape()
	width := float64(assets.PlayerImage.Bounds().Dx())
	height := float64(assets.PlayerImage.Bounds().Dy())
	box.Set([]box2d.B2Vec2{
		box2d.MakeB2Vec2(0, 0),
		box2d.MakeB2Vec2(width, 0),
		box2d.MakeB2Vec2(width, height),
		box2d.MakeB2Vec2(0, height),
	}, 4)

	fixtureDef := box2d.MakeB2FixtureDef()
	fixtureDef.Shape = box
	fixtureDef.Density = 1
	body.CreateFixtureFromDef(&fixtureDef)
	p.body = body
}

type Inventory struct {
	items [inventorySize]*InventoryItem
}

type InventoryItem struct {
	item   *items.Item
	amount int
}

func (i *InventoryItem) AddInventoryItem(o *InventoryItem) *InventoryItem {
	if i.item.Id != o.item.Id {
		return o
	}
	total := i.amount + o.amount
	if total <= i.item.MaxStackSize {
		i.amount = total
		return nil
	}
	i.amount = i.item.MaxStackSize
	remainder := total - i.item.MaxStackSize
	return &InventoryItem{
		item:   i.item,
		amount: remainder,
	}
}
