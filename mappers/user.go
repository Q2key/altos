package mappers

import (
	"altos/entities"
	"altos/models"
)

func NewGetUsersApiResponseDTO(user entities.User) *models.User {
	return &models.User{
		Email:     user.LastName,
		FirstName: user.LastName,
		ID:        0,
		LastName:  user.LastName,
	}
}
