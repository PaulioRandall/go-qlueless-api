package openapi

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	p "github.com/PaulioRandall/go-qlueless-assembly-api/internal/pkg"
)

var spec map[string]interface{} = nil

// OpenAPIHandler handles requests for the services OpenAPI specification
func OpenAPIHandler(res http.ResponseWriter, req *http.Request) {
	p.LogRequest(req)
	p.AppendCORSHeaders(&res, "GET, HEAD, OPTIONS")

	switch req.Method {
	case "GET":
		_GET_Spec(&res, req)
	case "HEAD":
		fallthrough
	case "OPTIONS":
		p.AppendJSONHeader(&res, "vnd.oai.openapi")
		res.WriteHeader(http.StatusOK)
	default:
		p.MethodNotAllowed(&res, req)
	}
}

// _GET_Spec generates responses for obtaining the OpenAPI specification
func _GET_Spec(res *http.ResponseWriter, req *http.Request) {
	if spec == nil {
		log.Println("[BUG] OpenAPI specification not loaded")
		p.WriteServerError(res, req)
		return
	}

	p.AppendJSONHeader(res, "vnd.oai.openapi")
	(*res).WriteHeader(http.StatusOK)
	json.NewEncoder(*res).Encode(spec)
}

// LoadJson loads the OpenAPI specification from a file
func LoadSpec() {

	path := "./openapi.json"
	bytes, err := ioutil.ReadFile(path)
	if p.LogIfErr(err) {
		spec = nil
		return
	}

	err = json.Unmarshal(bytes, &spec)
	if p.LogIfErr(err) {
		spec = nil
	}
}
