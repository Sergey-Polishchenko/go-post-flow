package commentapp

import (
	"context"
	"log"

	"github.com/Sergey-Polishchenko/go-post-flow/internal/core/comment"
)

func (app *CommentApp) GetReplies(ctx context.Context, id string) ([]*comment.Comment, error) {
	if id == "" {
		return nil, ErrInvalidInput
	}

	comment, err := app.repo.GetReplies(ctx, id)
	if err != nil {
		return nil, err
	}

	log.Printf("Replies of Comment(id: %s) getted", id)

	return comment, nil
}
