package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Response struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func addJsonHeaders(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}

func handler(w http.ResponseWriter, r *http.Request) {

	data := Response{
		FirstName: "Johnny",
		LastName:  "Bravo",
	}

	addJsonHeaders(w)

	json.NewEncoder(w).Encode(data)
	log.Println(r.Host)
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
