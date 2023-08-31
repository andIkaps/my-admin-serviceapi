package utils

import (
	"fmt"
	"net/smtp"
	"os"
)

func Email(to []string, message []byte) {
	userFrom := os.Getenv("SMTP_USER")
	userPass := os.Getenv("SMTP_PASS")
	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort := os.Getenv("SMTP_PORT")

	auth := smtp.PlainAuth("", userFrom, userPass, smtpHost)

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, userFrom, to, message)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Email Sent Successfully!")
}
