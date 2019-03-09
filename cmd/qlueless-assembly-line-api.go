package main

import (
  "encoding/json"
  "log"
  "net/http"
)

type Response struct {
  F string `json:"f"`
  S  string `json:"s"`
}

func addJsonHeaders(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}

func handler(w http.ResponseWriter, r *http.Request) {

  data := Response {
    F: "Johnny",
    S:  "Bravo",
  }

  addJsonHeaders(w)

  json.NewEncoder(w).Encode(data)
  log.Println(r.Host)
}

func main() {
  http.HandleFunc("/", handler)
  log.Fatal(http.ListenAndServe(":8080", nil))
}
