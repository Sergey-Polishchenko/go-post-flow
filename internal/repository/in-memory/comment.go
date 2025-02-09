package inmemory

import (
	"fmt"
	"time"

	"github.com/Sergey-Polishchenko/go-post-flow/internal/delivery/graph/model"
	reperrors "github.com/Sergey-Polishchenko/go-post-flow/internal/errors"
)

func (s *InMemoryStorage) CreateComment(input model.CommentInput) (*model.Comment, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	author, exists := s.Users[input.AuthorID]
	if !exists {
		return nil, reperrors.ErrAuthorNotFound
	}

	post, exists := s.Posts[input.PostID]
	if !exists {
		return nil, nil
	}

	comment := &model.Comment{
		ID:        fmt.Sprintf("comment_%d", len(s.Comments)+1),
		Text:      input.Text,
		Author:    author,
		Post:      post,
		CreatedAt: time.Now().Format(time.RFC3339),
		Children:  []*model.Comment{},
	}

	if input.ParentID != nil {
		parent, exists := s.Comments[*input.ParentID]
		if !exists {
			return nil, reperrors.ErrParentCommentNotFound
		}
		parent.Children = append(parent.Children, comment)
	} else {
		post.Comments = append(post.Comments, comment)
	}

	s.Comments[comment.ID] = comment
	return comment, nil
}

func (s *InMemoryStorage) GetChildren(commentID string) ([]*model.Comment, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	comment, exists := s.Comments[commentID]
	if !exists {
		return nil, reperrors.ErrCommentNotFound
	}

	if len(comment.Children) == 0 {
		return nil, reperrors.ErrCommentChildrenNotFound
	}

	return comment.Children, nil
}

func (s *InMemoryStorage) RegisterCommentChannel(postID string, ch chan *model.Comment) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.commentChannels[postID] = append(s.commentChannels[postID], ch)
}

func (s *InMemoryStorage) UnregisterCommentChannel(postID string, ch chan *model.Comment) {
	s.mu.Lock()
	defer s.mu.Unlock()
	channels := s.commentChannels[postID]
	for i, c := range channels {
		if c == ch {
			s.commentChannels[postID] = append(channels[:i], channels[i+1:]...)
			return
		}
	}
}

func (s *InMemoryStorage) BroadcastComment(comment *model.Comment) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	channels := s.commentChannels[comment.Post.ID]
	for _, ch := range channels {
		select {
		case ch <- comment:
		default:
		}
	}
}
