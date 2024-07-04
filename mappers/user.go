package mappers

import (
	"altos/entities"
	"altos/models"
)

type GetUsersDTO struct {
	users []entities.User
}

func NewGetUsersDTO(users *[]entities.User) *GetUsersDTO {
	return &GetUsersDTO{users: *users}
}

func (dto *GetUsersDTO) ToApiResponseDTO() []*models.User {
	usersResponse := make([]*models.User, 0)
	for i := 0; i < len(dto.users); i++ {
		u := dto.users[i]
		usersResponse = append(usersResponse, &models.User{
			Email:     u.Email,
			FirstName: u.FirstName,
			ID:        u.Id,
			LastName:  u.LastName,
		})
	}

	return usersResponse
}
