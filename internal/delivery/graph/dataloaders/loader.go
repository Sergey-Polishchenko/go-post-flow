package dataloaders

import (
	"context"
	"net/http"
	"time"

	"github.com/Sergey-Polishchenko/go-post-flow/internal/delivery/graph/model"
)

type ctxKey string

const (
	UserLoaderKey    ctxKey = "userLoader"
	PostLoaderKey    ctxKey = "postLoader"
	CommentLoaderKey ctxKey = "commentLoader"
)

type storage interface {
	GetUsersByIDs(ctx context.Context, ids []string) ([]*model.User, error)
	GetPostsByIDs(ctx context.Context, ids []string) ([]*model.Post, error)
	GetCommentsByIDs(ctx context.Context, ids []string) ([]*model.Comment, error)
}

func Middleware(storage storage) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				userLoader := NewUserLoader(
					func(ctx context.Context, ids []string) ([]*model.User, error) {
						return storage.GetUsersByIDs(ctx, ids)
					},
					10*time.Millisecond,
					100,
				)

				postLoader := NewPostLoader(
					func(ctx context.Context, ids []string) ([]*model.Post, error) {
						return storage.GetPostsByIDs(ctx, ids)
					},
					10*time.Millisecond,
					100,
				)

				commentLoader := NewCommentLoader(
					func(ctx context.Context, ids []string) ([]*model.Comment, error) {
						return storage.GetCommentsByIDs(ctx, ids)
					},
					10*time.Millisecond,
					100,
				)

				ctx := r.Context()
				ctx = context.WithValue(ctx, UserLoaderKey, userLoader)
				ctx = context.WithValue(ctx, PostLoaderKey, postLoader)
				ctx = context.WithValue(ctx, CommentLoaderKey, commentLoader)

				next.ServeHTTP(w, r.WithContext(ctx))
			},
		)
	}
}
