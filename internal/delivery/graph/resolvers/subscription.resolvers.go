package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.64

import (
	"context"
	"fmt"

	"github.com/Sergey-Polishchenko/go-post-flow/internal/delivery/graph/generated"
	"github.com/Sergey-Polishchenko/go-post-flow/internal/delivery/graph/model"
)

// CommentAdded is the resolver for the commentAdded field.
func (r *subscriptionResolver) CommentAdded(ctx context.Context, postID string) (<-chan *model.Comment, error) {
	// TODO: implement comment subscription
	panic(fmt.Errorf("not implemented: CommentAdded - commentAdded"))
}

// Subscription returns generated.SubscriptionResolver implementation.
func (r *Resolver) Subscription() generated.SubscriptionResolver { return &subscriptionResolver{r} }

type subscriptionResolver struct{ *Resolver }
