package main

import (
	"log"
	"net/http"
	"os"

	"github.com/willyabimanyu/microservice-go/handlers"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)

	hh := handlers.NewHello(l)
	// gh := handlers.NewHGoodbye(l)

	sm := http.NewServeMux()
	sm.Handle("/", hh)

	// http.HandleFunc()

	http.ListenAndServe(":9090", sm)

}
