package main

import "github.com/hajimehoshi/ebiten/v2"

type bullet struct {
	img         *ebiten.Image
	x           int
	y           int
	speed       int
	movingRight bool
	movingLeft  bool
	movingUp    bool
	movingDown  bool
}

func newBullet(x, y int, u, d, l, r bool) *bullet {
	img := ebiten.NewImage(1, 1)
	return &bullet{
		x:           x,
		y:           y,
		img:         img,
		speed:       5,
		movingRight: r,
		movingLeft:  l,
		movingUp:    u,
		movingDown:  d,
	}
}

func (b *bullet) Update() {
	if b.movingDown {
		b.y += b.speed
	} else if b.movingUp {
		b.y -= b.speed
	}

	if b.movingLeft {
		b.x -= b.speed
	} else if b.movingRight {
		b.x += b.speed
	}

}

func (b *bullet) Draw(screen *ebiten.Image) {
	opts := &ebiten.DrawImageOptions{}
	opts.ColorM.Translate(float64(0x99), float64(0x99), float64(0x0), float64(0xFF))
	opts.GeoM.Translate(float64(b.x), float64(b.y))
	screen.DrawImage(b.img, opts)
}
