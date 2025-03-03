package comment

import (
	"github.com/Sergey-Polishchenko/go-post-flow/internal/core/id"
	"github.com/Sergey-Polishchenko/go-post-flow/internal/core/user"
)

// Comment represents a comment in the system.
type Comment struct {
	ID     id.Identifier
	Author *user.User
	text   CommentText
}

// New creates a validated Comment instance.
func New(id id.Identifier, author *user.User, text CommentText) (*Comment, error) {
	if author == nil {
		return nil, ErrNilAuthor
	}

	if err := text.IsValid(); err != nil {
		return nil, &InvalidCommentTextError{err}
	}

	return &Comment{ID: id, Author: author, text: text}, nil
}

// Text returns the comment's text (read-only).
func (c *Comment) Text() CommentText {
	return c.text
}
