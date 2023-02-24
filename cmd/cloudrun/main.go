package main

import (
	"fmt"
	"go-aws-ecr-ecs/handlers"
	"log"
	"net/http"
	"os"
)

func main() {
	applicationName := os.Getenv("APPLICATION-NAME")
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	if len(applicationName) == 0 {
		log.Println("Running application locally")
	} else {
		log.Printf("Running application %s\n", applicationName)
	}
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			handlers.HelloHTTP(w, r)
		} else {
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})
	if len(port) == 0 {
		port = "8080"
	}
	addr := fmt.Sprintf("%s:%s", host, port)
	log.Printf("Running application in %s\n", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		panic("Error listening " + err.Error())
	}
}
