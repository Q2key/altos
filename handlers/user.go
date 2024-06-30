package handlers

import (
	"fmt"
	"net/http"
)

func GetUsers(w http.ResponseWriter, req *http.Request) {
	_, err := fmt.Fprintf(w, "hello world\n")
	if err != nil {
		return
	}
}
