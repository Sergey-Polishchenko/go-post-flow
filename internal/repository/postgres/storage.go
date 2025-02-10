package postgres

import (
	"database/sql"
	"fmt"
	"sync"

	_ "github.com/lib/pq"

	"github.com/Sergey-Polishchenko/go-post-flow/internal/delivery/graph/model"
	"github.com/Sergey-Polishchenko/go-post-flow/internal/repository/postgres/query"
)

type PostgresStorage struct {
	mu              sync.RWMutex
	db              *sql.DB
	commentChannels map[string][]chan *model.Comment
	queries         query.QueryCache
}

func NewStorage(connStr string) (*PostgresStorage, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}
	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}
	pg := &PostgresStorage{
		db:              db,
		commentChannels: make(map[string][]chan *model.Comment),
		queries:         *query.NewQueryCache(),
	}
	return pg, nil
}
