package postgres

import (
	"context"

	"github.com/lib/pq"

	"github.com/Sergey-Polishchenko/go-post-flow/internal/transport/graph/model"
	"github.com/Sergey-Polishchenko/go-post-flow/pkg/errors"
)

func (s *PostgresStorage) CreateUser(input model.UserInput) (*model.User, error) {
	query, err := s.queries.LoadQuery("user", "create")
	if err != nil {
		return nil, err
	}

	var userID string
	if err := s.db.QueryRow(query, input.Name).Scan(&userID); err != nil {
		return nil, &errors.SQLCreatingError{Value: err}
	}

	return &model.User{ID: userID, Name: input.Name}, nil
}

func (s *PostgresStorage) GetUsersByIDs(ctx context.Context, ids []string) ([]*model.User, error) {
	if len(ids) == 0 {
		return []*model.User{}, nil
	}

	query, err := s.queries.LoadQuery("user", "get_by_ids")
	if err != nil {
		return nil, err
	}

	rows, err := s.db.QueryContext(ctx, query, pq.Array(ids))
	if err != nil {
		return nil, &errors.SQLQueryError{Value: err}
	}
	defer rows.Close()

	usersMap := make(map[string]*model.User)
	for rows.Next() {
		var user model.User
		if err := rows.Scan(&user.ID, &user.Name); err != nil {
			return nil, &errors.SQLScaningError{Value: err}
		}
		usersMap[user.ID] = &user
	}

	if err = rows.Err(); err != nil {
		return nil, &errors.SQLIterationError{Value: err}
	}

	result := make([]*model.User, len(ids))
	for i, id := range ids {
		result[i] = usersMap[id]
	}

	return result, nil
}

func (s *PostgresStorage) getUser(id string) (*model.User, error) {
	users, err := s.GetUsersByIDs(context.Background(), []string{id})
	if err != nil || len(users) == 0 || users[0] == nil {
		return nil, errors.ErrUserNotFound
	}
	return users[0], nil
}
