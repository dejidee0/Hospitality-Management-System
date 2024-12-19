package routes

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"hms/config"
	"hms/models"
	"io"
	"log"
	"net/http"

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
