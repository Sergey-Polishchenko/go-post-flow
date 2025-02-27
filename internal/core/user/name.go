package user

import (
	"github.com/Sergey-Polishchenko/go-post-flow/internal/core/validation"
)

// UserName represents a user's name.
// It must adhere to the following rules:
//   - Length: 1-100 characters (checked via IsValid()).
//   - Allowed characters: UTF-8 (no control characters).
type UserName string

// IsValid checks the validity of the name.
//
// Returns:
//   - true: If the name length is between 1-100 characters.
//   - false: Otherwise.
func (name UserName) IsValid() bool {
	return validation.NewLengthValidator(1, 100).Validate(string(name))
}
