package orders

import (
	"strconv"
	"strings"
	"testing"

	shr "github.com/PaulioRandall/qlueless-assembly-line-api/internal/pkg"
	"github.com/stretchr/testify/assert"
)

func checkIsNumber(t *testing.T, s string, m ...interface{}) {
	n, err := strconv.Atoi(s)
	assert.Nil(t, err, "Expected string to be a number")
	s2 := strconv.Itoa(n)
	assert.Equal(t, s, s2, "Expected stringified number to equal the original string")
}

func checkNotBlank(t *testing.T, s string, m ...interface{}) {
	v := strings.TrimSpace(s)
	assert.NotEmpty(t, v, m)
}

func checkWorkItem(t *testing.T, w shr.WorkItem) {
	checkNotBlank(t, w.Title, "WorkItem.Title")
	checkNotBlank(t, w.Description, "WorkItem.Description")
	checkNotBlank(t, w.Work_item_id, "WorkItem.Work_item_id")
	checkIsNumber(t, w.Work_item_id, "WorkItem.Work_item_id")
	checkNotBlank(t, w.Tag_id, "WorkItem.Tag_id")
	checkNotBlank(t, w.Status_id, "WorkItem.Status_id")
}

func checkOrder(t *testing.T, o shr.WorkItem) {
	checkWorkItem(t, o)
}

func checkBatch(t *testing.T, b shr.WorkItem) {
	checkWorkItem(t, b)
	checkNotBlank(t, b.Parent_work_item_id, "WorkItem.Parent_work_item_id")
	checkIsNumber(t, b.Parent_work_item_id, "WorkItem.Parent_work_item_id")
}

// When invoked, should not return nil
func TestLoad_orders___1(t *testing.T) {
	assert.NotNil(t, Load_orders())
}

// When invoked, should return array with length > 0
func TestLoad_orders___2(t *testing.T) {
	var act []shr.WorkItem = Load_orders()
	for _, o := range act {
		checkOrder(t, o)
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
