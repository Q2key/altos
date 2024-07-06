package services

import (
	"altos/contracts"
	"altos/entities"
)

type UpdateUserService struct {
	userRepository contracts.UserRepository
}

func (service *UpdateUserService) Execute(input *contracts.UpdateUseInput) *entities.User {
	user := entities.User{}
	return service.userRepository.Create(&user)
}
