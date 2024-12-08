package database

import (
	"database/sql"
	"hms/config"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func GetDB() *sql.DB {
	// db, err := sql.Open("sqlite3", "data/hms.db")
	db, err := sql.Open(config.DBDriver, config.DBUrl)

	if err != nil {
		log.Fatalf("Failed to connect to the database: %v\n", err)
	}
	if err = db.Ping(); err != nil {
		log.Printf("Not pinging...: %v", err)
		return nil
	}
	return db
}
