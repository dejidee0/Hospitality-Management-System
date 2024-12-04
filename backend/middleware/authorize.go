package middleware

import (
	"hms/config"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func Authorize() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token_string := ctx.GetHeader("Authorization")
		if token_string == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": http.StatusText(http.StatusUnauthorized),
			})
			ctx.Abort()
			return
		}
		token, err := jwt.Parse(token_string, func(t *jwt.Token) (interface{}, error) {
			return config.JWTKey, nil
		})
		if err != nil {
			log.Println("Hit" + err.Error())
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": http.StatusText(http.StatusUnauthorized),
			})
			ctx.Abort()
			return
		}
		claims, ok := token.Claims.(jwt.MapClaims)
		if ok && token.Valid {
			ctx.Set("user_id", claims["Id"])
			// fmt.Printf("Claims: %+v\n", claims)
			ctx.Next()
		} else {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": http.StatusText(http.StatusUnauthorized),
			})
			ctx.Abort()
		}

	}
}
