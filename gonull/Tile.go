package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Tile struct {
	// Boost
	bx, by float64

	img *ebiten.Image

	code  TileCode
	solid bool
}

func (t *Tile) draw(screen *ebiten.Image, ih *ImageHandler, fh *FontHandler, x, y float64) {
	if !t.solid {
		return
	}

	if t.img == nil {
		vector.DrawFilledRect(screen, float32(x), float32(y), tileSize, tileSize, color.RGBA{0, 0, 255, 0}, false)
		return
	}

	ih.DrawImage(screen, t.img, x, y, tileSize, tileSize, 0, &ebiten.DrawImageOptions{})
}
