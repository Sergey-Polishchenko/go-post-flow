package inmemory

import (
	"context"
	"fmt"

	"github.com/Sergey-Polishchenko/go-post-flow/internal/delivery/graph/model"
)

func (s *InMemoryStorage) CreateUser(input model.UserInput) (*model.User, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	user := &model.User{
		ID:   fmt.Sprintf("%d", len(s.Users)+1),
		Name: input.Name,
	}
	s.Users[user.ID] = user
	return user, nil
}

func (s *InMemoryStorage) GetUsersByIDs(ctx context.Context, ids []string) ([]*model.User, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	users := make([]*model.User, len(ids))
	for i, id := range ids {
		users[i] = s.Users[id]
	}
	return users, nil
}
