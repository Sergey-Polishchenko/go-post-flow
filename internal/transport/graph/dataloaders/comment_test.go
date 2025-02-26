package dataloaders

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/Sergey-Polishchenko/go-post-flow/internal/transport/graph/model"
)

func TestCommentLoader(t *testing.T) {
	tests := []struct {
		name         string
		ids          []string
		batchFn      func(context.Context, []string) ([]*model.Comment, error)
		wantComments []*model.Comment
		wantErrors   []error
	}{
		{
			name: "Load comments successfully",
			ids:  []string{"1", "2"},
			batchFn: func(ctx context.Context, ids []string) ([]*model.Comment, error) {
				return []*model.Comment{
					{ID: "1", Text: "Comment1"},
					{ID: "2", Text: "Comment2"},
				}, nil
			},
			wantComments: []*model.Comment{
				{ID: "1", Text: "Comment1"},
				{ID: "2", Text: "Comment2"},
			},
			wantErrors: []error{nil, nil},
		},
		{
			name: "Load comments with error",
			ids:  []string{"1", "2"},
			batchFn: func(ctx context.Context, ids []string) ([]*model.Comment, error) {
				return nil, assert.AnError
			},
			wantComments: []*model.Comment{nil, nil},
			wantErrors:   []error{assert.AnError, assert.AnError},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			loader := NewCommentLoader(tt.batchFn, 10*time.Millisecond, 100)
			comments, errors := loader.LoadAll(context.Background(), tt.ids)

			assert.Equal(t, tt.wantComments, comments)
			assert.Equal(t, tt.wantErrors, errors)
		})
	}
}
