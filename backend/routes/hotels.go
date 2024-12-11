package routes

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"hms/config"
	"hms/models"
	"hms/utils"
	"io"
	"log"
	"net/http"
	"strings"

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

func HotelBooking(ctx *gin.Context) {
	// get the booking detail (the data)
	var boookingData utils.BookingDetails
	err := ctx.Bind(&boookingData)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	// boookingData.Night = 1
	// boookingData.Room = 1

	// calculate the total
	totalAmount, err := boookingData.CalculateTotalAmount()
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"ammount": totalAmount,
		})
		return
	}
	emails := strings.Split(boookingData.Emails, ",")
	firstEmail := emails[0]
	// fmt.Printf("emails are: %v\nfirstEmail is: %s\n", emails, firstEmail)
	// validate email
	ok := utils.ValidateEmail(firstEmail)
	if !ok {
		log.Println("invalid email")
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":      "first email in the email list is invalid",
			"emails":     emails,
			"firstEmail": firstEmail,
		})
		return
	}
	// create a payment intent on paystack

	// // the payload to send to paystack
	payload := struct {
		Email  string `json:"email"`
		Amount int    `json:"amount"`
	}{
		Email:  firstEmail,
		Amount: int(totalAmount) * 100,
	}
	jsonData, err := json.Marshal(payload)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":      err.Error(),
			"emails":     emails,
			"firstEmail": firstEmail,
		})
		return
	}

	paystack_url := "https://api.paystack.co/transaction/initialize"

	req, err := http.NewRequest("POST", paystack_url, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// Add headers, including the Authorization header
	req.Header.Set("Content-Type", "application/json")
	// req.Header.Set("Authorization", "Bearer YOUR_ACCESS_TOKEN")
	req.Header.Set("Authorization", "Bearer "+config.PAYSTACK_SECRET_KEY_TEST)

	// Use http.Client to send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}
	var PayStaResp PaystackResponse
	err = json.Unmarshal(body, &PayStaResp)
	if err != nil {
		log.Println(err)
		return
	}
	if !PayStaResp.Status {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "paystack-error: " + PayStaResp.Message,
			"email": firstEmail,
		})
		return
	}

	// Print the status and the response body
	// fmt.Printf("Response status: %s\n", resp.Status)

	// fmt.Printf("Response body: %s\n", string(body))
	ctx.JSON(http.StatusOK, gin.H{
		"paystack-access-code": PayStaResp.Data.AccessCode,
		"authorization-url":    PayStaResp.Data.AuthorizationUrl,
		"email":                firstEmail,
	})

	// resp, err := http.Post(paystack_url, "application/json", bytes.NewBuffer(jsonData))
	// if err != nil {
	// 	log.Println(err)
	// 	ctx.JSON(http.StatusBadRequest, gin.H{
	// 		"error":      "error post request to paystack: " + err.Error(),
	// 		"emails":     emails,
	// 		"firstEmail": firstEmail,
	// 	})
	// 	return
	// }
	// defer resp.Body.Close()
	// // Read the response body
	// body, err := io.ReadAll(resp.Body)
	// if err != nil {
	// 	log.Println(err)
	// 	ctx.JSON(http.StatusBadRequest, gin.H{
	// 		"error":      "error reading the response body: " + err.Error(),
	// 		"emails":     emails,
	// 		"firstEmail": firstEmail,
	// 	})
	// 	return
	// }

	// // Print the status and the response body
	// // fmt.Printf("Response status: %s\n", resp.Status)
	// // fmt.Printf("Response body: %s\n", string(body))
	// ctx.JSON(http.StatusOK, gin.H{
	// 	"paystack-response": map[string]any{
	// 		"status": resp.Status,
	// 		"body":   string(body),
	// 	},
	// 	"data":       boookingData,
	// 	"amount":     totalAmount,
	// 	"emails":     emails,
	// 	"firstEmail": firstEmail,
	// })

	// save booking to database with the access code, possibly bookindID, booking number and acesss code in a separate table

	// send the access code to the frontend to complete payment

	// ctx.JSON(http.StatusOK, gin.H{
	// 	"data":       boookingData,
	// 	"amount":     totalAmount,
	// 	"emails":     emails,
	// 	"firstEmail": firstEmail,
	// })

}

type PaystackResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    struct {
		AuthorizationUrl string `json:"authorization_url"`
		AccessCode       string `json:"access_code"`
		Reference        string `json:"reference"`
	} `json:"data"` // Replace `any` with the specific type if known
}

// "authorization_url": "https://checkout.paystack.com/nkdks46nymizns7",
