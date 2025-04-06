package service

import (
	"context"
	"go-api-template/model"
	"go-api-template/repository"
	repositorymodel "go-api-template/repository/model"
	"go-api-template/service/mapper"

	"github.com/go-openapi/strfmt"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

type IUser interface {
	UserByID(ctx context.Context, userID strfmt.UUID4) (*model.UserByIDResponse, error)
}

type user struct {
	userRepository repository.IUser
	tracer         trace.Tracer
}

func NewUser(
	userRepository repository.IUser,
) IUser {
	return &user{
		userRepository: userRepository,
		tracer:         otel.Tracer("service/user"),
	}
}

func (service *user) UserByID(ctx context.Context, userID strfmt.UUID4) (*model.UserByIDResponse, error) {
	ctx, span := service.tracer.Start(ctx, "UserByID")
	defer span.End()

	span.SetAttributes(attribute.String("user.id", userID.String()))

	filter := repositorymodel.UsersFilter{
		ID: &userID,
	}

	user, err := service.userRepository.SelectUserByFilter(ctx, filter)
	if err != nil {
		span.RecordError(err)
		return nil, err
	}

	return &model.UserByIDResponse{
		User: mapper.ToUserModel(user),
	}, nil
}
