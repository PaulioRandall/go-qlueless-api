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
	Additional  string `json:"additional"`
}

type DictionaryResponse struct {
	Tags []TagEntry `json:"tags"`
}

func DictionaryHandler(w http.ResponseWriter, r *http.Request) {
	response := DictionaryResponse{
		Tags: createTags(),
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
