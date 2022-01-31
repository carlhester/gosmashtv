package main

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Game struct {
	player *Player
	debug  bool
}

func newGame() *Game {
	p := newPlayer()

	return &Game{
		player: p,
		debug:  true,
	}
}

func (g *Game) Start() {
}

func (g *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyQ) {
		return fmt.Errorf("quit")
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyF1) {
		g.debug = !g.debug
	}

	g.player.Update()

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0, 0, 0, 1})

	g.player.Draw(screen)
	if g.debug {
		px, py := g.player.Coords()
		ebitenutil.DebugPrint(screen, fmt.Sprintf("(%d, %d)", px, py))
	}
}

func (g *Game) Layout(w, h int) (screenWidth, screenHeight int) {
	return int(LEVELWIDTH), int(LEVELHEIGHT)
}
