package validation

import "errors"

var (
	ErrEmpty    = errors.New("string is empty")
	ErrTooShort = errors.New("string is too short")
	ErrTooLong  = errors.New("string is too long")
)
