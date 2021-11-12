package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Goodbye is a simple handler
type Goodbye struct {
	l *log.Logger
}

// NewGoodbye creates a new hello handler with the given logger
func NewGoodbye(l *log.Logger) *Goodbye {
	return &Goodbye{l}
}

// ServeHTTP implements the go http.Handler interface
// https://golang.org/pkg/net/http/#Handler
func (h *Goodbye) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	h.l.Println("Handle requests")

	// read the body
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		h.l.Println("Error reading body", err)

		http.Error(rw, "Unable to read request body", http.StatusBadRequest)
		return
	}

	// write the response
	fmt.Fprintf(rw, "Goodbye %s", b)
}
