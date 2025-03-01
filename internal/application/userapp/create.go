package userapp

import (
	"context"
	"log"

	"github.com/Sergey-Polishchenko/go-post-flow/internal/core/user"
)

func (app *UserApp) CreateUser(ctx context.Context, name string) (*user.User, error) {
	user, err := app.repo.Create(ctx, name)
	if err != nil {
		return nil, err
	}

	log.Printf("User(id: %s) created", user.ID())

	return user, nil
}
