package datasource

import (
	"altos/config"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func InitDataSource(config *config.Config) *sql.DB {
	const (
		host     = "localhost"
		port     = 5436
		user     = "postgres"
		password = "postgres"
		dbname   = "alto"
	)

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	//defer db.Close()

	return db
}
