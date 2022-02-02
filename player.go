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
	facingRight bool
	facingLeft  bool
	facingUp    bool
	facingDown  bool
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

func (p *Player) face(dir string) {
	p.facingDown = false
	p.facingUp = false
	p.facingLeft = false
	p.facingRight = false

	switch dir {
	case "up":
		p.facingUp = true
	case "down":
		p.facingDown = true
	case "left":
		p.facingLeft = true
	case "right":
		p.facingRight = true
	}

}

func (p *Player) Update() {
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowRight) {
		p.movingRight = true
		p.face("right")
	} else if inpututil.IsKeyJustReleased(ebiten.KeyArrowRight) {
		p.movingRight = false
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyArrowLeft) {
		p.movingLeft = true
		p.face("left")
	} else if inpututil.IsKeyJustReleased(ebiten.KeyArrowLeft) {
		p.movingLeft = false
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) {
		p.movingUp = true
		p.face("up")
	} else if inpututil.IsKeyJustReleased(ebiten.KeyArrowUp) {
		p.movingUp = false
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) {
		p.movingDown = true
		p.face("down")
	} else if inpututil.IsKeyJustReleased(ebiten.KeyArrowDown) {
		p.movingDown = false
	}

	if p.movingLeft {
		p.img = p.imgs[0]
		p.x -= p.speed
	}
	if p.movingRight {
		p.img = p.imgs[0]
		p.x += p.speed
	}
	if p.movingUp {
		p.img = p.imgs[1]
		p.y -= p.speed
	}
	if p.movingDown {
		p.img = p.imgs[1]
		p.y += p.speed
	}

	if p.x <= 0 {
		p.x = 0
	}
	if p.x >= LEVELWIDTH-SPRITESIZE {
		p.x = LEVELWIDTH - SPRITESIZE
	}

	if p.y <= 0 {
		p.y = 0
	}
	if p.y >= LEVELHEIGHT-SPRITESIZE {
		p.y = LEVELHEIGHT - SPRITESIZE
	}

}

func (p *Player) Coords() (x, y int) {
	return p.x, p.y
}

func (p *Player) Draw(screen *ebiten.Image) {
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(float64(p.x), float64(p.y))
	// sx, sy := p.img.Size()
	// opts.GeoM.Scale(float64(SPRITESIZE/sx), float64(SPRITESIZE/sy))
	screen.DrawImage(p.img, opts)
}
