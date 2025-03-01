package commentapp

import (
	"context"
	"log"

	"github.com/Sergey-Polishchenko/go-post-flow/internal/core/comment"
)

func (app *CommentApp) GetComment(ctx context.Context, id string) (*comment.Comment, error) {
	if id == "" {
		return nil, ErrInvalidInput
	}

	comment, err := app.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	log.Printf("Comment(id: %s) getted", id)

	return comment, nil
}
