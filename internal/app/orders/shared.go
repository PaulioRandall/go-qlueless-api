package orders

import (
	"strconv"
	"sync"

	. "github.com/PaulioRandall/go-qlueless-assembly-api/internal/pkg"
)

var orders map[string]WorkItem = nil
var orderLoader sync.Once

// LoadOrders loads all orders into the orders array
func LoadOrders() map[string]WorkItem {
	orderLoader.Do(createDummyOrders)
	return orders
}

// MapToOrder converts a map representing an order to an order struct
func MapToOrder(m map[string]interface{}) WorkItem {
	return WorkItem{
		Description:      ValueOrEmpty(m, "description"),
		WorkItemID:       ValueOrEmpty(m, "work_item_id"),
		ParentWorkItemID: ValueOrEmpty(m, "parent_work_item_id"),
		TagID:            ValueOrEmpty(m, "tag_id"),
		StatusID:         ValueOrEmpty(m, "status_id"),
		Additional:       ValueOrEmpty(m, "additional"),
	}
}

// AddOrder adds a new order to the data store returning the newly assigned ID
func AddOrder(o WorkItem) (string, error) {
	next := 1
	for k, _ := range orders {
		ID, err := strconv.Atoi(k)
		if err != nil {
			return "", nil
		}

		if ID > next {
			next = ID
		}
	}

	next++
	o.WorkItemID = strconv.Itoa(next)
	orders[o.WorkItemID] = o
	return o.WorkItemID, nil
}

func createDummyOrders() {
	orders = map[string]WorkItem{}
	orders["1"] = WorkItem{
		Description: "# Outline the saga\nCreate a rough outline of the new saga.",
		WorkItemID:  "1",
		TagID:       "mid",
		StatusID:    "in_progress",
	}
}
