package repository

import (
	"context"
	repositorymodel "go-api-template/repository/model"

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

func (mock *UserMock) SelectUserByFilter(ctx context.Context, filter repositorymodel.UsersFilter) (*repositorymodel.User, error) {
	args := mock.Called(ctx, filter)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*repositorymodel.User), args.Error(1)
}
