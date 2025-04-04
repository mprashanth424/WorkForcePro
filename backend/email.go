// email.go
package main

import (
	"fmt"
	"log"
	"net/smtp"
	// Correct import for the `database` package
)

// sendWelcomeEmail sends a welcome email to the new employee
func sendWelcomeEmail(toEmail, firstName, employeeID, password string) error {
	// Set up the email details
	from := ""       // Replace with your email
	passwordEmail := "" // Replace with your email password
	smtpHost := ""
	smtpPort := ""

	// Message to send
	subject := "Welcome to the Company"
	body := fmt.Sprintf(
		"Hello %s,\n\nWelcome to our company! Your Employee ID is %s and your temporary password is %s.\n\nPlease change your password after your first login.\n\nBest Regards,\nYour Company Team",
		firstName, employeeID, password)
	message := fmt.Sprintf("Subject: %s\r\n\r\n%s", subject, body)

	// Set up authentication information
	auth := smtp.PlainAuth("", from, passwordEmail, smtpHost)

	log.Println("Starting to send welcome email...")

	// Send email
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{toEmail}, []byte(message))
	if err != nil {
		log.Printf("Error while sending email to %s: %v", toEmail, err)
		return fmt.Errorf("failed to send welcome email to %s: %w", toEmail, err)
	}

	log.Printf("Welcome email sent successfully to: %s", toEmail)
	return nil
}
