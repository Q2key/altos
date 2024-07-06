package services

import (
	"altos/contracts"
	"altos/entities"
)

type CreateUsersService struct {
	userRepository contracts.UserRepository
}

func NewCreateUsersService(r contracts.UserRepository) contracts.CreateUserService {
	return &CreateUsersService{
		r,
	}
}

func (service *CreateUsersService) Execute(input *contracts.CreateUserServiceInput) *entities.User {
	user := entities.User{
		Email:     input.Email,
		FirstName: input.FirstName,
		LastName:  input.LastName,
	}
	return service.userRepository.Create(&user)
}
