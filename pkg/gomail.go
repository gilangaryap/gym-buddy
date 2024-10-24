package pkg

import (
	"log"
	"os"

	"gopkg.in/gomail.v2"
)

type EmailSender struct {
    Dialer *gomail.Dialer
}

func NewEmailSender() *EmailSender {
    host := "smtp.gmail.com"
    port := 587
    user := os.Getenv("USER_EMAIL")
    pass := os.Getenv("USER_PASS")

    dialer := gomail.NewDialer(host, port, user, pass)

    return &EmailSender{
        Dialer: dialer,
    }
}

func (es *EmailSender) Send(to, subject, body string) error {
    from := os.Getenv("FROM_GMAIL")

    m := gomail.NewMessage()
    m.SetHeader("From", from)
    m.SetHeader("To", to)
    m.SetHeader("Subject", subject)
    m.SetBody("text/plain", body)

    if err := es.Dialer.DialAndSend(m); err != nil {
        log.Fatalf("Failed to send email: %v", err)
        return err
    }

    return nil
}