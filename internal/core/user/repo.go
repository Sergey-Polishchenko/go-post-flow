package user

// UserRepository provides persistence for users.
//
// Implementations must guarantee:
//   - Idempotency of the Create method.
//   - Thread safety.
type UserRepository interface {
	Create(user *User) error
	Remove(id string) error
	GetById(id string) (*User, error)
}
