package orders

import (
	"testing"

	shr "github.com/PaulioRandall/qlueless-assembly-line-api/internal/pkg"
	"github.com/stretchr/testify/assert"
)

// When invoked, should not return nil
func TestLoad_orders___1(t *testing.T) {
	assert.NotNil(t, Load_orders())
}

// When invoked, should return array with length > 0
func TestLoad_orders___2(t *testing.T) {
	var act []shr.WorkItem = Load_orders()
	for _, o := range act {
		shr.CheckOrder(t, o)
	}
}

// When given some orders and the ID of an item that appears within those
// orders, does NOT return nil
func TestFind_order___1(t *testing.T) {
	createDummyOrders()
	id := "1"
	o := find_order(orders, id)
	assert.NotNil(t, o)
}

// When given some orders and the ID of an item that does NOT appear within
// those orders, returns nil
func TestFind_order___2(t *testing.T) {
	createDummyOrders()
	id := "99999999999"
	o := find_order(orders, id)
	assert.Nil(t, o)
}

// When given some orders and the ID of an item that appears within those
// orders, returns the expected order
func TestFind_order___3(t *testing.T) {
	createDummyOrders()
	id := "1"
	o := find_order(orders, id)
	assert.Equal(t, id, o.Work_item_id)
}
