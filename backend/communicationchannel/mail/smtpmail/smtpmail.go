package smtpmail

import (
	"log"
	"lsapp/util"
	"net/smtp"
)

type SMTP struct {
	Name string
}

func (m SMTP) SendMail(request util.Request) (response util.Response, err error) {
	// Sender's email address
	sender := "lovlesh0107@gmail.com"

	// Recipient's email address
	to := "lovleshbaghel0107@gmail.com"

	// SMTP server configuration
	smtpHost := "smtp.example.com"
	smtpPort := "587" // or "465" for SSL/TLS

	// Compose the email message
	subject := "Test Email"
	body := "This is a test email sent from Go."
	msg := []byte("To: " + to + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"\r\n" +
		body)

	// Send the email using SMTP
	err = smtp.SendMail(smtpHost+":"+smtpPort, nil, sender, []string{to}, msg)

	log.Println("Email sent successfully!")
	if err != nil {
		log.Println("Error:", err)
		return
	}

	log.Println("Email sent successfully!")
	return
}
