package stagui

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Button struct {
	X float32
	Y float32
	W float32
	H float32

	Name    string // Used to identify the button
	Enabled bool
	Pressed bool

	Text          string
	FontSize      float64
	BgColor       color.Color
	BgImg         *ebiten.Image
	PressedColor  color.Color
	DisabledColor color.Color
	TextColor     color.Color
}

func (b *Button) Draw(screen *ebiten.Image, ih ImageHandler, fh FontHandler) {
	if b.BgImg == nil {
		b.drawAsSolidColor(screen)
	} else {
		ih.DrawImage(screen, b.BgImg, float64(b.X), float64(b.Y), float64(b.W), float64(b.H), 0, &ebiten.DrawImageOptions{})
	}

	op := text.DrawOptions{}
	op.PrimaryAlign = text.AlignCenter
	op.ColorScale.ScaleWithColor(b.TextColor)
	fh.DrawText(screen, b.Text, b.FontSize, float64(b.X+b.W/2), float64(b.Y), fh.GetFont("button"), &op)
}

func (b *Button) drawAsSolidColor(screen *ebiten.Image) {
	if !b.Enabled {
		vector.DrawFilledRect(screen, b.X, b.Y, b.W, b.H, b.DisabledColor, false)
		return
	}

	if b.Pressed {
		vector.DrawFilledRect(screen, b.X, b.Y, b.W, b.H, b.PressedColor, false)
	} else {
		vector.DrawFilledRect(screen, b.X, b.Y, b.W, b.H, b.BgColor, false)
	}
}

func (b Button) CheckClick(x, y float32) bool {
	return b.Enabled &&
		b.X <= x && x <= b.X+b.W &&
		b.Y <= y && y <= b.Y+b.H
}
