package main

import (
	"hms/config"
	"hms/database"
	"hms/middleware"
	"hms/routes"
	"log"

	"github.com/gin-gonic/gin"
)

// connect and close database
func init() {
	db, err := database.GetDB()
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	log.Println("db connected successfully!")
	if err = db.Ping(); err != nil {
		log.Println("failed to ping database...:" + err.Error())
		return
	}
	log.Println("db pinged successfully!")
}

func main() {
	r := gin.Default()
	r.Use(middleware.CORSMiddleware())

	// index routes
	r.GET("/v1", routes.Index)
	r.GET("/v1/health", routes.Health)

	// this serves the data for the home page
	r.GET("/v1/index", routes.Home)

	// auth routes
	r.POST("/v1/auth/signup", routes.Signup)
	r.POST("/v1/auth/login", routes.Login)
	r.GET("/v1/auth/profile", middleware.Authorize(), routes.Profile)

	// reset password
	r.GET("/v1/auth/reset-password", routes.ResetPassword)
	r.POST("/v1/auth/change-password", routes.ChangePassword)

	if err := r.Run(config.Server_address); err != nil {
		log.Fatal(err)
	}
}
