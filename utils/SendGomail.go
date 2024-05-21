package utils

import (
	"bytes"
	"errors"
	"fmt"
	"text/template"

	"github.com/dogkerr/mailing-service/m/v2/structs"
	"gopkg.in/gomail.v2"
)

var templatePaths = map[string]string{
	"verification":   "templates/verification.html",
	"billing-notice": "templates/billing-notice.html",
}

func SendGomail(templateType structs.TemplateType, data structs.Data, subject string, to ...string) error {
	// Get data
	var validData interface{}

	if data.VerificationData != nil && templateType == structs.Verification {
		if err := data.VerificationData.Validate(); err != nil {
			fmt.Println("Error validating data:", err)
			return err
		}
		validData = data.VerificationData
	} else if data.BillingNoticeData != nil && templateType == structs.BillingNotice {
		if err := data.BillingNoticeData.Validate(); err != nil {
			fmt.Println("Error validating data:", err)
			return err
		}
		validData = data.BillingNoticeData
	} else {
		fmt.Println("Invalid data or template type")
		return errors.New("Invalid data or template type")
	}

	// Get html
	var body bytes.Buffer
	t, err := template.ParseFiles(templatePaths[string(templateType)])
	if err != nil {
		fmt.Println("Error parsing template:", err)
		return err
	}

	if err = t.Execute(&body, validData); err != nil {
		fmt.Println("Error executing template:", err)
		return err
	}

	// Send mail with gomail
	m := gomail.NewMessage()
	m.SetHeader("From", "dogker.dog@gmail.com")
	m.SetHeader("To", to...)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body.String())

	d := gomail.NewDialer("smtp.gmail.com", 587, "dogker.dog@gmail.com", "kquwwmxnyfnuvqjd")

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		fmt.Println("Error sending email:", err)
		return err
	}

	fmt.Printf("%s email successfully sent to %v\n", templateType, to)

	return nil
}
