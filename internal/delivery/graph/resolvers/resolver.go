package resolvers

import "github.com/Sergey-Polishchenko/go-post-flow/internal/repository"

type Resolver struct {
	storage repository.Storage
}

func NewResolver(storage repository.Storage) *Resolver {
	return &Resolver{storage: storage}
}
