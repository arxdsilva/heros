package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

var sqr *ebiten.Image
var opts ebiten.DrawImageOptions

func update(screen *ebiten.Image) (err error) {
	screen.Fill(color.NRGBA{0xff, 0x00, 0x00, 0xff})
	ebitenutil.DebugPrint(screen, "Hello world!")
	if sqr == nil {
		sqr, _ = ebiten.NewImage(16, 16, ebiten.FilterNearest)
	}
	sqr.Fill(color.White)
	// When the "up arrow key" is pressed..
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		opts.GeoM.Translate(0, -1)
		ebitenutil.DebugPrint(screen, "You're pressing the 'UP' button.")
	}
	// When the "down arrow key" is pressed..
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		opts.GeoM.Translate(0, 1)
		ebitenutil.DebugPrint(screen, "\nYou're pressing the 'DOWN' button.")
	}
	// When the "left arrow key" is pressed..
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		opts.GeoM.Translate(-1, 0)
		ebitenutil.DebugPrint(screen, "\n\nYou're pressing the 'LEFT' button.")
	}
	// When the "right arrow key" is pressed..
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		opts.GeoM.Translate(1, 0)
		ebitenutil.DebugPrint(screen, "\n\n\nYou're pressing the 'RIGHT' button.")
	}
	screen.DrawImage(sqr, &opts)
	return
}

func main() {
	ebiten.Run(update, 320, 240, 2, "Hello world!")
}
