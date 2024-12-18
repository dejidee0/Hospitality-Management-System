package models

import (
	"database/sql"
	"hms/database"
	"log"
	"time"
)

//
// id VARCHAR(36) PRIMARY KEY,
// name VARCHAR(255) NOT NULL,
// summary VARCHAR(MAX),
// category VARCHAR(100),
// date DATETIME,
// venue VARCHAR(255),
// price DECIMAL(10,2) DEFAULT 0,
// about VARCHAR(MAX),
// images VARCHAR(MAX), -- comma sepaterated urlofimages
// format VARCHAR(100), -- e.g conference, class, festival, party
// popular BIT DEFAULT 0,
// state VARCHAR(50)

type Event struct {
	Id       string    `json:"id,omitempty"`
	Name     string    `json:"name,omitempty"`
	Summary  string    `json:"summary,omitempty"`
	Category string    `json:"category,omitempty"`
	State    string    `json:"state,omitempty"`
	Venue    string    `json:"venue,omitempty"`
	Date     time.Time `json:"date,omitempty"`
	Format   string    `json:"format,omitempty"`
	Price    float64   `json:"price,omitempty"`
	Images   string    `json:"images,omitempty"`
	Popular  bool      `json:"popular,omitempty"`
	About    string    `json:"about,omitempty"`
}

func (h *Event) GetEventsByState(state string) []Event {
	// SELECT hotels.id, hotels.name, hotels.city,	hotels.rating, hotels.images, COALESCE(MIN(rooms.price_per_night), 0) AS min_price_per_night, COUNT(reviews.id) AS review_count FROM hotels LEFT JOIN reviews
	// ON hotels.id = reviews.hotel_id INNER JOIN rooms ON hotels.id = rooms.hotel_id WHERE hotels.popular = 1 GROUP BY hotels.id, hotels.name, hotels.city, hotels.rating, hotels.images ORDER BY min_price_per_night;`

	query := `SELECT events.id, events.name, events.date, events.venue, events.price, events.category, events.format, events.state, events.images FROM events WHERE events.state = @state;`

	rows, err := database.DB.Query(query, sql.Named("state", state))
	if err != nil {
		log.Println(err)
		return nil
	}
	defer rows.Close()

	var events []Event
	for rows.Next() {
		var event Event
		err = rows.Scan(
			&event.Id, &event.Name, &event.Date, &event.Venue, &event.Price,
			&event.Category, &event.Format, &event.State, &event.Images)
		if err == nil {
			events = append(events, event)
		} else {
			log.Println("error on scan: " + err.Error())
		}
	}

	return events
}
