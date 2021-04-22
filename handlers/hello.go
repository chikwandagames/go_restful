package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Hello handler
type Hello struct {
	l *log.Logger
}

// The benefit of the NewHello() dependancy injection comes when
// testing, in the test we can replace logger with something else
func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

// By attaching the ServerHTTP(ResponseWriter, *Request) function to
// Type Hello, Hello is now of type http.Handler
func (h *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	// ReadAll let you read from anything that implements io.Reader
	// Body implements io.ReadClosr
	d, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, "Ooops!", http.StatusBadRequest)
		return
	}

	// Using the logger
	h.l.Printf("Body contains: %s \n", d)
	// Write data back to the user
	fmt.Fprintf(rw, "Back to user: %s \n", d)
}
