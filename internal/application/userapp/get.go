package userapp

import (
	"context"
)

func (app *UserApp) GetUser(ctx context.Context, id string) (UserDTO, error) {
	userDTO, err := app.repo.GetByID(ctx, id)
	if err != nil {
		app.logger.Error("Failed to load user from repo", "error", err)
		return UserDTO{}, err
	}

	// TODO: business logic for User entity.
	//
	// user, err := user.User{
	//	 ID: user.UserID(userDTO.ID),
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
