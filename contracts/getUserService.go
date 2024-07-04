package contracts

import "altos/entities"

type GetUserServiceInput struct {
	where map[string]interface{}
}

type GetUserServiceOutput []entities.User

type GetUsersService interface {
	Execute(input *GetUserServiceInput) *[]entities.User
}
