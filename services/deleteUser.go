package services

import (
	"altos/contracts"
	"altos/entities"
)

type DeleteUserService struct {
	userRepository contracts.UserRepository
}

func (service *DeleteUserService) Execute(input *contracts.DeleteUserServiceInput) *entities.User {
	return service.userRepository.Delete(0)
}
