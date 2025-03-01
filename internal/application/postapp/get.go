package postapp

import (
	"context"
	"log"

	"github.com/Sergey-Polishchenko/go-post-flow/internal/core/post"
)

func (app *PostApp) GetPost(ctx context.Context, id string) (*post.Post, error) {
	if id == "" {
		return nil, ErrInvalidInput
	}

	post, err := app.repo.GetById(ctx, id)
	if err != nil {
		return nil, err
	}

	log.Printf("Post(id: %s) getted", id)

	return post, nil
}
