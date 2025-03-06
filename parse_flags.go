package main

import "flag"

func parseFlags(app *AppConfig) {

	flag.StringVar(&app.Port, "port", ":8675", "the port the application listens on")
	flag.BoolVar(&app.IsProduction, "is_production", false, "app is in production(true|false)")
	flag.StringVar(&app.SMTPHost, "smtp_host", "localhost", "the hostname for the mail server")
	flag.IntVar(&app.SMTPPort, "smtp_port", 1025, "the port for the mail server")
	flag.StringVar(&app.SMTPUsername, "smtp_username", "tesuser", "the smtp username")
	flag.StringVar(&app.SMTPUserpass, "smtp_userpass", "testpass", "the smtp password")
	flag.Parse()

}
