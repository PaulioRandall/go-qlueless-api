//usr/bin/go run $0 $@ ; exit
package main

import (
	"log"
	"net/http"

	res "github.com/PaulioRandall/qlueless-assembly-line-api/internal/app"
)

func main() {
	http.HandleFunc("/dictionaries", res.DictionaryHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
