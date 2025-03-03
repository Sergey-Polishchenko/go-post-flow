package userapp

import (
	"context"

	"github.com/Sergey-Polishchenko/go-post-flow/internal/core/id"
)

func (app *UserApp) GetUser(ctx context.Context, userID string) (UserDTO, error) {
	user, err := app.repo.GetByID(ctx, id.ID(userID))
	if err != nil {
		app.logger.Error("Failed to load user from repo", "error", err)
		return UserDTO{}, err
	}

	// TODO: business logic for User entity.

	app.logger.Info("User got", "id", user.ID.String())

	return UserDTO{
		ID:   user.ID.String(),
		Name: user.Name().String(),
	}, nil
}
