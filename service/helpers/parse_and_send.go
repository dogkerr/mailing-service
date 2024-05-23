package helpers

import (
	"bytes"
	"log"
	"text/template"

	"gopkg.in/gomail.v2"
)

func ParseAndSend(templatePath string, data interface{}, to string, subject string) error {
	// Get html
	var body bytes.Buffer
	t, err := template.ParseFiles(templatePath)
	if err != nil {
		log.Println("Error parsing template:", err)
		return err
	}

	if err = t.Execute(&body, data); err != nil {
		log.Println("Error executing template:", err)
		return err
	}

	// Send mail with gomail
	m := gomail.NewMessage()
	m.SetHeader("From", "dogker.dog@gmail.com")
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body.String())

	d := gomail.NewDialer("smtp.gmail.com", 587, "dogker.dog@gmail.com", "kquwwmxnyfnuvqjd")

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		log.Println("Error sending email:", err)
		return err
	}

	log.Printf("email successfully sent to %v\n", to)

	return nil
}
