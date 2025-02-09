package postgres

import (
	"database/sql"
	"fmt"

	"github.com/Sergey-Polishchenko/go-post-flow/internal/delivery/graph/model"
)

func (s *PostgresStorage) CreateUser(input model.UserInput) (*model.User, error) {
	var userID string
	query := `INSERT INTO users (name) VALUES ($1) RETURNING id`
	err := s.db.QueryRow(query, input.Name).Scan(&userID)
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}
	return &model.User{ID: userID, Name: input.Name}, nil
}

func (s *PostgresStorage) GetUser(id string) (*model.User, error) {
	var user model.User
	query := `SELECT id, name FROM users WHERE id = $1`
	err := s.db.QueryRow(query, id).Scan(&user.ID, &user.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}
		return nil, fmt.Errorf("failed to get user: %w", err)
	}
	return &user, nil
}
