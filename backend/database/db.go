package database

import (
	"database/sql"
	"errors"
	"hms/config"

	// _ "github.com/mattn/go-sqlite3"
	_ "github.com/denisenkom/go-mssqldb"
)

func GetDB() (*sql.DB, error) {
	// db, err := sql.Open("sqlite3", "data/hms.db")
	db, err := sql.Open(config.DBDriver, config.DBUrl)

	if err != nil {
		return nil, errors.New("failed to connect to the database: " + err.Error())
	}
	return db, nil
}
