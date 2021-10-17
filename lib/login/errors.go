package login

import (
	"errors"
)

var (
	ErrNoLocalTokenContext      = errors.New("no local token context")
	ErrUnrecognizedTokenContext = errors.New("unrecognized token context format")
)
