package post

import "context"

// PostRepository defines persistence operations for Posts.
type PostRepository interface {
	Create(ctx context.Context, authorID string, title string, content string) (*Post, error)
	Remove(ctx context.Context, id string) error
	GetByID(ctx context.Context, id string) (*Post, error)
	List(ctx context.Context, offset int, limit int) ([]*Post, error)
}
