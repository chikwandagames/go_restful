package main

import (
	"log"
	"net/http"
	"os"

	"github.com/chikwandagames/go_restful.git/handlers"
)

func main() {
	// Define a Logger, the logger can be used to log to a file, Stdout etc,
	// product-api is a prefix,
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	// Ref to the handler
	hh := handlers.NewHello(l)
	gh := handlers.NewGoodBye(l)

	// Register the new handler with the server
	// SeverMux is a multiplexer, and it a handler
	// Multiplexer - selects between several input signals
	// and forwards the selected input to a single output line
	sm := http.NewServeMux()
	sm.Handle("/", hh)
	sm.Handle("/goodbye", gh)

	http.ListenAndServe(":9090", sm)
}
