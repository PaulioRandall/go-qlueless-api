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

func DictionaryHandler(w http.ResponseWriter, r *http.Request) {

	item := TagEntry{
		Title:       "Low priority",
		Description: "Low priority work items.",
		Tag_id:      "low",
		Additional:  "colour: #0000FF",
	}

	shr.AppendStdHeaders(w)

	json.NewEncoder(w).Encode(item)
	log.Println(r.Host)
}
