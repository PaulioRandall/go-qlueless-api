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
// OUTDATED
type Reply struct {
	Request  *http.Request  `json:"-"`
	Response *http.Response `json:"-"`
	Message  string         `json:"message,omitempty"`
	Self     string         `json:"self,omitempty"`
	Data     interface{}    `json:"data,omitempty"`
	Hints    string         `json:"hints,omitempty"`
}

// A WorkItem represents and is a genralisation of orders and batches
type WorkItem struct {
	Title               string `json:"title"`
	Description         string `json:"description"`
	Work_item_id        string `json:"work_item_id"`
	Parent_work_item_id string `json:"parent_work_item_id,omitempty"`
	Tag_id              string `json:"tag_id"`
	Status_id           string `json:"status_id"`
	Additional          string `json:"additional,omitempty"`
}

// IsBlank returns true if the string is empty or only contains whitespace
func IsBlank(s string) bool {
	v := strings.TrimSpace(s)
	if v == "" {
		return true
	}
	return false
}

// LogRequest logs the details of a request such as the URL
func LogRequest(r *http.Request) {
	if r.URL.RawQuery == "" {
		log.Println(r.URL.Path)
	} else {
		log.Println(r.URL.Path + "?" + r.URL.RawQuery)
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
func Http_500(w http.ResponseWriter) {
	r := Reply{
		Message: "Sorry, something went wrong on our end",
	}

	AppendJSONHeaders(w)
	w.WriteHeader(500)
	json.NewEncoder(w).Encode(r)
}

// Http_4xx sets up the response as a 4xx error
func Http_4xx(w http.ResponseWriter, status int, message string, hints string) {
	if status < 400 && status > 499 {
		panic("Status code must be between 400 and 499")
	}

	r := Reply{
		Message: message,
		Hints:   hints,
	}

	AppendJSONHeaders(w)
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(r)
}

// wrapData returns a nil if the client has not requested a wrapped response.
// If they have, a list of the requested meta information properties will be
// returned
func wrapData(r *http.Request) ([]string, error) {
	v := r.URL.Query()["wrap_with"]
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
func AppendJSONHeaders(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
}

// WriteJsonReply writes either the reply or the reply data using the
// ResponseWriter and appends the required JSON headers
func WriteJsonReply(message string, data interface{}, hints string, w http.ResponseWriter, r *http.Request) {
	v, err := wrapData(r)
	if err != nil {
		Http_4xx(w, 400, err.Error(), "Valid values [message|self|data|hints] e.g. 'wrap_with=message.data'")
		return
	}

	AppendJSONHeaders(w)
	w.WriteHeader(http.StatusOK)

	if v == nil {
		json.NewEncoder(w).Encode(data)
		return
	}

	reply := Reply{}
	for _, p := range v {
		switch p {
		case "message":
			reply.Message = message
		case "self":
			reply.Self = r.URL.String()
		case "data":
			reply.Data = data
		case "hints":
			reply.Hints = hints
		}
	}

	json.NewEncoder(w).Encode(reply)
}

// FindWorkItem finds the WorkItem with the specified work_item_id else returns
// nil
func FindWorkItem(items []WorkItem, id string) *WorkItem {
	for _, v := range items {
		if id == v.Work_item_id {
			return &v
		}
	}
	return nil
}
