package redisrepo

import (
	"github.com/redis/go-redis/v9"

	"github.com/Sergey-Polishchenko/go-post-flow/internal/application/commentapp"
	"github.com/Sergey-Polishchenko/go-post-flow/internal/application/postapp"
	"github.com/Sergey-Polishchenko/go-post-flow/internal/application/userapp"
)

type Factory struct {
	client *redis.Client
}

func New(client *redis.Client) *Factory {
	return &Factory{client: client}
}

func (f *Factory) NewUserRepo() userapp.UserRepository {
	return &UserRepository{f.client}
}

func (f *Factory) NewCommentRepo() commentapp.CommentRepository {
	return &CommentRepository{f.client}
}

func (f *Factory) NewPostRepo() postapp.PostRepository {
	return &PostRepository{f.client}
}
