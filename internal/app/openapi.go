// Package internal/app contains non-reusable internal application code
package app

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	shr "github.com/PaulioRandall/qlueless-assembly-line-api/internal/pkg"
)

// OpenAPIHandler implements the Go web server Handler interface to return the
// services OpenAPI specification
func OpenAPIHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Host)

	var err error

	go_path := os.Getenv("GOPATH")
	path := go_path +
		"/src/github.com/PaulioRandall/qlueless-assembly-line-api" +
		"/api/openapi/openapi.json"
	bytes, err := ioutil.ReadFile(path)
	shr.Check(err)

	var spec map[string]interface{}
	err = json.Unmarshal(bytes, &spec)
	shr.Check(err)

	shr.AppendJSONHeaders(&w)
	json.NewEncoder(w).Encode(spec)
}
