package orders

import (
	"sync"

	. "github.com/PaulioRandall/go-qlueless-assembly-api/internal/pkg"
)

var orders []WorkItem
var orderLoader sync.Once

// LoadOrders loads all orders into the orders array
func LoadOrders() []WorkItem {
	orderLoader.Do(createDummyOrders)
	return orders
}

func createDummyOrders() {
	orders = []WorkItem{
		WorkItem{
			Title:       "Outline the saga",
			Description: "Create a rough outline of the new saga.",
			WorkItemID:  "1",
			TagID:       "mid",
			StatusID:    "in_progress",
		},
	}
}
