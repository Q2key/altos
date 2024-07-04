package repos

import (
	"altos/contracts"
	"altos/entities"
)

type UserJsonRepository struct {
}

func (repo *UserJsonRepository) GetAll() *[]entities.User {
	return &[]entities.User{
		{FirstName: "FirstName-1", LastName: "lastName-1"},
		{FirstName: "FirstName-2", LastName: "LastName-2"},
		{FirstName: "FirstName-3", LastName: "LastName-3"},
	}
}

func (repo *UserJsonRepository) Create(user *entities.User) *entities.User {
	return &entities.User{FirstName: "FirstName-1", LastName: "lastName-1"}
}

func (repo *UserJsonRepository) GetOne(id string) *entities.User {
	return &entities.User{FirstName: "FirstName-1", LastName: "lastName-1"}
}

func (repo *UserJsonRepository) DeleteOne(id string) bool {
	return true
}

func NewUserJsonRepository() contracts.UserRepository {
	return &UserJsonRepository{}
}
