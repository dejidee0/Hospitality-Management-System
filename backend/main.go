package main

import (
	"hms/config"
	"hms/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// index routes
	r.GET("/", routes.Index)
	r.GET("/health", routes.Health)

	// auth routes
	r.POST("/auth/signup", routes.Signup)
	// r.GET("/auth/profile/:user_name", routes.Profile)
	// r.POST("/auth/login", routes.Login)

	if err := r.Run(config.Server_address); err != nil {
		log.Fatal(err)
	}
}
