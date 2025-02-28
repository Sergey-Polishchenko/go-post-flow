// Package post implements core post domain models and operations.
package post

import "github.com/Sergey-Polishchenko/go-post-flow/internal/core/user"

// Post represents a post in the system.
type Post struct {
	id      string
	title   PostTitle
	content PostContent
	Author  *user.User
}

// New creates a validated Post instance.
func New(id string, author *user.User, title PostTitle, content PostContent) (*Post, error) {
	if id == "" {
		return nil, ErrNilId
	}

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
		id:      id,
		title:   title,
		content: content,
		Author:  author,
	}, nil
}

// Id returns the post's id (read-only).
func (post *Post) Id() string {
	return post.id
}

// Title returns the post's title (read-only).
func (post *Post) Title() PostTitle {
	return post.title
}

// Content returns the post's content (read-only).
func (post *Post) Content() PostContent {
	return post.content
}
