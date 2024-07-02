package contracts

import "altos/entities"

type UserRepository interface {
	GetAll() []entities.User
}
