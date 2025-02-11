package dataloaders

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/Sergey-Polishchenko/go-post-flow/internal/delivery/graph/model"
)

func TestUserLoader(t *testing.T) {
	tests := []struct {
		name       string
		ids        []string
		batchFn    func(context.Context, []string) ([]*model.User, error)
		wantUsers  []*model.User
		wantErrors []error
	}{
		{
			name: "Load users successfully",
			ids:  []string{"1", "2"},
			batchFn: func(ctx context.Context, ids []string) ([]*model.User, error) {
				return []*model.User{
					{ID: "1", Name: "User1"},
					{ID: "2", Name: "User2"},
				}, nil
			},
			wantUsers: []*model.User{
				{ID: "1", Name: "User1"},
				{ID: "2", Name: "User2"},
			},
			wantErrors: []error{nil, nil},
		},
		{
			name: "Load users with error",
			ids:  []string{"1", "2"},
			batchFn: func(ctx context.Context, ids []string) ([]*model.User, error) {
				return nil, assert.AnError
			},
			wantUsers:  []*model.User{nil, nil},
			wantErrors: []error{assert.AnError, assert.AnError},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			loader := NewUserLoader(tt.batchFn, 10*time.Millisecond, 100)
			users, errors := loader.LoadAll(context.Background(), tt.ids)

			assert.Equal(t, tt.wantUsers, users)
			assert.Equal(t, tt.wantErrors, errors)
		})
	}
}
