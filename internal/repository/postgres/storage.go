package postgres

import (
	"database/sql"
	"sync"

	_ "github.com/lib/pq"

	"github.com/Sergey-Polishchenko/go-post-flow/internal/pkg/broadcast"
	"github.com/Sergey-Polishchenko/go-post-flow/internal/repository/postgres/query"
	"github.com/Sergey-Polishchenko/go-post-flow/pkg/errors"
)

type PostgresStorage struct {
	mu      sync.RWMutex
	db      *sql.DB
	queries query.QueryCache
	*broadcast.Broadcast
}

func NewStorage(connStr string) (*PostgresStorage, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, &errors.DatabaseConnectingError{Value: err}
	}
	if err = db.Ping(); err != nil {
		return nil, errors.ErrPingDatabase
	}
	pg := &PostgresStorage{
		db:        db,
		queries:   *query.NewQueryCache(),
		Broadcast: broadcast.NewBroadcast(),
	}
	return pg, nil
}
