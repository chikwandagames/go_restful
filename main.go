package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	// Handlers
	// To run
	// Curl -v localhost:9090
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		// ReadAll let you read from anything that implements io.Reader
		// Body implements io.ReadClosr
		d, err := ioutil.ReadAll(r.Body)
		if err != nil {
			// rw.WriteHeader(http.StatusBadRequest)
			// rw.Write([]byte("Ooops!"))

			// or
			http.Error(rw, "Ooops!", http.StatusBadRequest)
			return
		}

		log.Printf("Body contains: %s \n", d)
		// Write data back to the user
		fmt.Fprintf(rw, "Back to user: %s \n", d)
	})

	// To run, curl -v localhost:9090/goodbye
	http.HandleFunc("/goodbye", func(http.ResponseWriter, *http.Request) {
		log.Println("Bye")
	})

	http.ListenAndServe(":9090", nil)
}
