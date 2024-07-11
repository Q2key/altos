package main

import (
	"altos/cmd/altos/docs"
	"altos/config"
	"altos/datasource"
	"altos/handlers"
	"altos/repos"
	"altos/services"
	"fmt"
	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger/v2"
	"net/http"

	"github.com/go-chi/chi/v5/middleware"
)

// @title Altos API
// @version 1.0

//go:generate swag init --dir "../../cmd/altos,../../handlers,../../mappers" --parseDependency
func main() {

	cfg := config.NewConfig()

	dataSource := datasource.InitDataSource(cfg)
	userRepository := repos.NewUserPGRepo(dataSource)
	getUsersService := services.NewGetUsersService(userRepository)
	createUserService := services.NewCreateUsersService(userRepository)
	usersHandler := handlers.NewUserHandler(getUsersService, createUserService)

	address := fmt.Sprintf("%s:%s", cfg.ServerHost, cfg.ServerPort)

	docs.SwaggerInfo.Host = address
	docs.SwaggerInfo.BasePath = "/"

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	//init swagger/swag
	r.Get("/swagger/*", httpSwagger.Handler(httpSwagger.URL(fmt.Sprintf("http://localhost:%s/swagger/doc.json", cfg.ServerPort))))

	r.Get("/users", usersHandler.GetUsers)
	r.Post("/user", usersHandler.CreateUser)

	err := http.ListenAndServe(address, r)
	if err != nil {
		return
	}
}
