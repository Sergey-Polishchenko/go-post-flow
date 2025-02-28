package comment

import "github.com/Sergey-Polishchenko/go-post-flow/internal/core/validation"

// CommentText represents the text of a comment.
type CommentText string

// IsValid checks the validity of the comment text.
func (text CommentText) IsValid() error {
	return validation.NewLengthValidator(1, 2000).Validate(string(text))
}
