package main

import (
	"null/stagerror"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth  = 1280
	screenHeight = 640
	halfWidth    = screenWidth >> 1
	halfHeight   = screenHeight >> 1
	tileSize     = screenWidth >> 4
)

const IS_RELEASE bool = true

var showDebug bool = false
var takingScreenShot bool = false

var runTime uint64

func main() {
	g := newGame()

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("*NULL")

	args := os.Args
	if len(args) == 2 {
		// Attempt replay-mode
		g.phase = REPLAY
		g.player = newPlayer(g.image.GetImage("player"))

		var serr *stagerror.Error
		g.level, serr = g.loadLevel("level")
		if serr != nil {
			serr.SaveToLog(IS_RELEASE)
			return
		}

		serr = loadInputs(args[1])
		if serr != nil {
			serr.SaveToLog(IS_RELEASE)
			return
		}
	}

	err := ebiten.RunGame(&g)
	if err != nil {
		panic(err)
	}
}
