package user

import (
	"github.com/Sergey-Polishchenko/go-post-flow/internal/core/validation"
)

const (
	MinUserNameLength = 1
	MaxUserNameLength = 100
)

// UserName represents a user's name.
type UserName string

// IsValid checks the validity of the name.
func (name UserName) IsValid() error {
	return validation.NewLengthValidator(MinUserNameLength, MaxUserNameLength).
		Validate(string(name))
}
