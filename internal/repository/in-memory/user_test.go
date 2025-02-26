package inmemory

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Sergey-Polishchenko/go-post-flow/internal/transport/graph/model"
)

func TestCreateUser(t *testing.T) {
	storage := NewStorage()

	tests := []struct {
		name    string
		input   model.UserInput
		want    *model.User
		wantErr bool
	}{
		{
			name:  "Create user successfully",
			input: model.UserInput{Name: "John Doe"},
			want:  &model.User{ID: "1", Name: "John Doe"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := storage.CreateUser(tt.input)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want.Name, got.Name)
				assert.Equal(t, tt.want.ID, got.ID)
			}
		})
	}
}

func TestGetUsersByIDs(t *testing.T) {
	storage := NewStorage()
	user1, _ := storage.CreateUser(model.UserInput{Name: "User1"})
	user2, _ := storage.CreateUser(model.UserInput{Name: "User2"})

	tests := []struct {
		name    string
		ids     []string
		want    []*model.User
		wantErr bool
	}{
		{
			name: "Get users by IDs",
			ids:  []string{user1.ID, user2.ID},
			want: []*model.User{user1, user2},
		},
		{
			name:    "Non-existent user ID",
			ids:     []string{"non-existent"},
			want:    []*model.User{nil},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := storage.GetUsersByIDs(context.Background(), tt.ids)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}
