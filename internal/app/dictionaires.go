// Package internal/app contains non-reusable internal application code
package app

import (
	"encoding/json"
	"log"
	"net/http"

	shr "github.com/PaulioRandall/qlueless-assembly-line-api/internal/pkg"
)

// A TagEntry represents a single tag dictionary entry
type TagEntry struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Tag_id      string `json:"tag_id"`
	Additional  string `json:"additional,omitempty"`
}

// A StatusEntry represents a single status dictionary entry
type StatusEntry struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Status_id   string `json:"status_id"`
	Additional  string `json:"additional,omitempty"`
}

// A WorkItemTypeEntry represents a single work item type dictionary entry
type WorkItemTypeEntry struct {
	Title             string `json:"title"`
	Description       string `json:"description"`
	Work_item_type_id string `json:"work_item_type_id"`
	Additional        string `json:"additional,omitempty"`
}

// A DictionaryData holds all dictionaries for the service
type DictionaryData struct {
	Tags            []TagEntry          `json:"tags"`
	Statuses        []StatusEntry       `json:"statuses"`
	Work_item_types []WorkItemTypeEntry `json:"work_item_types"`
}

// DictionaryHandler implements the Go web server Handler interface to return a
// full map of all dictionaries held by the system
func DictionaryHandler(w http.ResponseWriter, r *http.Request) {
	response := shr.Reply{
		Message: "All service dictionaries and their entries",
		Data: DictionaryData{
			Tags:            createTags(),
			Statuses:        createStatuses(),
			Work_item_types: createWorkItemTypes(),
		},
	}

	shr.AppendJSONHeaders(w)
	json.NewEncoder(w).Encode(response)
	log.Println(r.Host)
}

// createTags returns a hardcoded array of all tag dictionary entries
func createTags() []TagEntry {
	return []TagEntry{
		TagEntry{
			Title:       "Low priority",
			Description: "Low priority work items.",
			Tag_id:      "low",
			Additional:  "colour: #0000FF",
		},
		TagEntry{
			Title:       "Mid priority",
			Description: "Mid priority work items.",
			Tag_id:      "mid",
			Additional:  "colour: #00FF00",
		},
		TagEntry{
			Title:       "High priority",
			Description: "High priority work items.",
			Tag_id:      "high",
			Additional:  "colour: #FF0000",
		},
	}
}

// createStatuses returns a hardcoded array of all status dictionary entries
func createStatuses() []StatusEntry {
	return []StatusEntry{
		StatusEntry{
			Title: "Potential",
			Description: "Potential work items that may or may not be worked on," +
				" i.e. a decision has yet to be made whether the item should exist" +
				" or maybe when it should be done hasn't been decided.",
			Status_id: "potential",
		},
		StatusEntry{
			Title: "Queued",
			Description: "Work items that have been started but are not currently" +
				" being worked on, they are waiting in a queue at a workstation.",
			Status_id: "queued",
		},
		StatusEntry{
			Title: "In Progress",
			Description: "Work items that have been started and are currently being" +
				" worked on.",
			Status_id: "in_progress",
		},
		StatusEntry{
			Title: "Dispatched",
			Description: "Work items that have been completed but are NOT yet" +
				" generating value (being used, available to customers, etc).",
			Status_id: "dispatched",
		},
		StatusEntry{
			Title: "Delivered",
			Description: "Work items that have been completed and are generating" +
				" value (being used, available to customers, etc).",
			Status_id: "delivered",
		},
	}
}

// createWorkItemTypes returns a hardcoded array of all work item type
// dictionary entries
func createWorkItemTypes() []WorkItemTypeEntry {
	return []WorkItemTypeEntry{
		WorkItemTypeEntry{
			Title: "Order",
			Description: "An order to be processed that will be split up into" +
				" one or many batches. Each order will typically be done by one" +
				" person who will do the breaking up into batches whilst working" +
				" through the order, at the start or as and when needed.",
			Work_item_type_id: "order",
		},
		WorkItemTypeEntry{
			Title: "Batch",
			Description: "A batch is a single unit of work. In a production line" +
				" then the batch will be the processing of N number of items. In" +
				" software it will be a single VCS commit-push to the shared" +
				" repository.",
			Work_item_type_id: "batch",
		},
	}
}
