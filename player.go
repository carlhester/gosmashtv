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
	w           int
	h           int
	speed       int
	movingRight bool
	movingLeft  bool
	movingUp    bool
	movingDown  bool
	facingRight bool
	facingLeft  bool
	facingUp    bool
	facingDown  bool
	rect
	Bullets *bullets
}

func newPlayer() *Player {
	sprite, _, err := ebitenutil.NewImageFromFile("./assets/player.png")
	if err != nil {
		panic(err)
	}

	imgs := []*ebiten.Image{}
	img := sprite.SubImage(image.Rect(0, 0, 6, 6)).(*ebiten.Image)
	imgs = append(imgs, img)
	img = sprite.SubImage(image.Rect(7, 0, 13, 6)).(*ebiten.Image)
	imgs = append(imgs, img)
	img = sprite.SubImage(image.Rect(12, 0, 18, 6)).(*ebiten.Image)
	imgs = append(imgs, img)
	img = sprite.SubImage(image.Rect(17, 0, 23, 6)).(*ebiten.Image)
	imgs = append(imgs, img)

	w, h := imgs[0].Size()
	x := 20
	y := 50

	rect := rect{
		t: y,
		b: y + h,
		l: x,
		r: x + w,
	}

	return &Player{
		x:     x,
		y:     y,
		w:     w,
		h:     h,
		imgs:  imgs,
		img:   imgs[0],
		rect:  rect,
		speed: 4,
	}
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
	if inpututil.IsKeyJustPressed(ebiten.KeyF) {
		p.movingRight = true
		p.face("right")
	} else if inpututil.IsKeyJustReleased(ebiten.KeyF) {
		p.movingRight = false
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyS) {
		p.movingLeft = true
		p.face("left")
	} else if inpututil.IsKeyJustReleased(ebiten.KeyS) {
		p.movingLeft = false
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyE) {
		p.movingUp = true
		p.face("up")
	} else if inpututil.IsKeyJustReleased(ebiten.KeyE) {
		p.movingUp = false
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyD) {
		p.movingDown = true
		p.face("down")
	} else if inpututil.IsKeyJustReleased(ebiten.KeyD) {
		p.movingDown = false
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) {
		p.Bullets.bullets = append(p.Bullets.bullets, newBullet(p.x+(p.w/2), p.y, T, F, F, F))
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) {
		p.Bullets.bullets = append(p.Bullets.bullets, newBullet(p.x+(p.w/2), p.y, F, T, F, F))
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowLeft) {
		p.Bullets.bullets = append(p.Bullets.bullets, newBullet(p.x, p.y+(p.h/2), F, F, T, F))
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowRight) {
		p.Bullets.bullets = append(p.Bullets.bullets, newBullet(p.x, p.y+(p.h/2), F, F, F, T))
	}

	if p.movingLeft {
		p.img = p.imgs[3]
		p.x -= p.speed
	}
	if p.movingRight {
		p.img = p.imgs[1]
		p.x += p.speed
	}
	if p.movingUp {
		p.img = p.imgs[0]
		p.y -= p.speed
	}
	if p.movingDown {
		p.img = p.imgs[2]
		p.y += p.speed
	}

	if p.x <= 0 {
		p.x = 0
	}
	if p.x >= LEVELWIDTH-p.w {
		p.x = LEVELWIDTH - p.w
	}

	if p.y <= 0 {
		p.y = 0
	}
	if p.y >= LEVELHEIGHT-p.h {
		p.y = LEVELHEIGHT - p.h
	}

	p.w, p.h = p.img.Size()
	p.rect.t = p.y
	p.rect.b = p.y + p.h
	p.rect.l = p.x
	p.rect.r = p.x + p.w

}

func (p *Player) Coords() (x, y int) {
	return p.x, p.y
}

func (p *Player) Draw(screen *ebiten.Image) {
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(float64(p.x), float64(p.y))
	screen.DrawImage(p.img, opts)
}
