package commentapp

import "github.com/Sergey-Polishchenko/go-post-flow/internal/core/comment"

type CommentApp struct {
	repo comment.CommentRepository
}

func New(repo comment.CommentRepository) *CommentApp {
	return &CommentApp{repo: repo}
}
