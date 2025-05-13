package main

import (
	"null/stagerror"
	"os"
	"strconv"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
)

type Level struct {
	data [][]*Tile
}

func (l *Level) update(g *Game) {
	// Nothing to update (removed nulls)
}

func (l *Level) draw(screen *ebiten.Image, ih *ImageHandler, fh *FontHandler, yOffset float64) {
	for y, row := range l.data {
		for x, tile := range row {
			tile.draw(screen, ih, fh, float64(x*tileSize), float64(y*tileSize)-yOffset)
		}
	}
}

func (g *Game) loadLevel(fName string) (*Level, *stagerror.Error) {
	l := Level{}

	f, err := os.ReadFile("assets/levels/" + fName + ".txt")
	if err != nil {
		return nil, stagerror.New(67, "Couldn't read level file")
	}

	rows := strings.Split(string(f), "\r\n")
	// In case the file gets windowed
	if len(rows) == 1 {
		rows = strings.Split(string(f), "\n")
	}

	for y, row := range rows {
		l.data = append(l.data, []*Tile{})

		cols := strings.Split(row, ",")

		for x, col := range cols {
			code, err := strconv.Atoi(col)
			if err != nil {
				return nil, stagerror.New(68, "Bad number in level file")
			}

			if code < 0 || code > 255 {
				return nil, stagerror.New(69, "tile codes must be in the range 0 to 255 inclusive")
			}

			bCode := TileCode(code)

			switch bCode {
			case PLAYER_START:
				g.player.x = float64(x * tileSize)
				g.player.y = float64(y * tileSize)
			case NULL_SPAWN:
				// Don't crash
				bCode = NO_TILE
			}

			var img *ebiten.Image
			imgName := bCode.image()
			if imgName == "" {
				img = nil
			} else {
				img = g.image.GetImage(imgName)
			}

			t := Tile{
				code:  bCode,
				solid: bCode.solid(),
				bx:    bCode.horBoost(),
				by:    bCode.verBoost(),
				img:   img,
			}

			l.data[y] = append(l.data[y], &t)
		}
	}

	return &l, nil
}

func (l Level) getPlayerPosition() (float64, float64) {
	for y, row := range l.data {
		for x, col := range row {
			if col.code == 24 {
				return float64(x * tileSize), float64(y * tileSize)
			}
		}
	}

	return 0, 0
}
