package commentapp

import (
	"context"

	"github.com/Sergey-Polishchenko/go-post-flow/internal/core/comment"
)

// CommentRepository defines persistence operations for Comments.
type CommentRepository interface {
	Create(ctx context.Context, authorID string, text string) (*comment.Comment, error)
	Remove(ctx context.Context, id string) error
	GetByID(ctx context.Context, id string) (*comment.Comment, error)
	GetReplies(ctx context.Context, id string) ([]*comment.Comment, error)
	GetPostReplies(ctx context.Context, postID string) ([]*comment.Comment, error)
}

type CommentApp struct {
	repo CommentRepository
}

func New(repo CommentRepository) *CommentApp {
	return &CommentApp{repo: repo}
}
