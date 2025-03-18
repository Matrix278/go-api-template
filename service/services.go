package service

import (
	"go-api-template/configuration"
	"go-api-template/repository"
)

type Services struct {
	User IUser
}

func NewServices(
	_ *configuration.Env,
	repository *repository.Repositories,
) *Services {
	return &Services{
		User: NewUser(repository.User),
	}
}
