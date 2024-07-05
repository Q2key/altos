package main

import (
	"altos/config"
	"altos/datasource"
	"altos/handlers"
	"altos/repos"
	"altos/services"
	"fmt"
	"net"
	"net/http"
)

func main() {

	cfg := config.NewConfig()

	dataSource := datasource.InitDataSource(cfg)
	userRepository := repos.NewUserPGRepo(dataSource)
	getUsersService := services.NewGetUsersService(userRepository)
	usersHandler := handlers.NewUserHandler(getUsersService)

	address := fmt.Sprintf("%s:%s", cfg.ServerHost, cfg.ServerPort)

	http.HandleFunc("/health", usersHandler.Health)
	http.HandleFunc("/users", usersHandler.GetUsers)
	http.HandleFunc("/create", usersHandler.CreateUser)

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
