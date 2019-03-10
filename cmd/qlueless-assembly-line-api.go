///bin/true; exec /usr/bin/env go run "$0" "$@"
package main

import (
	"log"
	"net/http"

	res "github.com/PaulioRandall/qlueless-assembly-line-api/internal/app"
)

func main() {
	http.HandleFunc("/dictionaries", res.DictionaryHandler)
	http.HandleFunc("/orders", res.OrderHandler)
	http.HandleFunc("/batches", res.BatchHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
