package contracts

import "altos/entities"

type CreateUserServiceInput struct {
	FirstName string
	LastName  string
	Email     string
	Password  string
}

type CreateUserService interface {
	Execute(input *CreateUserServiceInput) *entities.User
}
