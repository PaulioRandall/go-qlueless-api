package openapi

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"

	. "github.com/PaulioRandall/go-qlueless-assembly-api/internal/pkg"
)

var spec map[string]interface{} = nil

// OpenAPIHandler handles requests for the services OpenAPI specification
func OpenAPIHandler(res http.ResponseWriter, req *http.Request) {
	LogRequest(req)

	if spec == nil {
		Write500Reply(&res, req)
		return
	}

	WriteReply(&res, req, spec)
}

// LoadJson loads the OpenAPI specification from a file
func LoadSpec() {

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
