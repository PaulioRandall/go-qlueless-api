package openapi

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	cookies "github.com/PaulioRandall/go-cookies/cookies"
	h "github.com/PaulioRandall/go-qlueless-api/internal/uhttp"
)

var spec map[string]interface{} = nil

// OpenAPIHandler handles requests for the services OpenAPI specification
func OpenAPIHandler(res http.ResponseWriter, req *http.Request) {
	h.LogRequest(req)
	h.AppendCORSHeaders(&res, "GET, OPTIONS")

	switch req.Method {
	case "GET":
		get(&res, req)
	case "OPTIONS":
		res.WriteHeader(http.StatusOK)
	default:
		res.WriteHeader(http.StatusMethodNotAllowed)
	}
}

// get generates responses for obtaining the OpenAPI specification
func get(res *http.ResponseWriter, req *http.Request) {
	if spec == nil {
		log.Println("[BUG] OpenAPI specification not loaded")
		h.WriteServerError(res, req)
		return
	}

	h.AppendJSONHeader(res, "vnd.oai.openapi")
	(*res).WriteHeader(http.StatusOK)
	json.NewEncoder(*res).Encode(spec)
}

// LoadJson loads the OpenAPI specification from a file
func LoadSpec() {

	path := "./openapi.json"
	bytes, err := ioutil.ReadFile(path)
	if cookies.LogIfErr(err) {
		spec = nil
		return
	}

	err = json.Unmarshal(bytes, &spec)
	if cookies.LogIfErr(err) {
		spec = nil
	}
}
