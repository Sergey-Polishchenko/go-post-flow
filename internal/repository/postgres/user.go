package postgres

import (
	"database/sql"
	"fmt"

	"github.com/Sergey-Polishchenko/go-post-flow/internal/delivery/graph/model"
)

func (s *PostgresStorage) CreateUser(input model.UserInput) (*model.User, error) {
	var userID string
	query, err := s.queries.LoadQuery("user", "create")
	if err != nil {
		return nil, fmt.Errorf("on loading query: %s", err)
	}
	if err := s.db.QueryRow(query, input.Name).Scan(&userID); err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}
	return &model.User{ID: userID, Name: input.Name}, nil
}

func (s *PostgresStorage) GetUser(id string) (*model.User, error) {
	var user model.User
	query, err := s.queries.LoadQuery("user", "get")
	if err != nil {
		return nil, fmt.Errorf("on loading query: %s", err)
	}
	if err := s.db.QueryRow(query, id).Scan(&user.ID, &user.Name); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}
		return nil, fmt.Errorf("failed to get user: %w", err)
	}
	return &user, nil
}
