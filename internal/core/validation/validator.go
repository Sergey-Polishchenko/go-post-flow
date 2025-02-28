// Package validation provides string validation utilities.
package validation

// StringValidator defines validation contract for string inputs
type StringValidator interface {
	Validate(s string) error
}
