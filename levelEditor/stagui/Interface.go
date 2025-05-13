package stagui

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type ImageHandler interface {
	DrawImage(*ebiten.Image, *ebiten.Image, float64, float64, float64, float64, float64, *ebiten.DrawImageOptions)
}

type FontHandler interface {
	DrawText(*ebiten.Image, string, float64, float64, float64, *text.GoTextFaceSource, *text.DrawOptions)
	GetFont(string) *text.GoTextFaceSource
}
