package storage

import (
	"context"

	"github.com/Sergey-Polishchenko/go-post-flow/internal/pkg/broadcast"
	inmemory "github.com/Sergey-Polishchenko/go-post-flow/internal/storage/in-memory"
	"github.com/Sergey-Polishchenko/go-post-flow/internal/storage/postgres"
	"github.com/Sergey-Polishchenko/go-post-flow/internal/transport/graph/model"
)

type Storage interface {
	CreatePost(input model.PostInput) (*model.Post, error)
	GetPosts(limit, offset *int) ([]*model.Post, error)
	GetPostsByIDs(ctx context.Context, ids []string) ([]*model.Post, error)

	CreateComment(input model.CommentInput) (*model.Comment, error)
	GetCommentsByIDs(ctx context.Context, ids []string) ([]*model.Comment, error)
	GetCommentsIDs(postID string) ([]string, error)
	GetChildrenIDs(commentID string) ([]string, error)

	CreateUser(input model.UserInput) (*model.User, error)
	GetUsersByIDs(ctx context.Context, ids []string) ([]*model.User, error)

	broadcast.Broadcaster
}

func LoadStorage(isInmemory bool, connStr string) (Storage, error) {
	if isInmemory {
		return inmemory.NewStorage(), nil
	}
	return postgres.NewStorage(connStr)
}
