package main

import "github.com/hajimehoshi/ebiten/v2"

type Game struct {
	player   *Player
	world    *World
	calendar *GameCalendar
}

func NewGame(p *Player) (*Game, error) {
	world := &World{}
	calendar := &GameCalendar{}
	return &Game{
		player:   p,
		world:    world,
		calendar: calendar,
	}, nil
}

func (g *Game) CurrentMap() (*WorldMap, error) {
	return g.world.GetMap(g.player.location.mapId)
}

func (g *Game) Update() (GameState, error) {
	g.calendar.Update()
	return g.HandleInput()
}

func (g *Game) HandleInput() (GameState, error) {
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		g.player.location.position.y -= playerBaseSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		g.player.location.position.y += playerBaseSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		g.player.location.position.x -= playerBaseSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		g.player.location.position.x += playerBaseSpeed
	}
	return nil, nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.DrawBackground(screen)
	g.DrawPlayer(screen)
	g.DrawUi(screen)
}

func (g *Game) DrawBackground(screen *ebiten.Image) {}

func (g *Game) DrawPlayer(screen *ebiten.Image) {
	position := g.player.location.position
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(position.x, position.y)
	screen.DrawImage(PlayerImage, op)
}

func (g *Game) DrawUi(screen *ebiten.Image) {
	const (
		calendarSize = 96
	)
	calendarPosition := Point2F{
		x: gameWidth - calendarSize,
		y: gameHeight - calendarSize,
	}
	g.calendar.Draw(screen, calendarPosition)
}
