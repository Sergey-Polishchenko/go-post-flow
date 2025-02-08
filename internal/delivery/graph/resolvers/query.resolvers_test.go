package resolvers

import (
	"context"
	"reflect"
	"testing"

	"github.com/Sergey-Polishchenko/go-post-flow/internal/delivery/graph/model"
	"github.com/Sergey-Polishchenko/go-post-flow/internal/delivery/graph/resolvers/mock"
)

func Test_queryResolver_Post(t *testing.T) {
	mockStorage := new(mock.MockStorage)
	mockPost := &model.Post{ID: "1", Title: "Test Post", Content: "This is a test post."}

	mockStorage.On("GetPost", "1").Return(mockPost, nil)

	r := &queryResolver{
		Resolver: &Resolver{
			storage: mockStorage,
		},
	}

	ctx := context.Background()
	got, err := r.Post(ctx, "1")
	if err != nil {
		t.Errorf("queryResolver.Post() error = %v, wantErr %v", err, false)
		return
	}

	if !reflect.DeepEqual(got, mockPost) {
		t.Errorf("queryResolver.Post() = %v, want %v", got, mockPost)
	}

	mockStorage.AssertExpectations(t)
}

func Test_queryResolver_Posts(t *testing.T) {
	mockStorage := new(mock.MockStorage)
	mockPosts := []*model.Post{
		{ID: "1", Title: "Test Post 1", Content: "This is a test post 1."},
		{ID: "2", Title: "Test Post 2", Content: "This is a test post 2."},
	}

	limit := 10
	offset := 0

	mockStorage.On("GetPosts", &limit, &offset).Return(mockPosts, nil)

	r := &queryResolver{
		Resolver: &Resolver{
			storage: mockStorage,
		},
	}

	ctx := context.Background()
	got, err := r.Posts(ctx, &limit, &offset)
	if err != nil {
		t.Errorf("queryResolver.Posts() error = %v, wantErr %v", err, false)
		return
	}

	if !reflect.DeepEqual(got, mockPosts) {
		t.Errorf("queryResolver.Posts() = %v, want %v", got, mockPosts)
	}

	mockStorage.AssertExpectations(t)
}
