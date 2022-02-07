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
	bullets *bullets
	debug   bool
	enemies *enemies
}

func newGame() *Game {
	p := newPlayer()

	levelEnemies := []*enemy{newEnemy(50, 50, false, false, false, false)}
	enemies := &enemies{enemies: levelEnemies}

	return &Game{
		player:  p,
		debug:   true,
		enemies: enemies,
		bullets: &bullets{player: p, enemies: enemies},
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
	g.bullets.update()
	g.enemies.update()

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0, 0, 0, 1})

	g.player.Draw(screen)
	if g.debug {
		ebitenutil.DebugPrint(screen, fmt.Sprintf("(T: %d, B: %d, L:%d, R:%d, shots: %d, enemies: %d)",
			g.player.rect.t,
			g.player.rect.b,
			g.player.rect.l,
			g.player.rect.r,
			len(g.bullets.all()),
			len(g.enemies.enemies),
		))
	}

	g.bullets.draw(screen)
	g.enemies.draw(screen)

}

func (g *Game) Layout(w, h int) (screenWidth, screenHeight int) {
	return int(LEVELWIDTH), int(LEVELHEIGHT)
}
