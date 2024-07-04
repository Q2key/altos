package contracts

type DeleteUserServiceInput struct {
	id string
}

type DeleteUserService interface {
	Execute(input DeleteUserServiceInput) bool
}
