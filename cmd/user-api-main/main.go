// swagger:meta
package main

import (
	"altos/api"
	"altos/config"
	"altos/datasource"
	"altos/entities"
	"altos/repos"
	"altos/services"
	"fmt"
	"net"
	"net/http"

	_ "github.com/go-openapi/spec"
	_ "github.com/go-openapi/swag"
)

type WhereParams = map[string]string

// swagger:parameters GetUsersParams
type GetUsersParams struct {
	Offset int
	Limit  int
}

type GetUsersResponse struct {
	Offset int
	Limit  int
	Total  int
	Users  []entities.User
}

func main() {
	cfg := config.NewConfig()
	dataSource := datasource.InitDataSource(cfg)
	userRepository := repos.NewUserPGRepo(dataSource)
	getUsersService := services.NewGetUsersService(userRepository)

	http.HandleFunc("/users", api.MakeUsersHandler(getUsersService))

	address := fmt.Sprintf("%s:%s", cfg.ServerHost, cfg.ServerPort)

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
