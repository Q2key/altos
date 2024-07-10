package main

import (
	"altos/cmd/altos/docs"
	"altos/config"
	"altos/datasource"
	"altos/handlers"
	"altos/repos"
	"altos/services"
	"fmt"
	httpSwagger "github.com/swaggo/http-swagger/v2"
	"net"
	"net/http"
)

// @title Altos API
// @version 1.0

//go:generate swag init --dir "../../cmd/altos,../../handlers,../../mappers" --parseDependency
func main() {

	cfg := config.NewConfig()

	dataSource := datasource.InitDataSource(cfg)
	userRepository := repos.NewUserPGRepo(dataSource)
	getUsersService := services.NewGetUsersService(userRepository)
	usersHandler := handlers.NewUserHandler(getUsersService)

	address := fmt.Sprintf("%s:%s", cfg.ServerHost, cfg.ServerPort)

	docs.SwaggerInfo.Host = address
	docs.SwaggerInfo.BasePath = "/"

	//init swagger/swag
	http.HandleFunc("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL(fmt.Sprintf("http://localhost:%s/swagger/doc.json", cfg.ServerPort)))) //The url pointing to API definition

	http.HandleFunc("/health", usersHandler.Health)
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
