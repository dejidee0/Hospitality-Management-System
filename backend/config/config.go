package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func LoadEnv() (err error) {
	err = godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	return
}

var _ = LoadEnv()

var host string = os.Getenv("SERVER_HOST")
var port string = os.Getenv("SERVER_PORT")
var Server_address = fmt.Sprintf("%s:%s", host, port)

var (
	DBDriver, DBUrl = func() (string, string) {
		driver := os.Getenv("DB_DRIVER")
		if driver == "" {
			return "sqlite3", "data/hms.db"
		}
		var dbHost = os.Getenv("DATABASE_HOST")
		var dbPort = os.Getenv("DATABASE_PORT")

		var dbName = os.Getenv("DATABASE_NAME")
		var dbUser = os.Getenv("DATABASE_USERNAME")
		var dbPassword = os.Getenv("DATABASE_PASSWORD")
		var url = fmt.Sprintf("server=%s;port=%s;database=%s;user id=%s;password=%s", dbHost, dbPort, dbName, dbUser, dbPassword)
		return driver, url
	}()
)

var JWTKey = []byte(os.Getenv("JWT_SECRET_KEY"))

var MailServer = os.Getenv("MAIL_SERVER")

var MailPort = getMailPort()
var MailUsername = os.Getenv("MAIL_USERNAME")
var MailPassword = os.Getenv("MAIL_PASSWORD")

func getMailPort() int {
	port, ok := os.LookupEnv("MAIL_PORT")
	if !ok {
		log.Fatal("error: MAIL_PORT environment variable not set")
	}
	MailPortInt, err := strconv.Atoi(port)
	if err != nil {
		log.Printf("error: MAIL_PORT environment variable must be an integer")
		log.Fatal(err)
	}
	return MailPortInt
}
