package repos

import (
	"altos/contracts"
	"altos/entities"
	"database/sql"
	"log"
)

type UserDbRepository struct {
	DB *sql.DB
}

func (repo *UserDbRepository) GetAll() *[]entities.User {
	sqlStatement := `SELECT u.id as Id,  u."firstName" as FirstName, u."lastName" as LastName, u.email as Email FROM public."user" as u`
	rows, err := repo.DB.Query(sqlStatement)
	if err != nil {
		panic(err)
	}

	users := make([]entities.User, 0)
	for rows.Next() {
		if rows.Err() != nil {
			log.Println("Log error:", rows.Err().Error())
		}

		var user entities.User
		if err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email); err != nil {
			log.Println(err.Error())
		}
		users = append(users, user)
	}

	return &users
}

func (repo *UserDbRepository) Create(user *entities.User) *entities.User {
	return &entities.User{FirstName: "FirstName-1", LastName: "lastName-1"}
}

func NewUserPGRepo(db *sql.DB) contracts.UserRepository {
	return &UserDbRepository{
		DB: db,
	}
}
