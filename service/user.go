package service

import (
	"go-api-template/model"
	"go-api-template/repository"
	repositorymodel "go-api-template/repository/model"
	"go-api-template/service/mapper"

	"github.com/gin-gonic/gin"
	"github.com/go-openapi/strfmt"
)

type IUser interface {
	UserByID(ctx *gin.Context, userID strfmt.UUID4) (*model.UserByIDResponse, error)
}

type user struct {
	userRepository repository.IUser
}

func NewUser(
	userRepository repository.IUser,
) IUser {
	return &user{
		userRepository: userRepository,
	}
}

func (service *user) UserByID(_ *gin.Context, userID strfmt.UUID4) (*model.UserByIDResponse, error) {
	filter := repositorymodel.UsersFilter{
		ID: &userID,
	}

	user, err := service.userRepository.SelectUserByFilter(filter)
	if err != nil {
		return nil, err
	}

	return &model.UserByIDResponse{
		User: mapper.ToUserModel(user),
	}, nil
}
