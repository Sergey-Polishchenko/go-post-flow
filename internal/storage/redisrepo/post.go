package redisrepo

import (
	"context"

	"github.com/redis/go-redis/v9"

	"github.com/Sergey-Polishchenko/go-post-flow/internal/application/postapp"
	"github.com/Sergey-Polishchenko/go-post-flow/internal/core/post"
)

var _ postapp.PostRepository = (*PostRepository)(nil)

type PostRepository struct {
	client *redis.Client
}

func (repo *PostRepository) Create(
	ctx context.Context,
	authorID string,
	title string,
	content string,
) (*post.Post, error) {
	return nil, nil
}

func (repo *PostRepository) Remove(ctx context.Context, id string) error {
	return nil
}

func (repo *PostRepository) GetByID(ctx context.Context, id string) (*post.Post, error) {
	return nil, nil
}

func (repo *PostRepository) List(ctx context.Context, offset int, limit int) ([]*post.Post, error) {
	return []*post.Post{}, nil
}
