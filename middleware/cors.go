package middleware

import (
	"go-api-template/configuration"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/cors"
)

// CORSMiddleware sets up CORS headers based on the configuration.
type CORSMiddleware struct {
	AllowedOrigins []string
}

func NewCORSMiddleware(cfg *configuration.New) *CORSMiddleware {
	return &CORSMiddleware{
		AllowedOrigins: cfg.AllowedOrigins,
	}
}

func (middleware *CORSMiddleware) Handler() gin.HandlerFunc {
	corsConfig := cors.New(cors.Options{
		AllowedOrigins: middleware.AllowedOrigins,
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"*"},
	})

	return func(ctx *gin.Context) {
		corsConfig.HandlerFunc(ctx.Writer, ctx.Request)

		if ctx.Request.Method != http.MethodOptions {
			ctx.Next()
		}
	}
}
