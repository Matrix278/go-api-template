package controller

import (
	"errors"
	"go-api-template/model"
	"go-api-template/model/commonerrors"
	"go-api-template/pkg/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

func StatusOK(ctx *gin.Context, data []byte) {
	var response interface{}
	if err := json.Decode(data, &response); err != nil {
		StatusInternalServerError(ctx, err)

		return
	}

	ctx.JSON(http.StatusOK, response)
}

func StatusOKWithResponseModel(ctx *gin.Context, response interface{}) {
	ctx.JSON(http.StatusOK, response)
}

func StatusOKWithOutDecode(ctx *gin.Context, message string) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": message,
		"code":    "OK",
	})
}

func StatusCreatedWithOutDecode(ctx *gin.Context, message string) {
	ctx.JSON(http.StatusCreated, gin.H{
		"message": message,
		"code":    "CREATED",
	})
}

func StatusPartialSuccess(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusMultiStatus, gin.H{
		"message": err.Error(),
		"code":    "PARTIAL_SUCCESS",
	})
}

func StatusInternalAPIErrorWithMessage(ctx *gin.Context, statusCode int, message string) {
	ctx.JSON(statusCode, gin.H{
		"message": message,
		"code":    "INTERNAL_API_ERROR",
	})
}

func StatusInternalServerError(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusInternalServerError, gin.H{
		"message": err.Error(),
		"code":    "INTERNAL_SERVER_ERROR",
	})
}

func StatusUnprocessableEntity(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusUnprocessableEntity, gin.H{
		"message": err.Error(),
		"code":    "UNPROCESSABLE_ENTITY",
	})
}

func StatusBadRequest(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusBadRequest, gin.H{
		"message": err.Error(),
		"code":    "BAD_REQUEST",
	})
}

func StatusBadRequestWithValidationErrorDetails(ctx *gin.Context, err error) {
	parsedErr := model.ParseError(err)
	ctx.JSON(http.StatusBadRequest, gin.H{
		"message": "Validation error",
		"code":    "BAD_REQUEST",
		"errors":  parsedErr.Errors,
	})
}

func StatusUnauthorized(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusUnauthorized, gin.H{
		"message": err.Error(),
		"code":    "UNAUTHORIZED",
	})
}

func StatusTooManyRequests(ctx *gin.Context, message string) {
	ctx.JSON(http.StatusTooManyRequests, gin.H{
		"message": message,
		"code":    "TOO_MANY_REQUESTS",
	})
}

func StatusForbidden(ctx *gin.Context, message string) {
	ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
		"message": message,
		"code":    "FORBIDDEN",
	})
}

// handleCommonErrors handles common errors and returns appropriate HTTP status codes.
func HandleCommonErrors(ctx *gin.Context, err error) {
	var commonErr *commonerrors.CommonError
	if errors.As(err, &commonErr) {
		StatusUnprocessableEntity(ctx, err)

		return
	}

	StatusInternalServerError(ctx, err)
}
