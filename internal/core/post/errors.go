package post

import (
	"errors"
	"fmt"
)

// InvalidPostTitleError is returned when attempting to create a Post with an invalid title.
// Contains detail field for an exception, which brings the error.
type InvalidPostTitleError struct {
	detail error
}

func (err *InvalidPostTitleError) Error() string {
	return fmt.Sprintf("%v: %v", ErrInvalidPostTitle, err.detail)
}

func (err *InvalidPostTitleError) Unwrap() error {
	return err.detail
}

// InvalidPostContentError is returned when attempting to create a Post with invalid content.
// Contains detail field for an exception, which brings the error.
type InvalidPostContentError struct {
	detail error
}

func (err *InvalidPostContentError) Error() string {
	return fmt.Sprintf("%v: %v", ErrInvalidPostContent, err.detail)
}

func (err *InvalidPostContentError) Unwrap() error {
	return err.detail
}

var (
	ErrNilAuthor          = errors.New("nil author")
	ErrInvalidPostTitle   = errors.New("invalid post title")
	ErrInvalidPostContent = errors.New("invalid post content")
)
