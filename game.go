package main

import (
	"github.com/ByteArena/box2d"
	"github.com/Dedda/rural-brewery/assets"
	"github.com/hajimehoshi/ebiten/v2"
	"log"
)

type Game struct {
	player   *Player
	world    *World
	calendar *GameCalendar
	b2dWorld *box2d.B2World
}

func NewGame(p *Player) (*Game, error) {
	world, err := LoadWorld()
	calendar := &GameCalendar{}
	b2dWorld := box2d.MakeB2World(box2d.MakeB2Vec2(0, 0))
	p.CreateBody(&b2dWorld)
	return &Game{
		player:   p,
		world:    world,
		calendar: calendar,
		b2dWorld: &b2dWorld,
	}, err
}

func (g *Game) CurrentMap() (WorldMap, error) {
	return g.world.GetMap(g.player.location.mapId)
}

func (g *Game) Update() (GameState, error) {
	g.b2dWorld.Step(1.0/60, 8, 3)
	g.player.location.position = Vec2FFromB2Vec2(g.player.body.GetPosition())
	g.calendar.Update()
	for id, m := range g.world.maps {
		err := m.Update(g, id == g.player.location.mapId)
		if err != nil {
			return nil, err
		}
	}
	return g.HandleInput()
}

func (g *Game) HandleInput() (GameState, error) {
	playerMovement := GetPlayerMovementInput().MulF64(playerBaseSpeed).MulF64(60)
	g.player.body.SetLinearVelocity(playerMovement.ToB2D())
	return nil, nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.DrawBackground(screen)
	g.DrawPlayer(screen)
	g.DrawUi(screen)
}

func (g *Game) DrawBackground(screen *ebiten.Image) {
	m, err := g.world.GetMap(g.player.location.mapId)
	if err != nil {
		log.Fatal(err)
	}
	m.Draw(screen, g)
}

func (g *Game) DrawPlayer(screen *ebiten.Image) {
	position := g.player.body.GetPosition()
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(position.X, position.Y)
	screen.DrawImage(assets.PlayerImage, op)
}

func (g *Game) DrawUi(screen *ebiten.Image) {
	const (
		calendarSize = 96
	)
	calendarPosition := Vec2F{
		X: gameWidth - calendarSize,
		Y: gameHeight - calendarSize,
	}
	g.calendar.Draw(screen, calendarPosition)
}
