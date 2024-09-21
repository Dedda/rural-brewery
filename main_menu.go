package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"image/color"
	"os"
)

const (
	mainMenuHeaderSize = 64
	mainMenuOptionSize = 48
)

type MainMenu struct {
	options  []MainMenuOption
	selected int
}

type MainMenuOption struct {
	title  string
	action func(m *MainMenu) (GameState, error)
}

func NewMainMenu() (*MainMenu, error) {
	m := &MainMenu{
		options: []MainMenuOption{
			{
				title: "Start Game",
				action: func(m *MainMenu) (GameState, error) {
					return m.StartGame()
				},
			},
			{
				title: "Quit",
				action: func(m *MainMenu) (GameState, error) {
					os.Exit(0)
					return nil, nil
				},
			},
		},
	}
	return m, nil
}

func (m *MainMenu) StartGame() (GameState, error) {
	p := NewPlayer("Dedda")
	return NewGame(p)
}

func (m *MainMenu) Up() {
	m.selected--
	if m.selected < 0 {
		m.selected = len(m.options) - 1
	}
}

func (m *MainMenu) Down() {
	m.selected++
	if m.selected >= len(m.options) {
		m.selected = 0
	}
}

func (m *MainMenu) Update() (GameState, error) {
	if inpututil.IsKeyJustPressed(ebiten.KeyDown) {
		m.Down()
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyUp) {
		m.Up()
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		return m.options[m.selected].action(m)
	}
	return nil, nil
}

func (m *MainMenu) Draw(screen *ebiten.Image) {
	DrawHeader(screen)
	m.DrawOptions(screen)
}

func DrawHeader(screen *ebiten.Image) {
	headerText := "Rural Brewery"
	headerFace := &text.GoTextFace{
		Source: JacquardaBastarda9FaceSource,
		Size:   mainMenuHeaderSize,
	}
	op := &text.DrawOptions{}
	op.GeoM.Translate(CenterTextHorizontally(&headerText, headerFace, gameWidth), 100)
	op.ColorScale.ScaleWithColor(color.White)
	text.Draw(screen, headerText, headerFace, op)
}

func (m *MainMenu) DrawOptions(screen *ebiten.Image) {
	titles := make([]string, len(m.options))
	for i, option := range m.options {
		titles[i] = option.title
	}
	face := &text.GoTextFace{
		Source: JacquardaBastarda9FaceSource,
		Size:   mainMenuOptionSize,
	}
	for i, title := range titles {
		if i == m.selected {
			title = "- " + title + " -"
		}
		op := &text.DrawOptions{}
		op.GeoM.Translate(CenterTextHorizontally(&title, face, gameWidth), 300+64*float64(i))
		op.ColorScale.ScaleWithColor(color.White)
		text.Draw(screen, title, face, op)
	}
}
