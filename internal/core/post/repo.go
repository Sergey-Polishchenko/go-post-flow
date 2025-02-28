package post

// PostRepository defines persistence operations for Posts.
type PostRepository interface {
	Create(post *Post) error
	Remove(id string) error
	GetById(id string) (*Post, error)
	List(offset int, limit int) ([]*Post, error)
}
