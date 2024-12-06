package mail

import (
	"fmt"
	"hms/config"

	"gopkg.in/gomail.v2"
)

func SendToken(token, email string) error {
	mailer := gomail.NewMessage()
	mailer.SetHeader("From", "noreply@hms.com")
	mailer.SetHeader("To", email)
	mailer.SetHeader("Subject", "HMS - PASSWORD RESET -TEST EMAIL")

	msg := "Here is the link to reset your password.\nwwww.findpeacefrontend.com/reset?reset_token=" + token

	mailer.SetBody("text/plain", msg)

	dialer := gomail.NewDialer(config.MailServer, config.MailPort, config.MailUsername, config.MailPassword)

	err := dialer.DialAndSend(mailer)
	if err != nil {
		return err
	}

	fmt.Println("Email sent successfully! to " + email)
	return nil
}
