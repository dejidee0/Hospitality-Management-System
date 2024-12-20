package models

import (
	"database/sql"
	"errors"
	"hms/database"
	"hms/utils"
	"log"
	"strings"
	"time"

	"github.com/google/uuid"
)

type EventBooking struct {
	EventBookingId string    `json:"event-booking-id,omitempty"`
	EventID        string    `json:"event-id,omitempty" binding:"required"`
	FirstName      string    `json:"first-name,omitempty" binding:"required"`
	LastName       string    `json:"last-name,omitempty" binding:"required"`
	Email          string    `json:"email,omitempty" binding:"required"`
	PaymentMethod  string    `json:"payment-method,omitempty"`
	Quantity       int       `json:"quantity,omitempty" binding:"required,min=1"`
	PromoCode      string    `json:"promo-code,omitempty"`
	TotalAmount    float64   `json:"total-amount,omitempty"`
	BookingNumber  int64     `json:"booking-number,omitempty"`
	AccessCode     string    `json:"access-code,omitempty"`
	Reference      string    `json:"tx-reference,omitempty"`
	EventName      string    `json:"event-name,omitempty"`
	EventDate      time.Time `json:"event-date,omitempty"`
	Venue          string    `json:"venue,omitempty"`
	Image          string    `json:"image,omitempty"`
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

// firstname VARCHAR(100) NOT NULL,
// booking_number INT IDENTITY(1,1),
// email VARCHAR(100) NOT NULL,
// quantity
// total_price DECIMAL(10,2) NOT NULL,
// event_id VARCHAR(36) NOT NULL, to get name of event, date, venue
// payment_status VARCHAR(20) DEFAULT 'pending' CHECK(payment_status IN ('pending', 'processing', 'paid', 'refunded')),
//
//

func (evb *EventBooking) GetBookingDetails(reference string) error {
	query := `SELECT event_bookings.firstname, event_bookings.booking_number, event_bookings.email, 
	event_bookings.quantity, event_bookings.total_price, events.name, 
	events.date, events.venue, events.images FROM event_bookings JOIN events 
	ON event_bookings.event_id = events.id WHERE event_bookings.paystack_reference = @reference;`

	// var data BookingType
	row := database.DB.QueryRow(query, sql.Named("reference", reference))
	err := row.Scan(&evb.FirstName, &evb.BookingNumber, &evb.Email, &evb.Quantity,
		&evb.TotalAmount, &evb.EventName, &evb.EventDate, &evb.Venue, &evb.Image,
	)
	if err != nil {
		log.Println(err)
		return err
	}
	evb.Image = strings.Split(evb.Image, ",")[0]
	return nil
}
