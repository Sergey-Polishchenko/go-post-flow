package postgres

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/Sergey-Polishchenko/go-post-flow/internal/delivery/graph/model"
	reperrors "github.com/Sergey-Polishchenko/go-post-flow/internal/errors"
)

func (s *PostgresStorage) CreateComment(input model.CommentInput) (*model.Comment, error) {
	var commentID string
	var createdAt time.Time
	query := `
        INSERT INTO comments (text, author_id, post_id, parent_id)
        VALUES ($1, $2, $3, $4)
        RETURNING id, created_at
    `
	err := s.db.QueryRow(query, input.Text, input.AuthorID, input.PostID, input.ParentID).
		Scan(&commentID, &createdAt)
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
		CreatedAt: createdAt.Format(time.RFC3339),
	}

	s.BroadcastComment(comment)
	return comment, nil
}

func (s *PostgresStorage) GetComment(id string) (*model.Comment, error) {
	query := `
        SELECT c.id, c.text, c.author_id, u.name AS author_name, c.post_id, c.created_at
        FROM comments c
        JOIN users u ON c.author_id = u.id
        WHERE c.id = $1
    `
	var comment model.Comment
	var authorID, authorName, postID string
	var createdAt time.Time

	err := s.db.QueryRow(query, id).Scan(
		&comment.ID,
		&comment.Text,
		&authorID,
		&authorName,
		&postID,
		&createdAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, reperrors.ErrCommentNotFound
		}
		return nil, fmt.Errorf("failed to get comment: %w", err)
	}

	post, err := s.GetPost(postID)
	if err != nil {
		return nil, fmt.Errorf("failed to get post: %w", err)
	}

	comment.Author = &model.User{ID: authorID, Name: authorName}
	comment.Post = post
	comment.CreatedAt = createdAt.Format(time.RFC3339)

	return &comment, nil
}

func (s *PostgresStorage) GetComments(postID string) ([]*model.Comment, error) {
	query := `
        SELECT c.id, c.text, c.author_id, u.name AS author_name, c.created_at, c.post_id
        FROM comments c
        JOIN users u ON c.author_id = u.id
        WHERE c.post_id = $1 AND c.parent_id IS NULL
    `
	rows, err := s.db.Query(query, postID)
	if err != nil {
		return nil, fmt.Errorf("failed to get comments: %w", err)
	}
	defer rows.Close()

	var comments []*model.Comment
	for rows.Next() {
		var comment model.Comment
		var authorID, authorName, postID string
		var createdAt time.Time
		err := rows.Scan(&comment.ID, &comment.Text, &authorID, &authorName, &createdAt, &postID)
		if err != nil {
			return nil, fmt.Errorf("failed to scan comment: %w", err)
		}

		post, err := s.GetPost(postID)
		if err != nil {
			return nil, fmt.Errorf("failed to get post: %w", err)
		}

		comment.Author = &model.User{ID: authorID, Name: authorName}
		comment.Post = post
		comment.CreatedAt = createdAt.Format(time.RFC3339)
		comments = append(comments, &comment)
	}
	return comments, nil
}

func (s *PostgresStorage) GetChildren(commentID string) ([]*model.Comment, error) {
	query := `
        SELECT c.id, c.text, c.author_id, u.name AS author_name, c.created_at, c.post_id
        FROM comments c
        JOIN users u ON c.author_id = u.id
        WHERE c.parent_id = $1
    `
	rows, err := s.db.Query(query, commentID)
	if err != nil {
		return nil, fmt.Errorf("failed to get children: %w", err)
	}
	defer rows.Close()

	var children []*model.Comment
	for rows.Next() {
		var comment model.Comment
		var authorID, authorName, postID string
		var createdAt time.Time
		err := rows.Scan(&comment.ID, &comment.Text, &authorID, &authorName, &createdAt, &postID)
		if err != nil {
			return nil, fmt.Errorf("failed to scan comment: %w", err)
		}

		post, err := s.GetPost(postID)
		if err != nil {
			return nil, fmt.Errorf("failed to get post: %w", err)
		}

		comment.Author = &model.User{ID: authorID, Name: authorName}
		comment.Post = post
		comment.CreatedAt = createdAt.Format(time.RFC3339)
		children = append(children, &comment)
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
			break
		}
	}
}

func (s *PostgresStorage) BroadcastComment(comment *model.Comment) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	for _, ch := range s.commentChannels[comment.Post.ID] {
		select {
		case ch <- comment:
		default:
			log.Println("Channel is full, skipping broadcast")
		}
	}
}
