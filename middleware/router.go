package middleware

import (
	"fmt"
	"go-api-template/configuration"
	"go-api-template/controller"
	"go-api-template/pkg/logger"

	"github.com/gin-gonic/gin"
	// swagger embed files
	// gin-swagger middleware
	_ "go-api-template/docs" // Import the Swagger docs package to register the generated documentation with the Swagger router.
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

	apiRouter.GET("/user/:user_id", controllers.User.UserByID).Use(authorizationHeaderRequired())

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
