package service

import (
	"go-api-template/configuration"
	"go-api-template/model"
	"go-api-template/repository"
	mapping "go-api-template/service/mappers"

	"github.com/gin-gonic/gin"
	"github.com/go-openapi/strfmt"
)

type IUser interface {
	UserByID(ctx *gin.Context, userID strfmt.UUID4) (*model.UserByIDResponse, error)
}

type user struct {
	cfg            *configuration.Config
	userRepository repository.IUser
}

func NewUser(
	cfg *configuration.Config,
	userRepository repository.IUser,
) IUser {
	return &user{
		cfg:            cfg,
		userRepository: userRepository,
	}
}

func (service *user) UserByID(_ *gin.Context, userID strfmt.UUID4) (*model.UserByIDResponse, error) {
	user, err := service.userRepository.SelectUserByID(userID)
	if err != nil {
		return nil, err
	}

	return &model.UserByIDResponse{
		User: mapping.ToUserModel(user),
	}, nil
}
