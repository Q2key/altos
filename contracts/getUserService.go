package contracts

import "altos/entities"

type GetUserService interface {
	Execute() []entities.User
}
