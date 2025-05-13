package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth  = 1280
	screenHeight = 640
	halfWidth    = screenWidth >> 1
	halfHeight   = screenHeight >> 1
	tileSize     = 40
	topBarHeight = 130
)

var IS_RELEASE bool = true
var showDebug bool = false
var takingScreenShot bool = false

func main() {
	g := newGame()

	ebiten.SetWindowTitle("*NULL Level Editor")
	ebiten.SetWindowSize(screenWidth, screenHeight)

	if err := ebiten.RunGame(&g); err != nil {
		log.Fatal(err)
	}
}
