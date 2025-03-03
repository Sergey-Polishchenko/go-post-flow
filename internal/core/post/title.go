package post

import "github.com/Sergey-Polishchenko/go-post-flow/internal/core/validation"

const (
	MinPostTitleLength = 1
	MaxPostTitleLength = 100
)

// PostTitle represents a post's title.
type PostTitle string

func (title PostTitle) String() string {
	return string(title)
}

// IsValid checks the validity of the title.
func (title PostTitle) IsValid() error {
	return validation.NewLengthValidator(MinPostTitleLength, MaxPostTitleLength).
		Validate(string(title))
}
