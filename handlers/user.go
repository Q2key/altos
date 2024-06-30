package handlers

import (
	"altos/repos"
	"fmt"
	"net/http"
)

func MakeGetUserHandler(repo *repos.UserRepo) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		repo.GetUsers()
		_, err := fmt.Fprintf(w, "hello world\n")
		if err != nil {
			return
		}
		return
	}
}

//func MakeDBHandler(fn func(db *sql.DB, w http.ResponseWriter, r *http.Request)) http.HandlerFunc {
//	return func(w http.ResponseWriter, r *http.Request) {
//		db := sql.DB{}
//		fn(&db, w, r)
//	}
//}

//http.HandleFunc("/foo/", MakeDBHandler(myDBHandler))
