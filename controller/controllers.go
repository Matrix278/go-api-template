package controller

import (
	"go-api-template/service"
)

type Controllers struct {
	User IUser
}

func NewControllers(services *service.Services) *Controllers {
	return &Controllers{
		User: NewUser(services.User),
	}
}
