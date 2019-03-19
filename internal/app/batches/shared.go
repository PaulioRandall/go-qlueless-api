package batches

import (
	. "github.com/PaulioRandall/go-qlueless-assembly-api/internal/pkg"
)

var batches = make(map[string]Thing)

// CreateDummyBatches creates some dummy batches for testing during these
// initial phases of development
func CreateDummyBatches() {
	batches["2"] = Thing{
		Description: "# Name the saga\nThink of a name for the saga.",
		ID:          "2",
		ParentID:    "1",
		State:       "ready",
	}
	batches["3"] = Thing{
		Description: "# Outline the first chapter.",
		ID:          "3",
		ParentID:    "1",
		State:       "finished",
		Additional:  "archive_note:Done but not a compelling start",
	}
	batches["4"] = Thing{
		Description: "# Outline the second chapter.",
		ID:          "4",
		ParentID:    "1",
		State:       "in_progress",
	}
}
