package service

import (
	"go-api-template/configuration"
	"go-api-template/repository"
)

type Services struct {
	User IUser
}

func NewServices(
	cfg *configuration.Config,
	repository *repository.Repositories,
) *Services {
	return &Services{
		User: NewUser(cfg, repository.User),
	}
}
