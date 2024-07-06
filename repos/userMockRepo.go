package repos

import (
	"altos/contracts"
	"altos/entities"
)

type UserJsonRepository struct {
}

func (r *UserJsonRepository) GetAll() *[]entities.User {
	return &[]entities.User{
		{FirstName: "FirstName-1", LastName: "lastName-1"},
		{FirstName: "FirstName-2", LastName: "LastName-2"},
		{FirstName: "FirstName-3", LastName: "LastName-3"},
	}
}

func (r *UserJsonRepository) Create(user *entities.User) *entities.User {
	return &entities.User{FirstName: "FirstName-1", LastName: "lastName-1"}
}

func (r *UserJsonRepository) Update(user *entities.User) *entities.User {
	return &entities.User{FirstName: "FirstName-1", LastName: "lastName-1"}
}

func (r *UserJsonRepository) Delete(userId int) *entities.User {
	return &entities.User{FirstName: "FirstName-1", LastName: "lastName-1"}
}

func NewUserJsonRepository() contracts.UserRepository {
	return &UserJsonRepository{}
}
