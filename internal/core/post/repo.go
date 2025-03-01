package post

import "context"

// PostRepository defines persistence operations for Posts.
type PostRepository interface {
	Create(ctx context.Context, authorId string, title string, content string) (*Post, error)
	Remove(ctx context.Context, id string) error
	GetById(ctx context.Context, id string) (*Post, error)
	List(ctx context.Context, offset int, limit int) ([]*Post, error)
}
