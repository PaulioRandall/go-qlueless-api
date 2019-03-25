package openapi

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	. "github.com/PaulioRandall/go-qlueless-assembly-api/internal/pkg"
)

var spec map[string]interface{} = nil

// OpenAPIHandler handles requests for the services OpenAPI specification
func OpenAPIHandler(res http.ResponseWriter, req *http.Request) {
	LogRequest(req)

	switch req.Method {
	case "GET":
		get_Spec(&res, req)
	case "HEAD":
		WriteEmptyJSONReply(&res, "vnd.oai.openapi")
	case "OPTIONS":
		WriteEmptyJSONReply(&res, "vnd.oai.openapi")
	default:
		MethodNotAllowed(&res, req)
	}
}

// get_Spec generates responses for obtaining the OpenAPI specification
func get_Spec(res *http.ResponseWriter, req *http.Request) {
	if spec == nil {
		log.Println("[BUG] OpenAPI specification not loaded")
		Write500Reply(res, req)
		return
	}

	WriteJSONReply(res, req, spec, "vnd.oai.openapi")
}

// LoadJson loads the OpenAPI specification from a file
func LoadSpec() {

	path := "../api/openapi/openapi.json"
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
