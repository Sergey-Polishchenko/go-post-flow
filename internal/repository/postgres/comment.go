package postgres

import (
	"fmt"
	"log"
	"time"

	"github.com/Sergey-Polishchenko/go-post-flow/internal/delivery/graph/model"
	reperrors "github.com/Sergey-Polishchenko/go-post-flow/internal/errors"
)

func (s *PostgresStorage) CreateComment(input model.CommentInput) (*model.Comment, error) {
	var commentID string
	query := `
        INSERT INTO comments (text, author_id, post_id, parent_id)
        VALUES ($1, $2, $3, $4)
        RETURNING id
    `
	err := s.db.QueryRow(query, input.Text, input.AuthorID, input.PostID, input.ParentID).
		Scan(&commentID)
	if err != nil {
		return nil, fmt.Errorf("failed to create comment: %w", err)
	}

	author, err := s.GetUser(input.AuthorID)
	if err != nil {
		return nil, fmt.Errorf("failed to get author: %w", err)
	}

	post, err := s.GetPost(input.PostID)
	if err != nil {
		return nil, fmt.Errorf("failed to get post: %w", err)
	}

	comment := &model.Comment{
		ID:        commentID,
		Text:      input.Text,
		Author:    author,
		Post:      post,
		CreatedAt: time.Now().Format(time.RFC3339),
	}

	if input.ParentID != nil {
		parent, err := s.GetComment(*input.ParentID)
		if err != nil {
			return nil, fmt.Errorf("failed to get parent comment: %w", err)
		}
		parent.Children = append(parent.Children, comment)
	} else {
		post.Comments = append(post.Comments, comment)
	}

	s.BroadcastComment(comment)

	return comment, nil
}

func (s *PostgresStorage) GetComment(id string) (*model.Comment, error) {
	query := `
			SELECT id, text, author_id, created_at
			FROM comments
			WHERE id = $1
	`
	rows, err := s.db.Query(query, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get comments: %w", err)
	}
	defer rows.Close()

	var comment model.Comment
	if err := rows.Scan(&comment.ID, &comment.Text, &comment.Author.ID, &comment.CreatedAt); err != nil {
		return nil, fmt.Errorf("failed to scan comment: %w", err)
	}

	return &comment, nil
}

func (s *PostgresStorage) GetComments(postID string) ([]*model.Comment, error) {
	query := `
        SELECT id, text, author_id, created_at
        FROM comments
        WHERE post_id = $1 AND parent_id IS NULL
    `
	rows, err := s.db.Query(query, postID)
	if err != nil {
		return nil, fmt.Errorf("failed to get comments: %w", err)
	}
	defer rows.Close()

	var comments []*model.Comment
	for rows.Next() {
		var comment model.Comment
		var authorID string
		if err := rows.Scan(&comment.ID, &comment.Text, &authorID, &comment.CreatedAt); err != nil {
			return nil, fmt.Errorf("failed to scan comment: %w", err)
		}

		author, err := s.GetUser(authorID)
		if err != nil {
			return nil, reperrors.ErrAuthorNotFound
		}
		comment.Author = author

		comments = append(comments, &comment)
	}

	return comments, nil
}

func (s *PostgresStorage) GetChildren(commentID string) ([]*model.Comment, error) {
	query := `
        SELECT id, text, author_id, created_at
        FROM comments
        WHERE parent_id = $1
    `
	rows, err := s.db.Query(query, commentID)
	if err != nil {
		return nil, reperrors.ErrCommentChildrenNotFound
	}
	defer rows.Close()

	var children []*model.Comment
	for rows.Next() {
		var comment model.Comment
		var authorID string
		if err := rows.Scan(&comment.ID, &comment.Text, &authorID, &comment.CreatedAt); err != nil {
			return nil, fmt.Errorf("failed to scan comment: %w", err)
		}

		author, err := s.GetUser(authorID)
		if err != nil {
			return nil, reperrors.ErrAuthorNotFound
		}
		comment.Author = author

		children = append(children, &comment)
	}

	if len(children) == 0 {
		return nil, reperrors.ErrCommentChildrenNotFound
	}

	return children, nil
}

func (s *PostgresStorage) RegisterCommentChannel(postID string, ch chan *model.Comment) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.commentChannels[postID] = append(s.commentChannels[postID], ch)
}

func (s *PostgresStorage) UnregisterCommentChannel(postID string, ch chan *model.Comment) {
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

func (s *PostgresStorage) BroadcastComment(comment *model.Comment) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	channels := s.commentChannels[comment.Post.ID]
	for _, ch := range channels {
		select {
		case ch <- comment:
		default:
			log.Println("Channel is full, skipping broadcast")
		}
	}
}
