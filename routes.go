package main

import (
	"encoding/json"
	"net/http"
)

func (a *AppConfig) routes() http.Handler {

	mux := http.NewServeMux()
	mux.HandleFunc("/ping", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		type Ping struct {
			Code int    `json:"code"`
			Msg  string `json:"msg"`
		}

		var p = Ping{
			Code: 200,
			Msg:  "all systems go",
		}

		if err := json.NewEncoder(w).Encode(p); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

	}))

	mux.HandleFunc("/api/v1/andrew-mccall/contact", app.HandlePostAndrewMcCallContact)

	return AddHeaders(HandlePreFlight(mux))

}
