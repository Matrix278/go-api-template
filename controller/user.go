package controller

import (
	"errors"
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
//	@Success		200		{object}	swagger.UserByIDResponseSwagger	"Get user by ID"
//	@Failure		400		{object}	swagger.StatusBadRequest
//	@Failure		403		{object}	swagger.StatusForbidden
//	@Failure		500		{object}	swagger.StatusInternalError
//	@Router			/users/{user_id} [get]
func (controller *user) UserByID(ctx *gin.Context) {
	// Validate path params
	userID := ctx.Param("user_id")
	if !strfmt.IsUUID4(userID) {
		StatusBadRequest(ctx, errors.New("invalid user_id"))
		return
	}

	response, err := controller.service.UserByID(ctx, strfmt.UUID4(userID))
	if err != nil {
		HandleCommonErrors(ctx, err)
		return
	}

	StatusOKWithResponseModel(ctx, response)
}
