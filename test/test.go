//usr/bin/env go run "$0" "$@"; exit "$?"

package main

import (
	"encoding/json"
	"log"
)

type Thing struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Agility string `json:"agility"`
}

func main() {
	b := []byte(`{
		"name": "Oliver",
		"age": "24",
		"strength": 292
	}`)

	var t Thing
	json.Unmarshal(b, &t)

	log.Println(t)
}
