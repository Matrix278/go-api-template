package middleware

import (
	"go-api-template/configuration"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SecurityHeadersMiddleware struct {
	cfg *configuration.New
}

func NewSecurityHeadersMiddleware(cfg *configuration.New) *SecurityHeadersMiddleware {
	return &SecurityHeadersMiddleware{
		cfg: cfg,
	}
}

func (middleware *SecurityHeadersMiddleware) Handler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		expectedHost := middleware.cfg.AppHost + ":" + middleware.cfg.AppPort
		if ctx.Request.Host != expectedHost {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": "Invalid host header",
				"code":  http.StatusBadRequest,
			})
			return
		}

		ctx.Header("X-Frame-Options", "DENY")
		ctx.Header("Content-Security-Policy", "default-src 'self'; connect-src *; font-src *; script-src-elem * 'unsafe-inline'; img-src * data:; style-src * 'unsafe-inline';")
		ctx.Header("X-XSS-Protection", "1; mode=block")
		ctx.Header("Strict-Transport-Security", "max-age=31536000; includeSubDomains; preload")
		ctx.Header("Referrer-Policy", "strict-origin")
		ctx.Header("X-Content-Type-Options", "nosniff")
		ctx.Header("Permissions-Policy", "geolocation=(),midi=(),sync-xhr=(),microphone=(),camera=(),magnetometer=(),gyroscope=(),fullscreen=(self),payment=()")

		ctx.Next()
	}
}
