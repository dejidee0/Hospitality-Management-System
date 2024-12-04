package main

import (
	"hms/config"
	"hms/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", routes.Index)
	r.GET("/health", routes.Health)

	if err := r.Run(config.Server_address); err != nil {
		log.Fatal(err)
	}
}
