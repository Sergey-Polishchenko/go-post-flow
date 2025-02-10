package postgres

import (
	"database/sql"
	"fmt"

	"github.com/Sergey-Polishchenko/go-post-flow/internal/delivery/graph/model"
	reperrors "github.com/Sergey-Polishchenko/go-post-flow/internal/errors"
)

func (s *PostgresStorage) CreatePost(input model.PostInput) (*model.Post, error) {
	var postID string
	query, err := s.queries.LoadQuery("post", "create")
	if err != nil {
		return nil, fmt.Errorf("on loading query: %s", err)
	}
	if err := s.db.QueryRow(
		query,
		input.Title,
		input.Content,
		input.AuthorID,
		input.AllowComments,
	).Scan(&postID); err != nil {
		return nil, fmt.Errorf("failed to create post: %w", err)
	}

	author, err := s.GetUser(input.AuthorID)
	if err != nil {
		return nil, fmt.Errorf("failed to get author: %w", err)
	}

	return &model.Post{
		ID:            postID,
		Title:         input.Title,
		Content:       input.Content,
		Author:        author,
		AllowComments: input.AllowComments,
	}, nil
}

func (s *PostgresStorage) GetPost(id string) (*model.Post, error) {
	var post model.Post
	var authorID string
	query, err := s.queries.LoadQuery("post", "get")
	if err != nil {
		return nil, fmt.Errorf("on loading query: %s", err)
	}
	if err := s.db.QueryRow(query, id).
		Scan(
			&post.ID,
			&post.Title,
			&post.Content,
			&authorID,
			&post.AllowComments,
		); err != nil {
		if err == sql.ErrNoRows {
			return nil, reperrors.ErrPostNotFound
		}
		return nil, fmt.Errorf("failed to get post: %w", err)
	}

	author, err := s.GetUser(authorID)
	if err != nil {
		return nil, fmt.Errorf("failed to get author: %w", err)
	}
	post.Author = author

	return &post, nil
}

func (s *PostgresStorage) GetPosts(limit, offset *int) ([]*model.Post, error) {
	query, err := s.queries.LoadQuery("post", "posts")
	if err != nil {
		return nil, fmt.Errorf("on loading query: %s", err)
	}
	rows, err := s.db.Query(query, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to get posts: %w", err)
	}
	defer rows.Close()

	var posts []*model.Post
	for rows.Next() {
		var post model.Post
		var authorID string
		if err := rows.Scan(&post.ID, &post.Title, &post.Content, &authorID, &post.AllowComments); err != nil {
			return nil, fmt.Errorf("failed to scan post: %w", err)
		}

		author, err := s.GetUser(authorID)
		if err != nil {
			return nil, fmt.Errorf("failed to get author: %w", err)
		}
		post.Author = author

		posts = append(posts, &post)
	}

	return posts, nil
}
