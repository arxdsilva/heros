package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten"
)

func updateHome(screen *ebiten.Image) (err error) {
	screen.Fill(color.NRGBA{0xff, 0x00, 0x00, 0xff})
	button, err := ebiten.NewImage(int(width/3), int(height/10), ebiten.FilterNearest)
	if err != nil {
		return
	}
	button.Fill(color.White)
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(float64((width/2)-(width/6)), float64((height/2)-(height/20)))
	screen.DrawImage(button, opts)
	x, y := ebiten.CursorPosition()
	if buttonClicked(x, y) {
		update = updateRoster
		// fmt.Println(update, updateRoster, updateHome)
		return ErrNextScreen
	}
	return
}
