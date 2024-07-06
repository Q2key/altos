package contracts

import "altos/entities"

type DeleteUserServiceInput struct {
	id string
}

type DeleteUserService interface {
	Execute(input *DeleteUserServiceInput) *entities.User
}
