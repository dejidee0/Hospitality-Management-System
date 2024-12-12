package utils

import (
	"database/sql"
	"errors"
	"hms/database"
	"log"
	"time"
)

type SignupData struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginData struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ChangePasswordData struct {
	Token    string `json:"token"`
	Password string `json:"password"`
}

type BookingDetails struct {
	RoomID         string    `json:"room-id"`
	GuestNames     string    `json:"guest-names"`
	PhoneNumbers   string    `json:"phone-numbers"`
	Emails         string    `json:"emails"`
	SpecialRequest string    `json:"special-requests"`
	PaymentMethod  string    `json:"payment-method"`
	Room           int       `quantity:"quantity"`
	PromoCode      string    `json:"promo-code"`
	Checkin        time.Time `json:"check-in"`
	CheckOut       time.Time `json:"check-out"`
	Night          int       `json:"number-of-night"`
}

func (b *BookingDetails) CalculateTotalAmount() (float64, error) {
	// panic("Not implemented")
	// total = (price_per_night * number_of_nights * quantity * ((100-promoRate)/100) * (100+taxrate)/100)
	price_per_night, err := GetPricePerNight(b.RoomID)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	number_of_nights := 1 //b.Night // to be caalculated from checkin and checkout dates
	quantity := 1         // b.Room
	promoRate := 0        // no promo 0 percent off. to be gotten from the promocode
	taxRate := 0          // no tax. 0 percent tax added. to be gotten from the database

	total := price_per_night * float64(number_of_nights) * float64(quantity) * float64((100-promoRate)/100) * float64((100+taxRate)/100)

	return total, nil
}

func GetPricePerNight(room_id string) (float64, error) {
	query := `SELECT price_per_night FROM rooms WHERE id = @room_id;`
	row := database.DB.QueryRow(query, sql.Named("room_id", room_id))
	var price float64
	err := row.Scan(&price)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, errors.New("no record for such room ID")
		}
		return 0, err
	}
	return price, nil
}
