package main

import (
	"database/sql"

	"gopkg.in/gomail.v2"
)

type AppConfig struct {
	Port         string
	IsProduction bool
	DB           *sql.DB
	Mailer       *gomail.Dialer
	MsgChan      chan EmailAndMessage
	ErrorChan    chan error
	DoneChan     chan bool
	SMTPHost     string
	SMTPPort     int
	SMTPUsername string
	SMTPUserpass string
}
