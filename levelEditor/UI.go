package main

import (
	"image/color"
	"nullEditor/stagui"
)

func (g *Game) initUI() {

	g.page = &stagui.Page{
		Title:  "",
		BgDraw: false,

		Buttons: []*stagui.Button{
			{
				X: 500,
				Y: 10,
				W: 100,
				H: 50,

				BgColor:   color.RGBA{48, 48, 48, 255},
				TextColor: color.White,
				FontSize:  16,

				Name: "save",
				Text: "save",

				Enabled: true,
			},
			{
				X: 610,
				Y: 10,
				W: 100,
				H: 50,

				BgColor:   color.RGBA{48, 48, 48, 255},
				TextColor: color.White,
				FontSize:  16,

				Name: "load",
				Text: "load",

				Enabled: true,
			},
			{
				X: 720,
				Y: 10,
				W: 100,
				H: 50,

				BgColor:   color.RGBA{48, 48, 48, 255},
				TextColor: color.White,
				FontSize:  16,

				Name: "codeUp",
				Text: ">",

				Enabled: true,
			},
			{
				X: 830,
				Y: 10,
				W: 100,
				H: 50,

				BgColor:   color.RGBA{48, 48, 48, 255},
				TextColor: color.White,
				FontSize:  16,

				Name: "codeDown",
				Text: "<",

				Enabled: true,
			},
			{
				X: 720,
				Y: 70,
				W: 100,
				H: 50,

				BgColor:   color.RGBA{48, 48, 48, 255},
				TextColor: color.White,
				FontSize:  16,

				Name: "codeUpBlock",
				Text: "=>",

				Enabled: true,
			},
			{
				X: 830,
				Y: 70,
				W: 100,
				H: 50,

				BgColor:   color.RGBA{48, 48, 48, 255},
				TextColor: color.White,
				FontSize:  16,

				Name: "codeDownBlock",
				Text: "<=",

				Enabled: true,
			},
			{
				X: 940,
				Y: 10,
				W: 100,
				H: 50,

				BgColor:   color.RGBA{48, 48, 48, 255},
				TextColor: color.White,
				FontSize:  16,

				Name: "clear",
				Text: "Clear",

				Enabled: true,
			},
			{
				X: 940,
				Y: 70,
				W: 100,
				H: 50,

				BgColor:   color.RGBA{48, 48, 48, 255},
				TextColor: color.White,
				FontSize:  16,

				Name: "addSection",
				Text: "Add Section",

				Enabled: true,
			},
		},

		InlineTextBoxes: []*stagui.InlineTextBox{
			{
				Name: "width",
				Text: "16",

				X: 500,
				Y: 70,
				W: 100,
				H: 50,

				BgColor:   color.RGBA{48, 48, 48, 255},
				TextColor: color.White,

				Active: false,
			},
			{
				Name: "height",
				Text: "16",

				X: 610,
				Y: 70,
				W: 100,
				H: 50,

				BgColor:   color.RGBA{48, 48, 48, 255},
				TextColor: color.White,

				Active: false,
			},
		},
	}
}
