package validation

import "unicode/utf8"

// LengthValidator validates string length constraints.
type LengthValidator struct {
	Min, Max int
}

// NewLengthValidator creates a validator with length bounds.
func NewLengthValidator(min, max int) LengthValidator {
	return LengthValidator{Min: min, Max: max}
}

func (v LengthValidator) Validate(s string) error {
	if utf8.RuneCountInString(s) == 0 {
		return ErrEmpty
	}
	if utf8.RuneCountInString(s) < v.Min {
		return ErrTooShort
	}
	if utf8.RuneCountInString(s) > v.Max {
		return ErrTooLong
	}
	return nil
}
