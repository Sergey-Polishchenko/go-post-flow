package commentapp

import (
	"context"
	"log"
)

func (app *CommentApp) RemoveComment(ctx context.Context, id string) error {
	if id == "" {
		return ErrInvalidInput
	}

	err := app.repo.Remove(ctx, id)
	if err != nil {
		return err
	}

	log.Printf("Comment(id: %s) removed", id)

	return nil
}
