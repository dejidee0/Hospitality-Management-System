package routes

import (
	"fmt"
	"hms/config"
	"hms/mail"
	"hms/models"
	"hms/utils"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func ResetPassword(ctx *gin.Context) {
	email := ctx.Query("email")
	// validate email
	ok := utils.ValidateEmail(email)

	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "email not valid",
		})
		return
	}
	var user models.User
	// and check if email exists in the users table
	err := user.GetUserByEmail(email)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "no account with this email",
		})
		return
	}

	// create token with the expiry date, put email in the claims
	// create jwt
	expirationTime := time.Now().Add(1 * time.Hour).Unix()

	claims := &struct {
		Id    string
		Email string
		jwt.StandardClaims
	}{
		Id:    user.Id,
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime,
			Issuer:    "HMS-FindPeace",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(config.JWTKey)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// update users table with the token
	err = user.UpdateResetPasswordToken(tokenString)
	if err != nil {
		log.Println(err)
	}

	// send token together with the link to update form to email
	// e.g wwww.findpeacefrontend.com/reset?reset_token=token

	err = mail.SendToken(tokenString, email)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("email not sent to (%s), try again!", email),
		})
		return
	}

	// success
	ctx.JSON(http.StatusOK, gin.H{
		"message": "email sent successfully",
	})
}

func ChangePassword(ctx *gin.Context) {
	var data utils.ChangePasswordData
	err := ctx.Bind(&data)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	// validate token, check if its still valid
	token, err := jwt.Parse(data.Token, func(t *jwt.Token) (interface{}, error) {
		return config.JWTKey, nil
	})
	if err != nil {
		log.Println("parsing and validating error: " + err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("Token is invalid, %s", err),
		})
		return
	}

	// get the email from the token

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		// ctx.Set("user_id", claims["Id"])
		// // fmt.Printf("Claims: %+v\n", claims)
		// ctx.Next()
		log.Println("error getting claim")
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Token is invalid",
		})
		return
	}
	email := claims["Email"].(string) // type assertion, make sure the result is string
	fmt.Printf("email from token is: %s\n", email)

	// hash the password
	hashedPass := models.HashPassword(data.Password)
	// upadte the pasword in the users table and invalidate the token in the users table
	var user models.User
	_ = user.UpdatePassword(email, hashedPass)

	// success
	ctx.JSON(http.StatusOK, gin.H{
		"message": "password updated successfully",
	})
}
