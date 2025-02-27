package validation

import "unicode/utf8"

// LengthValidator validates strings based on minimum and maximum length constraints.
// It implements the StringValidator interface.
//
// Fields:
//   - Min: The minimum allowed length of the string (inclusive).
//   - Max: The maximum allowed length of the string (inclusive).
type LengthValidator struct {
	Min, Max int
}

// NewLengthValidator creates a new LengthValidator with the specified minimum and maximum length limits.
//
// Parameters:
//   - min: The minimum allowed length of the string (inclusive).
//   - max: The maximum allowed length of the string (inclusive).
//
// Returns:
//   - LengthValidator: A validator instance configured with the provided length bounds.
func NewLengthValidator(min, max int) LengthValidator {
	return LengthValidator{Min: min, Max: max}
}

// Validate checks if the given string's length falls within the specified bounds.
//
// Parameters:
//   - s: The string to validate.
//
// Returns:
//   - bool: True if the string length is between Min and Max (inclusive), false otherwise.
//
// Notes:
//   - The length is calculated using UTF-8 rune counting, ensuring accurate validation for multi-byte characters.
func (v LengthValidator) Validate(s string) bool {
	return utf8.RuneCountInString(s) >= v.Min &&
		utf8.RuneCountInString(s) <= v.Max
}
