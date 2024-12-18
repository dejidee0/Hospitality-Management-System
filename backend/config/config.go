package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func LoadEnv() (err error) {
	if os.Getenv("ENV") == "dev" {
		err = godotenv.Load()
		if err != nil {
			log.Fatalf("error loading env: %v\n", err)
		}
		return
	}
	return nil
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

var PAYSTACK_SECRET_KEY_TEST = os.Getenv("PAYSTACK_SECRET_KEY_TEST")
var PAYSTACK_PUBLIC_KEY_TEST = os.Getenv("PAYSTACK_PUBLIC_KEY_TEST")

func getMailPort() int {
	fmt.Println("hit here 1")
	port := os.Getenv("MAIL_PORT")
	fmt.Println("the port is: " + port)
	// if !ok {
	// 	log.Fatal("error: MAIL_PORT environment variable not set o")
	// }
	MailPortInt, err := strconv.Atoi(port)
	if err != nil {
		log.Printf("error: MAIL_PORT environment variable must be an integer")
		log.Fatal(err)
	}
	return MailPortInt
}
