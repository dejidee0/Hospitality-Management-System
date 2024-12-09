package database

import (
	"database/sql"
	"hms/config"
	"log"

	// _ "github.com/mattn/go-sqlite3"
	_ "github.com/denisenkom/go-mssqldb"
)

func GetDB() *sql.DB {
	// db, err := sql.Open("sqlite3", "data/hms.db")
	db, err := sql.Open(config.DBDriver, config.DBUrl)

	if err != nil {
		log.Fatalf("failed to connect to the database: %v", err)
	}
	log.Println("db connected successfully!")
	if err = db.Ping(); err != nil {
		log.Fatalln("failed to ping database...:" + err.Error())
	}
	log.Println("db pinged successfully!")
	return db
}

var DB = GetDB()
