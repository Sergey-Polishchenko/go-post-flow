package dataloader

import (
	"context"
	"sync"
	"time"

	"github.com/Sergey-Polishchenko/go-post-flow/internal/delivery/graph/model"
	flowerrors "github.com/Sergey-Polishchenko/go-post-flow/internal/errors"
)

type PostLoader struct {
	batchFn  func(context.Context, []string) ([]*model.Post, error)
	wait     time.Duration
	maxBatch int
	cache    map[string]*model.Post
	mu       sync.Mutex
}

func NewPostLoader(
	batchFn func(context.Context, []string) ([]*model.Post, error),
	wait time.Duration,
	maxBatch int,
) *PostLoader {
	return &PostLoader{
		batchFn:  batchFn,
		wait:     wait,
		maxBatch: maxBatch,
		cache:    make(map[string]*model.Post),
	}
}

func (l *PostLoader) Load(ctx context.Context, id string) (*model.Post, error) {
	posts, errors := l.LoadAll(ctx, []string{id})
	return posts[0], errors[0]
}

func (l *PostLoader) LoadAll(ctx context.Context, ids []string) ([]*model.Post, []error) {
	posts := make([]*model.Post, len(ids))
	errors := make([]error, len(ids))

	l.mu.Lock()
	defer l.mu.Unlock()

	var missingIDs []string
	for i, id := range ids {
		if post, ok := l.cache[id]; ok {
			posts[i] = post
		} else {
			missingIDs = append(missingIDs, id)
		}
	}

	if len(missingIDs) == 0 {
		return posts, errors
	}

	fetchedPosts, err := l.batchFn(ctx, missingIDs)
	if err != nil {
		for i := range missingIDs {
			errors[i] = err
		}
		return posts, errors
	}

	for i, id := range missingIDs {
		if i < len(fetchedPosts) && fetchedPosts[i] != nil {
			l.cache[id] = fetchedPosts[i]
			posts[i] = fetchedPosts[i]
		} else {
			errors[i] = flowerrors.ErrPostNotFound
		}
	}

	return posts, errors
}
