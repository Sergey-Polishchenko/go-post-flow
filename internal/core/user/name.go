package user

import (
	"github.com/Sergey-Polishchenko/go-post-flow/internal/core/validation"
)

const (
	MinUserNameLength = 1
	MaxUserNameLength = 100
)

var usernameValidator = validation.NewLengthValidator(MinUserNameLength, MaxUserNameLength)

// UserName represents a user's name.
type UserName string

// String return UserName string representation.
func (name UserName) String() string {
	return string(name)
}

// IsValid checks the validity of the name.
func (name UserName) IsValid() error {
	return usernameValidator.Validate(string(name))
}
