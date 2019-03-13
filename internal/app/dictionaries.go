// Package internal/app contains non-reusable internal application code
package app

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sync"

	shr "github.com/PaulioRandall/qlueless-assembly-line-api/internal/pkg"
)

var data map[string]interface{}
var once_dict sync.Once

// DictionaryHandler handles requests for the service dictionaries
func DictionaryHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL.Path)

	once_dict.Do(loadJson)

	shr.AppendJSONHeaders(&w)
	json.NewEncoder(w).Encode(data)
}

// loadJson loads the dictionary response from a file
func loadJson() {
	go_path := os.Getenv("GOPATH")
	path := go_path +
		"/src/github.com/PaulioRandall/qlueless-assembly-line-api" +
		"/web/dictionaries.json"
	bytes, err := ioutil.ReadFile(path)
	shr.Check(err)

	err = json.Unmarshal(bytes, &data)
	shr.Check(err)
}
