package id

import "github.com/google/uuid"

type Identifier interface {
	String() string
}

type ID string

func New() ID {
	return ID(uuid.New().String())
}

func (id ID) String() string {
	return string(id)
}
