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
	. "github.com/PaulioRandall/go-qlueless-assembly-api/internal/pkg"
)

// Not_found_handler handles requests for which no handler could be found
func QluelessNotFoundHandler(res http.ResponseWriter, req *http.Request) {
	LogRequest(req)

	r := Reply{
		Req:     req,
		Res:     &res,
		Message: Str("Resource not found"),
	}

	AppendJSONHeaders(&r)
	json.NewEncoder(res).Encode(r)
}

// Main is the entry point for the web server
func main() {
	log.Println("[Go Qlueless Assembly API]: Starting application")

	ord.CreateDummyOrders()
	bat.CreateDummyBatches()
	dict.LoadDicts()

	gorilla := mux.NewRouter()

	gorilla.HandleFunc("/openapi", oai.OpenAPIHandler).Methods("GET")
	gorilla.HandleFunc("/dictionaries", dict.AllDictsHandler).Methods("GET")
	gorilla.HandleFunc("/orders", ord.AllOrdersHandler).Methods("GET")
	gorilla.HandleFunc("/order", ord.NewOrderHandler).Methods("POST", "OPTIONS")
	gorilla.HandleFunc("/orders/{order_id}", ord.SingleOrderHandler).Methods("GET")
	gorilla.HandleFunc("/batches", bat.AllBatchesHandler).Methods("GET")
	gorilla.HandleFunc("/batches/{batch_id}", bat.SingleBatchHandler).Methods("GET")

	gorilla.NotFoundHandler = http.HandlerFunc(QluelessNotFoundHandler)
	http.Handle("/", gorilla)

	log.Println("[Go Qlueless Assembly API]: Starting server")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
