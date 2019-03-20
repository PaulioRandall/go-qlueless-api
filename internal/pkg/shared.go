// Package internal/pkg contains reusable internal application code
package pkg

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

// A ReplyMeta represents the response that should be returned when the
// client has requested data be wrapped and meta information included
type ReplyMeta struct {
	Message string      `json:"message"`
	Self    string      `json:"self"`
	Data    interface{} `json:"data,omitempty"`
	Hints   string      `json:"hints,omitempty"`
}

// A Reply4XX represents the response that is returned for a client '4xx' error
type Reply4XX struct {
	Res     *http.ResponseWriter `json:"-"`
	Req     *http.Request        `json:"-"`
	Message string               `json:"message"`
	Self    string               `json:"self"`
	Hints   string               `json:"hints,omitempty"`
}

// IsInt returns true if the string contains an integer
func IsInt(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

// Str returns a pointer to the passed string, useful for getting the address of
// a string in one line
func Str(s string) *string {
	return &s
}

// DeleteStr removes a string from a string array
func DeleteStr(s []string, i int) []string {
	l := len(s) - 1
	s[i] = s[l]
	s[l] = "" // Set to zero value
	return s[:l]
}

// IsBlank returns true if the string is empty or only contains whitespace
func IsBlank(s string) bool {
	v := strings.TrimSpace(s)
	if v == "" {
		return true
	}
	return false
}

// RelURL creates the absolute relative URL of the request without any fragment
func RelURL(req *http.Request) string {
	r := req.URL.Path
	if req.URL.RawQuery != "" {
		r += "?" + req.URL.RawQuery
	}
	return r
}

// LogRequest logs the details of a request such as the URL
func LogRequest(req *http.Request) {
	log.Println(req.URL.String())
}

// Check is a shorthand function for panic if err is not nil
func Check(err error) {
	if err != nil {
		panic(err)
	}
}

// LogIfErr checks if the input err is NOT nil returning true if it is.
// When true the error is logged so all the calling handler needs to do is
// clean up then invoke Http_500(*http.ResponseWriter) before returning
func LogIfErr(err error) bool {
	if err != nil {
		log.Println(err)
		return true
	}
	return false
}

// Reply500 sets up the response with generic 500 error details. This method
// should be used when ever a 500 error needs to be returned
func Write500Reply(res *http.ResponseWriter, req *http.Request) {
	r := ReplyMeta{
		Message: "Bummer! Something went wrong on the server.",
		Self:    (*req).URL.String(),
	}

	AppendJSONHeaders(res)
	(*res).WriteHeader(500)
	json.NewEncoder(*res).Encode(r)
}

// Reply4XX sets up the response as a 4XX error
func Write4XXReply(status int, r *Reply4XX) {
	if status < 400 || status > 499 {
		log.Println("[BUG] Status code must be between 400 and 499")
		Write500Reply(r.Res, r.Req)
		return
	}

	if (*r).Message == "" {
		log.Println("[BUG] 4xx response message is missing")
		Write500Reply(r.Res, r.Req)
		return
	}

	if (*r).Self == "" {
		(*r).Self = RelURL(r.Req)
	}

	AppendJSONHeaders((*r).Res)
	(*r.Res).WriteHeader(status)
	json.NewEncoder(*r.Res).Encode(r)
}

// IsMetaReply returns true if the response should be wrapped and meta
// information included
func IsMetaReply(req *http.Request) bool {
	v := req.URL.Query()["meta"]
	if v == nil {
		return false
	}
	return true
}

// PrepResponseData returns the response data after wrapping it up and adding
// meta information but only if the client has requested it be so. Else the
// input data is returned unchanged
func PrepResponseData(req *http.Request, data interface{}, msg string) interface{} {
	if IsMetaReply(req) {
		return ReplyMeta{
			Message: msg,
			Self:    req.URL.String(),
			Data:    data,
		}
	} else {
		return data
	}
}

// AppendJSONHeaders appends the response headers for JSON requests to
// ResponseWriters
func AppendJSONHeaders(res *http.ResponseWriter) {
	(*res).Header().Set("Content-Type", "application/json; charset=utf-8")
	(*res).Header().Set("Access-Control-Allow-Origin", "*")
	(*res).Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	(*res).Header().Set("Access-Control-Allow-Headers", "*")
}

// WriteReply appends the required JSON headers and then writes the response
// data
func WriteReply(res *http.ResponseWriter, req *http.Request, data interface{}) {
	AppendJSONHeaders(res)
	(*res).WriteHeader(http.StatusOK)
	json.NewEncoder(*res).Encode(data)
}

// WriteEmptyReply appends the required JSON headers and sets status as OK
func WriteEmptyReply(res *http.ResponseWriter) {
	AppendJSONHeaders(res)
	(*res).WriteHeader(http.StatusOK)
}

// methodNotAllowed handles cases where a HTTP method has been used but is not
// handled by this particular endpoint
func MethodNotAllowed(res *http.ResponseWriter, req *http.Request) {
	reply := Reply4XX{
		Res:     res,
		Req:     req,
		Message: fmt.Sprintf("Method not allowed for this endpoint (%s)", req.Method),
	}
	Write4XXReply(405, &reply)
}
