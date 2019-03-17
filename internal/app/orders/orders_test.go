package orders

import (
	"testing"

	. "github.com/PaulioRandall/go-qlueless-assembly-api/internal/pkg"
	"github.com/stretchr/testify/assert"
)

// When invoked, should not return nil
func TestLoadOrders___1(t *testing.T) {
	assert.NotNil(t, LoadOrders())
}

// When invoked, should return array of valid orders
func TestLoadOrders___2(t *testing.T) {
	act := LoadOrders()
	for _, o := range *act {
		CheckOrder(t, o)
	}
}

// When a valid map is provided, a WorkItem is returned
func TestMapToOrder___1(t *testing.T) {
	m := make(map[string]interface{})
	m["description"] = "description"

	act := MapToOrder(m)
	assert.Equal(t, "description", act.Description)
	assert.Empty(t, act.WorkItemID)
	assert.Empty(t, act.ParentWorkItemID)
	assert.Empty(t, act.TagID)
	assert.Empty(t, act.StatusID)
	assert.Empty(t, act.Additional)
}

// When a valid map is provided, a WorkItem is returned
func TestMapToOrder___2(t *testing.T) {
	m := make(map[string]interface{})
	m["description"] = "description"
	m["work_item_id"] = "work_item_id"
	m["parent_work_item_id"] = "parent_work_item_id"
	m["tag_id"] = "tag_id"
	m["status_id"] = "status_id"
	m["additional"] = "abc: xyz; colour: black"

	act := MapToOrder(m)
	assert.Equal(t, "description", act.Description)
	assert.Equal(t, "work_item_id", act.WorkItemID)
	assert.Equal(t, "parent_work_item_id", act.ParentWorkItemID)
	assert.Equal(t, "tag_id", act.TagID)
	assert.Equal(t, "status_id", act.StatusID)
	assert.Equal(t, "abc: xyz; colour: black", act.Additional)
}

func createDummyOrder() WorkItem {
	return WorkItem{
		Description: "# Outline the saga\nCreate a rough outline of the new saga.",
		TagID:       "mid",
		StatusID:    "in_progress",
	}
}

// When given an order, returns an order ID
func TestAddOrder___1(t *testing.T) {
	*orders = make([]WorkItem, 0)
	o := createDummyOrder()
	act, err := AddOrder(o)
	assert.Nil(t, err)
	assert.NotEmpty(t, act)
	o.WorkItemID = act
	assert.Contains(t, *orders, o)
}

// When invoked twice with the same order, returns differnet IDs each time
func TestAddOrder___2(t *testing.T) {
	*orders = make([]WorkItem, 0)

	a := createDummyOrder()
	ID_1, err := AddOrder(a)
	assert.Nil(t, err)
	a.WorkItemID = ID_1

	b := createDummyOrder()
	ID_2, err := AddOrder(b)
	assert.Nil(t, err)
	b.WorkItemID = ID_2

	assert.NotEqual(t, ID_1, ID_2)
	assert.Contains(t, *orders, a)
	assert.Contains(t, *orders, b)
}
