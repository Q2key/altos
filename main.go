package main

import (
	"altos/datasource"
	"altos/handlers"
	"altos/repos"
	"altos/services"
	"fmt"
	"net"
	"net/http"
)

func main() {
	dataSource := datasource.InitDataSource()
	userRepository := repos.NewUserPGRepo(dataSource)
	getUsersService := services.NewGetUsersService(userRepository)

	http.HandleFunc("/users", handlers.MakeUsersHandler(getUsersService))

	port := 8090
	url := "localhost"
	address := fmt.Sprintf("%s:%d", url, port)

	l, err := net.Listen("tcp", address)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Listening on port %s ", address)
	err = http.Serve(l, nil)
	if err != nil {
		fmt.Println(err)
	}
}
