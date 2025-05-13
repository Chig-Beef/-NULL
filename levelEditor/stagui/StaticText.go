package stagui

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type StaticText struct {
	Text string

	X float64
	Y float64

	Color color.Color
	Font  *text.GoTextFaceSource
	Size  float64
	Align text.Align
}

func (st *StaticText) Draw(screen *ebiten.Image) {
	op := text.DrawOptions{}
	op.PrimaryAlign = st.Align
	op.GeoM.Translate(st.X, st.Y)
	text.Draw(
		screen,
		st.Text,
		&text.GoTextFace{
			Source: st.Font,
			Size:   st.Size,
		},
		&op,
	)
}
