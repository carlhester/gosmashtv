package main

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type enemy struct {
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

func newEnemy(x, y int, u, d, l, r bool) *enemy {
	sprite, _, err := ebitenutil.NewImageFromFile("./assets/enemy.png")
	if err != nil {
		panic(err)
	}
	img := sprite.SubImage(image.Rect(1, 1, 7, 7)).(*ebiten.Image)

	return &enemy{
		x:           x,
		y:           y,
		img:         img,
		movingRight: r,
		movingLeft:  l,
		movingUp:    u,
		movingDown:  d,
		active:      true,
		speed:       1,
	}
}

func (e *enemy) update(p *Player) {
	if p.x > e.x {
		e.x += 1
	}

	if p.x < e.x {
		e.x -= 1
	}

	if p.y > e.y {
		e.y += 1
	}

	if p.y < e.y {
		e.y -= 1
	}

	e.w, e.h = e.img.Size()
	e.rect.t = e.y
	e.rect.b = e.y + e.h
	e.rect.l = e.x
	e.rect.r = e.x + e.w
}

func (e *enemy) Draw(screen *ebiten.Image) {
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(float64(e.x), float64(e.y))
	screen.DrawImage(e.img, opts)
}
