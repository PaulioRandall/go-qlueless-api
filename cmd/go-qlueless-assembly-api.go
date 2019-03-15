//usr/bin/env go run "$0" "$@"; exit "$?"

package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	bat "github.com/PaulioRandall/go-qlueless-assembly-api/internal/app/batches"
	dict "github.com/PaulioRandall/go-qlueless-assembly-api/internal/app/dictionaries"
	oai "github.com/PaulioRandall/go-qlueless-assembly-api/internal/app/openapi"
	ord "github.com/PaulioRandall/go-qlueless-assembly-api/internal/app/orders"
	shr "github.com/PaulioRandall/go-qlueless-assembly-api/internal/pkg"
)

var reply shr.Reply = shr.Reply{
	Message: "Resource not found",
}

// Not_found_handler handles requests for which no handler could be found
func QluelessNotFoundHandler(w http.ResponseWriter, r *http.Request) {
	shr.LogRequest(r)
	shr.AppendJSONHeaders(w)
	json.NewEncoder(w).Encode(reply)
}

// Main is the entry point for the web server
func main() {
	log.Println("[Go Qlueless Assembly API]: Starting application")

	gorilla := mux.NewRouter()

	gorilla.HandleFunc("/openapi", oai.OpenAPIHandler)
	gorilla.HandleFunc("/dictionaries", dict.AllDictsHandler)
	gorilla.HandleFunc("/orders", ord.AllOrdersHandler)
	gorilla.HandleFunc("/orders/{order_id}", ord.SingleOrderHandler)
	gorilla.HandleFunc("/batches", bat.AllBatchesHandler)
	gorilla.HandleFunc("/batches/{batch_id}", bat.SingleBatchHandler)
	gorilla.NotFoundHandler = http.HandlerFunc(QluelessNotFoundHandler)

	http.Handle("/", gorilla)

	log.Println("[Go Qlueless Assembly API]: Starting server")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
