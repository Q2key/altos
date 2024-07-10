package handlers

import (
	"altos/contracts"
	"altos/entities"
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
	WriteOKWithJson[[]entities.User](w, *users)
}

// Health godoc
// @Summary       Health check
// @Router       /health [get]
func (h *UserHandler) Health(w http.ResponseWriter, r *http.Request) {
	WriteOKWithJson[map[string]string](w, map[string]string{"message": "pong"})
}
