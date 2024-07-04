package main

import (
	"altos/api"
	"altos/config"
	"altos/datasource"
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
	usersHandler := api.NewUserHandler(getUsersService)

	address := fmt.Sprintf("%s:%s", cfg.ServerHost, cfg.ServerPort)

	http.HandleFunc("/ping", usersHandler.Ping)
	http.HandleFunc("/users", usersHandler.GetUsers)

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
