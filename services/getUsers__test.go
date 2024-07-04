package services

import (
	"altos/contracts"
	"altos/repos"
	"testing"
)

func TestNewGetUsersServiceShouldReturnUsers(t *testing.T) {
	testMockRepo := repos.NewUserJsonRepository()
	userService := NewGetUsersService(testMockRepo)
	users := userService.Execute(&contracts.GetUserServiceInput{})

	if users == nil {
		t.Error("Expected users not to be nil")
		return
	}

	if len(*users) == 0 {
		t.Error("Expected users not to be empty")
		return
	}
}
