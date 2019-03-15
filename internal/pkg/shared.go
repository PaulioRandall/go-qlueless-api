// Package internal/pkg contains reusable internal application code
package pkg

import (
	"encoding/json"
	"log"
	"net/http"
)

// A Reply represents the top level JSON returned by all endpoints
type Reply struct {
	Message string      `json:"message"`
	Self    string      `json:"self,omitempty"`
	Data    interface{} `json:"data,omitempty"`
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

// Log_request logs the details of a request such as the URL
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
func Http_500(w *http.ResponseWriter) {
	r := Reply{
		Message: "Internal server error",
	}

	AppendJSONHeaders(w)
	json.NewEncoder(*w).Encode(r)
}

// Http_4xx sets up the response as a 4xx error
func Http_4xx(w *http.ResponseWriter, status int, message string) {
	if status < 400 && status > 499 {
		panic("Status code must be between 400 and 499")
	}

	r := Reply{
		Message: message,
	}

	(*w).WriteHeader(status)
	json.NewEncoder(*w).Encode(r)
	AppendJSONHeaders(w)
}

// WrapData returns true if the client has specified that the response data
// should be wrapped so meta information about the response can be included
func WrapData(r *http.Request) bool {
	wrap := r.URL.Query()["wrap"]
	if wrap != nil {
		return true
	}
	return false
}

// AppendJSONHeaders appends the response headers for JSON requests to
// ResponseWriters
func AppendJSONHeaders(w *http.ResponseWriter) {
	(*w).Header().Set("Content-Type", "application/json")
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

// WriteJsonReply writes either the reply or the reply data using the
// ResponseWriter and appends the required JSON headers
func WriteJsonReply(reply Reply, w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	AppendJSONHeaders(&w)

	if WrapData(r) {
		reply.Self = r.URL.String()
		json.NewEncoder(w).Encode(reply)
	} else {
		json.NewEncoder(w).Encode(reply.Data)
	}
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
