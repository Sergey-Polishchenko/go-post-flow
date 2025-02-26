package postgres

import (
	"context"

	"github.com/lib/pq"

	"github.com/Sergey-Polishchenko/go-post-flow/internal/delivery/graph/model"
	errors "github.com/Sergey-Polishchenko/go-post-flow/pkg/errors"
)

func (s *PostgresStorage) CreatePost(input model.PostInput) (*model.Post, error) {
	var postID string
	query, err := s.queries.LoadQuery("post", "create")
	if err != nil {
		return nil, err
	}
	if err := s.db.QueryRow(
		query,
		input.Title,
		input.Content,
		input.AuthorID,
		input.AllowComments,
	).Scan(&postID); err != nil {
		return nil, &errors.SQLCreatingError{Value: err}
	}

	author, err := s.getUser(input.AuthorID)
	if err != nil {
		return nil, errors.ErrAuthorNotFound
	}

	return &model.Post{
		ID:            postID,
		Title:         input.Title,
		Content:       input.Content,
		Author:        author,
		AllowComments: input.AllowComments,
	}, nil
}

func (s *PostgresStorage) GetPosts(limit, offset *int) ([]*model.Post, error) {
	defaultLimit := 100
	if limit == nil {
		limit = &defaultLimit
	}

	defaultOffset := 0
	if offset == nil {
		offset = &defaultOffset
	}

	query, err := s.queries.LoadQuery("post", "posts")
	if err != nil {
		return nil, err
	}

	rows, err := s.db.Query(query, *limit, *offset)
	if err != nil {
		return nil, &errors.SQLQueryError{Value: err}
	}
	defer rows.Close()

	var posts []*model.Post
	var authorIDs []string
	postsMap := make(map[string]*model.Post)

	for rows.Next() {
		var post model.Post
		var authorID string

		if err := rows.Scan(
			&post.ID,
			&post.Title,
			&post.Content,
			&authorID,
			&post.AllowComments,
		); err != nil {
			return nil, &errors.SQLScaningError{Value: err}
		}

		authorIDs = append(authorIDs, authorID)
		postsMap[authorID] = &post
		posts = append(posts, &post)
	}

	authors, err := s.GetUsersByIDs(context.Background(), authorIDs)
	if err != nil {
		return nil, err
	}

	for i, post := range posts {
		if authors[i] != nil {
			post.Author = authors[i]
		}
	}

	return posts, nil
}

func (s *PostgresStorage) GetPostsByIDs(ctx context.Context, ids []string) ([]*model.Post, error) {
	if len(ids) == 0 {
		return []*model.Post{}, nil
	}

	query, err := s.queries.LoadQuery("post", "get_by_ids")
	if err != nil {
		return nil, err
	}

	rows, err := s.db.QueryContext(ctx, query, pq.Array(ids))
	if err != nil {
		return nil, &errors.SQLQueryError{Value: err}
	}
	defer rows.Close()

	var posts []*model.Post
	authorIDs := make([]string, 0, len(ids))

	for rows.Next() {
		var post model.Post
		var authorID string

		if err := rows.Scan(
			&post.ID,
			&post.Title,
			&post.Content,
			&post.AllowComments,
			&authorID,
		); err != nil {
			return nil, &errors.SQLScaningError{Value: err}
		}

		post.Author = &model.User{ID: authorID}
		authorIDs = append(authorIDs, authorID)
		posts = append(posts, &post)
	}

	if err = rows.Err(); err != nil {
		return nil, &errors.SQLIterationError{Value: err}
	}

	users, err := s.GetUsersByIDs(ctx, authorIDs)
	if err != nil {
		return nil, err
	}

	for i, post := range posts {
		if users[i] != nil {
			post.Author = users[i]
		}
	}

	return posts, nil
}

func (s *PostgresStorage) getPost(id string) (*model.Post, error) {
	posts, err := s.GetPostsByIDs(context.Background(), []string{id})
	if err != nil || len(posts) == 0 || posts[0] == nil {
		return nil, errors.ErrPostNotFound
	}
	return posts[0], nil
}
