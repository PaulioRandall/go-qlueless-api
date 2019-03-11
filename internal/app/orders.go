// Package internal/app contains non-reusable internal application code
package app

import (
	"encoding/json"
	"log"
	"net/http"

	shr "github.com/PaulioRandall/qlueless-assembly-line-api/internal/pkg"
)

// OrderHandler implements the Go web server Handler interface to return a full
// list of all orders held by the system
func OrderHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Host)

	response := shr.Reply{
		Message: "Found dummy orders",
		Data:    createDummyOrders(),
	}

	shr.AppendJSONHeaders(&w)
	json.NewEncoder(w).Encode(response)
}

// createDummyOrders returns an array of dummy orders
func createDummyOrders() []shr.WorkItem {
	return []shr.WorkItem{
		shr.WorkItem{
			Title:        "Outline the saga",
			Description:  "Create a rough outline of the new saga.",
			Work_item_id: 1,
			Tag_id:       "mid",
			Status_id:    "in_progress",
		},
	}
}
