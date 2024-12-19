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

// fetch events that are popular and are from the given state
func (e *Event) GetPopularEventsIn(state string) ([]Event, error) {
	query := `SELECT events.id, events.name, events.date, events.venue, events.price,
	events.images, events.popular FROM events WHERE popular = 1 AND state = @state;`

	rows, err := database.DB.Query(query, sql.Named("state", state))
	if err != nil {
		// log.Println(err)
		return nil, err
	}
	defer rows.Close()

	var events []Event
	for rows.Next() {
		var event Event
		err = rows.Scan(
			&event.Id, &event.Name, &event.Date, &event.Venue, &event.Price,
			&event.Images, &event.Popular)
		if err == nil {
			events = append(events, event)
		} else {
			log.Println("error on scan: " + err.Error())
		}
	}

	return events, nil
}

func (e *Event) GetEventsByFormat(format string) []Event {
	query := `SELECT events.id, events.name, events.date, events.venue, events.price,
	events.images FROM events WHERE format = @format;`

	rows, err := database.DB.Query(query, sql.Named("format", format))
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
			&event.Images)
		if err == nil {
			events = append(events, event)
		} else {
			log.Println("error on scan: " + err.Error())
		}
	}

	return events
}

func (e *Event) GetEventsByCategory(category string) []Event {
	query := `SELECT events.id, events.name, events.date, events.venue, events.price,
	events.images FROM events WHERE category = @cat;`

	rows, err := database.DB.Query(query, sql.Named("cat", category))
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
			&event.Images)
		if err == nil {
			events = append(events, event)
		} else {
			log.Println("error on scan: " + err.Error())
		}
	}

	return events
}

func (e *Event) GetEventsByState(state string) []Event {
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

// this will populate data for the event pointed to by this method
func (e *Event) GetEventByID(id string) error {

	query := `SELECT events.id, events.name, events.date, events.summary, events.venue, events.about, 
	events.price, events.images, events.category FROM events WHERE events.id = @id;`

	row := database.DB.QueryRow(query, sql.Named("id", id))

	err := row.Scan(&e.Id, &e.Name, &e.Date, &e.Summary, &e.Venue, &e.About, &e.Price,
		&e.Images, &e.Category)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
