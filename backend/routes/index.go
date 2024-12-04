package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// index route, route to the index page. '/'
func Index(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Welcome home",
	})
}

// the health route tells if the server is running successfully or not
func Health(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Server running fine",
	})
}
