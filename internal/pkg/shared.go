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
	Data    interface{} `json:"data,omitempty"`
}

// A WorkItem represents and is a genralisation of orders and batches
type WorkItem struct {
	Title               string `json:"title"`
	Description         string `json:"description"`
	Work_item_id        int    `json:"work_item_id"`
	Parent_work_item_id int    `json:"parent_work_item_id,omitempty"`
	Tag_id              string `json:"tag_id"`
	Status_id           string `json:"status_id"`
	Additional          string `json:"additional,omitempty"`
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
func Log_if_err(err error) bool {
	if err != nil {
		log.Fatal(err)
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

	json.NewEncoder(*w).Encode(r)
	AppendJSONHeaders(w)
}

// AppendJSONHeaders appends the response headers for JSON requests to
// ResponseWriters
func AppendJSONHeaders(w *http.ResponseWriter) {
	(*w).Header().Set("Content-Type", "application/json")
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}
