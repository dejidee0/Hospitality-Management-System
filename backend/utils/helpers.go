package utils

import "time"

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

	return 0, nil
}

// id VARCHAR(36) PRIMARY KEY, -- this will be changed in sql server since i cant have two primary keys
// booking_number INT IDENTITY(1,1),
// room_id VARCHAR(36) NOT NULL,
// guest_names VARCHAR(512) NOT NULL, --comma separated names in case of multiple guests
// phone_number VARCHAR(15) NOT NULL,
// email VARCHAR(100) NOT NULL,
// check_in DATETIME NOT NULL,
// check_out DATETIME NOT NULL,
// price DECIMAL(10,2) NOT NULL,
// payment_status VARCHAR(20) DEFAULT 'pending' CHECK(payment_status IN ('pending', 'processing', 'paid', 'refunded')),
// created_at DATETIME DEFAULT GETDATE(),
// updated_at DATETIME,
// rooms INT DEFAULT 1,
// special_requests VARCHAR(1024),
// payment_method VARCHAR(20),
// FOREIGN KEY (room_id) REFERENCES rooms(id)
