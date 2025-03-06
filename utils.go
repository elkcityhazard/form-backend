package main

import (
	"bytes"
	"encoding/json"
	"html/template"
	"net/http"
)

func (a *AppConfig) HandleError(w http.ResponseWriter, r *http.Request, statusCode int, err error, data any) {

	errJSON := NewErrorMsg(statusCode, err.Error(), data)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if err = json.NewEncoder(w).Encode(errJSON); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func constructHTMLTemplate(data any) (string, error) {

	var tmplString string = `<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title>Document</title>
	</head>
	<body>

		{{with .Email }}
			<p><b>From: </b>{{ . }}</p>
		{{ end }}
		{{ with .Message }}
		<p><b>Message: </b>{{.}}</p>
		{{ end }}
	</body>
	</html>`

	tmpl, err := template.New("email").Parse(tmplString)

	if err != nil {
		return "", err
	}

	buf := &bytes.Buffer{}

	err = tmpl.Execute(buf, data)

	if err != nil {
		return "", err
	}

	return buf.String(), nil

}
