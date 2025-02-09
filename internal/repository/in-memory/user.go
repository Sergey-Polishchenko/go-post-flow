package inmemory

import (
	"errors"
	"fmt"

	"github.com/Sergey-Polishchenko/go-post-flow/internal/delivery/graph/model"
)

func (s *InMemoryStorage) CreateUser(input model.UserInput) (*model.User, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	user := &model.User{
		ID:   fmt.Sprintf("user_%d", len(s.Users)+1),
		Name: input.Name,
	}
	s.Users[user.ID] = user
	return user, nil
}

func (s *InMemoryStorage) GetUser(id string) (*model.User, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	user, exists := s.Users[id]
	if !exists {
		return nil, errors.New("user not found")
	}
	return user, nil
}
