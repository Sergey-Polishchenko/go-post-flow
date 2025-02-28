package comment

import "context"

// CommentRepository defines persistence operations for Comments.
type CommentRepository interface {
	Create(ctx context.Context, comment *Comment) error
	Remove(ctx context.Context, id string) error
	GetById(ctx context.Context, id string) (*Comment, error)
}
