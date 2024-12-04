package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authorize() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		auth := ctx.GetHeader("Authorization")
		if auth == "" {
			ctx.JSON(http.StatusForbidden, gin.H{
				"error": http.StatusText(http.StatusForbidden),
			})
			return
		}
		fmt.Printf("Authourization: %s\n", auth)

		fmt.Println("before profile starts 1")
		ctx.Next()
		fmt.Println("After profile end 3")
	}
}
