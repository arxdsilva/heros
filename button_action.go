package main

import (
	"github.com/hajimehoshi/ebiten"
)

func buttonClicked(x, y int) bool {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		// fmt.Println("lala")
		// inXrange := ((x>width/3)&&())
		return true
	}
	return false
}
