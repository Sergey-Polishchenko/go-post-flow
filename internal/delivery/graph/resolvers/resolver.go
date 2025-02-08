package resolvers

import "github.com/Sergey-Polishchenko/go-post-flow/internal/delivery/graph/model"

type Storage interface {
	GetPost(id string) (*model.Post, error)
	GetPosts(limit, offset *int) ([]*model.Post, error)
	GetUser(id string) (*model.User, error)

	GetComments(postID string, limit, offset *int) ([]*model.Comment, error)
	GetChildren(commentID string, limit, offset *int) ([]*model.Comment, error)

	CreatePost(input model.PostInput) (*model.Post, error)
	CreateComment(input model.CommentInput) (*model.Comment, error)
	CreateUser(input model.UserInput) (*model.User, error)
}

type Resolver struct {
	storage Storage
}

func NewResolver(storage Storage) *Resolver {
	return &Resolver{storage: storage}
}
