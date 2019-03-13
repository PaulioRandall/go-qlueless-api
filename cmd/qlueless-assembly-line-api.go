//usr/bin/env go run "$0" "$@"; exit "$?"

// Web server exposing access to a manufacturing themed TODO list database
package main

import (
	"log"
	"net/http"

	bat "github.com/PaulioRandall/qlueless-assembly-line-api/internal/app/batches"
	dict "github.com/PaulioRandall/qlueless-assembly-line-api/internal/app/dictionaries"
	oai "github.com/PaulioRandall/qlueless-assembly-line-api/internal/app/openapi"
	ord "github.com/PaulioRandall/qlueless-assembly-line-api/internal/app/orders"
)

// Main is the entry point for the web server
func main() {
	log.Println("[Qlueless Assembly Line API]: Starting application")
	http.HandleFunc("/openapi", oai.OpenAPI_handler)
	http.HandleFunc("/dictionaries", dict.All_dictionaries_handler)
	http.HandleFunc("/orders", ord.All_orders_handler)
	http.HandleFunc("/batches", bat.All_batches_handler)

	log.Println("[Qlueless Assembly Line API]: Starting server")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
