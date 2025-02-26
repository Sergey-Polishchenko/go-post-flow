package dataloaders

import (
	"context"
	"sync"
	"time"

	"github.com/Sergey-Polishchenko/go-post-flow/internal/delivery/graph/model"
	flowerrors "github.com/Sergey-Polishchenko/go-post-flow/pkg/errors"
)

type CommentLoader struct {
	batchFn  func(context.Context, []string) ([]*model.Comment, error)
	wait     time.Duration
	maxBatch int
	cache    map[string]*model.Comment
	mu       sync.Mutex
}

func NewCommentLoader(
	batchFn func(context.Context, []string) ([]*model.Comment, error),
	wait time.Duration,
	maxBatch int,
) *CommentLoader {
	return &CommentLoader{
		batchFn:  batchFn,
		wait:     wait,
		maxBatch: maxBatch,
		cache:    make(map[string]*model.Comment),
	}
}

func (l *CommentLoader) Load(ctx context.Context, id string) (*model.Comment, error) {
	comments, errors := l.LoadAll(ctx, []string{id})
	return comments[0], errors[0]
}

func (l *CommentLoader) LoadAll(ctx context.Context, ids []string) ([]*model.Comment, []error) {
	comments := make([]*model.Comment, len(ids))
	errors := make([]error, len(ids))

	l.mu.Lock()
	defer l.mu.Unlock()

	var missingIDs []string
	for i, id := range ids {
		if comment, ok := l.cache[id]; ok {
			comments[i] = comment
		} else {
			missingIDs = append(missingIDs, id)
		}
	}

	if len(missingIDs) == 0 {
		return comments, errors
	}

	fetchedComments, err := l.batchFn(ctx, missingIDs)
	if err != nil {
		for i := range missingIDs {
			errors[i] = err
		}
		return comments, errors
	}

	for i, id := range missingIDs {
		if i < len(fetchedComments) && fetchedComments[i] != nil {
			l.cache[id] = fetchedComments[i]
			comments[i] = fetchedComments[i]
		} else {
			errors[i] = flowerrors.ErrCommentNotFound
		}
	}

	return comments, errors
}
