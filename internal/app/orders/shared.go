package orders

import (
	"sync"

	shr "github.com/PaulioRandall/go-qlueless-assembly-api/internal/pkg"
)

var orders []shr.WorkItem
var orderLoader sync.Once

// LoadOrders loads all orders into the orders array
func LoadOrders() []shr.WorkItem {
	orderLoader.Do(createDummyOrders)
	return orders
}

func createDummyOrders() {
	orders = []shr.WorkItem{
		shr.WorkItem{
			Title:        "Outline the saga",
			Description:  "Create a rough outline of the new saga.",
			Work_item_id: "1",
			Tag_id:       "mid",
			Status_id:    "in_progress",
		},
	}
}
