package postapp

import (
	"context"
	"log"

	"github.com/Sergey-Polishchenko/go-post-flow/internal/core/post"
)

func (app *PostApp) GetList(ctx context.Context, offset, limit int) ([]*post.Post, error) {
	if offset < 0 || limit < 0 {
		return nil, ErrInvalidInput
	}

	list, err := app.repo.List(ctx, offset, limit)
	if err != nil {
		return nil, err
	}

	log.Printf("All Posts(from: %d, to: %d) getted", offset, offset+limit)

	return list, nil
}
