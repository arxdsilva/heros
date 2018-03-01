package main

import (
	"github.com/hajimehoshi/ebiten"
)

var width = 320
var height = 240
var update func(*ebiten.Image) error

func main() {
	update = updateHome
	ebiten.Run(update, width, height, 2, "Hello world!")
}
