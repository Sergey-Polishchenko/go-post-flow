package postgres

import (
	"context"
	"time"

	"github.com/lib/pq"

	"github.com/Sergey-Polishchenko/go-post-flow/internal/delivery/graph/model"
	"github.com/Sergey-Polishchenko/go-post-flow/pkg/errors"
)

func (s *PostgresStorage) CreateComment(input model.CommentInput) (*model.Comment, error) {
	if input.ParentID != nil {
		parent, err := s.getComment(*input.ParentID)
		if err != nil || parent == nil {
			return nil, errors.ErrParentCommentNotFound
		}
		if parent.Post == nil || parent.Post.ID != input.PostID {
			return nil, errors.ErrParentInOtherPost
		}
	}

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

	author, err := s.getUser(input.AuthorID)
	if err != nil || author == nil {
		return nil, errors.ErrAuthorNotFound
	}

	post, err := s.getPost(input.PostID)
	if err != nil || post == nil {
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

func (s *PostgresStorage) GetCommentsByIDs(
	ctx context.Context,
	ids []string,
) ([]*model.Comment, error) {
	if len(ids) == 0 {
		return []*model.Comment{}, nil
	}

	query, err := s.queries.LoadQuery("comment", "get_by_ids")
	if err != nil {
		return nil, err
	}

	rows, err := s.db.QueryContext(ctx, query, pq.Array(ids))
	if err != nil {
		return nil, &errors.SQLQueryError{Value: err}
	}
	defer rows.Close()

	commentsMap := make(map[string]*model.Comment)
	var postIDs []string
	var authorIDs []string

	for rows.Next() {
		var comment model.Comment
		var postID, authorID string
		var createdAt time.Time

		if err := rows.Scan(
			&comment.ID,
			&comment.Text,
			&authorID,
			&postID,
			&createdAt,
		); err != nil {
			return nil, &errors.SQLScaningError{Value: err}
		}

		comment.CreatedAt = createdAt.Format(time.RFC3339)
		comment.Post = &model.Post{ID: postID}
		comment.Author = &model.User{ID: authorID}
		commentsMap[comment.ID] = &comment
		postIDs = append(postIDs, postID)
		authorIDs = append(authorIDs, authorID)
	}

	posts, _ := s.GetPostsByIDs(ctx, postIDs)
	authors, _ := s.GetUsersByIDs(ctx, authorIDs)

	postMap := make(map[string]*model.Post)
	for _, post := range posts {
		if post != nil {
			postMap[post.ID] = post
		}
	}

	authorMap := make(map[string]*model.User)
	for _, author := range authors {
		if author != nil {
			authorMap[author.ID] = author
		}
	}

	for _, comment := range commentsMap {
		if post, ok := postMap[comment.Post.ID]; ok {
			comment.Post = post
		}
		if author, ok := authorMap[comment.Author.ID]; ok {
			comment.Author = author
		}
	}

	result := make([]*model.Comment, len(ids))
	for i, id := range ids {
		result[i] = commentsMap[id]
	}

	return result, nil
}

func (s *PostgresStorage) GetCommentsIDs(postID string) ([]string, error) {
	query, err := s.queries.LoadQuery("comment", "comments")
	if err != nil {
		return nil, err
	}

	rows, err := s.db.Query(query, postID)
	if err != nil {
		return nil, &errors.SQLQueryError{Value: err}
	}
	defer rows.Close()

	var ids []string
	for rows.Next() {
		var id string
		if err := rows.Scan(&id); err != nil {
			return nil, &errors.SQLScaningError{Value: err}
		}
		ids = append(ids, id)
	}

	return ids, nil
}

func (s *PostgresStorage) GetChildrenIDs(commentID string) ([]string, error) {
	query, err := s.queries.LoadQuery("comment", "children")
	if err != nil {
		return nil, err
	}

	rows, err := s.db.Query(query, commentID)
	if err != nil {
		return nil, &errors.SQLQueryError{Value: err}
	}
	defer rows.Close()

	var ids []string
	for rows.Next() {
		var id string
		if err := rows.Scan(&id); err != nil {
			return nil, &errors.SQLScaningError{Value: err}
		}
		ids = append(ids, id)
	}

	return ids, nil
}

func (s *PostgresStorage) getComment(id string) (*model.Comment, error) {
	comments, err := s.GetCommentsByIDs(context.Background(), []string{id})
	if err != nil || len(comments) == 0 || comments[0] == nil {
		return nil, errors.ErrCommentsNotFound
	}
	return comments[0], nil
}
