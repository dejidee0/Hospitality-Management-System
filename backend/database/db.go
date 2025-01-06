package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	// _ "github.com/mattn/go-sqlite3"
	_ "github.com/denisenkom/go-mssqldb"
)

func GetDB() *sql.DB {
	// db, err := sql.Open("sqlite3", "data/hms.db")
	dbHost := os.Getenv("DATABASE_HOST")
	dbPort := os.Getenv("DATABASE_PORT")
	dbName := os.Getenv("DATABASE_NAME")
	dbUser := os.Getenv("DATABASE_USERNAME")
	dbPassword := os.Getenv("DATABASE_PASSWORD")
	DBDriver := os.Getenv("DB_DRIVER")

	DBUrl := fmt.Sprintf("server=%s;port=%s;database=%s;user id=%s;password=%s", dbHost, dbPort, dbName, dbUser, dbPassword)

	db, err := sql.Open(DBDriver, DBUrl)

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
