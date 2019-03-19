package orders

import (
	"strconv"

	. "github.com/PaulioRandall/go-qlueless-assembly-api/internal/pkg"
)

var orders = map[string]Thing{}

// MapToOrder converts a map representing an order to an order struct
func mapToOrder(m map[string]interface{}) Thing {
	return Thing{
		Description: ValueOrEmpty(m, "description"),
		ID:          ValueOrEmpty(m, "id"),
		ParentID:    ValueOrEmpty(m, "parent_id"),
		State:       ValueOrEmpty(m, "state"),
		Additional:  ValueOrEmpty(m, "additional"),
	}
}

// AddOrder adds a new order to the data store returning the newly assigned ID
func addOrder(o Thing) (string, error) {
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
	o.ID = strconv.Itoa(next)
	orders[o.ID] = o
	return o.ID, nil
}

// CreateDummyOrders creates some dummy orders for testing during these initial
// phases of development
func CreateDummyOrders() {
	orders["1"] = Thing{
		Description: "# Outline the saga\nCreate a rough outline of the new saga.",
		ID:          "1",
		State:       "in_progress",
	}
}
