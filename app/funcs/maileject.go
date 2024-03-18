package funcs

import (
	"fmt"

	gomail "gopkg.in/mail.v2"
)

type EjectEmail struct {
	Email    string
	Password string
}

func (sender EjectEmail) Send(subject string, body string, reciever string) error {
	m := gomail.NewMessage()

	// Set E-Mail sender
	m.SetHeader("From", sender.Email)

	// Set E-Mail receiver
	m.SetHeader("To", reciever)

	// Set E-Mail subject
	m.SetHeader("Subject", subject)

	// Set E-Mail body.
	m.SetBody("text/html", body)

	// Settings for SMTP server
	d := gomail.NewDialer("smtp.gmail.com", 587, sender.Email, sender.Password)

	// Now send E-Mail
	if err := d.DialAndSend(m); err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("Email Sent!")
	return nil
}
