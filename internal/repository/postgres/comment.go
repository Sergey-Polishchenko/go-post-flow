package postgres

import (
	"database/sql"
	"time"

	"github.com/Sergey-Polishchenko/go-post-flow/internal/delivery/graph/model"
	"github.com/Sergey-Polishchenko/go-post-flow/internal/errors"
)

func (s *PostgresStorage) CreateComment(input model.CommentInput) (*model.Comment, error) {
	var commentID string
	var createdAt time.Time
	query, err := s.queries.LoadQuery("comment", "create")
	if err != nil {
		return nil, err
	}
	if err := s.db.QueryRow(
		query,
		input.Text,
		input.AuthorID,
		input.PostID,
		input.ParentID,
	).Scan(&commentID, &createdAt); err != nil {
		return nil, &errors.SQLCreatingError{Value: err}
	}

	author, err := s.GetUser(input.AuthorID)
	if err != nil {
		return nil, errors.ErrAuthorNotFound
	}

	post, err := s.GetPost(input.PostID)
	if err != nil {
		return nil, errors.ErrPostNotFound
	}
	if !post.AllowComments {
		return nil, errors.ErrCommentsNotAllowed
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
	query, err := s.queries.LoadQuery("comment", "get")
	if err != nil {
		return nil, err
	}
	var comment model.Comment
	var authorID, authorName, postID string
	var createdAt time.Time

	if err := s.db.QueryRow(query, id).Scan(
		&comment.ID,
		&comment.Text,
		&authorID,
		&authorName,
		&postID,
		&createdAt,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.ErrCommentNotFound
		}
		return nil, &errors.SQLScaningError{Value: err}
	}

	post, err := s.GetPost(postID)
	if err != nil {
		return nil, errors.ErrPostNotFound
	}

	comment.Author = &model.User{ID: authorID, Name: authorName}
	comment.Post = post
	comment.CreatedAt = createdAt.Format(time.RFC3339)

	return &comment, nil
}

func (s *PostgresStorage) GetComments(postID string) ([]*model.Comment, error) {
	query, err := s.queries.LoadQuery("comment", "comments")
	if err != nil {
		return nil, err
	}
	rows, err := s.db.Query(query, postID)
	if err != nil {
		return nil, errors.ErrCommentNotFound
	}
	defer rows.Close()

	var comments []*model.Comment
	for rows.Next() {
		var comment model.Comment
		var authorID, authorName, postID string
		var createdAt time.Time
		err := rows.Scan(&comment.ID, &comment.Text, &authorID, &authorName, &createdAt, &postID)
		if err != nil {
			return nil, &errors.SQLScaningError{Value: err}
		}

		post, err := s.GetPost(postID)
		if err != nil {
			return nil, errors.ErrPostNotFound
		}

		comment.Author = &model.User{ID: authorID, Name: authorName}
		comment.Post = post
		comment.CreatedAt = createdAt.Format(time.RFC3339)
		comments = append(comments, &comment)
	}
	return comments, nil
}

func (s *PostgresStorage) GetChildren(commentID string) ([]*model.Comment, error) {
	query, err := s.queries.LoadQuery("comment", "children")
	if err != nil {
		return nil, err
	}
	rows, err := s.db.Query(query, commentID)
	if err != nil {
		return nil, errors.ErrCommentNotFound
	}
	defer rows.Close()

	var children []*model.Comment
	for rows.Next() {
		var comment model.Comment
		var authorID, authorName, postID string
		var createdAt time.Time
		err := rows.Scan(&comment.ID, &comment.Text, &authorID, &authorName, &createdAt, &postID)
		if err != nil {
			return nil, &errors.SQLScaningError{Value: err}
		}

		post, err := s.GetPost(postID)
		if err != nil {
			return nil, errors.ErrPostNotFound
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
		}
	}
}
