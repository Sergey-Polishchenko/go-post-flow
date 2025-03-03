// Package user implements core user domain models and operations.
package user

// User represents system user account.
type User struct {
	ID   UserID
	name UserName
}

// New creates validated User instance.
func New(name UserName) (*User, error) {
	if err := name.IsValid(); err != nil {
		return nil, &InvalidUsernameError{err}
	}

	return &User{ID: NewUserID(), name: name}, nil
}

// Name returns the user's name(Read-only).
func (user *User) Name() UserName {
	return user.name
}
