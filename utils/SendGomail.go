package utils

import (
	"bytes"
	"fmt"
	"text/template"

	"github.com/dogkerr/mailing-service/m/v2/structs"
	"gopkg.in/gomail.v2"
)

var templatePaths = map[string]string{
	"verification":   "templates/verification.html",
	"billing-notice": "templates/billing-notice.html",
}

func SendGomail(templateType structs.TemplateType, data structs.Data, subject string, to ...string) {
	// Get data
	var validData interface{}

	if data.VerificationData != nil && templateType == structs.Verification {
		validData = data.VerificationData
	} else if data.BillingNoticeData != nil && templateType == structs.BillingNotice {
		validData = data.BillingNoticeData
	} else {
		fmt.Println("Invalid data or template type")
		return
	}

	// Get html
	var body bytes.Buffer
	t, err := template.ParseFiles(templatePaths[string(templateType)])
	if err != nil {
		fmt.Println("Error parsing template:", err)
		return
	}

	if err = t.Execute(&body, validData); err != nil {
		fmt.Println("Error executing template:", err)
		return
	}

	// Send mail with gomail
	m := gomail.NewMessage()
	m.SetHeader("From", "davidlou0810@gmail.com")
	m.SetHeader("To", to...)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body.String())

	d := gomail.NewDialer("smtp.gmail.com", 587, "davidlou0810@gmail.com", "bjjsivudkhxblbgh")

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		fmt.Println("Error sending email:", err)
		return
	}

	fmt.Printf("%s email successfully sent to %v", templateType, to)
}
