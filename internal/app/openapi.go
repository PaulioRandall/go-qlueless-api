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

var spec map[string]interface{}
var once_spec sync.Once

// OpenAPIHandler handles requests for the services OpenAPI specification
func OpenAPIHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL.Path)

	once_spec.Do(loadSpec)

	shr.AppendJSONHeaders(&w)
	json.NewEncoder(w).Encode(spec)
}

// loadJson loads the dictionary response from a file
func loadSpec() {
	go_path := os.Getenv("GOPATH")
	path := go_path +
		"/src/github.com/PaulioRandall/qlueless-assembly-line-api" +
		"/api/openapi/openapi.json"
	bytes, err := ioutil.ReadFile(path)
	shr.Check(err)

	err = json.Unmarshal(bytes, &spec)
	shr.Check(err)
}
