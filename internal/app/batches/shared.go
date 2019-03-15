package batches

import (
	"sync"

	shr "github.com/PaulioRandall/go-qlueless-assembly-api/internal/pkg"
)

var batches []shr.WorkItem
var batchLoader sync.Once

// LoadBatches loads all batches into the batches array and then returns the
// array
func LoadBatches() []shr.WorkItem {
	batchLoader.Do(createDummyBatches)
	return batches
}

// createDummyBatches creates some dummy batches
func createDummyBatches() {
	batches = []shr.WorkItem{
		shr.WorkItem{
			Title:               "Name the saga",
			Description:         "Think of a name for the saga.",
			Work_item_id:        "2",
			Parent_work_item_id: "1",
			Tag_id:              "mid",
			Status_id:           "potential",
		},
		shr.WorkItem{
			Title:               "Outline the first chapter",
			Description:         "Outline the first chapter.",
			Work_item_id:        "3",
			Parent_work_item_id: "1",
			Tag_id:              "mid",
			Status_id:           "delivered",
			Additional:          "archive_note:Done but not a compelling start",
		},
		shr.WorkItem{
			Title:               "Outline the second chapter",
			Description:         "Outline the second chapter.",
			Work_item_id:        "4",
			Parent_work_item_id: "1",
			Tag_id:              "mid",
			Status_id:           "in_progress",
		},
	}
}
