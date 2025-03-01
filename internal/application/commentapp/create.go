package commentapp

import (
	"context"
	"log"

	"github.com/Sergey-Polishchenko/go-post-flow/internal/core/comment"
)

func (app *CommentApp) CreateComment(
	ctx context.Context,
	authorID string,
	text string,
) (*comment.Comment, error) {
	comment, err := app.repo.Create(ctx, authorID, text)
	if err != nil {
		return nil, err
	}

	log.Printf("Comment(id: %s) created", comment.ID())

	return comment, nil
}
