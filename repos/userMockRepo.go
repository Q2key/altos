package repos

import (
	"altos/contracts"
	"altos/entities"
	"database/sql"
)

type UserJsonRepository struct {
}

func (repo *UserJsonRepository) GetAll() []entities.User {
	return []entities.User{
		{FirstName: "Dima", LastName: "Totskii"},
	}
}

func NewUserJsonRepository(_ *sql.DB) contracts.UserRepository {
	return &UserJsonRepository{}
}
