package post

import "context"

// PostRepository defines persistence operations for Posts.
type PostRepository interface {
	Create(ctx context.Context, post *Post) error
	Remove(ctx context.Context, id string) error
	GetById(ctx context.Context, id string) (*Post, error)
	List(ctx context.Context, offset int, limit int) ([]*Post, error)
}
