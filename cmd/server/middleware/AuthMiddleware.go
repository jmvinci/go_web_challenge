package middlewares

import (
	"log"
	"net/http"
	"os"

	"github.com/bootcamp-go/desafio-go-web/pkg/web"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	requiredToken := os.Getenv("API_TOKEN")
	if requiredToken == "" {
		log.Fatal("Please set API_TOKEN environment variable")
	}
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("TOKEN")
		if token == "" {
			ctx.JSON(http.StatusUnauthorized, web.Response{Message: "API token required", Data: nil})
			ctx.Abort()
			return
		}
		if token != requiredToken {
			ctx.JSON(http.StatusUnauthorized, web.Response{Message: "Invalid token", Data: nil})
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
