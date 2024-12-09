package models

import (
	"hms/database"
	"log"
	"time"
)

type Blog struct {
	Id           string
	Title        string
	Author       string `json:"author,omitempty"`
	DisplayImage string
	CreatedAt    time.Time
}

func (b *Blog) GetRecentBlogs() ([]Blog, error) {

	db, err := database.GetDB()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer db.Close()

	query := `SELECT TOP 3 id, title, display_image, created_at FROM blogs ORDER BY created_at DESC;`

	rows, err := db.Query(query)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()

	var blogs []Blog

	for rows.Next() {
		var blog Blog
		err = rows.Scan(
			&blog.Id, &blog.Title, &blog.DisplayImage, &blog.CreatedAt,
		)
		if err == nil {
			blogs = append(blogs, blog)
		} else {
			log.Println("error on scan: " + err.Error())
		}
	}
	return blogs, nil

}
