package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

func updateRoster(screen *ebiten.Image) (err error) {
	screen.Fill(color.NRGBA{0x00, 0x00, 0x00, 0x00})
	ebitenutil.DebugPrint(screen, "Hello world!")
	return
}
