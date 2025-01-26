package repository

import (
	repositorymodel "go-api-template/repository/model"

	"github.com/go-openapi/strfmt"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/mock"
)

type UserMock struct {
	mock.Mock
}

var _ IUser = &UserMock{}

func (mock *UserMock) Begin() (*sqlx.Tx, error) {
	args := mock.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*sqlx.Tx), args.Error(1)
}

func (mock *UserMock) SelectUserByID(userID strfmt.UUID4) (*repositorymodel.User, error) {
	args := mock.Called(userID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*repositorymodel.User), args.Error(1)
}
