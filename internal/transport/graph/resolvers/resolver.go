package resolvers

import (
	"github.com/Sergey-Polishchenko/go-post-flow/internal/application/userapp"
)

type Resolver struct {
	userApp *userapp.UserApp
	comLim  int
}

func NewResolver(userapp *userapp.UserApp) *Resolver {
	return &Resolver{
		userApp: userapp,
	}
}
