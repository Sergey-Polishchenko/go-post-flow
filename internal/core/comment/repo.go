package comment

// CommentRepository defines persistence operations for Comments.
type CommentRepository interface {
	Create(comment *Comment) error
	Remove(id string) error
	GetById(id string) (*Comment, error)
}
