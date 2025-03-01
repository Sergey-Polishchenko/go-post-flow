package userapp

import (
	"context"
	"log"

	"github.com/Sergey-Polishchenko/go-post-flow/internal/core/user"
)

func (app *UserApp) GetUser(ctx context.Context, id string) (*user.User, error) {
	if id == "" {
		return nil, ErrInvalidInput
	}

	user, err := app.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	log.Printf("User(id: %s) getted", id)

	return user, nil
}
