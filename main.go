package main

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	SPRITESIZE  = 6
	LEVELWIDTH  = SPRITESIZE * 48
	LEVELHEIGHT = SPRITESIZE * 36
)

var (
	WHITE = color.RGBA{0xff, 0xff, 0xff, 0xff}
	BLACK = color.RGBA{0x0, 0x0, 0x0, 0xff}
	GREEN = color.RGBA{0x0, 0xff, 0x0, 0xff}
)

func main() {

	ebiten.SetWindowSize(int(LEVELWIDTH*3), int(LEVELHEIGHT*3))
	ebiten.SetWindowTitle("EbitGame")

	game := newGame()
	game.Start()

	if err := ebiten.RunGame(game); err != nil {
		fmt.Printf("%+v\n", err)
	}
}
