package contracts

import "altos/entities"

type UserRepository interface {
	GetAll() *[]entities.User
	Create(user *entities.User) *entities.User
	Update(user *entities.User) *entities.User
	Delete(userId int) *entities.User
}
