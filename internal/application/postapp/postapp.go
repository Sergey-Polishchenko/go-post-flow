package postapp

import (
	"github.com/Sergey-Polishchenko/go-post-flow/internal/core/post"
)

type PostApp struct {
	repo post.PostRepository
}

func New(repo post.PostRepository) *PostApp {
	return &PostApp{repo: repo}
}
