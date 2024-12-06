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
