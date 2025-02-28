package comment

import (
	"errors"
	"strings"
	"testing"

	"github.com/Sergey-Polishchenko/go-post-flow/internal/core/user"
	"github.com/Sergey-Polishchenko/go-post-flow/internal/core/validation"
)

func TestNew(t *testing.T) {
	validUser := &user.User{}
	validText := CommentText("Valid comment text")

	tests := []struct {
		name      string
		author    *user.User
		text      CommentText
		wantedErr error
	}{
		{
			name:      "Correct",
			author:    validUser,
			text:      validText,
			wantedErr: nil,
		},
		{
			name:      "Nil Author",
			author:    nil,
			text:      validText,
			wantedErr: ErrNilAuthor,
		},
		{
			name:      "Empty Text",
			author:    validUser,
			text:      CommentText(""),
			wantedErr: &InvalidCommentTextError{validation.ErrEmpty},
		},
		{
			name:      "Too Long Text",
			author:    validUser,
			text:      CommentText(strings.Repeat("s", 2001)),
			wantedErr: &InvalidCommentTextError{validation.ErrTooLong},
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name,
			func(t *testing.T) {
				_, err := New(tt.author, tt.text)

				if tt.wantedErr == nil && err != nil {
					t.Fatalf("unexpected error: %v", err)
				}

				if tt.wantedErr != nil {
					var target *InvalidCommentTextError
					if errors.As(tt.wantedErr, &target) && !errors.As(err, &target) {
						t.Errorf("expected InvalidCommentTextError, got %T", err)
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
