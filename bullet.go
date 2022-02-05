package main

import "github.com/hajimehoshi/ebiten/v2"

type bullet struct {
	img         *ebiten.Image
	x           int
	y           int
	w           int
	h           int
	speed       int
	movingRight bool
	movingLeft  bool
	movingUp    bool
	movingDown  bool
	rect        rect
	active      bool
}

func newBullet(x, y int, u, d, l, r bool) *bullet {
	img := ebiten.NewImage(1, 1)

	w, h := img.Size()

	rect := rect{
		t: y,
		b: y + h,
		l: x,
		r: x + w,
	}

	return &bullet{
		x:           x,
		y:           y,
		w:           w,
		h:           h,
		img:         img,
		speed:       5,
		movingRight: r,
		movingLeft:  l,
		movingUp:    u,
		movingDown:  d,
		rect:        rect,
		active:      true,
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

	if b.x <= 0 {
		b.active = false
	}
	if b.x >= LEVELWIDTH {
		b.active = false
	}

	if b.y <= 0 {
		b.active = false
	}
	if b.y >= LEVELHEIGHT-b.h {
		b.active = false
	}

	b.w, b.h = b.img.Size()
	b.rect.t = b.y
	b.rect.b = b.y + b.h
	b.rect.l = b.x
	b.rect.r = b.x + b.w

}

func (b *bullet) Draw(screen *ebiten.Image) {
	opts := &ebiten.DrawImageOptions{}
	opts.ColorM.Translate(float64(0x99), float64(0x99), float64(0x0), float64(0xFF))
	opts.GeoM.Translate(float64(b.x), float64(b.y))
	screen.DrawImage(b.img, opts)
}
