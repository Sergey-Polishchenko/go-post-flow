package comment

import "github.com/Sergey-Polishchenko/go-post-flow/internal/core/user"

// Comment represents a comment in the system.
type Comment struct {
	Author *user.User
	text   CommentText
}

// New creates a validated Comment instance.
func New(author *user.User, text CommentText) (*Comment, error) {
	if author == nil {
		return nil, ErrNilAuthor
	}

	if err := text.IsValid(); err != nil {
		return nil, &InvalidCommentTextError{err}
	}

	return &Comment{Author: author, text: text}, nil
}

// Text returns the comment's text (read-only).
func (c *Comment) Text() CommentText {
	return c.text
}
