package resolvers

import "github.com/Sergey-Polishchenko/go-post-flow/internal/storage"

type Resolver struct {
	storage storage.Storage
	comLim  int
}

func NewResolver(storage storage.Storage) *Resolver {
	return &Resolver{
		storage: storage,
		comLim:  2000,
	}
}
