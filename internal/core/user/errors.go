package user

import "errors"

// ErrInvalidUsername is returned when attempting to create a User with an invalid name.
var ErrInvalidUsername = errors.New("invalid username")
