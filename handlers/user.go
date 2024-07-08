package handlers

import (
	"altos/contracts"
	"altos/entities"
	"encoding/json"
	"net/http"
)

type UsersResponse *[]entities.User

type UserHandler struct {
	getUsersService contracts.GetUsersService
}

func NewUserHandler(getUsersService contracts.GetUsersService) *UserHandler {
	return &UserHandler{
		getUsersService: getUsersService,
	}
}

// GetUsers godoc
// @Summary      List accounts
// @Description  get accounts
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Param        q    query     string  false  "name search by q"  Format(email)
// @Success      200  {array}   mappers.GetUsersDTO
// @Router       /users [get]
func (h *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	users := h.getUsersService.Execute(&contracts.GetUserServiceInput{})
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(users)
	if err != nil {
		return
	}
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "pong"})
}

func (h *UserHandler) Health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "pong"})
}
