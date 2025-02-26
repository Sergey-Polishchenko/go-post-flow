package inmemory

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewStorage(t *testing.T) {
	storage := NewStorage()
	assert.NotNil(t, storage)
	assert.NotNil(t, storage.Posts)
	assert.NotNil(t, storage.Comments)
	assert.NotNil(t, storage.Users)
	assert.NotNil(t, storage.Broadcast)
}

func intPtr(i int) *int {
	return &i
}
