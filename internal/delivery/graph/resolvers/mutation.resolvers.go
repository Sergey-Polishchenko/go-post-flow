package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.64

import (
	"context"

	"github.com/Sergey-Polishchenko/go-post-flow/internal/delivery/graph/generated"
	"github.com/Sergey-Polishchenko/go-post-flow/internal/delivery/graph/model"
	flowerrors "github.com/Sergey-Polishchenko/go-post-flow/internal/errors"
)

// CreatePost is the resolver for the createPost field.
func (r *mutationResolver) CreatePost(ctx context.Context, input model.PostInput) (*model.Post, error) {
	return r.storage.CreatePost(input)
}

// CreateComment is the resolver for the createComment field.
func (r *mutationResolver) CreateComment(ctx context.Context, input model.CommentInput) (*model.Comment, error) {
	if len(input.Text) >= r.comLim {
		return nil, flowerrors.ErrCommentTooLong
	}

	comment, err := r.storage.CreateComment(input)
	if err != nil {
		return nil, err
	}

	r.storage.BroadcastComment(comment)

	return comment, nil
}

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input model.UserInput) (*model.User, error) {
	return r.storage.CreateUser(input)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
