package main

type Phase byte

const (
	TITLE Phase = iota
	SETTINGS
	CREDITS
	INSTRUCTIONS
	GAME
	WIN
	REPLAY
)
