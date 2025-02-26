package inmemory

import (
	"context"
	"fmt"
	"time"

	"github.com/Sergey-Polishchenko/go-post-flow/internal/transport/graph/model"
	"github.com/Sergey-Polishchenko/go-post-flow/pkg/errors"
)

func (s *InMemoryStorage) CreateComment(input model.CommentInput) (*model.Comment, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	author, exists := s.Users[input.AuthorID]
	if !exists {
		return nil, errors.ErrAuthorNotFound
	}

	post, exists := s.Posts[input.PostID]
	if !exists {
		return nil, errors.ErrPostNotFound
	}
	if !post.AllowComments {
		return nil, errors.ErrCommentsNotAllowed
	}

	comment := &model.Comment{
		ID:        fmt.Sprintf("%d", len(s.Comments)+1),
		Text:      input.Text,
		Author:    author,
		Post:      post,
		Children:  []*model.Comment{},
		CreatedAt: time.Now().Format(time.RFC3339),
	}

	if input.ParentID != nil {
		parent, exists := s.Comments[*input.ParentID]
		if !exists {
			return nil, errors.ErrParentCommentNotFound
		}
		if parent.Post.ID != input.PostID {
			return nil, errors.ErrParentInOtherPost
		}
		parent.Children = append(parent.Children, comment)
	} else {
		post.Comments = append(post.Comments, comment)
	}

	s.Comments[comment.ID] = comment
	return comment, nil
}

func (s *InMemoryStorage) GetCommentsByIDs(
	ctx context.Context,
	ids []string,
) ([]*model.Comment, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	comments := make([]*model.Comment, len(ids))
	for i, id := range ids {
		comments[i] = s.Comments[id]
	}
	return comments, nil
}

func (s *InMemoryStorage) GetCommentsIDs(postID string) ([]string, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	post, exists := s.Posts[postID]
	if !exists {
		return nil, errors.ErrPostNotFound
	}

	ids := make([]string, 0)
	for _, c := range post.Comments {
		ids = append(ids, c.ID)
	}

	if len(ids) == 0 {
		return nil, errors.ErrCommentsNotFound
	}

	return ids, nil
}

func (s *InMemoryStorage) GetChildrenIDs(commentID string) ([]string, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	comment, exists := s.Comments[commentID]
	if !exists {
		return nil, errors.ErrCommentNotFound
	}

	ids := make([]string, 0)
	for _, c := range comment.Children {
		ids = append(ids, c.ID)
	}

	if len(ids) == 0 {
		return nil, errors.ErrCommentChildrenNotFound
	}

	return ids, nil
}
