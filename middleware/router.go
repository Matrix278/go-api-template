package middleware

import (
	"fmt"
	"go-api-template/configuration"
	"go-api-template/controller"
	"go-api-template/pkg/logger"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware

	_ "go-api-template/docs" // Import the Swagger docs package to register the generated documentation with the Swagger router.
)

func NewRouter(cfg *configuration.New, controllers *controller.Controllers) (*gin.Engine, error) {
	if cfg.AppEnv == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()

	// Initialize middlewares
	middleware := NewMiddlewares(cfg)
	router.Use(middleware.SecurityHeaders)
	router.Use(middleware.CORS)

	// Serve swagger docs
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

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
