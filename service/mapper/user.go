package mapper

import (
	"go-api-template/model"
	repositorymodel "go-api-template/repository/model"
)

func ToUserModel(user *repositorymodel.User) *model.User {
	return &model.User{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}
