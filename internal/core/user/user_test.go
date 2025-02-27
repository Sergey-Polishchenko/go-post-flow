package user

import (
	"strings"
	"testing"
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
			wantedErr: ErrInvalidUsername,
		},
		{
			name:      "Too long",
			username:  UserName(strings.Repeat("s", 101)),
			wantedErr: ErrInvalidUsername,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name,
			func(t *testing.T) {
				_, err := New(tt.username)
				if err != tt.wantedErr {
					t.Errorf("New() error = %v, wantedErr %v", err, tt.wantedErr)
					return
				}
			},
		)
	}
}
