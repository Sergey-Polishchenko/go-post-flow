package user

import (
	"errors"
	"fmt"
)

// InvalidUsernameError is returned when attempting to create a User with an invalid name.
// Contains detail field for an exception, which brings the error
type InvalidUsernameError struct {
	detail error
}

func (err *InvalidUsernameError) Error() string {
	return fmt.Sprintf("%v: %v", ErrInvalidUsername, err.detail)
}

func (err *InvalidUsernameError) Unwrap() error {
	return err.detail
}

var ErrInvalidUsername = errors.New("invalid username")
