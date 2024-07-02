package main

import (
	"altos/api"
	"altos/datasource"
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

	http.HandleFunc("/health", api.Health)
	http.HandleFunc("/users", api.MakeUsersHandler(getUsersService))

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
