// Package post implements core post domain models and operations.
package post

import (
	"github.com/Sergey-Polishchenko/go-post-flow/internal/core/id"
	"github.com/Sergey-Polishchenko/go-post-flow/internal/core/user"
)

// Post represents a post in the system.
type Post struct {
	ID      id.Identifier
	title   PostTitle
	content PostContent
	Author  *user.User
}

// New creates a validated Post instance.
func New(id id.Identifier, author *user.User, title PostTitle, content PostContent) (*Post, error) {
	if author == nil {
		return nil, ErrNilAuthor
	}

	if err := title.IsValid(); err != nil {
		return nil, &InvalidPostTitleError{err}
	}

	if err := content.IsValid(); err != nil {
		return nil, &InvalidPostContentError{err}
	}

	return &Post{
		ID:      id,
		title:   title,
		content: content,
		Author:  author,
	}, nil
}

// Title returns the post's title (read-only).
func (post *Post) Title() PostTitle {
	return post.title
}

// Content returns the post's content (read-only).
func (post *Post) Content() PostContent {
	return post.content
}
