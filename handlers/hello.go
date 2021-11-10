package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Hello struct {
	l *log.Logger
}

func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

func (h *Hello) ServerHTTP(rw http.ResponseWriter, r *http.Request) {
	h.l.Println("Hellow World")
	d, err := ioutil.ReadAll(r.Body)

	if err != nil {
		// rw.WriteHeader(http.StatusBadRequest)
		http.Error(rw, "Oops", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(rw, "Hellow %s\n", d)

}
