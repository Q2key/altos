package services

import (
	"altos/contracts"
	"altos/entities"
)

type GetUsersService struct {
	userRepository contracts.UserRepository
}

func NewGetUsersService(r contracts.UserRepository) contracts.GetUserService {
	return &GetUsersService{
		r,
	}
}

func (service *GetUsersService) Execute() []entities.User {
	return service.userRepository.GetAll()
}
