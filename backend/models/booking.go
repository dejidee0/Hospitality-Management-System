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

// {
//     "room-id": "room1-uuid-1234",
//     "guest-names": "qwertyu,qwertyu",
//     "phone-numbers": "345678,345678",
//     "emails": "imole@gmail.com,sdfghj,mike@ymail.com",
//     "special-requests": "dfghjkjhgfdssdfgh",
//     "payment-method": "cash",
//     "Room": 15,
//     "promo-code": "wertyui",
//     "check-in": "2024-12-13T15:00:00Z",
//     "check-out": "2024-12-17T15:00:00Z",
//     "number-of-night": 4
//   }

type HotelBooking struct {
	Id             string    `json:"id"`
	RoomID         string    `json:"room-id" binding:"required"`
	GuestNames     string    `json:"guest-names" binding:"required"`
	PhoneNumbers   string    `json:"phone-numbers"`
	Emails         string    `json:"emails" binding:"required"`
	Email          string    `json:"-"`
	SpecialRequest string    `json:"special-requests"`
	PaymentMethod  string    `json:"payment-method"`
	Room           int       `quantity:"quantity" binding:"required,min=1"`
	PromoCode      string    `json:"promo-code"`
	Checkin        time.Time `json:"check-in" binding:"required"`
	CheckOut       time.Time `json:"check-out" binding:"required"`
	Night          int       `json:"number-of-night" binding:"required,min=1"`
	TotalAmount    float64   `json:"total-amount"`
	BookingNumber  int64     `json:"booking-number"`
	AccessCode     string    `json:"access-code"`
	Reference      string    `json:"tx-reference"`
}

func (b *HotelBooking) Save() error {
	id := uuid.New().String()
	query := `INSERT INTO hotel_bookings (id, room_id, guest_names, phone_number, 
	email, check_in, check_out, price, rooms, special_requests, payment_method) OUTPUT INSERTED.booking_number 
	VALUES(@id, @room_id, @guest_names, @phone_number, @email, @check_in, @check_out, @price, @rooms, @special_requests, @payment_method);`

	row := database.DB.QueryRow(query, sql.Named("id", id), sql.Named("room_id", b.RoomID), sql.Named("guest_names", b.GuestNames),
		sql.Named("phone_number", b.PhoneNumbers), sql.Named("email", b.Emails), sql.Named("check_in", b.Checkin),
		sql.Named("check_out", b.CheckOut), sql.Named("price", b.TotalAmount), sql.Named("rooms", b.Room),
		sql.Named("special_requests", b.SpecialRequest), sql.Named("payment_method", b.PaymentMethod),
	)
	var bookinNum int64
	err := row.Scan(&bookinNum)
	if err != nil {
		log.Println(err)
		return err
	}
	b.BookingNumber = bookinNum
	b.Id = id
	return nil
}

func (b *HotelBooking) ValidateBookingAndCalculateTotalAmount() error {
	emails := strings.Split(b.Emails, ",")

	ok := utils.ValidateEmail(emails[0])
	if !ok {
		log.Println("first email in the email list is invalid")
		return errors.New("first email in the email list is invalid")
	}
	b.Email = emails[0]

	totalAmount, err := b.CalculateTotalAmount()
	if err != nil {
		log.Println(err)
		return err
	}
	b.TotalAmount = totalAmount

	return nil

}

func (b *HotelBooking) SaveReference(access_code, reference string) error {
	id := uuid.New().String()

	query := `INSERT INTO paystack_bookings (id, booking_number, booking_id, amount, type, reference, access_code) 
	VALUES(@id, @booking_number, @booking_id, @amount, @type, @reference, @access_code);`

	_, err := database.DB.Exec(query, sql.Named("id", id), sql.Named("booking_number", b.BookingNumber), sql.Named("booking_id", b.Id),
		sql.Named("amount", b.TotalAmount), sql.Named("type", "hotel"), sql.Named("reference", reference),
		sql.Named("access_code", access_code),
	)
	if err != nil {
		log.Println(err)
		return err
	}
	b.AccessCode = access_code
	b.Reference = reference
	return nil
}

func (b *HotelBooking) GetBookingDetails(booking_id string) (*BookingType, error) {
	query := `SELECT hotel_bookings.guest_names, hotel_bookings.booking_number, hotel_bookings.email, 
	hotel_bookings.check_in, hotel_bookings.check_out, hotel_bookings.rooms, hotel_bookings.price AS total_price,
	rooms.name, rooms.amenities, rooms.price_per_night, rooms.images, hotels.name FROM hotel_bookings JOIN rooms 
	ON hotel_bookings.room_id = rooms.id JOIN hotels ON rooms.hotel_id = hotels.id WHERE hotel_bookings.id = @booking_id;`

	var data BookingType
	row := database.DB.QueryRow(query, sql.Named("booking_id", booking_id))
	err := row.Scan(&data.GuestNames, &data.BookingNumber, &data.Email, &data.CheckIn, &data.CheckOut,
		&data.Rooms, &data.TotalPrice, &data.RoomName, &data.Amenities, &data.PricePerNight, &data.Images, &data.HotelName,
	)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	data.Email = strings.Split(data.Email, ",")[0]
	data.GuestNames = strings.Split(data.GuestNames, ",")[0]
	data.Images = strings.Split(data.Images, ",")[0]
	return &data, nil
}

type BookingType struct {
	GuestNames    string    `json:"guest-name"`
	BookingNumber int64     `json:"booking-number"`
	Email         string    `json:"email"`
	CheckIn       time.Time `json:"check-in"`
	CheckOut      time.Time `json:"check-out"`
	Rooms         int       `json:"rooms"`
	TotalPrice    float64   `json:"total-price"`
	RoomName      string    `json:"room-name"`
	Amenities     string    `json:"amenities"`
	PricePerNight float64   `json:"price-per-night"`
	Images        string    `json:"image"`
	HotelName     string    `json:"hotel-name"`
	TaxRate       string    `json:"tax-rate"`
	Discount      float64   `json:"discount"`
}

func (b *HotelBooking) CalculateTotalAmount() (float64, error) {
	// panic("Not implemented")
	// total = (price_per_night * number_of_nights * quantity * ((100-promoRate)/100) * (100+taxrate)/100)
	price_per_night, err := getPricePerNight(b.RoomID)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	number_of_nights := b.Night // to be caalculated from checkin and checkout dates
	quantity := b.Room
	promoRate := 0 // no promo 0 percent off. to be gotten from the promocode
	taxRate := 0   // no tax. 0 percent tax added. to be gotten from the database

	total := price_per_night * float64(number_of_nights) * float64(quantity) * float64((100-promoRate)/100) * float64((100+taxRate)/100)

	return total, nil
}

func getPricePerNight(room_id string) (float64, error) {
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
