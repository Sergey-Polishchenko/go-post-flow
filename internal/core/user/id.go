package user

import "github.com/google/uuid"

type UserID string

func NewUserID() UserID {
	return UserID(uuid.New().String())
}

func (uid UserID) String() string {
	return string(uid)
}
