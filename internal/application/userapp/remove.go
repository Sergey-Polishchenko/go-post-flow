package userapp

import (
	"context"

	"github.com/Sergey-Polishchenko/go-post-flow/internal/core/id"
)

func (app *UserApp) RemoveUser(ctx context.Context, userID string) error {
	err := app.repo.Remove(ctx, id.ID(userID))
	if err != nil {
		app.logger.Error("Failed to remove user from repo", "error", err)
		return err
	}

	app.logger.Info("User got", "id", userID)

	return nil
}
