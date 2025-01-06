package routes

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"hms/models"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func EventsIndex(ctx *gin.Context) {
	// state := ctx.Query("state")

	// fetch popular hotels in state=lagos
	var event models.Event
	popular_events, err := event.GetPopularEventsIn("lagos")
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "error fetching popular events in " + "lagos",
		})
		return
	}
	// get online, muisc and business events
	online_events := event.GetEventsByFormat("online")
	music_events := event.GetEventsByCategory("music")
	business_events := event.GetEventsByCategory("business")

	ctx.JSON(http.StatusOK, gin.H{
		"data": map[string]interface{}{
			"popular_events":  popular_events,
			"online_events":   online_events,
			"music_events":    music_events,
			"business_events": business_events,
		},
	})

}

func EventsSearch(ctx *gin.Context) {
	state := ctx.Query("state")
	var event models.Event
	events := event.GetEventsByState(state)

	ctx.JSON(http.StatusOK, gin.H{
		"events": events,
	})
}

// event details
func EventDetail(ctx *gin.Context) {
	event_id := ctx.Params.ByName("event_id")
	if event_id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid hotel_id",
		})
		return
	}
	var event models.Event

	err := event.GetEventByID(event_id)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"err": "no event with such id",
			})
		} else {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}
		return
	}

	// get events in the same category
	similarEvents := event.GetEventsByCategory(event.Category)
	limit := 5
	if len(similarEvents) < limit {
		limit = len(similarEvents)
	}
	similar := similarEvents[:limit]
	ctx.JSON(http.StatusOK, gin.H{
		"event":   event,
		"similar": similar,
	})
}

func EventBooking(ctx *gin.Context) {
	var PAYSTACK_SECRET_KEY_TEST = os.Getenv("PAYSTACK_SECRET_KEY_TEST")

	// get the booking detail (the data)
	var boookingData models.EventBooking
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

	// create a payment intent on paystack

	// // the payload to send to paystack
	payload := struct {
		FirstName string `json:"firstname"`
		LastName  string `json:"lastname"`
		Email     string `json:"email"`
		Amount    int    `json:"amount"`
		// CallbackUrl string `json:"callback_url"`
	}{
		FirstName: boookingData.FirstName,
		LastName:  boookingData.LastName,
		Email:     boookingData.Email,
		Amount:    int(boookingData.TotalAmount * 100),
		// CallbackUrl: config.Server_address + "/v1/hotels/booking/verify",
	}
	jsonData, err := json.Marshal(payload)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
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
	req.Header.Set("Authorization", "Bearer "+PAYSTACK_SECRET_KEY_TEST)

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
	err = boookingData.Save(PayStaResp.Data.AccessCode, PayStaResp.Data.Reference)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "error saving booking: " + err.Error(),
		})
		return
	}
	// err = boookingData.SaveReference(PayStaResp.Data.AccessCode, PayStaResp.Data.Reference)
	// if err != nil {
	// 	ctx.JSON(http.StatusInternalServerError, gin.H{
	// 		"error": "error saving reference number: " + err.Error(),
	// 	})
	// 	return
	// }

	// send the access code to the frontend to complete payment

	ctx.JSON(http.StatusOK, gin.H{
		"paystack-access-code": boookingData.AccessCode,
		"authorization-url":    PayStaResp.Data.AuthorizationUrl,
		"email":                boookingData.Email,
		"booking-number":       boookingData.BookingNumber,
		"tx-reference":         boookingData.Reference,
		"total-amount":         boookingData.TotalAmount,
		"event-booking-id":     boookingData.EventBookingId,
	})
}

func EventBookingVerify(ctx *gin.Context) {
	var PAYSTACK_SECRET_KEY_TEST = os.Getenv("PAYSTACK_SECRET_KEY_TEST")

	reference := ctx.Query("reference")

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
	req.Header.Set("Authorization", "Bearer "+PAYSTACK_SECRET_KEY_TEST)

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

	var booking models.EventBooking

	err = booking.GetBookingDetails(reference)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "payment is successful but error getting booking details",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "payment is successful!",
		"data":    booking,
	})
}
