package service

import (
	"errors"
	"go-api-template/configuration"
	"go-api-template/model"
	"go-api-template/model/commonerrors"
	"go-api-template/repository"
	"go-api-template/service/mapper"

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
		if errors.Is(err, commonerrors.ErrUserNotFound) {
			return nil, commonerrors.ErrUserNotFound
		}

		return nil, err
	}

	return &model.UserByIDResponse{
		User: mapper.ToUserModel(user),
	}, nil
}
