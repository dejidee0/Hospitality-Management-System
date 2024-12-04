package routes

import (
	"hms/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// var users = make(map[string]models.User)

func Signup(ctx *gin.Context) {
	var data utils.SignupData
	err := ctx.Bind(&data)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = utils.ValidateSignupData(&data)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "User created successfully",
		"user_id": "234df",
		"user":    &data,
	})

}

// func Profile(ctx *gin.Context) {
// 	user_name := ctx.Param("user_name")

// 	user, ok := users[user_name]
// 	if !ok {
// 		log.Println("no user with: " + user_name)
// 		ctx.JSON(http.StatusBadRequest, gin.H{
// 			"error": "no user with: " + user_name,
// 		})
// 		return
// 	}

// 	ctx.JSON(http.StatusCreated, gin.H{
// 		"user": &user,
// 	})
// }
