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
	AppendCORSHeaders(&res, "GET, HEAD, OPTIONS")

	switch req.Method {
	case "GET":
		get_Spec(&res, req)
	case "HEAD":
		fallthrough
	case "OPTIONS":
		AppendJSONHeader(&res, "vnd.oai.openapi")
		res.WriteHeader(http.StatusOK)
	default:
		MethodNotAllowed(&res, req)
	}
}

// get_Spec generates responses for obtaining the OpenAPI specification
func get_Spec(res *http.ResponseWriter, req *http.Request) {
	if spec == nil {
		log.Println("[BUG] OpenAPI specification not loaded")
		WriteServerError(res, req)
		return
	}

	AppendJSONHeader(res, "vnd.oai.openapi")
	(*res).WriteHeader(http.StatusOK)
	json.NewEncoder(*res).Encode(spec)
}

// LoadJson loads the OpenAPI specification from a file
func LoadSpec() {

	path := "./openapi.json"
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
