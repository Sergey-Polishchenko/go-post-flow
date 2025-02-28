package post

import "github.com/Sergey-Polishchenko/go-post-flow/internal/core/validation"

const (
	MinPostContentLength = 1
	MaxPostContentLength = 2000
)

// PostContent represents a post's content.
type PostContent string

// IsValid checks the validity of the content.
func (content PostContent) IsValid() error {
	return validation.NewLengthValidator(MinPostContentLength, MaxPostContentLength).
		Validate(string(content))
}
