package main

import (
	"image/color"
	"math"
	"nullEditor/stagui"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Game struct {
	font  *FontHandler
	image *ImageHandler

	curMousePos [2]int

	offset      [2]int
	cameraSpeed int

	page *stagui.Page

	curCodeSelection TileCode
	curTileType      string

	level Level
}

func (g *Game) Update() error {
	x, y := ebiten.CursorPosition()
	g.curMousePos = [2]int{x, y}

	var pressed string
	// var button *stagui.Button
	// var slider *stagui.Slider
	var inlineTextBox *stagui.InlineTextBox

	pressed, _, _, inlineTextBox = g.page.Update(g.curMousePos)

	var codeChanged bool

	var w, h int

	switch pressed {
	case "save":
		g.level.save()
		return nil
	case "load":
		l, err := loadLevel(g)
		if err == nil {
			g.level = l
		} else {
			err.SaveToLog(IS_RELEASE)
		}
		return nil
	case "clear":
		l := newLevel(len(g.level.data[0]), len(g.level.data))
		g.level = l
		return nil
	case "addSection":
		g.level.addSection()
		return nil
	case "codeUp":
		g.curCodeSelection++
		codeChanged = true
	case "codeDown":
		g.curCodeSelection--
		codeChanged = true
	case "codeUpBlock":
		g.curCodeSelection += 8
		codeChanged = true
	case "codeDownBlock":
		g.curCodeSelection -= 8
		codeChanged = true
	case "width", "height":
		var err error
		w, err = strconv.Atoi(inlineTextBox.Text)
		if err != nil {
			return nil
		}
		h, err = strconv.Atoi(inlineTextBox.Text)
		if err != nil {
			return nil
		}
	}

	if codeChanged {
		g.curTileType = getRelevantCodeText(g.curCodeSelection)
		return nil
	}

	if w > 0 && h > 0 {
		g.level.resize(w, h)
		return nil
	}

	if ebiten.IsMouseButtonPressed(ebiten.MouseButton0) {
		// Placing a tile
		g.mouseHold()

		if inpututil.IsMouseButtonJustPressed(ebiten.MouseButton0) {
			g.mouseClick()
		}
	}

	// Right click for erasing
	if ebiten.IsMouseButtonPressed(ebiten.MouseButton2) {
		col := int(math.Floor(float64(g.curMousePos[0]+g.offset[0]) / tileSize))
		row := int(math.Floor(float64(g.curMousePos[1]-topBarHeight+g.offset[1]) / tileSize))

		if row < len(g.level.data) && row >= 0 {
			if col < len(g.level.data[row]) && col >= 0 {
				g.level.data[row][col] = 0
			}
		}
	}

	g.keyPress()

	return nil
}

func (g *Game) keyPress() {
	codeChanged := false

	// Changing the code
	// Technically this doesn't care about
	// overflows and underflows, and that's
	// fine because it means that if we use
	// all 256 codes that you can quickly
	// wrap to the last code, so I'm
	// keeping it in for now
	if inpututil.IsKeyJustPressed(ebiten.KeyR) {
		g.curCodeSelection++
		codeChanged = true
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyF) {
		g.curCodeSelection--
		codeChanged = true
	}

	// For jumping a section of codes
	if inpututil.IsKeyJustPressed(ebiten.KeyU) {
		g.curCodeSelection += 8
		codeChanged = true
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyJ) {
		g.curCodeSelection -= 8
		codeChanged = true
	}

	if codeChanged {
		g.curTileType = getRelevantCodeText(g.curCodeSelection)
	}

	// For moving the camera
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		g.offset[0] -= g.cameraSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		g.offset[0] += g.cameraSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		g.offset[1] -= g.cameraSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		g.offset[1] += g.cameraSpeed
	}
}

func (g *Game) mouseHold() {
	x := g.curMousePos[0]
	y := g.curMousePos[1]

	col := int(float64(x+g.offset[0]) / tileSize)
	row := int(float64(y+g.offset[1]-topBarHeight) / tileSize)

	if row < 0 || row >= len(g.level.data) {
		return
	}

	if col < 0 || col >= len(g.level.data[row]) {
		return
	}

	g.level.data[row][col] = g.curCodeSelection
}

func (g *Game) mouseClick() {
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.page.Draw(screen, g.image, g.font, screenWidth, screenHeight, halfWidth)

	for row := range len(g.level.data) {
		for col := range len(g.level.data[row]) {
			x := col*tileSize - g.offset[0]
			y := row*tileSize + topBarHeight - g.offset[1]
			g.drawRelevantImage(screen, g.level.data[row][col], float64(x), float64(y))
		}
	}
}

func getRelevantCodeText(code TileCode) string {
	switch code {
	case NO_TILE:
		return "Nothing"
	case BLOCK:
		return "Block"
	case NULL_SPAWN:
		return "Null Spawn"
	case PLAYER_START:
		return "Player Start"
	case UPWARD_BOOST:
		return "Upward Boost"
	case LEVEL_END:
		return "Level End"
	default:
		return "Unknown"
	}
}

func (g *Game) drawRelevantImage(screen *ebiten.Image, code TileCode, x, y float64) {
	switch code {
	case NO_TILE: // Nothing
		vector.DrawFilledRect(screen, float32(x+1), float32(y+1), tileSize-2, tileSize-2, color.RGBA{16, 16, 16, 255}, false)

	case BLOCK:
		vector.DrawFilledRect(screen, float32(x+1), float32(y+1), tileSize-2, tileSize-2, color.RGBA{255, 255, 255, 255}, false)
	case NULL_SPAWN:
		vector.DrawFilledRect(screen, float32(x+1), float32(y+1), tileSize-2, tileSize-2, color.RGBA{196, 196, 196, 255}, false)
	case PLAYER_START:
		vector.DrawFilledRect(screen, float32(x+1), float32(y+1), tileSize-2, tileSize-2, color.RGBA{64, 64, 255, 255}, false)
	case UPWARD_BOOST:
		vector.DrawFilledRect(screen, float32(x+1), float32(y+1), tileSize-2, tileSize-2, color.RGBA{255, 255, 0, 255}, false)
	case LEVEL_END:
		vector.DrawFilledRect(screen, float32(x+1), float32(y+1), tileSize-2, tileSize-2, color.RGBA{0, 0, 255, 255}, false)

	// Bad Code
	default:
		vector.DrawFilledRect(screen, float32(x), float32(y), tileSize, tileSize, color.RGBA{255, 0, 0, 255}, false)
	}
}

func (*Game) Layout(int, int) (int, int) {
	return screenWidth, screenHeight
}

func newGame() Game {
	g := Game{}

	g.font = &FontHandler{}
	g.font.loadFonts()

	g.image = &ImageHandler{}
	g.image.initImages()

	g.initUI()

	g.level = newLevel(16, 16)

	g.cameraSpeed = 5

	g.curTileType = getRelevantCodeText(g.curCodeSelection)

	return g
}
