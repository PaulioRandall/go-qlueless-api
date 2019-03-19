//usr/bin/env go run "$0" "$@"; exit "$?"

package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	oai "github.com/PaulioRandall/go-qlueless-assembly-api/internal/app/openapi"
	thg "github.com/PaulioRandall/go-qlueless-assembly-api/internal/app/things"
	. "github.com/PaulioRandall/go-qlueless-assembly-api/internal/pkg"
)

// Not_found_handler handles requests for which no handler could be found
func QluelessNotFoundHandler(res http.ResponseWriter, req *http.Request) {
	LogRequest(req)

	r := Reply4XX{
		Res:     &res,
		Req:     req,
		Message: "Resource not found",
	}

	Write4XXReply(404, &r)
}

// Main is the entry point for the web server
func main() {
	log.Println("[Go Qlueless Assembly API]: Starting application")

	oai.LoadSpec()
	thg.CreateDummyThings()

	gorilla := mux.NewRouter()

	gorilla.HandleFunc("/openapi", oai.OpenAPIHandler).Methods("GET")
	gorilla.HandleFunc("/things", thg.AllThingsHandler).Methods("GET")
	gorilla.HandleFunc("/thing", thg.NewThingHandler).Methods("POST", "OPTIONS")
	gorilla.HandleFunc("/things/{id}", thg.SingleThingHandler).Methods("GET")

	gorilla.NotFoundHandler = http.HandlerFunc(QluelessNotFoundHandler)
	http.Handle("/", gorilla)

	log.Println("[Go Qlueless Assembly API]: Starting server")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
