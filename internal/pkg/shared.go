// Package internal/pkg contains reusable internal application code
package pkg

import (
	"encoding/json"
	"errors"
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
	Title            string `json:"title"`
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

// wrapData returns a nil if the client has not requested a wrapped response.
// If they have, a list of the requested meta information properties will be
// returned
func wrapData(r *Reply) ([]string, error) {
	v := (*r.Req).URL.Query()["wrap_with"]
	if v != nil {
		if len(v) > 1 {
			return nil, errors.New("Multiple 'wrap_with' query parameters not allowed")
		}
		return wrapWith(v[0])
	}
	return nil, nil
}

// wrapWith splits a dot `.` delimitered string of meta information properties
// and validates each one
func wrapWith(p string) ([]string, error) {
	v := strings.Split(p, ".")
	if len(v) == 0 {
		return nil, errors.New("At least one value within 'wrap_with' must be provided")
	}
	for _, p := range v {
		if IsBlank(p) {
			return nil, errors.New("Blank values within 'wrap_with' are not allowed")
		}
		if !isWrapperProp(p) {
			return nil, errors.New("Invalid value within 'wrap_with': " + p)
		}
	}
	return v, nil
}

// isWrapperProp returns true if the input string is a meta information property
// name
func isWrapperProp(p string) bool {
	p = strings.ToLower(p)
	switch p {
	case "message":
		fallthrough
	case "self":
		fallthrough
	case "data":
		fallthrough
	case "hints":
		return true
	}
	return false
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
	v, err := wrapData(r)
	if err != nil {
		r.Hints = Str("Valid values [message|self|data|hints] e.g. 'wrap_with=message.data'")
		r.Message = Str(err.Error())
		Http_4xx(r, 400)
		return
	}

	AppendJSONHeaders(r)
	(*r.Res).WriteHeader(http.StatusOK)

	if v == nil {
		json.NewEncoder((*r.Res)).Encode(data)
		return
	}

	for _, p := range v {
		switch p {
		case "message":
			r.Message = message
		case "self":
			r.Self = Str(r.Req.URL.String())
		case "data":
			r.Data = data
		case "hints":
			r.Hints = hints
		}
	}

	json.NewEncoder(*r.Res).Encode(r)
}

// FindWorkItem finds the WorkItem with the specified work_item_id else returns
// nil
func FindWorkItem(items *[]WorkItem, id string) *WorkItem {
	for _, v := range *items {
		if id == v.WorkItemID {
			return &v
		}
	}
	return nil
}
