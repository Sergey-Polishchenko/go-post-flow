package inmemory

import (
	"github.com/Sergey-Polishchenko/go-post-flow/internal/delivery/graph/model"
)

func (s *InMemoryStorage) CreatePost(input model.PostInput) (*model.Post, error) {
	return nil, nil
}

func (s *InMemoryStorage) GetPost(id string) (*model.Post, error) {
	return nil, nil
}

func (s *InMemoryStorage) GetComments(postID string, limit, offset *int) ([]*model.Comment, error) {
	return nil, nil
}
