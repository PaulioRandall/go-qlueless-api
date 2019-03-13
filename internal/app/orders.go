// Package internal/app contains non-reusable internal application code
package app

import (
	"net/http"
	"sync"

	shr "github.com/PaulioRandall/qlueless-assembly-line-api/internal/pkg"
)

var orders []shr.WorkItem
var order_loader sync.Once

// OrderHandler handles requests for all orders currently within the service
func OrderHandler(w http.ResponseWriter, r *http.Request) {
	shr.Log_request(r)

	order_loader.Do(loadOrders)
	if orders == nil {
		shr.Http_500(&w)
		return
	}

	reply := shr.Reply{
		Message: "Found all orders",
		Data:    orders,
	}

	shr.WriteJsonReply(reply, w, r)
}

// loadOrders loads all orders into the orders array
func loadOrders() {
	orders = []shr.WorkItem{
		shr.WorkItem{
			Title:        "Outline the saga",
			Description:  "Create a rough outline of the new saga.",
			Work_item_id: 1,
			Tag_id:       "mid",
			Status_id:    "in_progress",
		},
	}
}
