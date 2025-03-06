package main

import (
	"fmt"
	"html"
	"log"
	"net/http"

	"gopkg.in/gomail.v2"
)

var app AppConfig = AppConfig{}

func main() {
	parseFlags(&app)
	errChan := make(chan error)
	mailDoneChan := make(chan bool)
	msgChan := make(chan EmailAndMessage)
	app.MsgChan = msgChan
	app.DoneChan = mailDoneChan
	app.ErrorChan = errChan
	app.DB = nil
	app.Mailer = gomail.NewDialer(app.SMTPHost, app.SMTPPort, app.SMTPUsername, app.SMTPUserpass)

	go app.listenForMail()

	srv := &http.Server{
		Addr:    app.Port,
		Handler: app.routes(),
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Fatalln(err)
	}

}

func (a *AppConfig) listenForMail() {

	for {
		select {
		case msg := <-a.MsgChan:

			m := gomail.NewMessage()
			m.SetHeader("From", "no-reply@andrew-mccall.com")
			m.SetAddressHeader("To", "andrew@andrew-mccall.com", "Andrew M McCall")
			m.SetHeader("Subject", "andrew-mccall.com web form submission")
			data := make(map[string]any)

			data["Email"] = html.EscapeString(msg.Email)
			data["Message"] = html.EscapeString(msg.Message)

			tmpl, err := constructHTMLTemplate(data)

			if err != nil {
				log.Println(err)
			}

			m.SetBody("text/plain", fmt.Sprintf("Email: %s - Message: %s", msg.Email, msg.Message))
			m.AddAlternative("text/html", tmpl)

			err = app.Mailer.DialAndSend(m)

			if err != nil {
				log.Println(err)
			}
		case err := <-a.ErrorChan:
			log.Println(err)

		case <-a.DoneChan:
			return
		}
	}

}
