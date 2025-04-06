package middleware

import (
	"fmt"
	"go-api-template/configuration"
	"go-api-template/controller"
	"go-api-template/pkg/logger"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
)

func NewRouter(cfg *configuration.Env, controllers *controller.Controllers) (*gin.Engine, error) {
	if cfg.AppEnv == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()
	if cfg.Telemetry.Enabled {
		router.Use(otelgin.Middleware(cfg.Telemetry.ServiceName))
	}

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
