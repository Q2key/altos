package services

import (
	"altos/contracts"
	"altos/entities"
)

type GetUsersService struct {
	userRepository contracts.UserRepository
}

func NewGetUsersService(r contracts.UserRepository) contracts.GetUsersService {
	return &GetUsersService{
		r,
	}

}

func (service *GetUsersService) Execute(input *contracts.GetUserServiceInput) *[]entities.User {
	return service.userRepository.GetAll()
}
