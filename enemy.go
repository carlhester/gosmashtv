package main

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type enemy struct {
	img *ebiten.Image
	x   int
	y   int
	// w           int
	// h           int
	// speed       int
	movingRight bool
	movingLeft  bool
	movingUp    bool
	movingDown  bool
	// rect        rect
	active bool
}

func newEnemy(x, y int, u, d, l, r bool) *enemy {
	sprite, _, err := ebitenutil.NewImageFromFile("./assets/slime.png")
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
	}
}

func (e *enemy) Update() {
}

func (e *enemy) Draw(screen *ebiten.Image) {
	opts := &ebiten.DrawImageOptions{}
	// opts.ColorM.Translate(float64(0x00), float64(0x22), float64(0x0), float64(0xFF))
	opts.GeoM.Translate(float64(e.x), float64(e.y))
	screen.DrawImage(e.img, opts)
}
