//usr/bin/env go run "$0" "$@"; exit "$?"

package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Main is the entry point for the web server
func main() {
	log.Println("[Go Qlueless Assembly API]: Starting application")

	s := Server{
		router: mux.NewRouter(),
	}

	s.preload()
	s.routes()

	log.Println("[Go Qlueless Assembly API]: Starting server")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
