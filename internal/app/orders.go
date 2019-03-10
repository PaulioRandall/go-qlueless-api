package app

import (
	"encoding/json"
	"log"
	"net/http"

	shr "github.com/PaulioRandall/qlueless-assembly-line-api/internal/pkg"
)

func OrderHandler(w http.ResponseWriter, r *http.Request) {
	response := shr.Reply{
		Message: "Found dummy orders",
		Data:    createDummyOrder(),
	}

	shr.AppendJSONHeaders(w)
	json.NewEncoder(w).Encode(response)
	log.Println(r.Host)
}

func createDummyOrder() []shr.WorkItem {
	return []shr.WorkItem{
		shr.WorkItem{
			Title:        "Outline the saga",
			Description:  "Create a rough outline of the new saga.",
			Work_item_id: 1,
			Tag_id:       "mid",
			Status_id:    "in_progress",
		},
	}
}
