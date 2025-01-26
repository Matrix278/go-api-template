package middleware

import (
	"fmt"
	"go-api-template/configuration"
	"go-api-template/controller"
	"go-api-template/model"
	"go-api-template/pkg/logger"

	// Import the Swagger docs package to register the generated documentation with the Swagger router.
	_ "go-api-template/docs"

	"github.com/gin-gonic/gin"
)

func NewRouter(cfg *configuration.Config, controllers *controller.Controllers) (*gin.Engine, error) {
	if cfg.AppEnv == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()

	if err := model.InitValidation(); err != nil {
		return nil, fmt.Errorf("failed to initialize validation: %v", err)
	}

	// Serve swagger ReDoc HTML
	router.StaticFile("/docs", "./docs/index.html")

	apiRouter := router.Group(cfg.APIPath)
	apiRouter.Use(CORSMiddleware())
	apiRouter.Use(logger.RequestsLogHandler())

	apiRouter.GET("/users/:user_id", authorizationHeaderRequired(), controllers.User.UserByID)

	return router, nil
}

// private
func authorizationHeaderRequired() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("Authorization")
		if token == "" {
			controller.StatusUnauthorized(ctx, fmt.Errorf("authorization header required"))
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}
