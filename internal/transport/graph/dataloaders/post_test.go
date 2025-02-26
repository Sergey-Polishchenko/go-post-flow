package dataloaders

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/Sergey-Polishchenko/go-post-flow/internal/transport/graph/model"
)

func TestPostLoader(t *testing.T) {
	tests := []struct {
		name       string
		ids        []string
		batchFn    func(context.Context, []string) ([]*model.Post, error)
		wantPosts  []*model.Post
		wantErrors []error
	}{
		{
			name: "Load posts successfully",
			ids:  []string{"1", "2"},
			batchFn: func(ctx context.Context, ids []string) ([]*model.Post, error) {
				return []*model.Post{
					{ID: "1", Title: "Post1"},
					{ID: "2", Title: "Post2"},
				}, nil
			},
			wantPosts: []*model.Post{
				{ID: "1", Title: "Post1"},
				{ID: "2", Title: "Post2"},
			},
			wantErrors: []error{nil, nil},
		},
		{
			name: "Load posts with error",
			ids:  []string{"1", "2"},
			batchFn: func(ctx context.Context, ids []string) ([]*model.Post, error) {
				return nil, assert.AnError
			},
			wantPosts:  []*model.Post{nil, nil},
			wantErrors: []error{assert.AnError, assert.AnError},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			loader := NewPostLoader(tt.batchFn, 10*time.Millisecond, 100)
			posts, errors := loader.LoadAll(context.Background(), tt.ids)

			assert.Equal(t, tt.wantPosts, posts)
			assert.Equal(t, tt.wantErrors, errors)
		})
	}
}
