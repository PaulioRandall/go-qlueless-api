package pkg

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

var Things ThingStore = NewThingStore()

// A WrappedReply represents the response that should be returned when the
// client has requested data be wrapped and meta information included
type WrappedReply struct {
	Message string      `json:"message"`
	Self    string      `json:"self"`
	Data    interface{} `json:"data,omitempty"`
	Hints   string      `json:"hints,omitempty"`
}

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

// Reply500 sets up the response with generic 500 error details. This method
// should be used when ever a 500 error needs to be returned
func Write500Reply(res *http.ResponseWriter, req *http.Request) {
	r := WrappedReply{
		Message: "Bummer! Something went wrong on the server.",
		Self:    (*req).URL.String(),
	}

	AppendJSONHeader(res, "")
	(*res).WriteHeader(500)
	json.NewEncoder(*res).Encode(r)
}

// CheckStatusBetween validates that the status code is between the max an min
func CheckStatusBetween(res *http.ResponseWriter, req *http.Request, status int, minInc int, maxInc int) bool {
	if status < minInc || status > maxInc {
		log.Printf("[BUG] Status code must be between %d and %d\n", minInc, maxInc)
		Write500Reply(res, req)
		return false
	}
	return true
}

// CheckReplyMetaMessage validates that the ReplyMeta.Message is not empty
func CheckReplyMetaMessage(res *http.ResponseWriter, req *http.Request, r WrappedReply) bool {
	if r.Message == "" {
		log.Println("[BUG] error response message is missing")
		Write500Reply(res, req)
		return false
	}
	return true
}

// Write4XXReply writes the response for a 4XX error
func Write4XXReply(res *http.ResponseWriter, req *http.Request, status int, r WrappedReply) {
	if !CheckStatusBetween(res, req, status, 400, 499) {
		return
	}

	if !CheckReplyMetaMessage(res, req, r) {
		return
	}

	if r.Self == "" {
		r.Self = RelURL(req)
	}

	AppendJSONHeader(res, "")
	(*res).WriteHeader(status)
	json.NewEncoder(*res).Encode(r)
}

// WrapReply returns true if the response should be wrapped and meta
// information included
func WrapReply(req *http.Request) bool {
	v := req.URL.Query()["wrap"]
	if v == nil {
		return false
	}
	return true
}

// PrepResponseData returns the response data after wrapping it up and adding
// meta information but only if the client has requested it be so. Else the
// input data is returned unchanged
func PrepResponseData(req *http.Request, data interface{}, msg string) interface{} {
	if WrapReply(req) {
		return WrappedReply{
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

// WriteReply appends the required headers and then writes the response data
func WriteReply(res *http.ResponseWriter, data *[]byte, contentType string) {
	(*res).Header().Set("Content-Type", contentType)
	(*res).WriteHeader(http.StatusOK)
	(*res).Write(*data)
}

// WriteEmptyReply appends the required headers without writing any data
func WriteEmptyReply(res *http.ResponseWriter, contentType string) {
	(*res).Header().Set("Content-Type", contentType)
	(*res).WriteHeader(http.StatusOK)
}

// WriteJSONReply appends the required JSON headers and then writes the response
// data
func WriteJSONReply(res *http.ResponseWriter, req *http.Request, data interface{}, extensionType string) {
	AppendJSONHeader(res, extensionType)
	(*res).WriteHeader(http.StatusOK)
	json.NewEncoder(*res).Encode(data)
}

// WriteEmptyJSONReply appends the required JSON headers and sets status as OK
func WriteEmptyJSONReply(res *http.ResponseWriter, extensionType string) {
	AppendJSONHeader(res, extensionType)
	(*res).WriteHeader(http.StatusOK)
}

// methodNotAllowed handles cases where a HTTP method has been used but is not
// handled by this particular endpoint
func MethodNotAllowed(res *http.ResponseWriter, req *http.Request) {
	r := WrappedReply{
		Message: fmt.Sprintf("Method not allowed for this endpoint (%s)", req.Method),
	}
	Write4XXReply(res, req, 405, r)
}
