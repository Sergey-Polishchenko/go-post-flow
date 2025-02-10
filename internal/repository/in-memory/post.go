package inmemory

import (
	"context"
	"fmt"
	"sort"

	"github.com/Sergey-Polishchenko/go-post-flow/internal/delivery/graph/model"
	"github.com/Sergey-Polishchenko/go-post-flow/internal/errors"
)

func (s *InMemoryStorage) CreatePost(input model.PostInput) (*model.Post, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	author, exists := s.Users[input.AuthorID]
	if !exists {
		return nil, errors.ErrAuthorNotFound
	}

	post := &model.Post{
		ID:            fmt.Sprintf("%d", len(s.Posts)+1),
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

	sort.Slice(posts, func(i, j int) bool {
		return posts[i].ID < posts[j].ID
	})

	return posts, nil
}

func (s *InMemoryStorage) GetPostsByIDs(ctx context.Context, ids []string) ([]*model.Post, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	posts := make([]*model.Post, len(ids))
	for i, id := range ids {
		posts[i] = s.Posts[id]
	}
	return posts, nil
}
