package main

import (
	"altos/handlers"
	"net/http"
)

func main() {
	http.HandleFunc("/users", handlers.GetUsers)

	err := http.ListenAndServe(":8090", nil)
	if err != nil {
		return
	}

}
