// Package internal/pkg contains reusable internal application code
package pkg

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

// A Reply represents the top level JSON returned by all endpoints
type Reply struct {
	Req     *http.Request        `json:"-"`
	Res     *http.ResponseWriter `json:"-"`
	Message *string              `json:"message,omitempty"`
	Self    *string              `json:"self,omitempty"`
	Data    interface{}          `json:"data,omitempty"`
	Hints   *string              `json:"hints,omitempty"`
}

// A WorkItem represents and is a genralisation of orders and batches
type WorkItem struct {
	Description      string `json:"description"`
	WorkItemID       string `json:"work_item_id"`
	ParentWorkItemID string `json:"parent_work_item_id,omitempty"`
	TagID            string `json:"tag_id"`
	StatusID         string `json:"status_id"`
	Additional       string `json:"additional,omitempty"`
}

// Str returns a pointer to the passed string, useful for getting the address of
// a string in one line
func Str(s string) *string {
	return &s
}

// IsBlank returns true if the string is empty or only contains whitespace
func IsBlank(s string) bool {
	v := strings.TrimSpace(s)
	if v == "" {
		return true
	}
	return false
}

// ValueOrEmpty returns the value of the parameter or an empty string
func ValueOrEmpty(m map[string]interface{}, k string) string {
	v, ok := m[k].(string)
	if ok {
		return v
	}
	return ""
}

// LogRequest logs the details of a request such as the URL
func LogRequest(req *http.Request) {
	if req.URL.RawQuery == "" {
		log.Println(req.URL.Path)
	} else {
		log.Println(req.URL.String())
	}
}

// Check is a shorthand function for panic if err is not nil
func Check(err error) {
	if err != nil {
		panic(err)
	}
}

// Log_if_err checks if the input err is NOT nil returning true if it is.
// When true the error is logged so all the calling handler needs to do is
// clean up then invoke Http_500(*http.ResponseWriter) before returning
func LogIfErr(err error) bool {
	if err != nil {
		log.Println(err)
		return true
	}
	return false
}

// Http_500 sets up the response with generic 500 error details. This method
// should be used when ever a 500 error needs to be returned
func Http_500(r *Reply) {
	reply := Reply{
		Message: Str("Sorry, something went wrong at our end"),
	}

	AppendJSONHeaders(r)
	(*r.Res).WriteHeader(500)
	json.NewEncoder(*r.Res).Encode(reply)
}

// Http_4xx sets up the response as a 4xx error
func Http_4xx(r *Reply, status int) {
	if status < 400 && status > 499 {
		panic("Status code must be between 400 and 499")
	}

	if r.Message == nil {
		log.Println("[BUG] Reply.Message required for error messages")
		Http_500(r)
		return
	}

	reply := Reply{
		Message: r.Message,
		Self:    r.Self,
		Hints:   r.Hints,
	}

	AppendJSONHeaders(r)
	(*r.Res).WriteHeader(status)
	json.NewEncoder(*r.Res).Encode(reply)
}

// AppendJSONHeaders appends the response headers for JSON requests to
// ResponseWriters
func AppendJSONHeaders(r *Reply) {
	(*r.Res).Header().Set("Content-Type", "application/json; charset=utf-8")
	(*r.Res).Header().Set("Access-Control-Allow-Origin", "*")
	(*r.Res).Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	(*r.Res).Header().Set("Access-Control-Allow-Headers", "*")
}

// WriteJsonReply writes either the reply or the reply data using the
// ResponseWriter and appends the required JSON headers
func WriteJsonReply(r *Reply, message *string, data interface{}, hints *string) {

	AppendJSONHeaders(r)
	(*r.Res).WriteHeader(http.StatusOK)

	v := (*r.Req).URL.Query()["wrap"]
	if v == nil {
		json.NewEncoder(*r.Res).Encode(data)
		return
	}

	r.Message = message
	r.Self = Str(r.Req.URL.String())
	r.Data = data
	r.Hints = hints

	json.NewEncoder(*r.Res).Encode(r)
}
