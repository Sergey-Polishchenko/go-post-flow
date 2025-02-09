package repository

import (
	"github.com/Sergey-Polishchenko/go-post-flow/internal/delivery/graph/model"
	inmemory "github.com/Sergey-Polishchenko/go-post-flow/internal/repository/in_memory"
)

type Storage interface {
	GetPost(id string) (*model.Post, error)
	GetPosts(limit, offset *int) ([]*model.Post, error)
	GetUser(id string) (*model.User, error)

	GetComments(postID string) ([]*model.Comment, error)
	GetChildren(commentID string) ([]*model.Comment, error)

	CreatePost(input model.PostInput) (*model.Post, error)
	CreateComment(input model.CommentInput) (*model.Comment, error)
	CreateUser(input model.UserInput) (*model.User, error)

	BroadcastComment(comment *model.Comment)
	RegisterCommentChannel(postID string, ch chan *model.Comment)
	UnregisterCommentChannel(postID string, ch chan *model.Comment)
}

func LoadStorage(isInmemory bool) Storage {
	if isInmemory {
		return inmemory.NewStorage()
	}
	return nil
}
