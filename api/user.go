package api

import (
	"altos/contracts"
	"encoding/json"
	"net/http"
)

func MakeUsersHandler(userService contracts.GetUserService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		users := userService.Execute()

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		err := json.NewEncoder(w).Encode(users)
		if err != nil {
			return
		}
	}
}

func Health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	if err != nil {
		return
	}
}
