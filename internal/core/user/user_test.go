package user

import (
	"errors"
	"strings"
	"testing"

	"github.com/Sergey-Polishchenko/go-post-flow/internal/core/validation"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name      string
		username  UserName
		wantedErr error
	}{
		{
			name:      "Correct",
			username:  UserName("Human"),
			wantedErr: nil,
		},
		{
			name:      "Empty",
			username:  UserName(""),
			wantedErr: &InvalidUsernameError{validation.ErrEmpty},
		},
		{
			name:      "Too long",
			username:  UserName(strings.Repeat("s", 101)),
			wantedErr: &InvalidUsernameError{validation.ErrTooLong},
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name,
			func(t *testing.T) {
				_, err := New(tt.username)

				if tt.wantedErr == nil && err != nil {
					t.Fatalf("unexpected error: %v", err)
				}

				if tt.wantedErr != nil {
					var target *InvalidUsernameError
					if !errors.As(err, &target) {
						t.Errorf("expected InvalidUsernameError, got %T", err)
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
