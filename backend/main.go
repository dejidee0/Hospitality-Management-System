package main

import (
	"hms/config"
	"hms/middleware"
	"hms/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(middleware.CORSMiddleware())

	// index routes
	r.GET("/", routes.Index)
	r.GET("/health", routes.Health)

	// auth routes
	r.POST("/auth/signup", routes.Signup)
	r.POST("/auth/login", routes.Login)
	r.GET("/auth/profile", middleware.Authorize(), routes.Profile)

	// reset password
	r.GET("/auth/reset-password", routes.ResetPassword)
	r.POST("/auth/change-password", routes.ChangePassword)

	// r.GET("/auth/profile/:user_name", routes.Profile)
	// r.POST("/auth/login", routes.Login)

	if err := r.Run(config.Server_address); err != nil {
		log.Fatal(err)
	}
}
