package main

import (
	"hms/config"
	"hms/database"
	"hms/middleware"
	"hms/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// close database before main ends
	defer database.DB.Close()

	r := gin.Default()
	r.Use(middleware.CORSMiddleware())

	// index routes
	r.GET("/v1", routes.Index)
	r.GET("/v1/health", routes.Health)

	// auth routes
	r.POST("/v1/auth/signup", routes.Signup)
	r.POST("/v1/auth/login", routes.Login)
	r.GET("/v1/auth/profile", middleware.Authorize(), routes.Profile)

	// reset password
	r.GET("/v1/auth/reset-password", routes.ResetPassword)
	r.POST("/v1/auth/change-password", routes.ChangePassword)

	// this serves the data for the home page
	r.GET("/v1/hotels/index", routes.HotelsIndex)
	r.GET("/v1/hotels/search", routes.HotelsSearch)
	r.GET("/v1/hotels/:hotel_id", routes.HotelDetail)
	r.POST("/v1/hotels/booking", routes.HotelBooking)
	r.GET("/v1/hotels/booking/verify", routes.HotelBookingVerify)

	if err := r.Run(config.Server_address); err != nil {
		log.Fatal(err)
	}
}
