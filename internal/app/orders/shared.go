package orders

import (
	"sync"

	shr "github.com/PaulioRandall/qlueless-assembly-line-api/internal/pkg"
)

var orders []shr.WorkItem
var order_loader sync.Once

// loadOrders loads all orders into the orders array
func Load_orders() []shr.WorkItem {
	order_loader.Do(createDummyOrders)
	return orders
}

func createDummyOrders() {
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
