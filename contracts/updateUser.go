package contracts

import "altos/entities"

type UpdateUseInput struct {
	Id        *string `json:"id"`
	FirstName *string `json:"firstName"`
	LastName  *string `json:"lastName"`
	Email     *string `json:"email"`
}

type UpdateUserService interface {
	Execute(*UpdateUseInput) *entities.User
}
