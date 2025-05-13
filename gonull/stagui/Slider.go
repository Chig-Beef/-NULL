package stagui

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Slider struct {
	X float32
	Y float32
	W float32
	H float32

	Value float64

	Name    string // Used to identify the slider
	Pressed bool

	LineColor   color.Color
	SliderColor color.Color
}

func (s *Slider) Draw(screen *ebiten.Image) {
	// Bar
	vector.DrawFilledRect(screen, s.X, s.Y+s.H/4, s.W, s.H/2, s.LineColor, false)

	// Slidy thing
	vector.DrawFilledRect(screen, s.X+float32(s.Value)*(s.W-20), s.Y, 20, s.H, s.SliderColor, false)
}

func (s *Slider) CheckCollide(x, y float32) bool {
	return s.X <= x && x <= s.X+s.W &&
		s.Y <= y && y <= s.Y+s.H
}

func (s *Slider) Update(curMousePos [2]int) bool {
	x := float32(curMousePos[0])
	y := float32(curMousePos[1])

	if !ebiten.IsMouseButtonPressed(ebiten.MouseButton0) {
		return false
	}

	if !s.CheckCollide(x, y) {
		return false
	}

	dx := x - s.X

	s.Value = float64(dx / (s.W - 20))
	if s.Value > 1 {
		s.Value = 1
	}

	// Changed
	return true
}
