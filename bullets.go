package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type bullets struct {
	bullets []*bullet
	player  *Player
	enemies *enemies
}

func (b *bullets) all() []*bullet {
	return b.bullets
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
