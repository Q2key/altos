package main

import (
	"altos/handlers"
	"altos/repos"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"net/http"
)

func initDataSource() *sql.DB {
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

	defer db.Close()

	return db
}

func main() {
	ds := initDataSource()
	userRepo := repos.NewUserRepo(ds)

	http.HandleFunc("/users", handlers.MakeGetUserHandler(userRepo))

	fmt.Println("Listening on port: 8090")
	err := http.ListenAndServe(":8090", nil)
	if err != nil {
		return
	}
}
