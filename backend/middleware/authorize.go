package middleware

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func Authorize() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		JWTKEY := []byte(os.Getenv("JWT_SECRET_KEY"))

		token_string := ctx.GetHeader("Authorization")
		if token_string == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": http.StatusText(http.StatusUnauthorized),
			})
			ctx.Abort()
			return
		}
		token, err := jwt.Parse(token_string, func(t *jwt.Token) (interface{}, error) {
			return JWTKEY, nil
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
