package utils

import (
	"bytes"
	"fmt"
	"text/template"

	"github.com/dogkerr/mailing-service/m/v2/structs"
	"gopkg.in/gomail.v2"
)

var templatePaths = map[string]string{
	"verification":   "./../templates/verification.html",
	"billing-notice": "./../templates/billing-notice.html",
}

func SendGomail(templateType structs.TemplateType, data structs.VerificationData, to string, subject string) {
	fmt.Println(data.Name)

	// Get html
	var body bytes.Buffer
	t, err := template.ParseFiles(templatePaths[string(templateType)])
	t.Execute(&body, data)

	if err != nil {
		return
	}

	// Send mail with gomail
	m := gomail.NewMessage()
	m.SetHeader("From", "davidlou0810@gmail.com")
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body.String())

	d := gomail.NewDialer("smtp.gmail.com", 587, "davidlou0810@gmail.com", "bjjsivudkhxblbgh")

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}
