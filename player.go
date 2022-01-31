package main

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"

	_ "image/png"
)

type Player struct {
	imgs        []*ebiten.Image
	img         *ebiten.Image
	x           int
	y           int
	speed       int
	movingRight bool
	movingLeft  bool
	movingUp    bool
	movingDown  bool
}

func newPlayer() *Player {
	sprite, _, err := ebitenutil.NewImageFromFile("./assets/slime.png")
	if err != nil {
		panic(err)
	}

	img := sprite.SubImage(image.Rect(1, 1, 7, 7)).(*ebiten.Image)
	img2 := sprite.SubImage(image.Rect(9, 2, 15, 7)).(*ebiten.Image)

	imgs := []*ebiten.Image{}
	imgs = append(imgs, img)
	imgs = append(imgs, img2)

	return &Player{
		x:     0,
		y:     0,
		imgs:  imgs,
		img:   imgs[0],
		speed: 4,
	}
}

func (p *Player) Type() string {
	return "player"
}

func (p *Player) Update() {
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowRight) {
		p.movingRight = true
	} else if inpututil.IsKeyJustReleased(ebiten.KeyArrowRight) {
		p.movingRight = false
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyArrowLeft) {
		p.movingLeft = true
	} else if inpututil.IsKeyJustReleased(ebiten.KeyArrowLeft) {
		p.movingLeft = false
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) {
		p.movingUp = true
	} else if inpututil.IsKeyJustReleased(ebiten.KeyArrowUp) {
		p.movingUp = false
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) {
		p.movingDown = true
	} else if inpututil.IsKeyJustReleased(ebiten.KeyArrowDown) {
		p.movingDown = false
	}

	if p.movingLeft {
		p.x -= p.speed
		p.img = p.imgs[0]
	}
	if p.movingRight {
		p.x += p.speed
		p.img = p.imgs[0]
	}
	if p.movingUp {
		p.y -= p.speed
		p.img = p.imgs[1]
	}
	if p.movingDown {
		p.y += p.speed
		p.img = p.imgs[1]
	}

}

func (p *Player) Coords() (x, y int) {
	return p.x, p.y
}

func (p *Player) Draw(screen *ebiten.Image) {
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(float64(p.x), float64(p.y))
	screen.DrawImage(p.img, opts)
}
