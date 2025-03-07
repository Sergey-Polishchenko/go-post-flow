package userapp

import (
	"context"

	"github.com/Sergey-Polishchenko/go-post-flow/internal/core/id"
	"github.com/Sergey-Polishchenko/go-post-flow/internal/core/user"
)

func (app *UserApp) CreateUser(ctx context.Context, createUserDTO CreateUserDTO) (UserDTO, error) {
	id := id.New()
	user, err := user.New(id, user.UserName(createUserDTO.Name))
	if err != nil {
		app.logger.Error("Failed to validate user", "error", err)
		return UserDTO{}, err
	}

	if err = app.repo.Save(ctx, user); err != nil {
		app.logger.Error("Failed to save user in repo", "error", err)
		return UserDTO{}, err
	}

	app.logger.Info("User created", "id", user.ID.String())

	return UserDTO{
		ID:   user.ID.String(),
		Name: user.Name().String(),
	}, nil
}
