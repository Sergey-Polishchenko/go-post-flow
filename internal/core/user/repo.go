package user

import "context"

// UserRepository defines persistence operations for Users.
type UserRepository interface {
	Create(ctx context.Context, user *User) error
	Remove(ctx context.Context, id string) error
	GetById(ctx context.Context, id string) (*User, error)
}
