package inmemory

import (
	"fmt"

	"github.com/Sergey-Polishchenko/go-post-flow/internal/delivery/graph/model"
	reperrors "github.com/Sergey-Polishchenko/go-post-flow/internal/repository/errors"
	"github.com/Sergey-Polishchenko/go-post-flow/internal/utils"
)

func (s *InMemoryStorage) CreatePost(input model.PostInput) (*model.Post, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	author, exists := s.Users[input.AuthorID]
	if !exists {
		return nil, reperrors.ErrAuthorNotFound
	}

	post := &model.Post{
		ID:            fmt.Sprintf("post_%d", len(s.Posts)+1),
		Title:         input.Title,
		Content:       input.Content,
		Author:        author,
		AllowComments: input.AllowComments,
	}

	s.Posts[post.ID] = post
	return post, nil
}

func (s *InMemoryStorage) GetPosts(limit, offset *int) ([]*model.Post, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	posts := make([]*model.Post, 0, len(s.Posts))
	for _, post := range s.Posts {
		posts = append(posts, post)
	}

	return utils.ApplyPagination(posts, limit, offset), nil
}

func (s *InMemoryStorage) GetPost(id string) (*model.Post, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	post, exists := s.Posts[id]
	if !exists {
		return nil, reperrors.ErrPostNotFound
	}

	return post, nil
}

func (s *InMemoryStorage) GetComments(postID string, limit, offset *int) ([]*model.Comment, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	post, exists := s.Posts[postID]
	if !exists {
		return nil, reperrors.ErrPostNotFound
	}

	return utils.ApplyPagination(post.Comments, limit, offset), nil
}
