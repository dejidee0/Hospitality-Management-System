package models

import (
	"database/sql"
	"hms/database"
	"log"
)

type Hotel struct {
	Id              string
	Name            string
	City            string
	State           string
	Address         string
	Rating          string
	AvgRating       int
	Amenities       string
	Price_per_night float64
	Images          string
	ReviewCount     int
}

// type PopularHotel struct {
// 	Id              string
// 	Name            string
// 	City            string
// 	Rating          string
// 	Price_per_night float64
// 	Images          string
// 	ReviewCount     int
// 	Popular         bool
// }

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

	query := `SELECT hotels.id, hotels.name, hotels.city, hotels.state, hotels.rating, hotels.images, 
	COALESCE(MIN(rooms.price_per_night), 0) AS min_price_per_night,	COUNT(reviews.id) AS review_count, 
	AVG(ISNULL(reviews.rating, 0)) AS avg_rating, STRING_AGG(rooms.amenities, ', ') AS amenities FROM hotels LEFT JOIN reviews ON hotels.id = reviews.hotel_id 
	INNER JOIN rooms ON hotels.id = rooms.hotel_id WHERE hotels.state = @state GROUP BY hotels.id, hotels.name, hotels.city, hotels.state, hotels.rating, hotels.images;`

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
			&hotel.Id, &hotel.Name, &hotel.City, &hotel.State, &hotel.Rating,
			&hotel.Images, &hotel.Price_per_night, &hotel.ReviewCount,
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
