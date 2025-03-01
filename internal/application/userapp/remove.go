package userapp

import (
	"context"
	"log"
)

func (app *UserApp) RemoveUser(ctx context.Context, id string) error {
	if id == "" {
		return ErrInvalidInput
	}

	err := app.repo.Remove(ctx, id)
	if err != nil {
		return err
	}

	log.Printf("User(id: %s) removed", id)

	return nil
}
