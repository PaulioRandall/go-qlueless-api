// Package internal/pkg contains reusable internal application code
package pkg

import "net/http"

// A Reply represents the top level JSON returned by all endpoints
type Reply struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
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

// AppendJSONHeaders appends the response headers for JSON requests to
// ResponseWriters
func AppendJSONHeaders(w *http.ResponseWriter) {
	(*w).Header().Set("Content-Type", "application/json")
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}
