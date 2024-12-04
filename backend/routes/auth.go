package routes

import (
	"fmt"
	"hms/config"
	"hms/models"
	"hms/utils"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func Signup(ctx *gin.Context) {
	// getting post data
	var data utils.SignupData
	err := ctx.Bind(&data)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// create and validate user
	user, err := models.NewUser(data.Email, data.Password)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// add user to db
	err = user.Save()
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	// return success
	ctx.JSON(http.StatusCreated, gin.H{"message": "user created successfully"})
}

func Login(ctx *gin.Context) {
	// getting post data
	var data utils.LoginData
	err := ctx.Bind(&data)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	// validate email and password

	// authenticate user
	user, err := models.Authenticate(data.Email, data.Password)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	// create jwt
	expirationTime := time.Now().Add(1 * time.Hour).Unix()

	claims := &JwtClaims{
		Id:   user.Id,
		Name: user.Name,
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

	// return success
	ctx.JSON(http.StatusAccepted, gin.H{
		"message": "user logged in successfully",
		"token":   tokenString,
	})
}

type JwtClaims struct {
	Id   string
	Name string
	jwt.StandardClaims
}

func Profile(ctx *gin.Context) {
	fmt.Println("profile starts 2")

	// getting post data

	// return success
	ctx.JSON(http.StatusAccepted, gin.H{
		"message": "Welcome user!",
	})
}
