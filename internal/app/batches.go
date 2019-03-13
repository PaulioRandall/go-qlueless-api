// Package internal/app contains non-reusable internal application code
package app

import (
	"net/http"

	shr "github.com/PaulioRandall/qlueless-assembly-line-api/internal/pkg"
)

var batches []shr.WorkItem

// BatchHandler handles requests for all batches currently within the service
func BatchHandler(w http.ResponseWriter, r *http.Request) {
	shr.Log_request(r)

	shr.Loader.Do(loadBatches)
	if batches == nil {
		shr.Http_500(&w)
		return
	}

	reply := shr.Reply{
		Message: "Found all batches",
		Data:    batches,
	}

	shr.WriteJsonReply(reply, w, r)
}

// loadBatches loads all batches into the batches array
func loadBatches() {
	batches = []shr.WorkItem{
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
