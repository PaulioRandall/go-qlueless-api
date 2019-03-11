// Package internal/app contains non-reusable internal application code
package app

import (
	"encoding/json"
	"log"
	"net/http"

	shr "github.com/PaulioRandall/qlueless-assembly-line-api/internal/pkg"
)

// BatchHandler implements the Go web server Handler interface to return a full
// list of all batches held by the system
func BatchHandler(w http.ResponseWriter, r *http.Request) {
	response := shr.Reply{
		Message: "Found dummy batches",
		Data:    createDummyBatches(),
	}

	shr.AppendJSONHeaders(w)
	json.NewEncoder(w).Encode(response)
	log.Println(r.Host)
}

// createDummyBatches returns an array of dummy batches
func createDummyBatches() []shr.WorkItem {
	return []shr.WorkItem{
		shr.WorkItem{
			Title:               "Name the saga",
			Description:         "Think of a name for the saga.",
			Work_item_id:        2,
			Parent_work_item_id: 1,
			Tag_id:              "mid",
			Status_id:           "potential",
		},
		shr.WorkItem{
			Title:               "Outline the first chapter",
			Description:         "Outline the first chapter.",
			Work_item_id:        3,
			Parent_work_item_id: 1,
			Tag_id:              "mid",
			Status_id:           "delivered",
			Additional:          "archive_note:Done but not a compelling start",
		},
		shr.WorkItem{
			Title:               "Outline the second chapter",
			Description:         "Outline the second chapter.",
			Work_item_id:        4,
			Parent_work_item_id: 1,
			Tag_id:              "mid",
			Status_id:           "in_progress",
		},
	}
}
