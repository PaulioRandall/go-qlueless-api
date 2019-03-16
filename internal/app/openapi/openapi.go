package openapi

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"sync"

	. "github.com/PaulioRandall/go-qlueless-assembly-api/internal/pkg"
)

var spec map[string]interface{}
var specLoader sync.Once

// OpenAPIHandler handles requests for the services OpenAPI specification
func OpenAPIHandler(res http.ResponseWriter, req *http.Request) {
	LogRequest(req)
	r := Reply{
		Req: req,
		Res: &res,
	}

	specLoader.Do(loadSpec)

	if spec == nil {
		Http_500(&r)
		return
	}

	AppendJSONHeaders(&r)
	json.NewEncoder(res).Encode(spec)
}

// loadJson loads the dictionary response from a file
func loadSpec() {

	go_path := os.Getenv("GOPATH")
	path := go_path +
		"/src/github.com/PaulioRandall/go-qlueless-assembly-api" +
		"/api/openapi/openapi.json"

	bytes, err := ioutil.ReadFile(path)
	if LogIfErr(err) {
		spec = nil
		return
	}

	err = json.Unmarshal(bytes, &spec)
	if LogIfErr(err) {
		spec = nil
	}
}
