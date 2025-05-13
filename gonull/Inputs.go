package main

import (
	"null/stagerror"
	"os"
	"strings"
)

type InputCode byte

const (
	LEFT InputCode = iota
	RIGHT
	JUMP
	BOOST
	NUM_INPUT_CODES
)

var inputCodeNames [NUM_INPUT_CODES]string = [NUM_INPUT_CODES]string{
	"LEFT",
	"RIGHT",
	"JUMP",
	"BOOST",
}

func (ic InputCode) String() string {
	if ic >= NUM_INPUT_CODES {
		return ""
	}
	return inputCodeNames[ic]
}

var runInputs = [][]InputCode{}

func saveInputs() {
	out := []byte{}

	for _, frame := range runInputs {
		for _, i := range frame {
			out = append(out, []byte(i.String())...)
			out = append(out, ' ')
		}
		out = append(out, '\n')
	}

	_ = os.WriteFile("run.nulldemo", out, 0644)
}

func loadInputs(fileName string) *stagerror.Error {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return stagerror.New(1, "Couldn't read nulldemo file")
	}

	// Windows
	lines := strings.Split(string(data), "\r\n")
	if len(lines) < 2 {
		lines = strings.Split(string(data), "\n")
	}

	for _, line := range lines {
		runInputs = append(runInputs, []InputCode{})

		if line == "" {
			continue
		}

		inputs := strings.Split(line, " ")

		for n := 0; n < len(inputs); n++ {
			i := inputs[n]
			var b InputCode
			switch i {
			case "LEFT":
				b = LEFT
			case "RIGHT":
				b = RIGHT
			case "JUMP":
				b = JUMP
			case "BOOST":
				b = BOOST
			case "":
				continue
			}
			runInputs[len(runInputs)-1] = append(runInputs[len(runInputs)-1], b)
		}
	}

	return nil
}
