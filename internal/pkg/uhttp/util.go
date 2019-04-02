package uhttp

import (
	"fmt"
	"log"
	"net/http"

	w "github.com/PaulioRandall/go-qlueless-assembly-api/internal/pkg/wrapped"
)

// LogRequest logs the details of a request such as the URL
func LogRequest(req *http.Request) {
	log.Printf("(%s) %s\n", req.Method, req.URL.String())
}

// RelURL creates the absolute relative URL of the request without any fragment
func RelURL(req *http.Request) string {
	r := req.URL.Path
	if req.URL.RawQuery != "" {
		r += "?" + req.URL.RawQuery
	}
	return r
}

// CheckStatusBetween validates that the status code is between the max an min
func CheckStatusBetween(res *http.ResponseWriter, req *http.Request, status int, minInc int, maxInc int) bool {
	if status < minInc || status > maxInc {
		log.Printf("[BUG] Status code must be between %d and %d\n", minInc, maxInc)
		WriteServerError(res, req)
		return false
	}
	return true
}

// CheckNotEmpty validates that a require string is not empty
func CheckNotEmpty(res *http.ResponseWriter, req *http.Request, name string, m string) bool {
	if m == "" {
		log.Println("[BUG] error missing '" + name + "'")
		WriteServerError(res, req)
		return false
	}
	return true
}

// PrepResponseData returns the response data after wrapping it up and adding
// meta information but only if the client has requested it be so. Else the
// input data is returned unchanged
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

// AppendCORSHeaders appends the CORS response headers for requests to
// ResponseWriters
func AppendCORSHeaders(res *http.ResponseWriter, httpMethods string) {
	(*res).Header().Set("Access-Control-Allow-Origin", "*")
	(*res).Header().Set("Access-Control-Allow-Methods", httpMethods)
	(*res).Header().Set("Access-Control-Allow-Headers", "*")
}

// AppendJSONHeader appends the response headers for JSON requests to
// ResponseWriters
func AppendJSONHeader(res *http.ResponseWriter, extensionType string) {
	var ct string
	if extensionType != "" {
		ct = fmt.Sprintf("application/%s+json; charset=utf-8", extensionType)
	} else {
		ct = "application/json; charset=utf-8"
	}
	(*res).Header().Set("Content-Type", ct)
}
