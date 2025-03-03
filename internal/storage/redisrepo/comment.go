package redisrepo

import (
	"context"

	"github.com/redis/go-redis/v9"

	"github.com/Sergey-Polishchenko/go-post-flow/internal/application/commentapp"
	"github.com/Sergey-Polishchenko/go-post-flow/internal/core/comment"
)

var _ commentapp.CommentRepository = (*CommentRepository)(nil)

type CommentRepository struct {
	client *redis.Client
}

func (repo *CommentRepository) Create(
	ctx context.Context,
	authorID string,
	text string,
) (*comment.Comment, error) {
	return nil, nil
}

func (repo *CommentRepository) Remove(ctx context.Context, id string) error {
	return nil
}

func (repo *CommentRepository) GetByID(ctx context.Context, id string) (*comment.Comment, error) {
	return nil, nil
}

func (repo *CommentRepository) GetReplies(
	ctx context.Context,
	id string,
) ([]*comment.Comment, error) {
	return []*comment.Comment{}, nil
}

func (repo *CommentRepository) GetPostReplies(
	ctx context.Context,
	postID string,
) ([]*comment.Comment, error) {
	return []*comment.Comment{}, nil
}
