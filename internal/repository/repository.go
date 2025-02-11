package repository

import (
	"context"

	"github.com/Sergey-Polishchenko/go-post-flow/internal/delivery/graph/model"
	inmemory "github.com/Sergey-Polishchenko/go-post-flow/internal/repository/in-memory"
	"github.com/Sergey-Polishchenko/go-post-flow/internal/repository/postgres"
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

	BroadcastComment(comment *model.Comment)
	RegisterCommentChannel(postID string, ch chan *model.Comment)
	UnregisterCommentChannel(postID string, ch chan *model.Comment)
}

func LoadStorage(isInmemory bool, connStr string) (Storage, error) {
	if isInmemory {
		return inmemory.NewStorage(), nil
	}
	return postgres.NewStorage(connStr)
}
