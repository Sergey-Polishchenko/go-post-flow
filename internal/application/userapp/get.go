package userapp

import (
	"context"

	"github.com/Sergey-Polishchenko/go-post-flow/internal/core/id"
)

func (app *UserApp) GetUser(ctx context.Context, userID string) (UserDTO, error) {
	userDTO, err := app.repo.GetByID(ctx, id.ID(userID))
	if err != nil {
		app.logger.Error("Failed to load user from repo", "error", err)
		return UserDTO{}, err
	}

	// TODO: business logic for User entity.
	//
	// user, err := user.User{
	//	 ID: id.ID(userDTO.ID),
	//	 Name: user.UserName(userDTO.Name)),
	// }
	// if err != nil {
	// 	return UserDTO{}, err
	// }
	//
	// Any business logic for user...

	app.logger.Info("User got", "id", userDTO.ID)

	return userDTO, nil
}
