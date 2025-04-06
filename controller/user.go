package controller

import (
	"go-api-template/model/commonerrors"
	"go-api-template/service"

	"github.com/gin-gonic/gin"
	"github.com/go-openapi/strfmt"
)

type IUser interface {
	UserByID(ctx *gin.Context)
}

type user struct {
	service service.IUser
}

func NewUser(service service.IUser) IUser {
	return &user{
		service: service,
	}
}

// UserByID godoc
//
//	@Summary		Get user by ID
//	@Description	Get user by ID
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			user_id	path		string							true	"User ID"	format(uuid)
//	@Success		200		{object}	model.UserByIDResponseSwagger	"Get user by ID"
//	@Failure		400		{object}	model.BadRequestResponse
//	@Failure		403		{object}	model.ForbiddenResponse
//	@Failure		500		{object}	model.InternalErrorResponse
//	@Router			/users/{user_id} [get]
func (controller *user) UserByID(ctx *gin.Context) {
	// Validate path params
	userID := ctx.Param("user_id")
	if !strfmt.IsUUID4(userID) {
		StatusBadRequest(ctx, commonerrors.ErrInvalidUserID)

		return
	}

	response, err := controller.service.UserByID(ctx, strfmt.UUID4(userID))
	if err != nil {
		HandleCommonErrors(ctx, err)

		return
	}

	StatusOKWithResponseModel(ctx, response)
}
