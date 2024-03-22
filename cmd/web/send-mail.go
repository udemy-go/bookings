package main

import (
	"log"
	"time"

	"github.com/thiruthanikaiarasu/udemy-go/bookings/internal/models"
	mail "github.com/xhit/go-simple-mail"
)

func listenForMail() {
	// the mail have to be sent in the background process, it may take some time, but i don't want to wait for that seconds
	go func() {
		for {
			msg := <-app.MailChan
			sendMsg(msg)
		}
	}()

}

func sendMsg(m models.MailData) {
	server := mail.NewSMTPClient()
	server.Host = "localhost"
	server.Port = 1025 // actual port can be used is 25, 587, 465
	server.KeepAlive = false
	server.ConnectTimeout = 10 * time.Second
	server.SendTimeout = 10 * time.Second

	client, err := server.Connect()
	if err != nil {
		errorLog.Println(err)
	}

	email := mail.NewMSG()
	email.SetFrom(m.From).AddTo(m.To).SetSubject(m.Subject)
	email.SetBody(mail.TextHTML, m.Content)

	err = email.Send(client)
	if err != nil {
		log.Println(err)
	} else {
		log.Println("Email sent!")
	}
}
