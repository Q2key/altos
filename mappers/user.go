package mappers

import (
	"altos/entities"
)

type GetUsersDTO struct {
	users []entities.User
}

func NewGetUsersDTO(users *[]entities.User) *GetUsersDTO {
	return &GetUsersDTO{users: *users}
}
