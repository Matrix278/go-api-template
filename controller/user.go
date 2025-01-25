package controller

import (
	"errors"
	"go-api-template/service"

	"github.com/gin-gonic/gin"
	"github.com/go-openapi/strfmt"
)

type User struct {
	service service.IUser
}

func NewUser(
	service service.IUser,
) *User {
	return &User{
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
//	@Param			id	path		string					true	"User ID"
//	@Success		200	{object}	model.UserByIDResponse	"Get user by ID"
//	@Failure		400	{object}	swagger.StatusBadRequest
//	@Failure		403	{object}	swagger.StatusForbidden
//	@Failure		500	{object}	swagger.StatusInternalError
//	@Router			/user/{id} [get]
func (controller *User) UserByID(ctx *gin.Context) {
	// Validate path params
	userID := ctx.Param("user_id")
	if !strfmt.IsUUID4(userID) {
		StatusBadRequest(ctx, errors.New("invalid user_id"))
		return
	}

	response, err := controller.service.UserByID(ctx, strfmt.UUID4(userID))
	if err != nil {
		StatusInternalServerError(ctx, err)
		return
	}

	StatusOKWithResponseModel(ctx, response)
}
