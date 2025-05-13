package stagui

import (
	"slices"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func handleKey() (string, ebiten.Key) {
	var text string

	keys := inpututil.AppendJustPressedKeys([]ebiten.Key{})
	if len(keys) == 0 {
		return "None", ebiten.Key(0)
	}

	k := keys[0]

	if slices.Contains(inpututil.AppendPressedKeys([]ebiten.Key{}), ebiten.KeyShift) {
		return handleKeyWithShift(k)
	}

	if slices.Contains(inpututil.AppendPressedKeys([]ebiten.Key{}), ebiten.KeyControl) {
		return "Ctrl", k
	}

	switch k {
	case ebiten.KeySlash:
		text = "/"
	case ebiten.KeyBackslash:
		text = "\\"
	case ebiten.KeyQuote:
		text = "'"
	case ebiten.KeyTab:
		text = "  "
	case ebiten.KeyBracketLeft:
		text = "["
	case ebiten.KeyBracketRight:
		text = "]"
	case ebiten.KeySpace:
		text = " "
	case ebiten.KeyEqual:
		text = "="
	case ebiten.KeyMinus:
		text = "-"
	case ebiten.KeySemicolon:
		text = ";"
	case ebiten.KeyComma:
		text = ","
	case ebiten.KeyPeriod:
		text = "."
	case ebiten.KeyEnter:
		text = "\n"
	case ebiten.Key0:
		text = "0"
	case ebiten.Key1:
		text = "1"
	case ebiten.Key2:
		text = "2"
	case ebiten.Key3:
		text = "3"
	case ebiten.Key4:
		text = "4"
	case ebiten.Key5:
		text = "5"
	case ebiten.Key6:
		text = "6"
	case ebiten.Key7:
		text = "7"
	case ebiten.Key8:
		text = "8"
	case ebiten.Key9:
		text = "9"
	case ebiten.KeyMetaLeft:
		text = "None"
	case ebiten.KeyMetaRight:
		text = "None"
	default:
		text = strings.ToLower(k.String())
	}

	return text, k
}

func handleKeyWithShift(k ebiten.Key) (string, ebiten.Key) {
	var text string
	switch k {
	case ebiten.KeySlash:
		text = "?"
	case ebiten.KeyBackslash:
		text = "|"
	case ebiten.KeyQuote:
		text = "\""
	case ebiten.KeyTab:
		text = "  "
	case ebiten.KeyBracketLeft:
		text = "{"
	case ebiten.KeyBracketRight:
		text = "}"
	case ebiten.KeySpace:
		text = " "
	case ebiten.KeyEqual:
		text = "+"
	case ebiten.KeyMinus:
		text = "_"
	case ebiten.KeySemicolon:
		text = ":"
	case ebiten.KeyEnter:
		text = "\n"
	case ebiten.KeyComma:
		text = "<"
	case ebiten.KeyPeriod:
		text = ">"
	case ebiten.Key0:
		text = ")"
	case ebiten.Key1:
		text = "!"
	case ebiten.Key2:
		text = "@"
	case ebiten.Key3:
		text = "#"
	case ebiten.Key4:
		text = "$"
	case ebiten.Key5:
		text = "%"
	case ebiten.Key6:
		text = "^"
	case ebiten.Key7:
		text = "&"
	case ebiten.Key8:
		text = "*"
	case ebiten.Key9:
		text = "("
	case ebiten.KeyShiftRight:
		text = "None"
	case ebiten.KeyShiftLeft:
		text = "None"
	case ebiten.KeyMetaLeft:
		text = "None"
	case ebiten.KeyMetaRight:
		text = "None"
	default:
		text = k.String()
	}

	return text, k
}
