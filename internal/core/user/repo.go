package user

// UserRepository defines persistence operations for Users.
type UserRepository interface {
	Create(user *User) error
	Remove(id string) error
	GetById(id string) (*User, error)
}
