package dto

import (
	"context"
)

type userBase struct {
	Name         string `json:"name" validate:"min=3,max=100"`
	Phone_number string `json:"phone_number" validate:"omitempty,min=8,max=15,startswith=+"`
}

type CreateUserRequest struct {
	userBase
}

type UserResponse struct {
	ID int `json:"id"`
	userBase
}

type UserRepository interface {
	SaveUser(ctx context.Context, user *CreateUserRequest) error
}

type ProductDTO struct {
	id          int
	title       string
	description string
	price       int
	category    string
}
