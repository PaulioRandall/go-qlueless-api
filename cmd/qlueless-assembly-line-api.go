//usr/bin/env go run "$0" "$@"; exit "$?"

// Web server exposing access to a manufacturing themed TODO list database
package main

import (
	"log"
	"net/http"

	res "github.com/PaulioRandall/qlueless-assembly-line-api/internal/app"
)

// Main is the entry point for the web server
func main() {
	log.Println("[Qlueless Assembly Line API]: Starting application")
	http.HandleFunc("/openapi", res.OpenAPIHandler)
	http.HandleFunc("/dictionaries", res.DictionaryHandler)
	http.HandleFunc("/orders", res.OrderHandler)
	http.HandleFunc("/batches", res.BatchHandler)

	log.Println("[Qlueless Assembly Line API]: Starting server")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
