package main

import (
	"nullEditor/stagerror"
	"os"
	"strconv"
	"strings"
)

type Level struct {
	data [][]TileCode
}

func newLevel(width, height int) Level {
	l := Level{}

	l.data = make([][]TileCode, height)

	for row := range height {
		l.data[row] = make([]TileCode, width)
	}

	return l
}

func (l Level) save() error {
	levelData := l.stringify()

	f, err := os.Create("level.txt")
	if err != nil {
		return err
	}

	_, err = f.WriteString(levelData)
	if err != nil {
		return err
	}

	return nil
}

func loadLevel(g *Game) (Level, *stagerror.Error) {
	l := Level{}

	f, err := os.ReadFile("level.txt")
	if err != nil {
		return l, stagerror.New(0, "Couldn't read from level")
	}

	data := [][]TileCode{}

	rows := strings.Split(string(f), "\n")

	for y, row := range rows {
		cols := strings.Split(row, ",")

		data = append(data, []TileCode{})

		for _, col := range cols {
			tile, err := strconv.Atoi(col)
			if err != nil {
				return l, stagerror.New(0, "Bad number in level")
			}

			if tile < 0 || tile > 255 {
				return l, stagerror.New(0, "level data contained non-byte")
			}

			data[y] = append(data[y], TileCode(tile))
		}
	}

	l.data = data

	g.page.InlineTextBoxes[0].Text = strconv.Itoa(len(l.data[0]))
	g.page.InlineTextBoxes[1].Text = strconv.Itoa(len(l.data))

	return l, nil
}

func (l Level) stringify() string {
	output := ""

	for _, row := range l.data {
		for _, tile := range row {
			output += strconv.Itoa(int(tile)) + ","
		}
		output = output[:len(output)-1] + "\n"
	}
	output = output[:len(output)-1]

	return output
}

func (l *Level) resize(w, h int) {
	if h > len(l.data) {
		l.data = append(l.data, make([][]TileCode, h-len(l.data))...)
	} else {
		l.data = l.data[:h]
	}

	if w > len(l.data[0]) {
		for i := 0; i < len(l.data); i++ {
			l.data[i] = append(l.data[i], make([]TileCode, w-len(l.data[i]))...)
		}
	} else {
		for r := range l.data {
			l.data[r] = l.data[r][:w]
		}

		// newData := [][]byte{}

		// for r := range l.data {
		// 	newData = append(newData, []byte{})
		// 	for c := range l.data[r] {
		// 		newData[r] = append(newData[r], l.data[r][c])
		// 	}
		// }

		// l.data = newData
	}
}

func (l *Level) addSection() {
	newData := make([][]TileCode, len(l.data)+16)

	for i := range newData {
		newData[i] = make([]TileCode, len(l.data[0]))
	}

	for i := 16; i < len(newData); i++ {
		copy(newData[i], l.data[i-16])
	}

	l.data = newData
}
