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

func (repo *UserDbRepository) GetAll() []entities.User {
	sqlStatement := `SELECT u."firstName" FirstName, u."lastName" LastName FROM public."user" u`
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
		if err := rows.Scan(&user.FirstName, &user.LastName); err != nil {
			log.Println(err.Error())
		}
		users = append(users, user)
	}

	log.Println(users)

	return users
}

func NewUserPGRepo(db *sql.DB) contracts.UserRepository {
	return &UserDbRepository{
		DB: db,
	}
}
