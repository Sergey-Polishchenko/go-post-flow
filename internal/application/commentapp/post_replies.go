package commentapp

import (
	"context"
	"log"

	"github.com/Sergey-Polishchenko/go-post-flow/internal/core/comment"
)

func (app *CommentApp) GetPostReplies(
	ctx context.Context,
	postID string,
) ([]*comment.Comment, error) {
	if postID == "" {
		return nil, ErrInvalidInput
	}

	comment, err := app.repo.GetReplies(ctx, postID)
	if err != nil {
		return nil, err
	}

	log.Printf("Replies of Post(id: %s) getted", postID)

	return comment, nil
}
