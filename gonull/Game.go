package main

import (
	"bytes"
	"image/png"
	"null/stagerror"
	"null/stagui"
	"os"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Game struct {
	font  *FontHandler
	image *ImageHandler
	audio *AudioHandler

	title, settings, credits, instructions, win *stagui.Page

	level *Level

	phase       Phase
	curMousePos [2]int
	player      *Player
}

func (g *Game) Update() error {
	g.handleDebug()
	g.handleFullScreen()
	g.handleScreenShot()
	g.audio.updateAudio()

	x, y := ebiten.CursorPosition()
	g.curMousePos = [2]int{x, y}

	var pressed string
	var button *stagui.Button
	var slider *stagui.Slider

	switch g.phase {
	case TITLE:
		pressed, _, _ = g.title.Update(g.curMousePos)
		switch pressed {
		case "start":
			g.phase = GAME
			g.player = newPlayer(g.image.GetImage("player"))

			var err *stagerror.Error
			g.level, err = g.loadLevel("level")
			if err != nil {
				err.SaveToLog(IS_RELEASE)
				g.phase = TITLE
			}
			runTime = 0
			runInputs = [][]InputCode{}
		case "settings":
			g.phase = SETTINGS
		case "credits":
			g.phase = CREDITS
		case "instructions":
			g.phase = INSTRUCTIONS
		case "quit":
			os.Exit(0)
		}
	case SETTINGS:
		pressed, button, slider = g.settings.Update(g.curMousePos)
		switch pressed {
		case "mute":
			button.Pressed = !button.Pressed
			g.audio.toggleMute()
		case "volume":
			g.audio.setVolume(slider.Value)
		case "effectVolume":
			g.audio.setEffectVolume(slider.Value)
		case "back":
			g.phase = TITLE
		}
	case CREDITS:
		pressed, _, _ = g.credits.Update(g.curMousePos)
		switch pressed {
		case "back":
			g.phase = TITLE
		}
	case INSTRUCTIONS:
		pressed, _, _ = g.instructions.Update(g.curMousePos)
		switch pressed {
		case "back":
			g.phase = TITLE
		}
	case GAME:
		runTime++
		runInputs = append(runInputs, []InputCode{})
		g.player.update(g, g.level)
		g.level.update(g)
	case REPLAY:
		runTime++
		g.player.updateAsReplay(g, g.level)
		g.level.update(g)
	case WIN:
		pressed, _, _ = g.win.Update(g.curMousePos)
		switch pressed {
		case "back":
			g.phase = TITLE
		case "save":
			saveInputs()
		}
	}

	return nil
}

func (g *Game) handleFullScreen() {
	if !(inpututil.IsKeyJustPressed(ebiten.KeyF) && ebiten.IsKeyPressed(ebiten.KeyGraveAccent)) {
		return
	}

	ebiten.SetFullscreen(!ebiten.IsFullscreen())
}

func (g *Game) handleScreenShot() {
	if inpututil.IsKeyJustPressed(ebiten.KeyS) && ebiten.IsKeyPressed(ebiten.KeyGraveAccent) {
		takingScreenShot = true
	}
}

func (g *Game) handleDebug() {
	if !inpututil.IsKeyJustPressed(ebiten.KeyGraveAccent) {
		return
	}

	showDebug = !showDebug

	if !ebiten.IsKeyPressed(ebiten.KeyShift) {
		return
	}

	if ebiten.IsVsyncEnabled() {
		ebiten.SetTPS(1_000)
		ebiten.SetVsyncEnabled(false)
	} else {
		ebiten.SetTPS(60)
		ebiten.SetVsyncEnabled(true)
	}
}

func (g *Game) Draw(screen *ebiten.Image) {
	switch g.phase {
	case TITLE:
		g.title.Draw(screen, g.image, g.font, screenWidth, screenHeight, halfWidth)
	case SETTINGS:
		g.settings.Draw(screen, g.image, g.font, screenWidth, screenHeight, halfWidth)
	case CREDITS:
		g.credits.Draw(screen, g.image, g.font, screenWidth, screenHeight, halfWidth)
	case INSTRUCTIONS:
		g.instructions.Draw(screen, g.image, g.font, screenWidth, screenHeight, halfWidth)
	case GAME:
		g.player.draw(screen, g)
		g.level.draw(screen, g.image, g.font, g.player.y+g.player.h/2-halfHeight)
	case REPLAY:
		g.player.draw(screen, g)
		g.level.draw(screen, g.image, g.font, g.player.y+g.player.h/2-halfHeight)
	case WIN:
		g.win.Draw(screen, g.image, g.font, screenWidth, screenHeight, halfWidth)
	}

	g.drawDebugScreen(screen)
	g.takeScreenShot(screen)
}

func (g *Game) takeScreenShot(screen *ebiten.Image) {
	if !takingScreenShot {
		return
	}

	takingScreenShot = false

	var out bytes.Buffer
	err := png.Encode(&out, screen)
	if err != nil {
		stagerror.New(64, "Couldn't convert to png when screenshotting").SaveToLog(IS_RELEASE)
		return
	}

	err = os.WriteFile("ss.png", out.Bytes(), 0644)
	if err != nil {
		stagerror.New(65, "Couldn't convert save screenshot to file").SaveToLog(IS_RELEASE)
	}
}

func (g *Game) winPhase() {
	g.phase = WIN
	g.createWinPage(g.saveNewTime())
}

func (g *Game) drawDebugScreen(screen *ebiten.Image) {
	if !showDebug {
		return
	}

	ebitenutil.DebugPrint(
		screen,
		"TPS: "+strconv.FormatFloat(ebiten.ActualTPS(), 'f', 3, 64)+"\n"+
			"FPS: "+strconv.FormatFloat(ebiten.ActualFPS(), 'f', 3, 64),
	)
}

func (*Game) Layout(int, int) (int, int) {
	return screenWidth, screenHeight
}

func (g *Game) saveNewTime() uint64 {

	data, err := os.ReadFile("time.txt")
	var oldTime uint64
	if err != nil || len(data) != 8 {
		err = os.WriteFile("time.txt", splitUInt64(uint64(runTime)), 0644)
		if err != nil {
			stagerror.New(66, "Couldn't save score to disk").SaveToLog(IS_RELEASE)
		}
	} else {
		oldTime = joinUInt64(data)
		if oldTime < runTime {
			return oldTime
		}

		err = os.WriteFile("time.txt", splitUInt64(uint64(runTime)), 0644)
		if err != nil {
			stagerror.New(66, "Couldn't save score to disk").SaveToLog(IS_RELEASE)
		}
	}

	return runTime
}

func newGame() Game {
	g := Game{}

	g.audio = &AudioHandler{}
	g.audio.init()

	g.audio.playAudio("theme0")

	g.font = &FontHandler{}
	g.font.loadFonts()

	g.image = &ImageHandler{}
	g.image.initImages()

	g.initUI()

	return g
}
