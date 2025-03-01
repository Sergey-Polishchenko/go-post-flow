package commentapp

import (
	"context"
	"log"

	"github.com/Sergey-Polishchenko/go-post-flow/internal/core/comment"
)

func (app *CommentApp) CreateComment(
	ctx context.Context,
	authorId string,
	text string,
) (*comment.Comment, error) {
	comment, err := app.repo.Create(ctx, authorId, text)
	if err != nil {
		return nil, err
	}

	log.Printf("Comment(id: %s) created", comment.Id())

	return comment, nil
}
