package userapp

import (
	"context"

	"github.com/Sergey-Polishchenko/go-post-flow/internal/core/user"
)

// UserRepository defines persistence operations for Users.
type UserRepository interface {
	Create(ctx context.Context, name string) (*user.User, error)
	Remove(ctx context.Context, id string) error
	GetByID(ctx context.Context, id string) (*user.User, error)
}

type UserApp struct {
	repo UserRepository
}

func New(repo UserRepository) *UserApp {
	return &UserApp{repo: repo}
}
