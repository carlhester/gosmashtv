package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type bullets struct {
	bullets []*bullet
	player  *Player
	enemies *enemies
}

func (b *bullets) all() []*bullet {
	return b.bullets
}

func (b *bullets) handleInput() {
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) {
		b.bullets = append(b.bullets, newBullet(b.player.x+(b.player.w/2), b.player.y, T, F, F, F))
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) {
		b.bullets = append(b.bullets, newBullet(b.player.x+(b.player.w/2), b.player.y, F, T, F, F))
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowLeft) {
		b.bullets = append(b.bullets, newBullet(b.player.x, b.player.y+(b.player.h/2), F, F, T, F))
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowRight) {
		b.bullets = append(b.bullets, newBullet(b.player.x, b.player.y+(b.player.h/2), F, F, F, T))
	}

}

func (b *bullets) refreshActive() {
	activeBullets := []*bullet{}
	for _, bullet := range b.bullets {
		if bullet.active {
			activeBullets = append(activeBullets, bullet)
		}
	}

	b.bullets = activeBullets
}

func (b *bullets) update() {
	b.handleInput()

	for _, bullet := range b.bullets {
		bullet.Update(b.enemies.enemies)
	}

	b.refreshActive()

}

func (b *bullets) draw(screen *ebiten.Image) {
	for _, bullet := range b.bullets {
		bullet.Draw(screen)
	}
}
