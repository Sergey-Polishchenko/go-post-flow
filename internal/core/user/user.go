// Package user provides domain models and operations related to user management.
// It includes the definition of the User entity, validation rules for usernames,
// and basic operations such as creating and retrieving user information.
//
// Key Components:
//   - User: Represents a user in the system with a validated username.
//   - UserName: A validated string type representing a user's name.
//   - New: A constructor for creating valid User instances.
//   - UserRepository: An interface for persisting and retrieving User entities.
//
// Validation Rules:
//   - Usernames must be between 1 and 100 characters long.
//   - Usernames must consist of valid UTF-8 characters (no control characters).
//
// Errors:
//   - ErrInvalidUsername: Returned when a username fails validation.
package user

// User represents a user account in the system.
// The object guarantees the validity of its name (UserName) through checks in the constructor New.
// The structure is immutable: once created, its fields cannot be modified.
type User struct {
	name UserName
}

// New creates a new User object.
//
// Parameters:
//   - name: A valid username (type UserName).
//
// Returns:
//   - *User: A pointer to the created user if the name passes validation.
//   - error: An error (ErrInvalidUsername) if the name is invalid.
func New(name UserName) (*User, error) {
	if !name.IsValid() {
		return nil, ErrInvalidUsername
	}

	return &User{name: name}, nil
}

// Name returns the user's name.
// The method provides read-only access to the name field.
func (user *User) Name() UserName {
	return user.name
}
