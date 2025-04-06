package middleware

import (
	"errors"
	"go-api-template/configuration"
	"go-api-template/controller"
	"go-api-template/pkg/logger"

	"github.com/gin-gonic/gin"
)

func NewRouter(cfg *configuration.Env, controllers *controller.Controllers) (*gin.Engine, error) {
	if cfg.AppEnv == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()

	// Initialize middlewares
	middleware := NewMiddlewares(cfg)
	router.Use(middleware.SecurityHeaders)
	router.Use(middleware.CORS)

	// Serve swagger ReDoc HTML
	router.StaticFile("/docs", "./docs/index.html")

	apiRouter := router.Group(cfg.APIPath)
	apiRouter.Use(logger.RequestsLogHandler())

	apiRouter.GET("/users/:user_id", authorizationHeaderRequired(), controllers.User.UserByID)

	return router, nil
}

// private.
func authorizationHeaderRequired() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("Authorization")
		if token == "" {
			controller.StatusUnauthorized(ctx, errors.New("authorization header required"))
			ctx.Abort()

			return
		}

		ctx.Next()
	}
}
