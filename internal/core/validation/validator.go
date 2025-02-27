// Package validation provides utilities for validating strings based on specific rules.
// It includes interfaces and implementations for string validation, such as length constraints.
//
// Key Components:
//   - StringValidator: An interface for defining custom string validation logic.
//   - LengthValidator: A concrete implementation of StringValidator that validates strings based on minimum and maximum length limits.
//   - NewLengthValidator: A constructor for creating LengthValidator instances with specified length bounds.
//
// Notes:
//   - LengthValidator uses UTF-8 rune counting, ensuring accurate validation for multi-byte characters.
package validation

// StringValidator defines the contract for validating strings.
// Implementations of this interface can provide custom validation logic.
type StringValidator interface {
	// Validate checks if the given string meets the validation criteria.
	//
	// Parameters:
	//   - s: The string to validate.
	//
	// Returns:
	//   - bool: True if the string is valid, false otherwise.
	Validate(s string) bool
}
