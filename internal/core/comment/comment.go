package comment

import "github.com/Sergey-Polishchenko/go-post-flow/internal/core/user"

// Comment represents a comment in the system.
type Comment struct {
	id     string
	Author *user.User
	text   CommentText
}

// New creates a validated Comment instance.
func New(id string, author *user.User, text CommentText) (*Comment, error) {
	if id == "" {
		return nil, ErrNilId
	}

	if author == nil {
		return nil, ErrNilAuthor
	}

	if err := text.IsValid(); err != nil {
		return nil, &InvalidCommentTextError{err}
	}

	return &Comment{id: id, Author: author, text: text}, nil
}

// ID returns the comment's id (read-only).
func (c *Comment) ID() string {
	return c.id
}

// Text returns the comment's text (read-only).
func (c *Comment) Text() CommentText {
	return c.text
}
