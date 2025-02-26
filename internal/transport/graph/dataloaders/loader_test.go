package dataloaders

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Sergey-Polishchenko/go-post-flow/internal/transport/graph/model"
)

type mockStorage struct{}

func (m *mockStorage) GetUsersByIDs(ctx context.Context, ids []string) ([]*model.User, error) {
	return []*model.User{
		{ID: "1", Name: "User1"},
	}, nil
}

func (m *mockStorage) GetPostsByIDs(ctx context.Context, ids []string) ([]*model.Post, error) {
	return []*model.Post{
		{ID: "1", Title: "Post1"},
	}, nil
}

func (m *mockStorage) GetCommentsByIDs(
	ctx context.Context,
	ids []string,
) ([]*model.Comment, error) {
	return []*model.Comment{
		{ID: "1", Text: "Comment1"},
	}, nil
}

func TestMiddleware(t *testing.T) {
	storage := &mockStorage{}
	middleware := Middleware(storage)

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userLoader := r.Context().Value(UserLoaderKey).(*UserLoader)
		postLoader := r.Context().Value(PostLoaderKey).(*PostLoader)
		commentLoader := r.Context().Value(CommentLoaderKey).(*CommentLoader)

		assert.NotNil(t, userLoader)
		assert.NotNil(t, postLoader)
		assert.NotNil(t, commentLoader)

		w.WriteHeader(http.StatusOK)
	})

	req := httptest.NewRequest("GET", "/", nil)
	rec := httptest.NewRecorder()

	middleware(handler).ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
}
