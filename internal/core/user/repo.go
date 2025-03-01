package user

import "context"

// UserRepository defines persistence operations for Users.
type UserRepository interface {
	Create(ctx context.Context, name string) (*User, error)
	Remove(ctx context.Context, id string) error
	GetByID(ctx context.Context, id string) (*User, error)
}
