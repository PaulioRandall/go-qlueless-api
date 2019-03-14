package pkg

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogIfErr___1(t *testing.T) {
	act := LogIfErr(nil)
	assert.False(t, act)
	// Output:
	//
}

func TestLogIfErr___2(t *testing.T) {
	var err error = errors.New("Computer says no!")
	act := LogIfErr(err)
	assert.True(t, act)
	// Output:
	// Computer says no!
}

func dummyWorkItems() []WorkItem {
	return []WorkItem{
		WorkItem{
			Title:        "Outline the saga",
			Description:  "Create a rough outline of the new saga.",
			Work_item_id: "1",
			Tag_id:       "mid",
			Status_id:    "in_progress",
		},
		WorkItem{
			Title:               "Name the saga",
			Description:         "Think of a name for the saga.",
			Work_item_id:        "2",
			Parent_work_item_id: "1",
			Tag_id:              "mid",
			Status_id:           "potential",
		},
		WorkItem{
			Title:               "Outline the first chapter",
			Description:         "Outline the first chapter.",
			Work_item_id:        "3",
			Parent_work_item_id: "1",
			Tag_id:              "mid",
			Status_id:           "delivered",
			Additional:          "archive_note:Done but not a compelling start",
		},
		WorkItem{
			Title:               "Outline the second chapter",
			Description:         "Outline the second chapter.",
			Work_item_id:        "4",
			Parent_work_item_id: "1",
			Tag_id:              "mid",
			Status_id:           "in_progress",
		},
	}
}

// When given some WorkItems and the ID of an item that appears within those
// WorkItems, does NOT return nil
func TestFindWorkItem___1(t *testing.T) {
	items := dummyWorkItems()
	var w *WorkItem
	w = FindWorkItem(items, "1")
	assert.NotNil(t, w)
	w = FindWorkItem(items, "3")
	assert.NotNil(t, w)
}

// When given some batches and the ID of an item that does NOT appear within
// those batches, returns nil
func TestFindWorkItem___2(t *testing.T) {
	items := dummyWorkItems()
	id := "99999999999"
	w := FindWorkItem(items, id)
	assert.Nil(t, w)
}

// When given some batches and the ID of an item that appears within those
// batches, returns the expected batch
func TestFindWorkItem___3(t *testing.T) {
	items := dummyWorkItems()
	id := "2"
	w := FindWorkItem(items, id)
	assert.Equal(t, id, w.Work_item_id)
}
