package login

import (
	"errors"
)

var (
	ErrNoLocalSessionFile       = errors.New("no local session found")
	ErrUnrecognizedTokenContext = errors.New("unrecognized token context format")
)
