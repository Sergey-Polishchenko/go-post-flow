package comment

import (
	"errors"
	"fmt"
)

// InvalidCommentTextError is returned when attempting to create a Comment with invalid text.
// Contains detail field for an exception, which brings the error.
type InvalidCommentTextError struct {
	detail error
}

func (err *InvalidCommentTextError) Error() string {
	return fmt.Sprintf("%v: %v", ErrInvalidCommentText, err.detail)
}

func (err *InvalidCommentTextError) Unwrap() error {
	return err.detail
}

var (
	ErrNilAuthor          = errors.New("nil author")
	ErrInvalidCommentText = errors.New("invalid comment text")
)
