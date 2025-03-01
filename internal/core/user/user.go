// Package user implements core user domain models and operations.
package user

// User represents system user account.
type User struct {
	id   string
	name UserName
}

// New creates validated User instance.
func New(id string, name UserName) (*User, error) {
	if id == "" {
		return nil, ErrNilId
	}

	if err := name.IsValid(); err != nil {
		return nil, &InvalidUsernameError{err}
	}

	return &User{id: id, name: name}, nil
}

// ID returns the user's id(Read-only).
func (user *User) ID() string {
	return user.id
}

// Name returns the user's name(Read-only).
func (user *User) Name() UserName {
	return user.name
}
