package main

type TileCode byte

const (
	NO_TILE TileCode = iota
	BLOCK
	NULL_SPAWN
	PLAYER_START
	UPWARD_BOOST
	END_LEVEL
	NUM_TILE_CODES
)

var tileSolid = [NUM_TILE_CODES]bool{
	NO_TILE:      false,
	BLOCK:        true,
	NULL_SPAWN:   false,
	PLAYER_START: false,
	UPWARD_BOOST: true,
	END_LEVEL:    true,
}

func (code TileCode) solid() bool {
	// Bad code
	if code >= NUM_TILE_CODES {
		return true
	}

	return tileSolid[code]
}

var tileImage = [NUM_TILE_CODES]string{
	NO_TILE:      "",
	BLOCK:        "block",
	NULL_SPAWN:   "",
	PLAYER_START: "",
	UPWARD_BOOST: "upwardBoost",
	END_LEVEL:    "",
}

func (code TileCode) image() string {
	// Bad code
	if code >= NUM_TILE_CODES {
		return ""
	}

	return tileImage[code]
}

func (code TileCode) horBoost() float64 {
	switch code {
	case UPWARD_BOOST:
		return 0
	default:
		return 0
	}
}

func (code TileCode) verBoost() float64 {
	switch code {
	case UPWARD_BOOST:
		return -5
	default:
		return 0
	}
}
