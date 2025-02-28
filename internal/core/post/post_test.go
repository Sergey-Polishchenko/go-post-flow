package post

import (
	"errors"
	"strings"
	"testing"

	"github.com/Sergey-Polishchenko/go-post-flow/internal/core/user"
	"github.com/Sergey-Polishchenko/go-post-flow/internal/core/validation"
)

func TestNew(t *testing.T) {
	validUser := &user.User{}
	validTitle := PostTitle("Valid Title")
	validContent := PostContent("Valid Content")

	tests := []struct {
		name      string
		author    *user.User
		title     PostTitle
		content   PostContent
		wantedErr error
	}{
		{
			name:      "Correct",
			author:    validUser,
			title:     validTitle,
			content:   validContent,
			wantedErr: nil,
		},
		{
			name:      "Nil Author",
			author:    nil,
			title:     validTitle,
			content:   validContent,
			wantedErr: ErrNilAuthor,
		},
		{
			name:      "Empty Title",
			author:    validUser,
			title:     PostTitle(""),
			content:   validContent,
			wantedErr: &InvalidPostTitleError{validation.ErrEmpty},
		},
		{
			name:      "Too Long Title",
			author:    validUser,
			title:     PostTitle(strings.Repeat("s", 101)),
			content:   validContent,
			wantedErr: &InvalidPostTitleError{validation.ErrTooLong},
		},
		{
			name:      "Empty Content",
			author:    validUser,
			title:     validTitle,
			content:   PostContent(""),
			wantedErr: &InvalidPostContentError{validation.ErrEmpty},
		},
		{
			name:      "Too Long Content",
			author:    validUser,
			title:     validTitle,
			content:   PostContent(strings.Repeat("s", 2001)),
			wantedErr: &InvalidPostContentError{validation.ErrTooLong},
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name,
			func(t *testing.T) {
				_, err := New(tt.author, tt.title, tt.content)

				if tt.wantedErr == nil && err != nil {
					t.Fatalf("unexpected error: %v", err)
				}

				if tt.wantedErr != nil {
					var target interface{}

					switch tt.wantedErr.(type) {
					case *InvalidPostTitleError:
						target = &InvalidPostTitleError{}
					case *InvalidPostContentError:
						target = &InvalidPostContentError{}
					default:
						target = tt.wantedErr
					}

					if !errors.As(err, &target) {
						t.Errorf("expected %T, got %T", tt.wantedErr, err)
					}
				}

				if !errors.Is(errors.Unwrap(err), errors.Unwrap(tt.wantedErr)) {
					t.Errorf(
						"cause mismatch: got %v, want %v",
						errors.Unwrap(err),
						errors.Unwrap(tt.wantedErr),
					)
				}
			},
		)
	}
}
