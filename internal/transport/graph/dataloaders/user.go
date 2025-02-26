package dataloaders

import (
	"context"
	"sync"
	"time"

	"github.com/Sergey-Polishchenko/go-post-flow/internal/transport/graph/model"
	flowerrors "github.com/Sergey-Polishchenko/go-post-flow/pkg/errors"
)

type UserLoader struct {
	batchFn  func(context.Context, []string) ([]*model.User, error)
	wait     time.Duration
	maxBatch int
	cache    map[string]*model.User
	mu       sync.Mutex
}

func NewUserLoader(
	batchFn func(context.Context, []string) ([]*model.User, error),
	wait time.Duration,
	maxBatch int,
) *UserLoader {
	return &UserLoader{
		batchFn:  batchFn,
		wait:     wait,
		maxBatch: maxBatch,
		cache:    make(map[string]*model.User),
	}
}

func (l *UserLoader) Load(ctx context.Context, id string) (*model.User, error) {
	users, errors := l.LoadAll(ctx, []string{id})
	return users[0], errors[0]
}

func (l *UserLoader) LoadAll(ctx context.Context, ids []string) ([]*model.User, []error) {
	users := make([]*model.User, len(ids))
	errors := make([]error, len(ids))

	l.mu.Lock()
	defer l.mu.Unlock()

	var missingIDs []string
	for i, id := range ids {
		if user, ok := l.cache[id]; ok {
			users[i] = user
		} else {
			missingIDs = append(missingIDs, id)
		}
	}

	if len(missingIDs) == 0 {
		return users, errors
	}

	fetchedUsers, err := l.batchFn(ctx, missingIDs)
	if err != nil {
		for i := range missingIDs {
			errors[i] = err
		}
		return users, errors
	}

	for i, id := range missingIDs {
		if i < len(fetchedUsers) && fetchedUsers[i] != nil {
			l.cache[id] = fetchedUsers[i]
			users[i] = fetchedUsers[i]
		} else {
			errors[i] = flowerrors.ErrUserNotFound
		}
	}

	return users, errors
}
