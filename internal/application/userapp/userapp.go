package userapp

import "github.com/Sergey-Polishchenko/go-post-flow/internal/core/user"

type UserApp struct {
	repo user.UserRepository
}

func New(repo user.UserRepository) *UserApp {
	return &UserApp{repo: repo}
}
