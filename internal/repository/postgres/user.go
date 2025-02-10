package postgres

import (
	"database/sql"

	"github.com/Sergey-Polishchenko/go-post-flow/internal/delivery/graph/model"
	"github.com/Sergey-Polishchenko/go-post-flow/internal/errors"
)

func (s *PostgresStorage) CreateUser(input model.UserInput) (*model.User, error) {
	var userID string
	query, err := s.queries.LoadQuery("user", "create")
	if err != nil {
		return nil, err
	}
	if err := s.db.QueryRow(query, input.Name).Scan(&userID); err != nil {
		return nil, &errors.SQLCreatingError{Value: err}
	}
	return &model.User{ID: userID, Name: input.Name}, nil
}

func (s *PostgresStorage) GetUser(id string) (*model.User, error) {
	var user model.User
	query, err := s.queries.LoadQuery("user", "get")
	if err != nil {
		return nil, err
	}
	if err := s.db.QueryRow(query, id).Scan(&user.ID, &user.Name); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.ErrUserNotFound
		}
		return nil, &errors.SQLScaningError{Value: err}
	}
	return &user, nil
}
