package errors

import "errors"

var (
	InvalidLengthError = errors.New("invalid length")
)

var (
	NoMoreBytesError = errors.New("no more bytes to read")
)
