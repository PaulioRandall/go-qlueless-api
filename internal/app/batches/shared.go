package batches

import (
	"sync"

	. "github.com/PaulioRandall/go-qlueless-assembly-api/internal/pkg"
)

var batches map[string]WorkItem
var batchLoader sync.Once

// LoadBatches loads all batches into the batches array and then returns the
// array
func LoadBatches() map[string]WorkItem {
	batchLoader.Do(createDummyBatches)
	return batches
}

// createDummyBatches creates some dummy batches
func createDummyBatches() {
	batches = make(map[string]WorkItem)
	batches["2"] = WorkItem{
		Description:      "# Name the saga\nThink of a name for the saga.",
		WorkItemID:       "2",
		ParentWorkItemID: "1",
		TagID:            "mid",
		StatusID:         "potential",
	}
	batches["3"] = WorkItem{
		Description:      "# Outline the first chapter.",
		WorkItemID:       "3",
		ParentWorkItemID: "1",
		TagID:            "mid",
		StatusID:         "delivered",
		Additional:       "archive_note:Done but not a compelling start",
	}
	batches["4"] = WorkItem{
		Description:      "# Outline the second chapter.",
		WorkItemID:       "4",
		ParentWorkItemID: "1",
		TagID:            "mid",
		StatusID:         "in_progress",
	}
}
