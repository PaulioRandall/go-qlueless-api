//usr/bin/env go run "$0" "$@"; exit "$?"

package main

import (
	"log"
	"net/http"
)

// Main is the entry point for the web server
func main() {
	log.Println("[Go Qlueless Assembly API]: Starting application")

	s := QServer{}

	s.preload()
	defer s.Close()
	s.routes()

	log.Println("[Go Qlueless Assembly API]: Starting server")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
