package service

import (
	"context"
	"go-api-template/model"

	"github.com/go-openapi/strfmt"
	"github.com/stretchr/testify/mock"
)

type UserMock struct {
	mock.Mock
}

var _ IUser = &UserMock{}

func (mock *UserMock) UserByID(ctx context.Context, userID strfmt.UUID4) (*model.UserByIDResponse, error) {
	args := mock.Called(ctx, userID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.UserByIDResponse), args.Error(1)
}
