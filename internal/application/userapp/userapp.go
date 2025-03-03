package userapp

import (
	"context"

	"github.com/Sergey-Polishchenko/go-post-flow/internal/core/id"
	"github.com/Sergey-Polishchenko/go-post-flow/internal/core/user"
	"github.com/Sergey-Polishchenko/go-post-flow/internal/pkg/logging"
)

// UserRepository defines persistence operations for Users.
type UserRepository interface {
	Save(ctx context.Context, user *user.User) error
	Remove(ctx context.Context, id id.Identifier) error
	GetByID(ctx context.Context, id id.Identifier) (*user.User, error)
}

type UserApp struct {
	repo   UserRepository
	logger logging.Logger
}

func New(repo UserRepository, logger logging.Logger) *UserApp {
	return &UserApp{repo: repo, logger: logger}
}
