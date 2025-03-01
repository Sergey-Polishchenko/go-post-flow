package postapp

import (
	"context"
	"log"
)

func (app *PostApp) RemovePost(ctx context.Context, id string) error {
	if id == "" {
		return ErrInvalidInput
	}

	err := app.repo.Remove(ctx, id)
	if err != nil {
		return err
	}

	log.Printf("Post(id: %s) removed", id)

	return nil
}
