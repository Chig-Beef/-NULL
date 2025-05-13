package main

import (
	"bytes"
	"nullEditor/stagerror"

	"github.com/hajimehoshi/ebiten/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type FontHandler struct {
	fonts map[string]*text.GoTextFaceSource
}

func (fh *FontHandler) GetFont(name string) *text.GoTextFaceSource {
	f, ok := fh.fonts[name]
	if ok {
		return f
	}

	stagerror.New(49, "Couldn't find this font: "+name).SaveToLog(IS_RELEASE)

	// Will hard error if default doesn't exist
	return fh.fonts["default"]
}

func (fh *FontHandler) DrawText(screen *ebiten.Image, str string, size, x, y float64, drawFont *text.GoTextFaceSource, op *text.DrawOptions) {
	op.GeoM.Translate(x, y)
	text.Draw(
		screen,
		str,
		&text.GoTextFace{
			Source: drawFont,
			Size:   size,
		},
		op,
	)
}

func (fh *FontHandler) loadFonts() {
	fh.fonts = make(map[string]*text.GoTextFaceSource)

	s, err := text.NewGoTextFaceSource(bytes.NewReader(fonts.MPlus1pRegular_ttf))
	if err != nil {
		// Panic because it's vital
		stagerror.New(48, "Couldn't load fonts").SaveToLog(IS_RELEASE)
	}

	fh.fonts["default"] = s
	fh.fonts["button"] = s
	fh.fonts["textBox"] = s
}
