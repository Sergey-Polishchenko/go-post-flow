package postapp

import (
	"context"
	"log"

	"github.com/Sergey-Polishchenko/go-post-flow/internal/core/post"
)

func (app *PostApp) CreatePost(
	ctx context.Context,
	authorId string,
	title string,
	content string,
) (*post.Post, error) {
	post, err := app.repo.Create(ctx, authorId, title, content)
	if err != nil {
		return nil, err
	}

	log.Printf("Post(id: %s) created", post.Id())

	return post, nil
}
