package repository

import (
	"go-api-template/model"

	"github.com/go-openapi/strfmt"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type IUser interface {
	SelectUserByID(userID strfmt.UUID4) (*model.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUser {
	return &userRepository{
		db: db,
	}
}

func (repository *userRepository) SelectUserByID(userID strfmt.UUID4) (*model.User, error) {
	var user model.User

	if err := repository.db.Where("id = ?", userID).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, model.ErrUserNotFound
		}

		return nil, errors.Wrap(err, "selecting user by ID failed")
	}

	return &user, nil
}
