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

func EnableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With")
}

func (h *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)
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
