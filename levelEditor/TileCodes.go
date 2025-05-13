package main

type TileCode byte

const (
	NO_TILE TileCode = iota
	BLOCK
	NULL_SPAWN
	PLAYER_START
	UPWARD_BOOST
	LEVEL_END
)
