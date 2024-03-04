package mail

import (
	"log"
	"net/smtp"
	"os"

	"github.com/joho/godotenv"
)

func Mail(name, from, subject, message string) error {
	err := godotenv.Load(".env")
	if err != nil {
		return err
	}

	//get hostname
	hostname := os.Getenv("ENDPOINT")
	usr := os.Getenv("SMTP_USR")
	pass := os.Getenv("SMTP_PASS")
	id := os.Getenv("ID")
	sender := os.Getenv("SEND_EMAIL")

	payload := buildPayload(name, from, sender, subject, message)

	//set up authentication info
	auth := smtp.PlainAuth(id, usr, pass, hostname)

	to := []string{sender}
	//send email
	if err := smtp.SendMail(hostname+":25", auth, sender, to, payload); err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func buildPayload(n, t, f, s, m string) []byte {
	// Construct the email payload with appropriate headers
	payload := "Subject: " + s + "\r\n"
	payload += "From: " + f + "\r\n"
	payload += "To: " + t + "\r\n"   // Add the recipient email address here if needed
	payload += "\r\n"                // Empty line before the message body
	payload += "Name: " + n + "\r\n" //add name to message
	payload += m + "\r\n"            // Message body
	return []byte(payload)
}
