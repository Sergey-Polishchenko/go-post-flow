package query

import (
	"sync"

	"github.com/Sergey-Polishchenko/go-post-flow/internal/errors"
)

type QueryCache struct {
	queries map[string]string
	mu      sync.RWMutex
}

func NewQueryCache() *QueryCache {
	return &QueryCache{
		queries: make(map[string]string),
	}
}

func (c *QueryCache) LoadQuery(target, method string) (string, error) {
	c.mu.RLock()
	if query, ok := c.queries[target+method]; ok {
		c.mu.RUnlock()
		return query, nil
	}
	c.mu.RUnlock()

	query, err := loadQuery(target, method)
	if err != nil {
		return "", &errors.SQLQueryLoadingError{Value: err}
	}

	c.mu.Lock()
	c.queries[target+method] = query
	c.mu.Unlock()

	return query, nil
}
