package postapp

import (
	"context"

	"github.com/Sergey-Polishchenko/go-post-flow/internal/core/post"
)

// PostRepository defines persistence operations for Posts.
type PostRepository interface {
	Create(ctx context.Context, authorID string, title string, content string) (*post.Post, error)
	Remove(ctx context.Context, id string) error
	GetByID(ctx context.Context, id string) (*post.Post, error)
	List(ctx context.Context, offset int, limit int) ([]*post.Post, error)
}

type PostApp struct {
	repo PostRepository
}

func New(repo PostRepository) *PostApp {
	return &PostApp{repo: repo}
}
