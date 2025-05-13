package main

import (
	"fmt"
	"image/color"
	"null/stagui"
	"time"

	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

func (g *Game) initUI() {
	g.title = &stagui.Page{
		Title:   "*NULL",
		BgColor: color.Black,
	}

	titleText := [4][2]string{
		{"start", "Start"},
		{"credits", "Credits"},
		{"instructions", "Instructions"},
		{"quit", "Quit"},
	}

	for i := range len(titleText) {
		g.title.Buttons = append(g.title.Buttons, &stagui.Button{
			X:         float32(halfWidth - 450/2),
			Y:         float32(160 + i*80),
			W:         450,
			H:         75,
			Enabled:   true,
			Name:      titleText[i][0],
			Text:      titleText[i][1],
			FontSize:  50,
			BgColor:   buttonColor,
			TextColor: color.White,
		})
	}

	g.title.Buttons = append(g.title.Buttons, &stagui.Button{
		X:         10,
		Y:         10,
		W:         50,
		H:         50,
		Name:      "settings",
		Enabled:   true,
		BgImg:     g.image.GetImage("settings"),
		TextColor: color.Black,
	})

	g.settings = &stagui.Page{
		Title:   "Settings",
		BgColor: color.Black,
		Buttons: []*stagui.Button{
			{
				X:            halfWidth - 150,
				Y:            150,
				W:            300,
				H:            100,
				Enabled:      true,
				Name:         "mute",
				Text:         "Mute",
				FontSize:     72,
				BgColor:      buttonColor,
				PressedColor: color.RGBA{32, 16, 16, 255},
				TextColor:    color.White,
			},
			{
				X:         halfWidth - 150,
				Y:         470,
				W:         300,
				H:         100,
				Enabled:   true,
				Name:      "back",
				Text:      "Back",
				FontSize:  72,
				BgColor:   buttonColor,
				TextColor: color.White,
			},
		},
		Sliders: []*stagui.Slider{
			{
				X:           halfWidth - 150,
				Y:           260,
				W:           300,
				H:           60,
				Value:       1,
				Name:        "volume",
				LineColor:   color.RGBA{32, 32, 32, 255},
				SliderColor: color.RGBA{128, 128, 128, 255},
			},
			{
				X:           halfWidth - 150,
				Y:           330,
				W:           300,
				H:           60,
				Value:       1,
				Name:        "effectVolume",
				LineColor:   buttonColor,
				SliderColor: color.RGBA{128, 128, 128, 255},
			},
		},
		Text: []*stagui.StaticText{
			{
				Text:  "Music",
				X:     halfWidth - 150,
				Y:     250,
				Color: color.White,
				Font:  g.font.GetFont("default"),
				Size:  50,
				Align: text.AlignEnd,
			},
			{
				Text:  "Effects",
				X:     halfWidth - 150,
				Y:     320,
				Color: color.White,
				Font:  g.font.GetFont("default"),
				Size:  50,
				Align: text.AlignEnd,
			},
		},
	}

	g.credits = &stagui.Page{
		Title:   "A Game By:",
		BgColor: color.Black,
		Buttons: []*stagui.Button{
			{
				X:         halfWidth - 200,
				Y:         400,
				W:         400,
				H:         75,
				Name:      "back",
				Enabled:   true,
				Text:      "Back",
				FontSize:  50,
				BgColor:   buttonColor,
				TextColor: color.White,
			},
		},
		Text: []*stagui.StaticText{
			{
				Text:  "Chig Beef",
				X:     halfWidth,
				Y:     150,
				Color: color.White,
				Font:  g.font.GetFont("default"),
				Size:  72,
				Align: text.AlignCenter,
			},
		},
	}

	creditsText := [3]string{
		"Lead Programmer and",
		"Game Director of",
		"STAG",
	}

	for i := range 3 {
		g.credits.Text = append(g.credits.Text, &stagui.StaticText{
			Text:  creditsText[i],
			X:     halfWidth,
			Y:     float64(225 + i*50),
			Color: color.White,
			Font:  g.font.GetFont("default"),
			Size:  50,
			Align: text.AlignCenter,
		})
	}

	g.instructions = &stagui.Page{
		Title:   "Instructions",
		BgColor: color.Black,
		Buttons: []*stagui.Button{
			{
				X:         halfWidth - 200,
				Y:         500,
				W:         400,
				H:         75,
				Name:      "back",
				Enabled:   true,
				Text:      "Back",
				FontSize:  50,
				BgColor:   buttonColor,
				TextColor: color.White,
			},
		},
	}

	instructionsText := [4]string{
		"Press `s` while holding ` to screenshot.",
		"Press `f` while holding ` to enter or exit fullscreen.",
		"WASD to move (and jump)",
		"Aim and click with mouse to dereference pointers",
	}

	for i := range len(instructionsText) {
		g.instructions.Text = append(g.instructions.Text, &stagui.StaticText{
			Text:  instructionsText[i],
			X:     halfWidth,
			Y:     float64(125 + i*50),
			Color: color.White,
			Font:  g.font.GetFont("default"),
			Size:  50,
			Align: text.AlignCenter,
		})
	}

	g.win = &stagui.Page{
		Title:   "You Won!",
		BgColor: color.Black,

		Buttons: []*stagui.Button{
			{
				X:         halfWidth - 150,
				Y:         360,
				W:         300,
				H:         100,
				Enabled:   true,
				Name:      "save",
				Text:      "Save",
				FontSize:  72,
				BgColor:   buttonColor,
				TextColor: color.White,
			},
			{
				X:         halfWidth - 150,
				Y:         470,
				W:         300,
				H:         100,
				Enabled:   true,
				Name:      "back",
				Text:      "Back",
				FontSize:  72,
				BgColor:   buttonColor,
				TextColor: color.White,
			},
		},
	}
}

func (g *Game) createWinPage(record uint64) {
	g.win.Text = []*stagui.StaticText{
		{
			Text:  "Your Time: " + fmt.Sprint(time.Duration(runTime*1000/60)*time.Millisecond),
			X:     halfWidth,
			Y:     150,
			Color: color.White,
			Font:  g.font.GetFont("default"),
			Size:  50,
			Align: text.AlignCenter,
		},

		{
			Text:  "Best Time: " + fmt.Sprint(time.Duration(record*1000/60)*time.Millisecond),
			X:     halfWidth,
			Y:     200,
			Color: color.White,
			Font:  g.font.GetFont("default"),
			Size:  50,
			Align: text.AlignCenter,
		},
	}
}
