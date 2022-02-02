package main

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Game struct {
	player  *Player
	bullets []*bullet
	debug   bool
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

	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		if g.player.facingDown {
			g.bullets = append(g.bullets, newBullet(g.player.x, g.player.y, false, true, false, false))
		} else if g.player.facingUp {
			g.bullets = append(g.bullets, newBullet(g.player.x, g.player.y, true, false, false, false))
		} else if g.player.facingLeft {
			g.bullets = append(g.bullets, newBullet(g.player.x, g.player.y, false, false, true, false))
		} else if g.player.facingRight {
			g.bullets = append(g.bullets, newBullet(g.player.x, g.player.y, false, false, false, true))
		}
	}

	g.player.Update()
	for _, bullet := range g.bullets {
		bullet.Update()
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0, 0, 0, 1})

	g.player.Draw(screen)
	if g.debug {
		px, py := g.player.Coords()
		ebitenutil.DebugPrint(screen, fmt.Sprintf("(%d, %d)", px, py))
		fmt.Println(g.player.img.Size())
	}

	for _, bullet := range g.bullets {
		bullet.Draw(screen)
	}
}

func (g *Game) Layout(w, h int) (screenWidth, screenHeight int) {
	return int(LEVELWIDTH), int(LEVELHEIGHT)
}
