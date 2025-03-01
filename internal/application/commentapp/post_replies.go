package commentapp

import (
	"context"
	"log"

	"github.com/Sergey-Polishchenko/go-post-flow/internal/core/comment"
)

func (app *CommentApp) GetPostReplies(
	ctx context.Context,
	postId string,
) ([]*comment.Comment, error) {
	if postId == "" {
		return nil, ErrInvalidInput
	}

	comment, err := app.repo.GetReplies(ctx, postId)
	if err != nil {
		return nil, err
	}

	log.Printf("Replies of Post(id: %s) getted", postId)

	return comment, nil
}
