package app

import (
	"encoding/json"
	"log"
	"net/http"

	shr "github.com/PaulioRandall/qlueless-assembly-line-api/internal/pkg"
)

type TagEntry struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Tag_id      string `json:"tag_id"`
	Additional  string `json:"additional,omitempty"`
}

type StatusEntry struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Status_id   string `json:"status_id"`
	Additional  string `json:"additional,omitempty"`
}

type DictionaryResponse struct {
	Tags     []TagEntry    `json:"tags"`
	Statuses []StatusEntry `json:"statuses"`
}

func DictionaryHandler(w http.ResponseWriter, r *http.Request) {
	response := DictionaryResponse{
		Tags:     createTags(),
		Statuses: createStatuses(),
	}

	shr.AppendStdHeaders(w)
	json.NewEncoder(w).Encode(response)
	log.Println(r.Host)
}

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

func createStatuses() []StatusEntry {
	return []StatusEntry{
		StatusEntry{
			Title:       "Potential",
			Description: "Potential work items that may or may not be worked on, i.e. a decision has yet to be made whether the item should exist or maybe when it should be done hasn't been decided.",
			Status_id:   "potential",
		},
		StatusEntry{
			Title:       "Queued",
			Description: "Work items that have been started but are not currently being worked on, they are waiting in a queue at a workstation.",
			Status_id:   "queued",
		},
		StatusEntry{
			Title:       "In Progress",
			Description: "Work items that have been started and are currently being worked on.",
			Status_id:   "in_progress",
		},
		StatusEntry{
			Title:       "Dispatched",
			Description: "Work items that have been completed but are NOT yet generating value (being used, available to customers, etc).",
			Status_id:   "dispatched",
		},
		StatusEntry{
			Title:       "Delivered",
			Description: "Work items that have been completed and are generating value (being used, available to customers, etc).",
			Status_id:   "delivered",
		},
	}
}
