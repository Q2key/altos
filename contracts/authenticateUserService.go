package contracts

type AuthenticateServiceInput struct {
	login    string
	password string
}

type AuthenticateUserService interface {
	Execute(input *AuthenticateServiceInput) bool
}
