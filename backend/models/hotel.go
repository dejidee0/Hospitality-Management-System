package models

import (
	"database/sql"
	"hms/database"
	"log"
	"time"
)

type Hotel struct {
	Id               string  `json:"id,omitempty"`
	Name             string  `json:"name,omitempty"`
	PropertyType     string  `json:"property-type,omitempty"`
	City             string  `json:"city,omitempty"`
	State            string  `json:"state,omitempty"`
	Address          string  `json:"address,omitempty"`
	Rating           string  `json:"start,omitempty"`
	AvgRating        int     `json:"average-rating,omitempty"`
	Amenities        string  `json:"amenities,omitempty"`
	Price_per_night  float64 `json:"price-per-night,omitempty"`
	Images           string  `json:"images,omitempty"`
	ReviewCount      int     `json:"reviews,omitempty"`
	Description      string  `json:"description,omitempty"`
	CleanlinessScore int     `json:"cleanliness-score,omitempty"`
	AmenitiesScore   int     `json:"amenities-score,omitempty"`
	LocationScore    int     `json:"location-score,omitempty"`
	ServicesScore    int     `json:"services-score,omitempty"`
}

type Room struct {
	Id           string  `json:"id"`
	Name         string  `json:"name"`
	Type         string  `json:"type"`
	Images       string  `json:"images"`
	Amenities    string  `json:"amenities"`
	Capacity     int     `json:"capacity"`
	AllowSmoking bool    `json:"allow-smoking"`
	Size         string  `json:"room-size"`
	Price        float64 `json:"price-per-night"`
}

type Review struct {
	Name    string    `json:"person-name"`
	Country string    `json:"country"`
	OneWord string    `json:"remark"`
	Body    string    `json:"body"`
	Rating  float64   `json:"rating"`
	Date    time.Time `json:"date"`
}

type Policy struct {
	CheckIn       string `json:"check-in"`
	CheckOut      string `json:"check-out"`
	AllowChildren bool   `json:"allow-children"`
	AllowPets     string `json:"allow-pets"`
	PaymentMethod string `json:"payments-available"`
}

func (h *Hotel) GetPopularHotels() ([]Hotel, error) {
	// db, err := database.GetDB()
	// if err != nil {
	// 	log.Println(err)
	// 	return nil, err
	// }
	// defer db.Close()

	// query := `
	// SELECT hotels.id, hotels.name, hotels.city,	hotels.rating, hotels.price_per_night,
	// hotels.images, hotels.popular, COUNT(reviews.id) AS review_count FROM hotels LEFT JOIN reviews
	// ON hotels.id = reviews.hotel_id WHERE hotels.popular = true GROUP BY hotels.id;`

	query := `
	SELECT hotels.id, hotels.name, hotels.city,	hotels.rating, hotels.images, COALESCE(MIN(rooms.price_per_night), 0) AS min_price_per_night, COUNT(reviews.id) AS review_count FROM hotels LEFT JOIN reviews
	ON hotels.id = reviews.hotel_id INNER JOIN rooms ON hotels.id = rooms.hotel_id WHERE hotels.popular = 1 GROUP BY hotels.id, hotels.name, hotels.city, hotels.rating, hotels.images ORDER BY min_price_per_night;`

	rows, err := database.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var hotels []Hotel

	for rows.Next() {
		var hotel Hotel
		err = rows.Scan(
			&hotel.Id, &hotel.Name, &hotel.City, &hotel.Rating,
			&hotel.Images, &hotel.Price_per_night, &hotel.ReviewCount,
		)
		if err == nil {
			hotels = append(hotels, hotel)
		} else {
			log.Println("error on scan: " + err.Error())
		}
	}
	return hotels, nil
}

func (h *Hotel) GetHotelsByState(state string) []Hotel {
	// SELECT hotels.id, hotels.name, hotels.city,	hotels.rating, hotels.images, COALESCE(MIN(rooms.price_per_night), 0) AS min_price_per_night, COUNT(reviews.id) AS review_count FROM hotels LEFT JOIN reviews
	// ON hotels.id = reviews.hotel_id INNER JOIN rooms ON hotels.id = rooms.hotel_id WHERE hotels.popular = 1 GROUP BY hotels.id, hotels.name, hotels.city, hotels.rating, hotels.images ORDER BY min_price_per_night;`

	query := `SELECT hotels.id, hotels.name, hotels.property_type, hotels.city, hotels.state, hotels.rating, hotels.images, MAX(hotels.address) As address,
	COALESCE(MIN(rooms.price_per_night), 0) AS min_price_per_night,	COUNT(reviews.id) AS review_count, 
	AVG(ISNULL(reviews.rating, 0)) AS avg_rating, STRING_AGG(rooms.amenities, ', ') AS amenities FROM hotels LEFT JOIN reviews ON hotels.id = reviews.hotel_id 
	INNER JOIN rooms ON hotels.id = rooms.hotel_id WHERE hotels.state = @state GROUP BY hotels.id, hotels.name, hotels.property_type, hotels.city, hotels.state, hotels.rating, hotels.images;`

	rows, err := database.DB.Query(query, sql.Named("state", state))
	if err != nil {
		log.Println(err)
		return nil
	}
	defer rows.Close()

	var hotels []Hotel
	for rows.Next() {
		var hotel Hotel
		err = rows.Scan(
			&hotel.Id, &hotel.Name, &hotel.PropertyType, &hotel.City, &hotel.State, &hotel.Rating,
			&hotel.Images, &hotel.Address, &hotel.Price_per_night, &hotel.ReviewCount,
			&hotel.AvgRating, &hotel.Amenities,
		)
		if err == nil {
			hotels = append(hotels, hotel)
		} else {
			log.Println("error on scan: " + err.Error())
		}
	}

	return hotels
}

// this will populate data for the hotel pointed to by this method
func (h *Hotel) GetHotelByID(id string) error {

	query := `SELECT hotels.id, hotels.name, hotels.city, hotels.rating, MAX(hotels.address) AS address, hotels.images, MAX(description) AS description,
	STRING_AGG(rooms.amenities, ', ') AS amenities, COUNT(reviews.id) AS review_count, AVG(ISNULL(reviews.rating, 0)) AS avg_rating,
	AVG(ISNULL(reviews.cleanliness, 0)) AS cleanliness_score, AVG(ISNULL(reviews.amenities, 0)) AS amenities_score, 
	AVG(ISNULL(reviews.location, 0)) AS location_score, AVG(ISNULL(reviews.services, 0)) AS services_score FROM hotels 
	LEFT JOIN reviews ON hotels.id = reviews.hotel_id JOIN rooms ON hotels.id = rooms.hotel_id WHERE hotels.id = @id 
	GROUP BY hotels.id, hotels.name, hotels.city, hotels.rating, hotels.images;`

	row := database.DB.QueryRow(query, sql.Named("id", id))

	err := row.Scan(&h.Id, &h.Name, &h.City, &h.Rating, &h.Address, &h.Images, &h.Description,
		&h.Amenities, &h.ReviewCount, &h.AvgRating, &h.CleanlinessScore, &h.AmenitiesScore,
		&h.LocationScore, &h.ServicesScore,
	)
	if err != nil {

		log.Println(err)
		return err
	}

	return nil
}

func (h *Hotel) GetRooms() []Room {

	// id VARCHAR(36) PRIMARY KEY,
	// hotel_id VARCHAR(36) NOT NULL,
	// images VARCHAR(2048), -- comma separated url, the first being the main one
	// name VARCHAR(255) NOT NULL, -- name of room e.g 'Executive One-Bedroom Suite'
	// type VARCHAR(255) NOT NULL, -- type of room e.g 'Double Room'
	// total_number INT NOT NULL, --number of rooms of this type
	// capacity INT NOT NULL, --number of guests that can stay in this room
	// room_size VARCHAR(50) NOT NULL, --how big is the room e.g '30 square meter'
	// allow_smoking BIT NOT NULL,
	// single_bed INT DEFAULT 0,
	// double_bed INT DEFAULT 0,
	// king_sized INT DEFAULT 0,
	// super_king_sized INT DEFAULT 0,
	// price_per_night DECIMAL(10,2) DEFAULT 0.0,
	// amenities VARCHAR(512),

	query := `SELECT id, name, type, images, allow_smoking, room_size, 
	amenities, capacity, price_per_night FROM rooms WHERE hotel_id = @hotel_id;`

	rows, err := database.DB.Query(query, sql.Named("hotel_id", h.Id))
	if err != nil {
		log.Println(err)
		return nil
	}
	defer rows.Close()

	var rooms []Room
	for rows.Next() {
		var room Room
		err := rows.Scan(&room.Id, &room.Name, &room.Type, &room.Images,
			&room.AllowSmoking, &room.Size, &room.Amenities, &room.Capacity,
			&room.Price)
		if err != nil {
			log.Println(err)
			continue
		}
		rooms = append(rooms, room)
	}
	return rooms
}

func (h *Hotel) GetReviews() []Review {
	// Name    string    `json:"person-name"`
	//
	//	Country string    `json:"country"`
	//	OneWord string    `json:"remark"`
	//	Body    string    `json:"body"`
	//	Date    time.Time `json:"date"`
	query := `SELECT name, country, rating, one_word, review_body, 
	created_at FROM reviews WHERE hotel_id = @hotel_id;`

	rows, err := database.DB.Query(query, sql.Named("hotel_id", h.Id))
	if err != nil {
		log.Println(err)
		return nil
	}
	defer rows.Close()

	var reviews []Review
	for rows.Next() {
		var review Review
		err := rows.Scan(&review.Name, &review.Country, &review.Rating, &review.OneWord,
			&review.Body, &review.Date)
		if err != nil {
			log.Println(err)
			continue
		}
		reviews = append(reviews, review)
	}
	return reviews
}

func (h *Hotel) GetPolicy() *Policy {
	// hotel_id VARCHAR(36) NOT NULL,
	// check_in VARCHAR(50) NOT NULL, -- e.g 'from 14:00 to 00:00'
	// check_out VARCHAR(50) NOT NULL, -- e.g from 01:00 to 12:00
	// allow_children BIT NOT NULL,
	// allow_pets VARCHAR(10) DEFAULT 'no' CHECK(allow_pets IN ('yes', 'no', 'request')),
	// cancel_booking VARCHAR(100),  -- when can guests cancel booking for free, e.g before 18:00 on the day of check-in
	// cancel_booking_charge VARCHAR(100),
	// currency VARCHAR(50),
	// payment_method VARCHAR(50), -- comma separated 'credit card,cash,transfer'
	// FOREIGN KEY (hotel_id) REFERENCES hotels(id)
	query := `SELECT check_in, check_out, allow_children, allow_pets, payment_method 
				FROM house_rules WHERE hotel_id = @hotel_id;`

	var policy Policy
	row := database.DB.QueryRow(query, sql.Named("hotel_id", h.Id))

	err := row.Scan(&policy.CheckIn, &policy.CheckOut, &policy.AllowChildren, &policy.AllowPets, &policy.PaymentMethod)
	if err != nil {
		log.Println(err)
		return nil
	}

	return &policy
}

func (h *Hotel) GetSimilar() []Hotel {
	query := `SELECT hotels.id, hotels.name, hotels.city, hotels.images, 
	COALESCE(MIN(rooms.price_per_night), 0) AS min_price_per_night,	COUNT(reviews.id) AS review_count, 
	AVG(ISNULL(reviews.rating, 0)) AS avg_rating FROM hotels LEFT JOIN reviews ON hotels.id = reviews.hotel_id 
	INNER JOIN rooms ON hotels.id = rooms.hotel_id WHERE hotels.city = @city GROUP BY hotels.id, hotels.name, hotels.city, hotels.images;`
	rows, err := database.DB.Query(query, sql.Named("city", h.City))
	if err != nil {
		log.Println(err)
		return nil
	}
	defer rows.Close()

	var hotels []Hotel
	for rows.Next() {
		var hotel Hotel
		err = rows.Scan(
			&hotel.Id, &hotel.Name, &hotel.City, &hotel.Images, &hotel.Price_per_night, &hotel.ReviewCount,
			&hotel.AvgRating,
		)
		if err != nil {
			log.Println("error on scan: " + err.Error())
			continue
		}
		hotels = append(hotels, hotel)
	}

	return hotels
}
