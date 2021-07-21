// using SendGrid's Go Library
// https://github.com/sendgrid/sendgrid-go
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func main() {
	from := mail.NewEmail("Chapelwood HOA Admin", "admin@chapelwoodhoa.dev")
	subject := "Test Email Message"
	to := mail.NewEmail("Stephen Daniel", "swd@pobox.com")
	plainTextContent := "This is a test email message.  Please ignore it."
	htmlContent := "<strong>This</strong> is a test email message.  <strong>Please</strong> ignore it."
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	response, err := client.Send(message)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}
