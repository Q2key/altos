package services

import (
	"altos/contracts"
)

type AuthenticateUserService struct {
	userRepository contracts.UserRepository
}

func NewAuthenticateUserService(r contracts.UserRepository) contracts.AuthenticateUserService {
	return &AuthenticateUserService{
		r,
	}
}

func (service *AuthenticateUserService) Execute(input *contracts.AuthenticateServiceInput) bool {
	return false
}
