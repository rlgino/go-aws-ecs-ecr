package handlers

import (
	"encoding/json"
	"fmt"
	"html"
	"log"
	"net/http"
)

// HelloHTTP is an HTTP Cloud Function with a request parameter.
func HelloHTTP(w http.ResponseWriter, r *http.Request) {
	var d struct {
		Name string `json:"name"`
	}
	if err := json.NewDecoder(r.Body).Decode(&d); err != nil {
		fprintf(w, "Error decoding data %s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	log.Printf("GET --> %v", d)
	if d.Name == "" {
		fprintf(w, "Hello, World!")
		return
	}
	fprintf(w, "Hello, %s!", html.EscapeString(d.Name))
}

func fprintf(w http.ResponseWriter, msg string, params ...any) {
	if _, err := fmt.Fprintf(w, msg, params...); err != nil {
		return
	}
}
