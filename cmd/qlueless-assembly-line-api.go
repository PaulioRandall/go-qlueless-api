//usr/bin/env go run "$0" "$@"; exit "$?"

package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	bat "github.com/PaulioRandall/qlueless-assembly-line-api/internal/app/batches"
	dict "github.com/PaulioRandall/qlueless-assembly-line-api/internal/app/dictionaries"
	oai "github.com/PaulioRandall/qlueless-assembly-line-api/internal/app/openapi"
	ord "github.com/PaulioRandall/qlueless-assembly-line-api/internal/app/orders"
	shr "github.com/PaulioRandall/qlueless-assembly-line-api/internal/pkg"
)

var reply shr.Reply = shr.Reply{
	Message: "Resource not found",
}

// Not_found_handler handles requests for which no handler could be found
func Not_found_handler(w http.ResponseWriter, r *http.Request) {
	shr.Log_request(r)
	shr.AppendJSONHeaders(&w)
	json.NewEncoder(w).Encode(reply)
}

// Main is the entry point for the web server
func main() {
	log.Println("[Qlueless Assembly Line API]: Starting application")

	gorilla := mux.NewRouter()

	gorilla.HandleFunc("/openapi", oai.OpenAPI_handler)
	gorilla.HandleFunc("/dictionaries", dict.All_dictionaries_handler)
	gorilla.HandleFunc("/orders", ord.All_orders_handler)
	gorilla.HandleFunc("/orders/{order_id}", ord.Single_order_handler)
	gorilla.HandleFunc("/batches", bat.All_batches_handler)
	gorilla.HandleFunc("/batches/{batch_id}", bat.Single_batch_handler)
	gorilla.NotFoundHandler = http.HandlerFunc(Not_found_handler)

	http.Handle("/", gorilla)

	log.Println("[Qlueless Assembly Line API]: Starting server")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
