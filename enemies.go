package main

import "github.com/hajimehoshi/ebiten/v2"

type enemies struct {
	enemies []*enemy
	p       *Player
}

func (e *enemies) update() {
	for _, enemy := range e.enemies {
		enemy.update(e.p)
	}
	e.refreshActive()
}

func (e *enemies) refreshActive() {
	activeEnemies := []*enemy{}
	for _, enemy := range e.enemies {
		if enemy.active {
			activeEnemies = append(activeEnemies, enemy)
		}
	}

	e.enemies = activeEnemies
}

func (e *enemies) draw(screen *ebiten.Image) {
	for _, enemy := range e.enemies {
		enemy.Draw(screen)
	}
}
