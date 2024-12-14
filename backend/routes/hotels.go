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
	var boookingData models.HotelBooking
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

	// validate booking and calculate the total

	err = boookingData.ValidateBookingAndCalculateTotalAmount()

	// totalAmount, err := boookingData.CalculateTotalAmount()
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"ammount": boookingData.TotalAmount,
		})
		return
	}
	// emails := strings.Split(boookingData.Emails, ",")
	// firstEmail := emails[0]
	// // fmt.Printf("emails are: %v\nfirstEmail is: %s\n", emails, firstEmail)
	// // validate email
	// ok := utils.ValidateEmail(firstEmail)
	// if !ok {
	// 	log.Println("invalid email")
	// 	ctx.JSON(http.StatusBadRequest, gin.H{
	// 		"error":      "first email in the email list is invalid",
	// 		"emails":     emails,
	// 		"firstEmail": firstEmail,
	// 	})
	// 	return
	// }

	// booking data to save to db
	// roomID := boookingData.RoomID

	// create a payment intent on paystack

	// // the payload to send to paystack
	payload := struct {
		Email  string `json:"email"`
		Amount int    `json:"amount"`
		// CallbackUrl string `json:"callback_url"`
	}{
		Email:  boookingData.Email,
		Amount: int(boookingData.TotalAmount) * 100,
		// CallbackUrl: config.Server_address + "/v1/hotels/booking/verify",
	}
	jsonData, err := json.Marshal(payload)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":      err.Error(),
			"emails":     boookingData.Emails,
			"firstEmail": boookingData.Email,
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
			"email": boookingData.Email,
		})
		return
	}

	// save booking to database with the access code, possibly bookindID, booking number, reference and acesss code in a separate table
	err = boookingData.Save()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "error getting booking number: " + err.Error(),
		})
		return
	}
	err = boookingData.SaveReference(PayStaResp.Data.AccessCode, PayStaResp.Data.Reference)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "error saving reference number: " + err.Error(),
		})
		return
	}

	// send the access code to the frontend to complete payment

	ctx.JSON(http.StatusOK, gin.H{
		"paystack-access-code": boookingData.AccessCode,
		"authorization-url":    PayStaResp.Data.AuthorizationUrl,
		"email":                boookingData.Email,
		"booking-number":       boookingData.BookingNumber,
		"tx-reference":         boookingData.Reference,
		"total-amount":         boookingData.TotalAmount,
		"booking-id":           boookingData.Id,
	})
}

func HotelBookingVerify(ctx *gin.Context) {
	reference := ctx.Query("reference")
	booking_id := ctx.Query("booking_id")

	// // the payload to send to paystack

	// paystack_url := "https://api.paystack.co/transaction/initialize"
	paystack_verify_url := "https://api.paystack.co/transaction/verify/" + reference

	req, err := http.NewRequest("GET", paystack_verify_url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error creating request: " + err.Error(),
		})
		return
	}

	// Add headers, including the Authorization header
	// req.Header.Set("Content-Type", "application/json")
	// req.Header.Set("Authorization", "Bearer YOUR_ACCESS_TOKEN")
	req.Header.Set("Authorization", "Bearer "+config.PAYSTACK_SECRET_KEY_TEST)

	// Use http.Client to send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error sending request: " + err.Error(),
		})
		return
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error reading response: " + err.Error(),
		})
		return
	}
	// {

	// 	"message": "Verification successful",
	// 	"data": {
	// 	  "id": 4099260516,
	// 	  "domain": "test",
	// 	  "status": "success",
	// 	  "reference": "re4lyvq3s3",
	// 	  "receipt_number": null,
	// 	  "amount": 40333,
	// 	  "message": null,
	// 	  "gateway_response": "Successful",
	// 	  "paid_at": "2024-08-22T09:15:02.000Z",
	// 	  "created_at": "2024-08-22T09:14:24.000Z",
	// 	  "channel": "card",
	// 	  "currency": "NGN",
	// 	  "ip_address": "197.210.54.33",
	// 	  "metadata": "",
	// 	  "log": {
	// 		"start_time": 1724318098,
	// 		"time_spent": 4,
	// 		"attempts": 1,
	// 		"errors": 0,
	// 		"success": true,
	// 		"mobile": false,
	// 		"input": [],
	// 		"history": [
	// 		  {
	// 			"type": "action",
	// 			"message": "Attempted to pay with card",
	// 			"time": 3
	// 		  },
	// 		  {
	// 			"type": "success",
	// 			"message": "Successfully paid with card",
	// 			"time": 4
	// 		  }
	// 		]
	// 	  },
	// 	  "fees": 10283,
	// 	  "fees_split": null,
	// 	  "authorization": {
	// 		"authorization_code": "AUTH_uh8bcl3zbn",
	// 		"bin": "408408",
	// 		"last4": "4081",
	// 		"exp_month": "12",
	// 		"exp_year": "2030",
	// 		"channel": "card",
	// 		"card_type": "visa ",
	// 		"bank": "TEST BANK",
	// 		"country_code": "NG",
	// 		"brand": "visa",
	// 		"reusable": true,
	// 		"signature": "SIG_yEXu7dLBeqG0kU7g95Ke",
	// 		"account_name": null
	// 	  },
	// 	  "customer": {
	// 		"id": 181873746,
	// 		"first_name": null,
	// 		"last_name": null,
	// 		"email": "demo@test.com",
	// 		"customer_code": "CUS_1rkzaqsv4rrhqo6",
	// 		"phone": null,
	// 		"metadata": null,
	// 		"risk_action": "default",
	// 		"international_format_phone": null
	// 	  },
	// 	  "plan": null,
	// 	  "split": {},
	// 	  "order_id": null,
	// 	  "paidAt": "2024-08-22T09:15:02.000Z",
	// 	  "createdAt": "2024-08-22T09:14:24.000Z",
	// 	  "requested_amount": 30050,
	// 	  "pos_transaction_data": null,
	// 	  "source": null,
	// 	  "fees_breakdown": null,
	// 	  "connect": null,
	// 	  "transaction_date": "2024-08-22T09:14:24.000Z",
	// 	  "plan_object": {},
	// 	  "subaccount": {}
	// 	}
	//   }

	var result struct {
		// Status  bool   `json:"status"`
		// Message string `json:"message"`
		Data struct {
			// Id        int64   `json:"id"`
			Status string `json:"status"`
			// Reference string  `json:"reference"`
			Amount float64 `json:"amount"`
			// Customer  struct {
			// 	Id        any    `json:"id"`
			// 	FirstName string `json:"first_name"`
			// 	LastName  string `json:"last_name"`
			// 	Email     string `json:"email"`
			// 	Phone     any    `json:"phone"`
			// } `json:"customer"`
		} `json:"data"`
	}

	if err := json.Unmarshal(body, &result); err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error unmarshalling JSON: " + err.Error(),
		})
		return
	}
	if result.Data.Status != "success" {
		fmt.Printf("payment with reference: %s not successful", reference)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("payment with reference: %s not successful", reference),
		})
		return
	}

	var booking models.HotelBooking

	data, err := booking.GetBookingDetails(booking_id)
	if err != nil {
		fmt.Printf("payment is successful but error getting booking details")
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "payment is successful but error getting booking details",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "payment successful!",
		"data":    data,
	})
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
