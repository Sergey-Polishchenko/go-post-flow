package comment

import "context"

// CommentRepository defines persistence operations for Comments.
type CommentRepository interface {
	Create(ctx context.Context, authorID string, text string) (*Comment, error)
	Remove(ctx context.Context, id string) error
	GetByID(ctx context.Context, id string) (*Comment, error)
	GetReplies(ctx context.Context, id string) ([]*Comment, error)
	GetPostReplies(ctx context.Context, postID string) ([]*Comment, error)
}
