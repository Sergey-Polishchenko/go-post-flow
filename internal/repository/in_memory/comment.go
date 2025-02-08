package inmemory

import (
	"github.com/Sergey-Polishchenko/go-post-flow/internal/delivery/graph/model"
)

func (s *InMemoryStorage) CreateComment(input model.CommentInput) (*model.Comment, error) {
	return nil, nil
}

func (s *InMemoryStorage) GetChildren(
	commentID string,
	limit, offset *int,
) ([]*model.Comment, error) {
	return nil, nil
}
