package main

import (
	"encoding/json"
	"errors"
	"html"
	"io"
	"log"
	"net/http"
	"net/mail"
	"strings"
)

func (a *AppConfig) HandlePostAndrewMcCallContact(w http.ResponseWriter, r *http.Request) {

	var reqBody *EmailAndMessage = &EmailAndMessage{}

	if r.Body != nil {

		defer r.Body.Close()

		body, err := io.ReadAll(r.Body)

		if err != nil {
			app.HandleError(w, r, 400, err, reqBody)
			return
		}

		err = json.Unmarshal(body, &reqBody)

		if err != nil {
			app.HandleError(w, r, 400, err, reqBody)
			return
		}

		reqBody.Email = html.EscapeString(reqBody.Email)
		reqBody.Message = html.EscapeString(reqBody.Message)
		reqBody.PhoneNumber = html.EscapeString(reqBody.PhoneNumber)

	}

	_, err := mail.ParseAddress(reqBody.Email)

	if err != nil {
		app.HandleError(w, r, 400, err, reqBody)
		return
	}

	if len(strings.TrimSpace(reqBody.Email)) == 0 {
		err = errors.New("invalid message content")

		app.HandleError(w, r, 400, err, reqBody)
		return
	}

	if reqBody.PhoneNumber != "" {
		log.Println("honeypot captured")
		err = errors.New("invalid message content")
		app.HandleError(w, r, http.StatusUnauthorized, err, reqBody)
		return
	}

	if len(strings.TrimSpace(reqBody.Message)) < 1 {
		err = errors.New("cannot send empty message")
		app.HandleError(w, r, http.StatusBadRequest, err, reqBody)
		return
	}

	a.MsgChan <- *reqBody

	if err = json.NewEncoder(w).Encode(reqBody); err != nil {
		app.HandleError(w, r, 500, err, reqBody)
		return
	}

}
