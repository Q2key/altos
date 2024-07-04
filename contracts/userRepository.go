package contracts

import "altos/entities"

type UserRepository interface {
	GetAll() *[]entities.User
	Create(user *entities.User) *entities.User
}
