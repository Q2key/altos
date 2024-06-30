package repos

import (
	"database/sql"
)

type UserRepo struct {
	DB *sql.DB
}

func (repo *UserRepo) GetUsers() {

}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{
		DB: db,
	}
}
