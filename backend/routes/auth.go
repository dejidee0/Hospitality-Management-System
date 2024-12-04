package routes

import (
	"hms/models"
	"hms/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// var users = make(map[string]models.User)

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
