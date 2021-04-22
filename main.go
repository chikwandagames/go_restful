package main

import (
	"log"
	"net/http"
)

func main() {
	// Handlers
	// To run
	// Curl -v localhost:9090
	http.HandleFunc("/", func(http.ResponseWriter, *http.Request) {
		log.Println("Hello world")
	})

	// To run, curl -v localhost:9090/goodbye
	http.HandleFunc("/goodbye", func(http.ResponseWriter, *http.Request) {
		log.Println("Bye")
	})

	http.ListenAndServe(":9090", nil)
}
