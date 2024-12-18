package mail

import (
	"fmt"
	"hms/config"
	"io/ioutil"
	"net/http"
	"strings"

	"gopkg.in/gomail.v2"
)

// CNAME mt80.imole.tech -> smtp.mailtrap.live
// TXT imole.tech -> v=spf1 include:_spf.smtp.mailtrap.live ~all
// CNAME rwmt1._domainkey.imole.tech -> rwmt1.dkim.smtp.mailtrap.live
// CNAME rwmt2._domainkey.imole.tech -> rwmt2.dkim.smtp.mailtrap.live
// TXT _dmarc.imole.tech -> v=DMARC1; p=none; rua=mailto:dmarc@smtp.mailtrap.live; ruf=mailto:dmarc@smtp.mailtrap.live; rf=afrf; pct=100
// CNAME mt-link.imole.tech -> t.mailtrap.live

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
		// fmt.Println(err)
		return err
	}

	fmt.Println("Email sent successfully! to " + email)
	return nil
}

func MailTrapGo() {

	url := "https://send.api.mailtrap.io/api/send"
	method := "POST"

	// str := `
	// {
	// 	"from":	{"email":"hello@demomailtrap.com","name":"Mailtrap Test"},
	// 	"to": 	[{"email":"dremkay71@gmail.com"}],
	// 	"subject":	"You are awesome!",
	// 	"text":	"Congrats for sending test email with Mailtrap!",
	// 	"category": "Integration Test"
	// }`

	payload := strings.NewReader(`{\"from\":{\"email\":\"hello@demomailtrap.com\",\"name\":\"Mailtrap Test\"},\"to\":[{\"email\":\"dremkay71@gmail.com\"}],\"subject\":\"You are awesome!\",\"text\":\"Congrats for sending test email with Mailtrap!\",\"category\":\"Integration Test\"}`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Authorization", "Bearer 81c69591c43eceac0b49ffc51a7dc072")
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
