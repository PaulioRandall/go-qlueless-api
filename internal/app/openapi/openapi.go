package openapi

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"sync"

	shr "github.com/PaulioRandall/qlueless-assembly-line-api/internal/pkg"
)

var spec map[string]interface{}
var specLoader sync.Once

// OpenAPIHandler handles requests for the services OpenAPI specification
func OpenAPIHandler(w http.ResponseWriter, r *http.Request) {
	shr.LogRequest(r)

	specLoader.Do(loadSpec)

	if spec == nil {
		shr.Http_500(&w)
		return
	}

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
	if shr.LogIfErr(err) {
		spec = nil
		return
	}

	err = json.Unmarshal(bytes, &spec)
	if shr.LogIfErr(err) {
		spec = nil
	}
}
