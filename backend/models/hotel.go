package models

import (
	"hms/database"
	"log"
)

type Hotel struct {
	Id              string
	Name            string
	City            string
	Rating          string
	Price_per_night float64
	Images          string
	ReviewCount     int
	Popular         bool
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
	db := database.GetDB()
	defer db.Close()

	// query := `
	// SELECT hotels.id, hotels.name, hotels.city,	hotels.rating, hotels.price_per_night,
	// hotels.images, hotels.popular, COUNT(reviews.id) AS review_count FROM hotels LEFT JOIN reviews
	// ON hotels.id = reviews.hotel_id WHERE hotels.popular = true GROUP BY hotels.id;`

	query := `
	SELECT hotels.id, hotels.name, hotels.city,	hotels.rating, hotels.images, COALESCE(MIN(rooms.price_per_night), 0) AS min_price_per_night, COUNT(reviews.id) AS review_count FROM hotels LEFT JOIN reviews
	ON hotels.id = reviews.hotel_id LEFT JOIN rooms ON hotels.id = rooms.hotel_id WHERE hotels.popular = 1 GROUP BY hotels.id;`

	rows, err := db.Query(query)
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
