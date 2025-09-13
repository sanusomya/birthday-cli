package reminder

import (
	"fmt"
	"net/smtp"
	"os"
)

func sendHtmlMail(name string, to string, date int8, month string, mobile int64) {

	auth := smtp.PlainAuth("", os.Getenv("mail"), os.Getenv("pass"), "smtp.gmail.com")
	subject := "Subject: Birthday Reminder\n"
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body := fmt.Sprintf("<html><body><h1>Reminder for %s's birthday</h1><p>Do wish them on %d %s</p><p>Contact them on %d</p></body></html>", name, date, month, mobile)
	msg := []byte(subject + mime + body)
	err := smtp.SendMail("smtp.gmail.com:587", auth, os.Getenv("from_mail"), []string{to}, msg)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println("Email Sent Successfully!")
}
