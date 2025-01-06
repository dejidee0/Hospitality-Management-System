package routes

import (
	"hms/models"
	"hms/utils"
	"log"
	"net/http"
	"os"
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
	JWTKEY := []byte(os.Getenv("JWT_SECRET_KEY"))

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
		Id:        user.Id,
		FirstName: user.FirstName,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime,
			Issuer:    "HMS-FindPeace",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(JWTKEY)
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
	Id        string
	FirstName string
	jwt.StandardClaims
}

func Profile(ctx *gin.Context) {

	// getting post data

	// return success
	id := ctx.GetString("user_id")
	ctx.JSON(http.StatusAccepted, gin.H{
		"message": "Welcome user!" + id,
	})
}
