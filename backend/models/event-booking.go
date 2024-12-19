package models

import (
	"database/sql"
	"errors"
	"hms/database"
	"hms/utils"
	"log"

	"github.com/google/uuid"
)

type EventBooking struct {
	EventBookingId string  `json:"event-booking-id"`
	EventID        string  `json:"event-id" binding:"required"`
	FirstName      string  `json:"first-name" binding:"required"`
	LastName       string  `json:"last-name" binding:"required"`
	Email          string  `json:"email" binding:"required"`
	PaymentMethod  string  `json:"payment-method"`
	Quantity       int     `json:"quantity" binding:"required,min=1"`
	PromoCode      string  `json:"promo-code"`
	TotalAmount    float64 `json:"total-amount"`
	BookingNumber  int64   `json:"booking-number"`
	AccessCode     string  `json:"access-code"`
	Reference      string  `json:"tx-reference"`
}

func (evb *EventBooking) ValidateBookingAndCalculateTotalAmount() error {
	ok := utils.ValidateEmail(evb.Email)
	if !ok {
		log.Println("email is invalid")
		return errors.New("email is invalid")
	}

	totalAmount, err := evb.CalculateTotalAmount()
	if err != nil {
		log.Println(err)
		return err
	}
	evb.TotalAmount = totalAmount

	return nil
}

func (evb *EventBooking) CalculateTotalAmount() (float64, error) {
	// panic("Not implemented")
	// total = (price_per_ticket * quantity * ((100-promoRate)/100) * (100+taxrate)/100)
	price_per_ticket, err := getEventPrice(evb.EventID)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	promoRate := 0 // no promo 0 percent off. to be gotten from the promocode
	taxRate := 0   // no tax. 0 percent tax added. to be gotten from the database

	total := price_per_ticket * float64(evb.Quantity) * float64((100-promoRate)/100) * float64((100+taxRate)/100)

	return total, nil
}

func getEventPrice(event_id string) (float64, error) {
	query := `SELECT price FROM events WHERE id = @event_id;`
	row := database.DB.QueryRow(query, sql.Named("event_id", event_id))
	var price float64
	err := row.Scan(&price)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, errors.New("no record for such event ID")
		}
		return 0, err
	}
	return price, nil
}

func (evb *EventBooking) Save(access_code, reference string) error {
	event_booking_id := uuid.New().String()
	query := `INSERT INTO event_bookings (id, event_id, firstname, lastname, 
	email, total_price, quantity, promocode_used, payment_method, paystack_access_code, paystack_reference) OUTPUT INSERTED.booking_number 
	VALUES(@id, @event_id, @firstname, @lastname, @email, @total_price, @quantity, 
	@promocode_used, @payment_method, @access_code, @reference);`

	row := database.DB.QueryRow(query, sql.Named("id", event_booking_id), sql.Named("event_id", evb.EventID), sql.Named("firstname", evb.FirstName),
		sql.Named("lastname", evb.LastName), sql.Named("email", evb.Email),
		sql.Named("total_price", evb.TotalAmount), sql.Named("quantity", evb.Quantity),
		sql.Named("promocode_used", evb.PromoCode), sql.Named("payment_method", evb.PaymentMethod),
		sql.Named("access_code", access_code), sql.Named("reference", reference),
	)
	var bookinNum int64
	err := row.Scan(&bookinNum)
	if err != nil {
		log.Println(err)
		return err
	}
	evb.BookingNumber = bookinNum
	evb.EventBookingId = event_booking_id
	evb.AccessCode = access_code
	evb.Reference = reference
	return nil
}
