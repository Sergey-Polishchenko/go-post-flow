package userapp

import (
	"context"
)

func (app *UserApp) RemoveUser(ctx context.Context, id string) error {
	err := app.repo.Remove(ctx, id)
	if err != nil {
		app.logger.Error("Failed to remove user from repo", "error", err)
		return err
	}

	app.logger.Info("User got", "id", id)

	return nil
}
