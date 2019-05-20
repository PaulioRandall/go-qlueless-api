package openapi

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"sync"

	cookies "github.com/PaulioRandall/go-cookies/cookies"
	uhttp "github.com/PaulioRandall/go-cookies/uhttp"
	writers "github.com/PaulioRandall/go-qlueless-api/shared/writers"
)

var spec map[string]interface{} = nil
var once sync.Once

var cors uhttp.CorsHeaders = uhttp.CorsHeaders{
	Origin:  "*",
	Headers: "*",
	Methods: "GET, OPTIONS",
}

// OpenAPIHandler handles requests for the services OpenAPI specification
func OpenAPIHandler(res http.ResponseWriter, req *http.Request) {
	uhttp.LogRequest(req)
	uhttp.UseCors(&res, &cors)

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
	once.Do(load)

	if spec == nil {
		log.Println("[BUG] OpenAPI specification not loaded")
		writers.WriteServerError(res, req)
		return
	}

	uhttp.UseUTF8Json(res, "vnd.oai.openapi")
	(*res).WriteHeader(http.StatusOK)
	json.NewEncoder(*res).Encode(spec)
}

// load loads the OpenAPI specification from a file
func load() {
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
