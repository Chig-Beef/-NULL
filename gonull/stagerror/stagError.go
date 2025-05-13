package stagerror

import (
	"os"
	"strconv"
)

type Error struct {
	Msg  string
	Code uint16 // Should be plenty of errors
}

func (serr *Error) String() string {
	return "\nSTAG done goofed\nError Code: " + strconv.Itoa(int(serr.Code)) + "\n\n" + serr.Msg
}

// Problem is we can't do anything when
// an error occurs
func (serr *Error) SaveToLog(isRelease bool) {
	if serr == nil {
		return
	}

	if serr.Code == 0 {
		return
	}

	// Make sure that when we hit an error
	// in development we know about it
	if !isRelease {
		panic(serr)
	}

	// However, if a user hits an error,
	// don't crash, we don't want it
	// ruining their experience. Instead,
	// we just log it for them to tell us
	// about

	data, err := os.ReadFile("errors.log")
	if err != nil {
		data = []byte{}
	}

	data = append(data, []byte(serr.String())...)

	err = os.WriteFile("errors.log", data, 0644)
}

func New(code uint16, msg string) *Error {
	return &Error{Code: code, Msg: msg}
}
