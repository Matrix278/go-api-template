package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/cors"
)

func CORSMiddleware() gin.HandlerFunc {
	corsConfig := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodOptions},
		AllowedHeaders: []string{"*"},
	})

	return func(ctx *gin.Context) {
		corsConfig.HandlerFunc(ctx.Writer, ctx.Request)

		if ctx.Request.Method != http.MethodOptions {
			ctx.Next()
		}
	}
}
