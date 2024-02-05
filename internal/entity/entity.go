package entity

import "github.com/google/uuid"

type Category struct {
	ID   string
	NAME string
}

func NewCategory(name string) *Category {
	return &Category{
		ID:   uuid.New().String(),
		NAME: name,
	}
}
