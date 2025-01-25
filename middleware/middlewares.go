package middleware

import (
	"go-api-template/configuration"

	"github.com/gin-gonic/gin"
)

type Middlewares struct {
	CORS            gin.HandlerFunc
	SecurityHeaders gin.HandlerFunc
}

func NewMiddlewares(cfg *configuration.New) *Middlewares {
	corsMiddleware := NewCORSMiddleware(cfg)
	securityHeadersMiddleware := NewSecurityHeadersMiddleware(cfg)

	return &Middlewares{
		CORS:            corsMiddleware.Handler(),
		SecurityHeaders: securityHeadersMiddleware.Handler(),
	}
}
