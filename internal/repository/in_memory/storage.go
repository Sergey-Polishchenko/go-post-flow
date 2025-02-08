package inmemory

import (
	"sync"

	"github.com/Sergey-Polishchenko/go-post-flow/internal/delivery/graph/model"
)

type InMemoryStorage struct {
	mu              sync.RWMutex
	Posts           map[string]*model.Post
	Comments        map[string]*model.Comment
	Users           map[string]*model.User
	commentChannels map[string][]chan *model.Comment
}

func NewStorage() *InMemoryStorage {
	return &InMemoryStorage{
		Posts:           make(map[string]*model.Post),
		Comments:        make(map[string]*model.Comment),
		Users:           make(map[string]*model.User),
		commentChannels: make(map[string][]chan *model.Comment),
	}
}
