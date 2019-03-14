package orders

import (
	"fmt"
	"strconv"
	"strings"
	"testing"

	shr "github.com/PaulioRandall/qlueless-assembly-line-api/internal/pkg"
)

var t *testing.T
var f_name string

func prep_test(tester *testing.T, func_name string) {
	t = tester
	f_name = func_name
}

func fail(exp interface{}, act interface{}) {
	var m = fmt.Sprintf("%s... Expected: %v... Actually: %v", f_name, exp, act)
	t.Errorf(m)
}

func check_not_empty(v string, v_name string) {
	if strings.TrimSpace(v) == "" {
		fail(fmt.Sprintf("%v to not be empty", v_name), "was empty")
	}
}

func check_is_number(v string, v_name string) {
	n, err := strconv.Atoi(v)
	if err != nil {
		fail(fmt.Sprintf("%v to be a number", v_name), "was not a number")
	}
	s := strconv.Itoa(n)
	if v != s {
		fail(fmt.Sprintf("%v to be a number", v_name), "contained additional characters")
	}
}

func check_work_item(w shr.WorkItem) {
	check_not_empty(w.Title, "'WorkItem.Title'")
	check_not_empty(w.Description, "'WorkItem.Description'")
	check_not_empty(w.Work_item_id, "'WorkItem.Work_item_id'")
	check_is_number(w.Work_item_id, "'WorkItem.Work_item_id'")
	check_not_empty(w.Tag_id, "'WorkItem.Tag_id'")
	check_not_empty(w.Status_id, "'WorkItem.Status_id'")
}

func check_order(o shr.WorkItem) {
	check_work_item(o)
}

func check_batch(b shr.WorkItem) {
	check_work_item(b)
	check_not_empty(b.Parent_work_item_id, "'WorkItem.Parent_work_item_id'")
	check_is_number(b.Parent_work_item_id, "'WorkItem.Parent_work_item_id'")
}

// When invoked, should not return nil
func TestLoad_orders___1(t *testing.T) {
	prep_test(t, "Load_orders()")
	act := Load_orders()
	if act == nil {
		fail("not nil", nil)
	}
}

// When invoked, should return array with length > 0
func TestLoad_orders___2(t *testing.T) {
	prep_test(t, "Load_orders()")
	var act []shr.WorkItem = Load_orders()
	for _, o := range act {
		check_order(o)
	}
}

// When given some orders and the ID of an item that appears within those
// orders, does NOT return nil
func TestFind_order___1(t *testing.T) {
	prep_test(t, "find_order()")
	createDummyOrders()
	id := "1"
	o := find_order(orders, id)
	if o == nil {
		fail("not nil", nil)
	}
}

// When given some orders and the ID of an item that does NOT appear within
// those orders, returns nil
func TestFind_order___2(t *testing.T) {
	prep_test(t, "find_order()")
	createDummyOrders()
	id := "99999999999"
	o := find_order(orders, id)
	if o != nil {
		fail(nil, "not nil")
	}
}

// When given some orders and the ID of an item that appears within those
// orders, returns the expected order
func TestFind_order___3(t *testing.T) {
	prep_test(t, "find_order()")
	createDummyOrders()
	id := "1"
	o := find_order(orders, id)
	if o.Work_item_id != id {
		fail(fmt.Sprintf("WorkItem with ID %v", id),
			fmt.Sprintf("WorkItem with ID %v", o.Work_item_id))
	}
}
