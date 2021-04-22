package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

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

	// Manually create a server, we need to handle a situation, by which a
	// server can get overloaded by too many requests
	s := &http.Server{
		Addr:         ":9090",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	// Wrap with goroutine to prevent blocking
	go func() {
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	// When execution goes beyond the above goroutine the s.Shutdown() will terminate
	// the program, we can prevent this using the os.Signal
	// Here we us the os.Signal package to register for signals Kill and Interrupt,
	// Interrup signal e.g. ctrl + c
	sigChan := make(chan os.Signal)
	// sig.Notify, will broadcast a message on this channel
	// whenever a os kill command, or os interrupt is received
	// Notify() takes a channel and signal
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	// Reading from a channel blocks until a message is available to be consumed
	// Which is what we want as long as the server is running, otherwise if we receive a message
	// in the channel it means an os.kill or os.Interrupt command has been recieved, at
	// s.Shutdown() will close the program
	sig := <-sigChan
	l.Println("Received terminate, graceful shutdown", sig)

	// Graceful shutdown,
	// Incase of server shutdown, Shutdown() waits until all requests that are
	// currently being handled by server have completed then shuts down,
	// Shutdown takes a context
	// Timeout Contex, with duration 30s, this will force handlers to close after 30s
	// if they are still working 30s after attempting graceful shutdown
	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)
}
