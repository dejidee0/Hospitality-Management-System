package routes

import (
	"hms/models"
	"hms/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Home(ctx *gin.Context) {

	// fetch popular hotels

	var hotel models.Hotel
	hotels, err := hotel.GetPopularHotels()
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "error fetching popular hotels",
		})
		return
	}
	// get trending destinations
	destinations := utils.GetTrendingDestinations()
	property_types := utils.GetPropertyTypes()

	// get blogs
	var blog models.Blog
	blogs, err := blog.GetRecentBlogs()
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "error fetching recent blogs",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": map[string]interface{}{
			"popular_hotels":        hotels,
			"trending_destinations": destinations,
			"property_types":        property_types,
			"blogs":                 blogs,
		},
	})
}
