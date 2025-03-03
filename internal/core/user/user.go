// Package user implements core user domain models and operations.
package user

type Identifier interface {
	String() string
}

// User represents system user account.
type User struct {
	ID   Identifier
	name UserName
}

// New creates validated User instance.
func New(id Identifier, name UserName) (*User, error) {
	if err := name.IsValid(); err != nil {
		return nil, &InvalidUsernameError{err}
	}

	return &User{ID: id, name: name}, nil
}

// Name returns the user's name(Read-only).
func (user *User) Name() UserName {
	return user.name
}
