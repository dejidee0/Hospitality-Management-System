package routes

import (
	"database/sql"
	"hms/models"
	"hms/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HotelsIndex(ctx *gin.Context) {

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

func HotelsSearch(ctx *gin.Context) {
	state := ctx.Query("state")
	var hotel models.Hotel
	hotels := hotel.GetHotelsByState(state)

	ctx.JSON(http.StatusOK, gin.H{
		"hotels": hotels,
	})
}

func HotelDetail(ctx *gin.Context) {
	hotel_id := ctx.Params.ByName("hotel_id")
	if hotel_id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid hotel_id",
		})
		return
	}
	var hotel models.Hotel

	err := hotel.GetHotelByID(hotel_id)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"hotel": nil,
			})
			return
		}
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	rooms := hotel.GetRooms()
	reviews := hotel.GetReviews()
	policy := hotel.GetPolicy()
	similarHotels := hotel.GetSimilar()

	ctx.JSON(http.StatusOK, gin.H{
		"hotel":    hotel,
		"rooms":    rooms,
		"reviews":  reviews,
		"policies": policy,
		"similar":  similarHotels,
	})
}
