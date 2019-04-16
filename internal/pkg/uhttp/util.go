package uhttp

import (
	"fmt"
	"log"
	"net/http"

	w "github.com/PaulioRandall/go-qlueless-assembly-api/internal/pkg/wrapped"
)

// LogRequest logs the details of a request such as the URL.
func LogRequest(req *http.Request) {
	log.Printf("(%s) %s\n", req.Method, req.URL.String())
}

// RelURL creates the absolute path of a requests URL without any fragment.
func RelURL(req *http.Request) string {
	r := req.URL.Path
	if req.URL.RawQuery != "" {
		r += "?" + req.URL.RawQuery
	}
	return r
}

// CheckNotEmpty validates that a required string is not empty writing a generic
// 500 response if it is.
func CheckNotEmpty(res *http.ResponseWriter, req *http.Request, name string, m string) bool {
	if m == "" {
		log.Println("[BUG] Error missing '" + name + "'")
		WriteServerError(res, req)
		return false
	}
	return true
}

// PrepResponseData prepares the response data read writing to the client. If
// the client has specified, the data is wrapped and meta information added else
// the input data is returned.
func PrepResponseData(req *http.Request, data interface{}, msg string) interface{} {
	if req.URL.Query()["wrap"] != nil {
		return w.WrappedReply{
			Message: msg,
			Self:    req.URL.String(),
			Data:    data,
		}
	} else {
		return data
	}
}

// AppendCORSHeaders appends the standard CORS response headers.
func AppendCORSHeaders(res *http.ResponseWriter, httpMethods string) {
	(*res).Header().Set("Access-Control-Allow-Origin", "*")
	(*res).Header().Set("Access-Control-Allow-Methods", httpMethods)
	(*res).Header().Set("Access-Control-Allow-Headers", "*")
}

// AppendJSONHeader appends the response headers for JSON requests.
func AppendJSONHeader(res *http.ResponseWriter, extensionType string) {
	var ct string
	if extensionType != "" {
		ct = fmt.Sprintf("application/%s+json; charset=utf-8", extensionType)
	} else {
		ct = "application/json; charset=utf-8"
	}
	(*res).Header().Set("Content-Type", ct)
}
